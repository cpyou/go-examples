package convert

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestIntToString(t *testing.T) {
	// int转string
	i := 100
	str := strconv.Itoa(i)
	fmt.Println(reflect.TypeOf(str), str)
}

func TestInt32ToString(t *testing.T) {
	// int32转string
	var i int32
	i = 100
	str := string(i)
	fmt.Println(reflect.TypeOf(str), str)
}

func TestInt64ToString(t *testing.T) {
	// int64转string
	var i int64
	i = 100
	str := strconv.FormatInt(i, 10)
	fmt.Println(reflect.TypeOf(str), str)
}

func TestFloat32ToString(t *testing.T) {
	// float32转string
	var i float32
	i = 100
	str := fmt.Sprintf("%f", i)
	fmt.Println(reflect.TypeOf(str), str)
}

func TestFloat64ToString(t *testing.T) {
	// float32转string
	var i float64
	i = 100.000000011
	//str := fmt.Sprintf("%f", i)
	str := strconv.FormatFloat(i, 'f', 10, 64)
	fmt.Println(reflect.TypeOf(str), str)
}
