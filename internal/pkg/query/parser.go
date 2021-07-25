package query

import (
	"reflect"
	"strings"
)

// ParseFilters take the params filter parse them,
// then create and return a Query string.
func (q *Query) parseFilters(filters map[string]map[string][]interface{}) string {
	query := "WHERE "

	// range trough filters to create WHERE and WHERE IN clauses.
	for key, value := range filters {
		for k, v := range value {
			// look up Operations from operations map and use '=' as default.
			opr, ok := q.Operations[k]
			if !ok {
				opr = q.Operations["eq"]
			}

			query += q.where(Where{
				Field:     key,
				Values:    v,
				Operation: opr,
			})
		}

		query += "AND "
	}

	// trim redundant AND / WHERE at the end of Query if exist any.
	query = strings.TrimSuffix(query, "AND ")
	query = strings.TrimSuffix(query, "WHERE ")

	return query
}

// parseArgs parse the arguments and convert them from
// string  to the desired type.
func (q *Query) parseArgs(args []interface{}) {
	for i := range args {
		if args[i] == "true" {
			args[i] = true
		} else if args[i] == "false" {
			args[i] = false
		}
	}
}

// sanitizeFilters range trough filters and delete invalid filters.
func (q *Query) sanitizeFilters(result interface{},
	selects []string) (map[string]map[string][]interface{}, []sort) {
	possibleFilters := make(map[string]string)

	for i := range selects {
		possibleFilters[dotSeparate(selects[i])] = selects[i]
	}

	tagToField := internalToExternalMapper(result)
	sanitizedFilters := q.sanitizeFilterQuery(possibleFilters, tagToField)

	for _, w := range q.wheres {
		m := make(map[string][]interface{})
		m[w.Operation] = w.Values
		sanitizedFilters[w.Field] = m
	}

	sanitizedSorts := q.sanitizeSortQuery(possibleFilters, tagToField)

	return sanitizedFilters, sanitizedSorts
}

func (q *Query) sanitizeSortQuery(possibleFilters, tagToField map[string]string) []sort {
	sorts := make([]sort, 0)

	for sortType, columns := range q.QueryFilters.Sort {
		for _, column := range columns {
			clm, ok := tagToField[column]
			if ok {
				possibleClm, ok := possibleFilters[clm]
				if ok && validSortType(sortType) {
					sorts = append(sorts, sort{
						strings.ToLower(sortType),
						possibleClm,
					})
				}
			}
		}
	}

	validSorts := make([]sort, 0)

	for _, s := range q.sorts {
		clm, ok := tagToField[s.field]
		if ok {
			possibleClm, ok := possibleFilters[clm]
			if ok && validSortType(s.sortType) {
				validSorts = append(validSorts, sort{sortType: s.sortType, field: possibleClm})
			}
		}
	}

	return shift(sorts, validSorts)
}

func (q *Query) sanitizeFilterQuery(possibleFilters, tagToField map[string]string) map[string]map[string][]interface{} {
	sanitizedFilters := make(map[string]map[string][]interface{})

	for key, values := range q.QueryFilters.Filter {
		field, ok := tagToField[key]
		if ok {
			filter, ok := possibleFilters[field]
			if ok {
				sanitizedFilters[filter] = convertStringSliceToInterface(values)
			}
		}
	}

	return sanitizedFilters
}

func internalToExternalMapper(result interface{}) map[string]string {
	tagToField := make(map[string]string)

	value := reflect.ValueOf(result)
	t := value.Type()

	underLyingStruct := t.Elem().Elem()

	for i := 0; i < underLyingStruct.NumField(); i++ {
		field := toSnakeCase(underLyingStruct.Field(i).Name)
		if field == "" {
			continue
		}

		if trimJSONTag(underLyingStruct.Field(i).Tag.Get("json")) != "" {
			tagToField[trimJSONTag(underLyingStruct.Field(i).Tag.Get("json"))] = field
		}
	}

	return tagToField
}

// dotSeparate check if string is in x.y format return y.
func dotSeparate(str string) string {
	if strings.Contains(str, ".") {
		return str[strings.Index(str, ".")+1:]
	}

	return str
}

func convertStringSliceToInterface(entryMap map[string][]string) map[string][]interface{} {
	result := make(map[string][]interface{})

	for key, value := range entryMap {
		s := make([]interface{}, 0)
		for i := range value {
			s = append(s, value[i])
		}

		result[key] = s
	}

	return result
}
