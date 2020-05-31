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

func alreadyExist(slice []string, pair string) bool {
	l1, l2 := string(pair[0]), string(pair[1])

	for i := 0; i < len(slice); i++ {

		if strings.Contains(slice[i], l1) || strings.Contains(slice[i], l2) {
			return true
		}
	}
	return false
}

func calculateIc(cipher, reflector string, permsRotor ValuesString, outputIc chan Result) {
	permsPosition := genAllPosition()
	cleanText(&cipher)
	bestIc := Result{ic: 0.039}

	for _, pos := range permsPosition {
		currentConfig := Config{
			Rotors:   []string{permsRotor.V1, permsRotor.V2, permsRotor.V3},
			Position: []string{pos.V1, pos.V2, pos.V3},
			Rings:    []int{15, 14, 18}, //15 14 18
			/*Plugboard: []string{"AB"},*/
			Reflector: reflector}
		text := EnigmaSim(currentConfig, cipher)

		ic := Ic(text)

		if ic > bestIc.ic {
			bestIc = Result{config: currentConfig, ic: ic, text: text}
		}
	}
	outputIc <- bestIc

}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

type IcPair struct {
	pair string
	ic   float64
}

func calculatePlugboard(cipher string, best Result) {

	var icPair []IcPair
	pairsPlug := genPlugboard()
	cleanText(&cipher)
	bestIc := Result{ic: best.ic}
	currentConfig := best.config
	for _, p := range pairsPlug {
		pair := p.V1 + p.V2

		currentConfig.Plugboard = []string{pair}

		text := EnigmaSim(currentConfig, cipher)

		ic := Ic(text)

		if ic > bestIc.ic {

			icPair = append(icPair, IcPair{pair: pair, ic: ic})
		}
	}

	avg := 0.0
	for _, p := range icPair {
		avg += p.ic
	}
	avg /= float64(len(icPair))
	for _, p := range icPair {
		if p.ic > 0.05 {
			println(p.pair, p.ic)
		}
	}

}
func crack(cipher string) {
	permsRotors := genAllRotors()
	outputIc := make(chan Result)
	defer close(outputIc)

	bestIc := Result{ic: 0.039}
	count := 0
	for _, r := range "BC" {
		for _, n := range permsRotors {
			count++
			go calculateIc(cipher, string(r), n, outputIc)
		}
	}

	for ic := range outputIc {
		count--
		if ic.ic > bestIc.ic {
			bestIc = ic
			fmt.Println(bestIc.config, bestIc.ic)
		}

		if count == 0 {
			break
		}
	}

	calculatePlugboard(cipher, bestIc)

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
	cleanText(&cipher)
	//fmt.Println(Ic(cipher))
	crack(cipher)

}

//AY IJ DR LB CK PU TV ZW
