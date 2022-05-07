package convert

import (
	"fmt"
	"strconv"
	"testing"

	"gotest.tools/v3/assert"
)

func TestStringToInt(t *testing.T) {
	i, err := strconv.Atoi("10000")
	fmt.Println(i, err)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, i, 10000, "StringToInt Error")
}

func TestStructToJsonDemo(t *testing.T) {
	tests := []People{
		{Name: "jqw", Age: 18},
		{Name: "asd", Age: 19},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			StructToJsonDemo(tt)
		})
	}
}

func TestJsonToStructDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: `{"name_title": "jqw","age_size": 12}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JsonToStructDemo(tt.name)
		})
	}
}

func TestJsonToMapDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: `{"name_title": "jqw","age_size": 12}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JsonToMapDemo(tt.name)
		})
	}
}

func TestMapToJsonDemo1(t *testing.T) {
	MapToJsonDemo1()
}

func TestMapToJsonDemo2(t *testing.T) {
	MapToJsonDemo2()
}

func TestMapToInterface(t *testing.T) {
	MapToInterface()
}
