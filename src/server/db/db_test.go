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
	user, err := GetUser("3")
	if err != nil {
		t.Fatal("Error during Get User Item")
	}
	fmt.Print(user.user_name)
}

func TestGetUserItem(t *testing.T) {
	item, err := GetItem("2")
	if err != nil {
		t.Fatal("Error during Get User Item")
	}
	fmt.Print(item.item_name)
}

func TestGetUserItems(t *testing.T) {
	item, err := GetUserItems("2")
	if err != nil {
		t.Fatal("Error during get All of User's Items")
	}
	if len(item) == 3 {
		fmt.Print("Great success!")
	}
}

func TestGetSwapRequest(t *testing.T) {
	swap, err := GetSwapRequest("2")
	if err != nil {
		t.Fatal("Error during get All of User's Items")
	}
	fmt.Print(swap.receiver_id)
}

func Test(t *testing.T) {
	err := DB.Ping()
	if err != nil {
		t.Fatal("Unable to ping the database")
	}
}
