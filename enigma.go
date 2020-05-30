package main

import (
	"github.com/emedvedev/enigma"
)

//Config part of the configuration for the wheels order
type Config struct {
	Rotors    []string
	Rings     []int
	Position  []string
	Plugboard []string
	Reflector string
}

//EnigmaSim Simulation of an Enigma
func EnigmaSim(config Config, cipher string) string {

	//Create ROTORS
	conf := make([]enigma.RotorConfig, len(config.Rotors))
	for index, rotor := range config.Rotors {
		ring := config.Rings[index]
		value := config.Position[index][0]
		conf[index] = enigma.RotorConfig{ID: rotor, Start: value, Ring: ring}
	}

	machine := enigma.NewEnigma(conf, config.Reflector, config.Plugboard)

	return machine.EncodeString(enigma.SanitizePlaintext(cipher))

}
