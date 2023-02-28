package utils

import (
	"bytes"
	"encoding/json"
	"time"
  "net/http"
)

func StructToJSON(data interface{}) ([]byte, error) {
  buf := new(bytes.Buffer)

  if err := json.NewEncoder(buf).Encode(data); err != nil {
    return nil, err
  }

  return buf.Bytes(), nil
}

func GetJson(url string, target interface{}) error {
  var myClient = &http.Client{Timeout: 10 * time.Second}
  r, err := myClient.Get(url)
  if err != nil {
    return err
  }
  defer r.Body.Close()
  return json.NewDecoder(r.Body).Decode(target)
}