package models

import "time"

type Education struct {
    ID                 uint              `gorm:"primaryKey" json:"id"`
    RecentHighSchool   RecentHighSchool  `gorm:"embedded" json:"recent_high_school"`
    Colleges           []College         `gorm:"foreignKey:EducationID;constraint:OnDelete:CASCADE" json:"colleges"`
    Grades            Grades            `gorm:"embedded" json:"grades"`
    CoursesTaken      []Course          `gorm:"foreignKey:EducationID;constraint:OnDelete:CASCADE" json:"courses_taken"`
    Honors            []Honor           `gorm:"foreignKey:EducationID;constraint:OnDelete:CASCADE" json:"honors"`
    FuturePlans       FuturePlans       `gorm:"embedded" json:"future_plans"`
}

type RecentHighSchool struct {
    Name                string    `gorm:"not null" json:"name"`
    DateOfEntry         time.Time `gorm:"not null" json:"date_of_entry"`
    BoardingSchool      bool      `gorm:"not null" json:"boarding_school"`
    WillGraduate        bool      `gorm:"not null" json:"will_graduate"`
    GraduationChange    string    `gorm:"not null" json:"graduation_change"`
}

type College struct {
    ID            uint      `gorm:"primaryKey" json:"id"`
    EducationID   uint      `gorm:"not null" json:"education_id"`
    Name          string    `gorm:"not null" json:"name"`
    StartDate     time.Time `gorm:"not null" json:"start_date"`
    EndDate       time.Time `gorm:"not null" json:"end_date"`
}

type Grades struct {
    ClassSize   int     `gorm:"not null" json:"class_size"`
    GPAScale    string  `gorm:"not null" json:"gpa_scale"`
    CGPA        float64 `gorm:"not null" json:"cgpa"`
}

type Course struct {
    ID            uint   `gorm:"primaryKey" json:"id"`
    EducationID   uint   `gorm:"not null" json:"education_id"`
    Subject       string `gorm:"not null" json:"subject"`
    CourseName    string `gorm:"not null" json:"course_name"`
    CourseLevel   string `gorm:"not null" json:"course_level"`
}

type Honor struct {
    ID                uint   `gorm:"primaryKey" json:"id"`
    EducationID       uint   `gorm:"not null" json:"education_id"`
    HonorTitle        string `gorm:"not null" json:"honor_title"`
    GradeLevel        string `gorm:"not null" json:"grade_level"`
    LevelOfRecognition string `gorm:"not null" json:"level_of_recognition"`
}

type FuturePlans struct {
    HighestDegree string `gorm:"not null" json:"highest_degree"`
    CareerInterest string `gorm:"not null" json:"career_interest"`
}
