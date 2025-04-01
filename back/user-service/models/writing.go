package models

type Writing struct {
    ID                         int    `json:"id"`
    UserID                     int    `json:"user_id"`
    PersonalEssay              string `json:"personal_essay"`
    AdditionalInformationEssay string `json:"additional_information_essay"`
}

// SQL below

// CREATE TABLE writing (
//     id SERIAL PRIMARY KEY,
//     user_id INT,
//     personal_essay TEXT,
//     additional_information_essay TEXT
// );
