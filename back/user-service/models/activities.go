package models

type Activities struct {
	ID            uint `gorm:"primaryKey" json:"id"`
	ProfileID     uint `gorm:"not null"   json:"profile_id"` // Foreign key to Profile
	HasActivities bool `json:"has_activities"`

	// If HasActivities is true, the user can fill in multiple activity entries:
	ActivityList []Activity `gorm:"foreignKey:ActivitiesID;constraint:OnDelete:CASCADE" json:"activity_list"`
}

// Each Activity record
type Activity struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	ActivitiesID uint   `gorm:"not null" json:"activities_id"`
	ActivityType string `gorm:"size:50"   json:"activity_type"`
	Position     string `gorm:"size:50"   json:"position"`
	Organization string `gorm:"size:100"  json:"organization"`
	Description  string `gorm:"size:150"  json:"description"`

	GradeLevels []string `gorm:"-"         json:"grade_levels"`
	Timing      []string `gorm:"-"         json:"timing"`

	Hours  string `gorm:"not null" json:"hours"`
	Weeks  string `gorm:"not null" json:"weeks"`
	Intend bool   `gorm:"not null" json:"intend"`
}
