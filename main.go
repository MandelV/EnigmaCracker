package main

import (
	"io/ioutil"
	"strings"
)

func main() {

	file, err := ioutil.ReadFile("cipher.txt")
	if err != nil {
		panic(err)
	}
	cipher := string(file)
	cipher = strings.ToUpper(cipher)
	cipher = strings.ReplaceAll(cipher, " ", "")
	cipher = strings.ReplaceAll(cipher, ".", "")
	cipher = strings.ReplaceAll(cipher, "\"", "")

	println(EnigmaSim(Config{Rotors: []string{"V", "I", "III"}, Position: []string{"A", "Q", "L"}, Rings: []int{1, 1, 1}, Plugboard: []string{"AB", "CD"}, Reflector: "B"}, cipher))
}

//UBSPYEITWFXMYNCSOTLCVCCWNWGAJBKREUY
