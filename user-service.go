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
	Resume  []Resume `bson:"resume,omitempty"`
}

type Resume struct {
	ResumeUrl     string `bson:"resume_url,required"`
	ResumeVersion string `bson:"resume_version,required"`
}

type FileMetadata struct {
	gorm.Model
	EmailId     string
	S3ObjectKey string
	FileVersion string
}
