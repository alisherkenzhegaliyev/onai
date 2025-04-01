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
