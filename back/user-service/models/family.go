package models

type Family struct {
    ID                uint      `gorm:"primaryKey"`
    Household         Household `gorm:"embedded"`
    Parent1           Parent    `gorm:"embedded;embeddedPrefix:parent1_"`
    Parent2           Parent    `gorm:"embedded;embeddedPrefix:parent2_"`
    Siblings          []Sibling `gorm:"foreignKey:FamilyID;constraint:OnDelete:CASCADE"`
}

type Household struct {
    MaritalStatus     string `gorm:"not null"` // Required
    PermanentHomeWith string `gorm:"not null"` // Required
}

type Parent struct {
    Role          string `gorm:"not null;check:role IN ('mother', 'father')"` // Only 'mother' or 'father'
    IsAlive       bool   `gorm:"not null"`
    FirstName     string `gorm:"not null"`
    SecondName    string `gorm:"not null"`
    HighestDegree string `gorm:"not null"`
}

type Sibling struct {
    ID         uint   `gorm:"primaryKey"`
    FamilyID   uint   `gorm:"index;not null"`
    FirstName  string `gorm:"not null"`
    SecondName string `gorm:"not null"`
    Age        int    `gorm:"not null"`
}
