package main

import "github.com/alexpantyukhin/combinations"

var num = []interface{}{"I", "II", "III", "IV", "V"}
var alpha = []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

type Values struct {
	V1 string
	V2 string
	V3 string
}

func genProduct(inter []interface{}) []Values {

	product, _ := combinations.NewProduct(inter, 3)

	var values []Values

	for product.Next() {
		v1 := product.Value()[0].(string)
		v2 := product.Value()[1].(string)
		v3 := product.Value()[2].(string)

		values = append(values, Values{V1: v1, V2: v2, V3: v3})

	}
	return values
}
func genPermutation(inter []interface{}) []Values {

	product, _ := combinations.NewPermutation(inter, 3)

	var values []Values

	for product.Next() {
		v1 := product.Value()[0].(string)
		v2 := product.Value()[1].(string)
		v3 := product.Value()[2].(string)

		values = append(values, Values{V1: v1, V2: v2, V3: v3})
	}
	return values
}
