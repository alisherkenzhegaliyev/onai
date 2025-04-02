package models

type Profile struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"` // DATE in DB, string in Go
	LegalSex    string `json:"legal_sex"`
	Citizen     string `json:"citizen"`
	AddressID   int    `json:"address_id"`
	ContactID   int    `json:"contact_id"`
}

type Address struct {
	ID      int    `json:"id"`
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
}

type ContactDetail struct {
	ID           int    `json:"id"`
	CountryCode  string `json:"country_code"`
	MobileNumber string `json:"mobile_number"`
}

type Language struct {
	ID          int    `json:"id"`
	ProfileID   int    `json:"profile_id"`
	Language    string `json:"language"`
	Proficiency string `json:"proficiency"`
}

// SQL below

// CREATE TABLE address (
//     id SERIAL PRIMARY KEY,
//     country VARCHAR(100),
//     city VARCHAR(100),
//     street VARCHAR(150)
// );

// CREATE TABLE contact_detail (
//     id SERIAL PRIMARY KEY,
//     country_code VARCHAR(15),
//     mobile_number VARCHAR(25)
// );

// CREATE TABLE profile (
//     id SERIAL PRIMARY KEY,
//     first_name VARCHAR(100),
//     last_name VARCHAR(100),
//     date_of_birth DATE,
//     legal_sex VARCHAR(15),
//     citizen VARCHAR(100),
//     address_id INT,
//     contact_id INT,
//     FOREIGN KEY (address_id) REFERENCES address(id) ON DELETE CASCADE,
//     FOREIGN KEY (contact_id) REFERENCES contact_detail(id) ON DELETE CASCADE
// );

// CREATE TABLE language (
//     id SERIAL PRIMARY KEY,
//     profile_id INT,
//     language VARCHAR(60),
//     proficiency VARCHAR(60),
//     FOREIGN KEY (profile_id) REFERENCES profile(id) ON DELETE CASCADE
// );
