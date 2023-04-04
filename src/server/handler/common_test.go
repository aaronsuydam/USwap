package handler

import (
	_ "net/http"
	//_ "net/http/httpnet"
	//_ "test"
)

/*

// Mock this test
func TestLogin(t *testing.T) {
	r, _ := http.NewRequest("PUT", "/login", nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	r = mux.SetURLVars(r, vars)
	LoginPost(w, r)
	fmt.Println("bleh2")
}

// Mock signup test
*/
