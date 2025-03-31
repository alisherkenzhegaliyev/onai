package models

type Profile struct {
    ID          int     `json:"id"`
    FirstName   string  `json:"first_name"`
    LastName    string  `json:"last_name"`
    DateOfBirth string  `json:"date_of_birth"` // DATE in DB, string in Go
    LegalSex    string  `json:"legal_sex"`
    Citizen     string  `json:"citizen"`
    AddressID   int     `json:"address_id"`
    ContactID   int     `json:"contact_id"`
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