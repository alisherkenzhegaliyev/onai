package models

type Testing struct {
	ID               uint `gorm:"primaryKey" json:"id"`
	ProfileID        uint `gorm:"not null"  json:"profile_id"`
	HasAcademicTests bool `gorm:"not null"  json:"has_academic_tests"`

	// If true, store multiple tests here, e.g. IELTS, TOEFL, SAT, etc.
	AcademicTests    []AcademicTest `gorm:"foreignKey:TestingID;constraint:OnDelete:CASCADE" json:"academic_tests"`
	HasNationalExams bool           `gorm:"not null"  json:"has_national_exams"`

	// If true, store a rating from 1 to 10
	NationalExamRating int `json:"national_exam_rating"`
}

// Similar to how Languages is separate in Profile model.
type AcademicTest struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	TestingID uint   `gorm:"not null"   json:"testing_id"`
	TestName  string `gorm:"not null"   json:"test_name"`
	Score     string `json:"score"`
	TestDate  string `json:"test_date"`
}
