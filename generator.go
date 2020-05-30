package main

import "github.com/alexpantyukhin/combinations"

var rotors = []interface{}{"I", "II", "III", "IV", "V"}
var positions = []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var rings = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}

//ValuesString :
type ValuesString struct {
	V1 string
	V2 string
	V3 string
}

//ValuesInt :
type ValuesInt struct {
	V1 int
	V2 int
	V3 int
}

func genAllRotors() []ValuesString {
	product, _ := combinations.NewPermutation(rotors, 3)

	var values []ValuesString

	for product.Next() {
		v1 := product.Value()[0].(string)
		v2 := product.Value()[1].(string)
		v3 := product.Value()[2].(string)

		values = append(values, ValuesString{V1: v1, V2: v2, V3: v3})
	}
	return values
}

func genAllPosition() []ValuesString {
	product, _ := combinations.NewProduct(positions, 3)

	var values []ValuesString

	for product.Next() {
		v1 := product.Value()[0].(string)
		v2 := product.Value()[1].(string)
		v3 := product.Value()[2].(string)

		values = append(values, ValuesString{V1: v1, V2: v2, V3: v3})

	}
	return values
}

func genAllRings() []ValuesInt {
	product, _ := combinations.NewProduct(rings, 3)

	var values []ValuesInt

	for product.Next() {
		v1 := product.Value()[0].(int)
		v2 := product.Value()[1].(int)
		v3 := product.Value()[2].(int)

		values = append(values, ValuesInt{V1: v1, V2: v2, V3: v3})

	}
	return values
}
