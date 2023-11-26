package models

import "gorm.io/gorm"

// Records user's preference of the companies and roles they want. "*" in the Company column means "any company".
// Scheduled cron job can run every day to get updates. Reduced cron job time intervals can be for paid plan.
type JobPreference struct {
	gorm.Model
	UserId         uint
	RolePreference string
	Company        string
	ExpectedCTC    int64
}

type Tabler interface {
	TableName() string
}

func (JobAvailable) TableName() string {
	return "jobs_available"
}

// Jobs scraped from LinkedIn periodically so that we don't have to scrape everytime. This is a sort of database so that users can access data post aggregation from LinkedIn.
type JobAvailable struct {
	gorm.Model
	JobProfile
	JobLocation
	JobCompany
}

type JobProfile struct {
	LinkedInInternalId string
	JobDescription     string
	ApplyUrl           string
	Title              string
	SeniorityLevel     string
	Industry           string
	EmploymentType     string
	JobFunctions       string
	TotalApplicants    int
}

type JobProfileResponse struct {
	JobProfile
	Location JobLocation
	Company  JobCompany
}
type JobLocation struct {
	Country    string
	Region     string
	City       string
	PostalCode string
	Latitude   string
	Longitude  string
}

type JobCompany struct {
	Name string
	Url  string
	Logo string
}
