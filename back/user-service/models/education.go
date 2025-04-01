package models

import "time"

type Education struct {
    ID     int `json:"id"`
    UserID int `json:"user_id"`
}

type Highschool struct {
    ID               int       `json:"id"`
    EducationID      int       `json:"education_id"`
    Name             string    `json:"name"`
    AddressID        int       `json:"address_id"`
    DateOfEntry      time.Time `json:"date_of_entry"`
    IsBoardingSchool bool      `json:"is_boarding_school"`
    WillGraduate     bool      `json:"will_graduate"`
    GraduationChange string    `json:"graduation_change"`
}

type Grade struct {
    EducationID int     `json:"education_id"`
    ClassSize   int     `json:"class_size"`
    GPAScale    int     `json:"gpa_scale"`
    CGPA        float64 `json:"cgpa"`
}

type College struct {
    ID          int       `json:"id"`
    EducationID int       `json:"education_id"`
    Name        string    `json:"name"`
    AddressID   int       `json:"address_id"`
    StartDate   time.Time `json:"start_date"`
    EndDate     time.Time `json:"end_date"`
    Degree      string    `json:"degree"`
}

type Course struct {
    ID              int    `json:"id"`
    EducationID     int    `json:"education_id"`
    SchedulingSystem string `json:"scheduling_system"`
    Subject         string `json:"subject"`
    CourseName      string `json:"course_name"`
    CourseLevel     string `json:"course_level"`
}

type Honor struct {
    ID                 int    `json:"id"`
    EducationID        int    `json:"education_id"`
    WillShareHonors    bool   `json:"will_share_honors"`
    HonorTitle         string `json:"honor_title"`
    GradeLevel         string `json:"grade_level"`
    LevelOfRecognition string `json:"level_of_recognition"`
}

type Plan struct {
    ID             int    `json:"id"`
    EducationID    int    `json:"education_id"`
    HighestDegree  string `json:"highest_degree"`
    CareerInterest string `json:"career_interest"`
}

// SQL below

// CREATE TABLE education (
//     user_id INT,
//     id SERIAL PRIMARY KEY
// );

// CREATE TABLE highschool(
//     education_id INT REFERENCES education(id) ON DELETE CASCADE,
//     id SERIAL PRIMARY KEY,
//     name VARCHAR(100),
// 	address_id INT REFERENCES address(id) ON DELETE CASCADE,
//     date_of_entry DATE,
//     is_boarding_school BOOLEAN,
//     will_graduate BOOLEAN,
//     graduation_change VARCHAR(50)
// );

// CREATE TABLE grade(
//     education_id INT REFERENCES education(id) ON DELETE CASCADE,
//     class_size INT,
//     gpa_scale INT,
//     cgpa FLOAT
// );

// CREATE TABLE college(
//     id SERIAL PRIMARY KEY,
//     education_id INT REFERENCES education(id) ON DELETE CASCADE,
//     name VARCHAR(100),
// 	address_id INT REFERENCES address(id) ON DELETE CASCADE,
//     start_date DATE,
//     end_date DATE,
// 	degree VARCHAR(100)
// );

// CREATE TABLE course(
//     id SERIAL PRIMARY KEY,
//     education_id INT REFERENCES education(id) ON DELETE CASCADE,
// 	scheduling_system VARCHAR(30),
//     subject VARCHAR(80),
//     course_name VARCHAR(100),
//     course_level VARCHAR(40)
// );

// CREATE TABLE honor(
//     id SERIAL PRIMARY KEY,
//     education_id INT REFERENCES education(id) ON DELETE CASCADE,
//     will_share_honors BOOLEAN,
//     honor_title VARCHAR(70),
//     grade_level VARCHAR(30),
//     level_of_recognition VARCHAR(60)
// );

// CREATE TABLE plan(
//     id SERIAL PRIMARY KEY,
//     education_id INT REFERENCES education(id) ON DELETE CASCADE,
//     highest_degree VARCHAR(50),
//     career_interest VARCHAR(70)
// );