package models

type ActivitiesSection struct {
	ID            uint `json:"id"`
	UserID        uint `json:"user_id"`
	HasActivities bool `json:"has_activities"`
}

// Each Activity record
type Activity struct {
	ID                  uint   `json:"id"`
	ActivitiesID        uint   `json:"activities_id"`
	ActivityType        string `json:"activity_type"`
	PositionDescription string `json:"position_description"`
	Organization        string `json:"organization"`
	ActivityDescription string `json:"activity_description"`
	Hours               string `json:"hours"`
	Weeks               string `json:"weeks"`
	Intend              bool   `json:"intend"`
}

//sql
// CREATE TABLE activities_section (
//     id SERIAL PRIMARY KEY,
//     user_id INT NOT NULL,
//     has_activities BOOLEAN NOT NULL
// );

// CREATE TABLE activity (
//     id SERIAL PRIMARY KEY,
//     activities_id INT NOT NULL REFERENCES activities_section(id) ON DELETE CASCADE,
//     activity_type VARCHAR(50),
//     position_description VARCHAR(50),
//     organization VARCHAR(100),
//     activity_description VARCHAR(150),
//     hours VARCHAR NOT NULL,
//     weeks VARCHAR NOT NULL,
//     intend BOOLEAN NOT NULL
// );
