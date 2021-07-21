package models

import "time"

type Customer struct {
	ID           ID           `json:"id"`
	Name         string       `json:"name"`
	LastName     string       `json:"last_name"`
	CompanyName  string       `json:"company_name"`
	Mobile       Mobile       `json:"mobile"`
	NationalCode NationalCode `json:"national_code"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type Admin struct {
	ID            ID             `json:"id,omitempty"`
	UserName      string         `json:"user_name,omitempty"`
	Password      string         `json:"password,omitempty"`
	Accessibility []ShopIdentity `json:"accessibility,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type Mobile string

type NationalCode string

func (m Mobile) Validate() error {
	return nil
}

//match, err := regexp.MatchString("^09[0-9]{9}$", number)
//if err != nil || !match {
//return structs.ErrMobile
//}
//
//return nil

func (m NationalCode) Validate() error {
	return nil
}

func NewMobile(mobileNumber string) (Mobile, error) {
	return Mobile(mobileNumber), nil
}

func NewNationalCode(nationalCode string) (NationalCode, error) {
	return NationalCode(nationalCode), nil
}
