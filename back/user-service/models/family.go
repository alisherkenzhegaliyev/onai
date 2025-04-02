package models

type Family struct {
	ID                int    `json:"id"`
	UserID            int    `json:"user_id"`
	MaritalStatus     string `json:"marital_status"`
	PermanentHomeWith string `json:"permanent_home_with"`
}

type Parent struct {
	ID            int    `json:"id"`
	FamilyID      int    `json:"family_id"`
	Role          string `json:"role"`
	FirstName     string `json:"first_name"`
	SecondName    string `json:"second_name"`
	HighestDegree string `json:"highest_degree"`
	Occupation    string `json:"occupation"`
}

type ParentContactDetails struct {
	ID           int    `json:"id"`
	ParentID     int    `json:"parent_id"`
	Email        string `json:"email"`
	CountryCode  string `json:"country_code"`
	MobileNumber string `json:"mobile_number"`
}

type ParentCollegeDetails struct {
	ID             int    `json:"id"`
	ParentID       int    `json:"parent_id"`
	CollegeName    string `json:"college_name"`
	DegreeReceived string `json:"degree_received"`
	YearReceived   string `json:"year_received"`
}

type Sibling struct {
	ID         int    `json:"id"`
	FamilyID   int    `json:"family_id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Age        int    `json:"age"`
}

// SQL below

// CREATE TABLE family (
//     id SERIAL PRIMARY KEY,
// 	user_id INT,
//     marital_status VARCHAR(50),
//     permanent_home_with VARCHAR(50)
// );

// CREATE TABLE parent (
//     id SERIAL PRIMARY KEY,
//     family_id INT,
//     role VARCHAR(10),
//     first_name VARCHAR(100),
//     second_name VARCHAR(100),
//     highest_degree VARCHAR(100),
//     occupation VARCHAR(100),

//     FOREIGN KEY (family_id) REFERENCES family(id) ON DELETE CASCADE
// );

// CREATE TABLE parent_contact_details (
//     id SERIAL PRIMARY KEY,
//     parent_id INT,
//     email VARCHAR(100),
//     country_code VARCHAR(15),
//     mobile_number VARCHAR(25),

//     FOREIGN KEY (parent_id) REFERENCES parent(id) ON DELETE CASCADE
// );

// CREATE TABLE parent_college_details (
//     id SERIAL PRIMARY KEY,
//     parent_id INT,
//     college_name VARCHAR(120),
//     degree_received VARCHAR(60),
//     year_received VARCHAR(30),

//     FOREIGN KEY (parent_id) REFERENCES parent(id) ON DELETE CASCADE
// );

// CREATE TABLE sibling (
//     id SERIAL PRIMARY KEY,
//     family_id INT,
//     first_name VARCHAR(100),
//     second_name VARCHAR(100),
//     age INT,

//     FOREIGN KEY (family_id) REFERENCES family(id) ON DELETE CASCADE
// );
