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

func TestGetUserItems(t *testing.T) {
	itemid, err := CreateItem("testitem1", "testitem1description", "testuser1", "testimagepath1")
	if err != nil {
		t.Fatal("Failed to create test user")
	}
	item, err := GetItem(itemid)
	if err != nil {
		t.Fatal(err)
	}
	if itemid != item.item_id {
		t.Fatal("Retrieved user does not match the passed in test user")
	}
}

func TestSwapRequestCreation(t *testing.T) {
	swapid, err := CreateSwapRequest("testuser1", "testitem1", "testuser2", "testitem2")
	if err != nil {
		t.Fatal("Failed to create test user")
	}
	swapRequest, err := GetSwapRequest(swapid)
	if err != nil {
		t.Fatal("Failed to get swap request")
	}
	if swapRequest.swap_id != swapid || swapRequest.sender_id != "testuser1" || swapRequest.sender_item_id != "testitem1" || swapRequest.receiver_id != "testuser2" || swapRequest.receiver_item_id != "testitem2" {
		t.Fatal("Retrieved swap request does not match the passed in test swap request")
	}
}

func TestSwapRequestAccept(t *testing.T) {
	userid2, err := CreateUser("testuser2", "testemail2@testemail.com", "testpassword2")
	if err != nil {
		t.Fatal("Error during get Swap")
	}
	userid3, _ := CreateUser("testuser3", "testemail3@testemail.com", "testpassword3")
	itemid2, _ := CreateItem("testitem2", "testitemdescription2", userid2, "fdjaifja")
	if err != nil {
		t.Fatal(err)
	}
	itemid3, _ := CreateItem("testitem3", "testitemdescription3", userid3, "fjdiajfia")
	swapid, err := CreateSwapRequest(userid2, itemid2, userid3, itemid3)
	if err != nil {
		t.Fatal("idfk its broken")
	}
	err = AcceptSwapRequest(swapid)
	if err != nil {
		t.Fatal(err)
	}
	item2New, err := GetItem(itemid2)
	if err != nil {
		t.Fatal("Error in getItem")
	}
	item3New, _ := GetItem(itemid3)
	if item2New.user_id != userid3 && item3New.user_id != userid2 {
		t.Fatal("Failed to swap user IDs on items")
	}
}

func TestSwapRequestDeny(t *testing.T) {}
