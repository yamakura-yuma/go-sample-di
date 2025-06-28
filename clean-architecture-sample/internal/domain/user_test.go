package domain

import "testing"

func TestUserStruct(t *testing.T) {
	u := User{ID: "1", Name: "Test", Email: "test@example.com"}
	if u.ID != "1" || u.Name != "Test" || u.Email != "test@example.com" {
		t.Errorf("User struct fields not set correctly: %+v", u)
	}
}
