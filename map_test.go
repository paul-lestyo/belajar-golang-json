package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id": "P0001", "name": "Laptop Gaming Jos", "price": 20000000 }`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapDecode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P0001",
		"name":  "Laptop Gaming Jos",
		"price": 2000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))

}
