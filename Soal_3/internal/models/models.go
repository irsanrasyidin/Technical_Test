package models

import "encoding/xml"

type UserData struct {
	ID       int    `gorm:"primaryKey;column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"-"`
}

func (UserData) TableName() string {
	return "users"
}

type SOAPRequest struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    SOAPBody `xml:"Body"`
}

type SOAPBody struct {
	GetUserDetails GetUserDetails `xml:"GetUserDetails"`
}

type GetUserDetails struct {
	UserID string `xml:"userID"`
}
