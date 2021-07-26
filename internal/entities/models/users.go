package models

import (
	"github.com/tahmooress/motor-shop/internal/pkg/query"
	"regexp"
)

type Customer struct {
	ID           ID             `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	LastName     string         `json:"last_name,omitempty"`
	CompanyName  string         `json:"company_name,omitempty"`
	Mobile       string         `json:"mobile,omitempty"`
	NationalCode string         `json:"national_code,omitempty"`
	CreatedAt    query.NullTime `json:"created_at,omitempty"`
	UpdatedAt    query.NullTime `json:"updated_at,omitempty"`
}

type Admin struct {
	ID        ID             `json:"id,omitempty" db:"admin_users.id"`
	UserName  string         `json:"user_name,omitempty"`
	Password  string         `json:"password,omitempty"`
	Shops     []Shop         `json:"shops,omitempty"`
	CreatedAt query.NullTime `json:"created_at,omitempty" db:"admin_users.created_at"`
	UpdatedAt query.NullTime `json:"updated_at,omitempty" db:"admin_users.updated_at"`
}

func MobileValidate(number string) error {
	match, err := regexp.MatchString("^09[0-9]{9}$", number)
	if err != nil || !match {
		return ErrMobile
	}

	return nil
}
