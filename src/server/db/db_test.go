package db

import (
	"fmt"
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

func TestGetUser(t *testing.T) {
	user, err := GetUser("0ec19c70-2491-4e13-b84f-10db98cea14d")
	if err != nil {
		t.Fatal("Error during Get User")
	}
	fmt.Print(user.user_name)
}

func TestGetUserItem(t *testing.T) {

	item, err := GetItem("74a8898b-d970-4359-5ad0-479d69d1081c")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(item.item_name)
}

func TestGetUserItems(t *testing.T) {

	CreateItem("testItem", "thisisatest", "2", "testpath")
	item, err := GetUserItems("2")
	if err != nil {
		t.Fatal("Error during get All of User's Items")
	}
	if len(item) == 4 {
		fmt.Print("Great success!")
	}
}

func TestGetSwapRequest(t *testing.T) {
	swap, err := GetSwapRequest("2")
	if err != nil {
		t.Fatal("Error during get Swap")
	}
	fmt.Print(swap.receiver_id)
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
