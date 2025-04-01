package models

type Test struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserTest struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
	TestID uint `json:"test_id"`
}

type TestDetail struct {
	ID         uint   `json:"id"`
	UserTestID uint   `json:"user_test_id"`
	FieldName  string `json:"field_name"`
	FieldValue string `json:"field_value"`
}

//SQL BELOW

// CREATE TABLE test (
//     id SERIAL PRIMARY KEY,
//     name VARCHAR(255) NOT NULL
// );

// CREATE TABLE user_test (
//     id SERIAL PRIMARY KEY,
//     user_id INT,
//     test_id varchar REFERENCES test(id) ON DELETE CASCADE
// );

// CREATE TABLE test_detail (
//     id SERIAL PRIMARY KEY,
//     user_test_id INT REFERENCES user_test(id) ON DELETE CASCADE,
//     field_name VARCHAR(255),
//     field_value TEXT
// );
