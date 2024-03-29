package db

import (
	"testing"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabaseConfig(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(err)
	}
	err = Initialize()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserTableCreation(t *testing.T) {
	//godotenv.Load("../.env")
	//Initialize()
	err := createUserTable()
	if err != nil {
		t.Fatal("Error during user table creation")
	}
}

func Test(t *testing.T) {
	//godotenv.Load("../.env")
	//Initialize()
	err := DB.Ping()
	if err != nil {
		t.Fatal("Unable to ping the database")
	}
}

func TestUserCreation(t *testing.T) {
	//godotenv.Load("../.env")
	//Initialize()
	userid, err := CreateUser("testuser1", "testemail1@testemail.com", "testpassword1")
	if err != nil {
		t.Fatal("Failed to create test user")
	}
	user, err := GetUser(userid)
	if err != nil {
		t.Fatal(err)
	}
	if user.User_id != userid || user.User_name != "testuser1" || user.User_email != "testemail1@testemail.com" || user.User_password != "testpassword1" {
		t.Fatal("Retrieved user does not match the passed in test user")
	}
}

func TestGetUserItems(t *testing.T) {
	//godotenv.Load("../.env")
	//Initialize()
	itemid1, _ := CreateItem("testitem1", "testitem1description", "testuser1", []byte("testimage1"))
	_, err := CreateItem("testitem2", "testitem2description", "testuser1", []byte("testimage2"))
	if err != nil {
		t.Fatal(err)
	}
	item1, err := GetItem(itemid1)
	if err != nil {
		t.Fatal(err)
	}
	if itemid1 != item1.Item_id {
		t.Fatal("Retrieved user does not match the passed in test user")
	}
	items, err := GetUserItems("testuser1")
	if err != nil {
		t.Fatal(err)
	}
	if len(items) <= 1 {
		t.Fatalf("Incorrect number of items obtained: %d", len(items))
	}
}

func TestSwapRequestCreation(t *testing.T) {
	//godotenv.Load("../.env")
	//Initialize()
	swapid, err := CreateSwapRequest("testuser1", "testitem1", "testuser2", "testitem2")
	if err != nil {
		t.Fatal("Failed to create test user")
	}
	swapRequest, err := GetSwapRequest(swapid)
	if err != nil {
		t.Fatal("Failed to get swap request")
	}
	if swapRequest.Swap_id != swapid || swapRequest.Sender_id != "testuser1" || swapRequest.Sender_item_id != "testitem1" || swapRequest.Receiver_id != "testuser2" || swapRequest.Receiver_item_id != "testitem2" {
		t.Fatal("Retrieved swap request does not match the passed in test swap request")
	}
}

func TestSwapRequestAccept(t *testing.T) {
	//godotenv.Load("../.env")
	//Initialize()
	userid2, err := CreateUser("testuser2", "testemail2@testemail.com", "testpassword2")
	if err != nil {
		t.Fatal("Error during get Swap")
	}
	userid3, _ := CreateUser("testuser3", "testemail3@testemail.com", "testpassword3")
	itemid2, _ := CreateItem("testitem2", "testitemdescription2", userid2, []byte("fdjaifja"))
	if err != nil {
		t.Fatal(err)
	}
	itemid3, _ := CreateItem("testitem3", "testitemdescription3", userid3, []byte("fjdiajfia"))
	swapid, err := CreateSwapRequest(userid2, itemid2, userid3, itemid3)
	if err != nil {
		t.Fatal("Failed to create swap request")
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
	if item2New.User_id != userid3 && item3New.User_id != userid2 {
		t.Fatal("Failed to swap user IDs on items")
	}
}

func TestSwapRequestDeny(t *testing.T) {
	//godotenv.Load("../.env")
	//Initialize()
	userid2, err := CreateUser("testuser2", "testemail2@testemail.com", "testpassword2")
	if err != nil {
		t.Fatal("Error during get Swap")
	}
	userid3, _ := CreateUser("testuser3", "testemail3@testemail.com", "testpassword3")
	itemid2, _ := CreateItem("testitem2", "testitemdescription2", userid2, []byte("fdjaifja"))
	if err != nil {
		t.Fatal(err)
	}
	itemid3, _ := CreateItem("testitem3", "testitemdescription3", userid3, []byte("fjdiajfia"))
	swapid, err := CreateSwapRequest(userid2, itemid2, userid3, itemid3)
	if err != nil {
		t.Fatal("Failed to create swap request")
	}
	err = RejectSwapRequest(swapid)
	if err != nil {
		t.Fatal(err)
	}
	item2New, err := GetItem(itemid2)
	if err != nil {
		t.Fatal("Error in getItem")
	}
	item3New, _ := GetItem(itemid3)
	if item2New.User_id != userid2 && item3New.User_id != userid3 {
		t.Fatal("Item relations to users were incorrectly modified in a reject swap request.")
	}
}

/* Current MySQL database version does not support Fulltext
func TestSearch(t *testing.T) {
	godotenv.Load("../.env")
	Initialize()
	items, err := SearchItems("test")
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range items {
		t.Logf("%s", item.item_name)
	}
}
*/

func TestGetItems(t *testing.T) {
	//godotenv.Load("../.env")
	//Initialize()
	_, err := GetItems()
	if err != nil {
		t.Fatal(err)
	}
}
