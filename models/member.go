package models

import (
	"time"
)

type Member struct {
	ID               int       `json:"id" gorm:"type:serial;primary_key"`
	Username         string    `json:"username" gorm:"type:varchar(64);unique;not null"`
	Passwd           string    `json:"passwd" gorm:"type:varchar(64);not null"`
	Email            string    `json:"email" gorm:"type:varchar(255);primary_key;not null"`
	Phone            string    `json:"phone" gorm:"type:varchar(255);unique;not null"`
	Address1         string    `json:"address1" gorm:"type:varchar(255)"`
	Address2         string    `json:"address2" gorm:"type:varchar(255)"`
	PostalCode       string    `json:"postal_code" gorm:"type:varchar(10)"`
	CountryID        int       `json:"country_id" gorm:"type:integer;default:999"`
	ProvinceID       int       `json:"province_id" gorm:"type:integer;default:999"`
	DistrictID       int       `json:"district_id" gorm:"type:integer;default:999"`
	SubdistrictID    int       `json:"subdistrict_id" gorm:"type:integer;default:999"`
	Organization     string    `json:"organization" gorm:"type:varchar(255)"`
	RegistrationTime time.Time `json:"registration_time" gorm:"type:timestamp;not null;default:current_timestamp"`
	ConfirmationTime time.Time `json:"confirmation_time" gorm:"type:timestamp"`
	UpdateTime       time.Time `json:"update_time" gorm:"type:timestamp"`
	Notes            string    `json:"notes" gorm:"type:text"`
}
