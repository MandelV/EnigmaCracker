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

func calculateIc(cipher, reflector string, permsRotor ValuesString, outputIc chan Result) {
	permsPosition := genAllPosition()
	cleanText(&cipher)
	bestIc := Result{ic: 0.039}

	for _, pos := range permsPosition {
		currentConfig := Config{
			Rotors:   []string{permsRotor.V1, permsRotor.V2, permsRotor.V3},
			Position: []string{pos.V1, pos.V2, pos.V3},
			Rings:    []int{1, 14, 18},
			//Plugboard: []string{"DR", "JX", "FW", "HS", "CL", "MU", "GY", "KV", "QZ", "BP"},
			Reflector: reflector}
		text := EnigmaSim(currentConfig, cipher)

		ic := Ic(text)

		if ic > bestIc.ic {
			bestIc = Result{config: currentConfig, ic: ic, text: text}
			outputIc <- bestIc
		}
	}

}
func crack(cipher string) {
	permsRotors := genAllRotors()
	outputIc := make(chan Result)
	defer close(outputIc)

	bestIc := Result{ic: 0.039}
	for _, r := range "BC" {
		for _, n := range permsRotors {
			go calculateIc(cipher, string(r), n, outputIc)
		}
	}

	for ic := range outputIc {

		if ic.ic > bestIc.ic {
			bestIc = ic
			fmt.Println(bestIc.config, bestIc.ic)
		}

	}

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
