package parse

import (
	"encoding/json"
	"fmt"
)

func test01() {
	empJson := `{
        "id" : 11,
        "name" : "Irshad",
        "department" : "IT",
        "designation" : "Product Manager"
	}`
	var result map[string]interface{}
	json.Unmarshal([]byte(empJson),&result)
	fmt.Println(result)
}
func Do() {
	test01()
}
