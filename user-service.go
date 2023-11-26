package models

import "gorm.io/gorm"

// Add more relevant fields
type User struct {
	gorm.Model
	LinkedinEmail  string
	AlternateEmail string
	Name           string
	PhoneNumber    string
}

type UserSkills struct {
	EmailId string   `bson:"_id,required"`
	Skills  []string `bson:"skills,omitempty"`
	Resume  string   `bson:"resume_url,omitempty"`
}
