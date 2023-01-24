package utils

import (
	"fmt"
)

// func StructToJSON(data interface{}) ([]byte, error) {
// 	buf := new(bytes.Buffer)

// 	if err := json.NewEncoder(buf).Encode(data); err != nil {
// 		return nil, err
// 	}

// 	return buf.Bytes(), nil
// }

func Fun() {
	fmt.Println("Hello from the other side")
}
