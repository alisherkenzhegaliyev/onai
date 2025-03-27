package models

type Family struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Household Household `gorm:"embedded" json:"household"`
	Parent1   Parent    `gorm:"embedded" json:"parent1"`
	Parent2   Parent    `gorm:"embedded" json:"parent2"`
	Siblings  []Sibling `gorm:"foreignKey:ProfileID;constraint:OnDelete:CASCADE" json:"siblings"`
}

type Household struct {
	MaritalStatus     string `gorm:"not null" json:"marital_status"`
	PermanentHomeWith string `gorm:"not null" json:"permanent_home_with"`
}

type Parent struct {
	Role          string `gorm:"not null" json:"role"` // mother / father
	IsAlive       bool   `gorm:"not null" json:"is_alive"`
	FirstName     string `gorm:"not null" json:"first_name"`
	SecondName    string `gorm:"not null" json:"second_name"`
	HighestDegree string `gorm:"not null" json:"highest_degree"`
}

type Sibling struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ProfileID  uint   `gorm:"not null" json:"profile_id"` // Foreign key
	FirstName  string `gorm:"not null" json:"first_name"`
	SecondName string `gorm:"not null" json:"second_name"`
	Age        int    `gorm:"not null" json:"age"`
}
