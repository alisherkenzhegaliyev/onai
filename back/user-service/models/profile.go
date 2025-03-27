package models

// Profile Model
type Profile struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	PersonalInfo   PersonalInfo   `gorm:"embedded" json:"personal_info"`
	Address        Address        `gorm:"embedded" json:"address"`
	ContactDetails ContactDetails `gorm:"embedded" json:"contact_details"`
	LegalSex       string         `gorm:"not null" json:"legal_sex"`
	Languages      []Language     `gorm:"foreignKey:ProfileID;constraint:OnDelete:CASCADE" json:"languages"`
	Citizen        string         `gorm:"not null" json:"citizen"`
}

// PersonalInfo Model (Embedded in Profile)
type PersonalInfo struct {
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	DOB       string `gorm:"not null" json:"date_of_birth"`
}

// Address Model (Embedded in Profile)
type Address struct {
	Country string `gorm:"not null" json:"country"`
	City    string `gorm:"not null" json:"city"`
}

// ContactDetails Model (Embedded in Profile)
type ContactDetails struct {
	CountryCode  string `gorm:"not null" json:"country_code"`
	MobileNumber string `gorm:"not null" json:"mobile_number"`
}

// Language Model (Separate table for multiple languages)
type Language struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ProfileID   uint   `gorm:"not null" json:"profile_id"` // Foreign key
	Language    string `gorm:"not null" json:"language"`
	Proficiency string `gorm:"not null" json:"proficiency"` // Can be Beginner, Intermediate, Fluent
}
