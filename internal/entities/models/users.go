package models

import (
	"regexp"
	"time"
)

type Customer struct {
	ID           ID        `json:"id"`
	Name         string    `json:"name"`
	LastName     string    `json:"last_name"`
	CompanyName  string    `json:"company_name"`
	Mobile       string    `json:"mobile"`
	NationalCode string    `json:"national_code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Admin struct {
	ID        ID        `json:"id,omitempty" db:"admin_users.id"`
	UserName  string    `json:"user_name,omitempty"`
	Password  string    `json:"password,omitempty"`
	Shops     []Shop    `json:"shops,omitempty"`
	CreatedAt time.Time `json:"created_at" db:"admin_users.created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"admin_users.updated_at"`
}

func MobileValidate(number string) error {
	match, err := regexp.MatchString("^09[0-9]{9}$", number)
	if err != nil || !match {
		return ErrMobile
	}

	return nil
}
