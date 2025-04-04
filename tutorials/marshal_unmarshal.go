package main

import (
	"encoding/json"
	"fmt"
)

type StructObject struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Convert() {
// func main() {
	structObject := []StructObject{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
		{Name: "Doe", Age: 40},
	}

	// Convert structObject to JSON
	jsonData, err := json.Marshal(structObject)

	if err != nil {
		fmt.Println("Error converting to JSON:", err)
	}

	fmt.Println("JSON data:", string(jsonData))

	jsonObject := `[{"name":"John","age":30},
					{"name":"Jane","age":25},
					{"name":"Doe","age":40}]`
	// Convert JSON to structObject
	var structObject2 []StructObject
	if err := json.Unmarshal([]byte(jsonObject), &structObject2); err != nil {
		fmt.Println("Error converting JSON to struct:", err)
	}
	fmt.Println("Struct data:", structObject2)
}
