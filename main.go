package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Result struct {
	ic     float64
	config Config
	text   string
}

func calculateIc(cipher string, permsRotor ValuesString, outputIc chan Result) {
	permsPosition := genAllPosition()
	bestIc := Result{ic: 0.039}
	for _, pos := range permsPosition {
		currentConfig := Config{
			Rotors:    []string{permsRotor.V1, permsRotor.V2, permsRotor.V3},
			Position:  []string{pos.V1, pos.V2, pos.V3},
			Rings:     []int{1, 1, 1},
			Plugboard: []string{"AB", "CD"},
			Reflector: "B"}
		text := EnigmaSim(currentConfig, cipher)

		ic := Ic(text)

		if ic > bestIc.ic {
			bestIc = Result{config: currentConfig, ic: ic, text: text}
		}
	}
	outputIc <- bestIc
}
func crack(cipher string) {
	permsRotors := genAllRotors()
	outputIc := make(chan Result)
	bestIc := Result{ic: 0.039}
	for _, n := range permsRotors {
		go calculateIc(cipher, n, outputIc)
	}

	for i := 0; i < 60; i++ {
		ic := <-outputIc
		if ic.ic > bestIc.ic {
			bestIc = ic
		}

	}

	fmt.Println(bestIc.config, bestIc.text)
}

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

	crack(cipher)
}
