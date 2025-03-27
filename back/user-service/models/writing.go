package models

type Writing struct {
    ID                         uint   `gorm:"primaryKey" json:"id"`
    PersonalEssay              string `gorm:"not null" json:"personal_essay"`
    AdditionalInformationEssay string `gorm:"not null" json:"additional_information_essay"`
}
