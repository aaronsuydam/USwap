package db

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabaseConfig(t *testing.T) {
	var err error
	DB, err = dbConnection()
	if err != nil {
		t.Fatal("Error during dbConnection")
	}
}

func TestUserTableCreation(t *testing.T) {
	err := createUserTable()
	if err != nil {
		t.Fatal("Error during user table creation")
	}
}

func Test(t *testing.T) {
	err := DB.Ping()
	if err != nil {
		t.Fatal("Unable to ping the database")
	}
}

func TestUserCreation(t *testing.T) {
	userid, err := CreateUser("testuser1", "testemail1@testemail.com", "testpassword1")
	if err != nil {
		t.Fatal("Failed to create test user")
	}
	user, err := GetUser(userid)
	if err != nil {
		t.Fatal(err)
	}
	if user.user_id != userid || user.user_name != "testuser1" || user.user_email != "testemail1@testemail.com" || user.user_password != "testpassword1" {
		t.Fatal("Retrieved user does not match the passed in test user")
	}
}

func TestSwapRequestCreation(t *testing.T) {}

func TestSwapRequestAccept(t *testing.T) {}

func TestSwapRequestDeny(t *testing.T) {}
