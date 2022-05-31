package convert

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name_title"`
	Age  int    `json:"age_size"`
}

func StructToJsonDemo(p People) {
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}

func JsonToStructDemo(jsonStr string) {
	//jsonStr := `
	//    {
	//		"name_title": "jqw",
	//		"age_size": 12
	//    }
	//    `
	var people People
	fmt.Println("JsonToStructDemo1", people)
	err := json.Unmarshal([]byte(jsonStr), &people)
	if err != nil {
		fmt.Println("JsonToStructDemo Error:", err)
		return
	}
	fmt.Println(people)
}

func JsonToMapDemo(jsonStr string) {
	//jsonStr := `
	//    {
	//		"name": "jqw",
	//		"age": 18
	//    }
	//    `
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	fmt.Println(mapResult)
}

func MapToJsonDemo1() {
	var mapInstances []map[string]interface{}
	instance1 := map[string]interface{}{"name": "John", "age": 10}
	instance2 := map[string]interface{}{"name": "Alex", "age": 12}
	mapInstances = append(mapInstances, instance1, instance2)

	jsonStr, err := json.Marshal(mapInstances)

	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(jsonStr))
}

func MapToJsonDemo2() {
	b, _ := json.Marshal(map[string]int{"test": 1, "try": 2})
	fmt.Println(string(b))
}

func MapToInterface() {
	bs := make(map[string]string)
	bs["name"] = "张三"
	bs["age"] = "12"
	var student interface{}

	student = bs
	fmt.Println(bs)
	fmt.Println(student)

	// interface转map
	a := student.(map[string]string)
	fmt.Println(a)
}

// ArrayToInterface converting a []string to a []interface{}
func ArrayToInterface(x []string) []interface{} {

	y := make([]interface{}, len(x))
	for i, v := range x {
		y[i] = v
	}
	return y

}

// InterfaceToArray converting a []interface{} to a []string
func InterfaceToArray(y []interface{}) []string {
	z := make([]string, len(y))
	for i, v := range y {
		z[i] = fmt.Sprint(v)
	}
	return z
}
