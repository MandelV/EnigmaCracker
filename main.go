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

	println(EnigmaSim(EnigmaConfig{wheels: *newWheelOrder(&Values{"I", "I", "I"}), Rings: "AAA", Position: "AAA", Reflector: "B"}, cipher))
}
