package main

import (
	"strings"

	"github.com/rjhacks/enigma/enigma"
)

//WheelOrder part of the configuration for the wheels order
type WheelOrder struct {
	R1 string
	R2 string
	R3 string
}

//EnigmaConfig configuration enigma
type EnigmaConfig struct {
	wheels    WheelOrder
	Rings     string
	Position  string
	Reflector string
	Plugboard string
}

func newWheelOrder(values *Values) *WheelOrder {
	wheelOrder := &WheelOrder{}
	wheelOrder.R1 = values.V1
	wheelOrder.R2 = values.V2
	wheelOrder.R3 = values.V3
	return wheelOrder
}

func createPlugboard(plug string) map[byte]byte {
	mappedpair := make(map[byte]byte)

	if plug == "" {
		return mappedpair
	}
	pairs := strings.Split(plug, " ")

	for _, pairs := range pairs {

		mappedpair[byte(pairs[0])] = byte(pairs[1])
	}
	return mappedpair
}

//EnigmaSim Simulation of an Enigma
func EnigmaSim(config EnigmaConfig, cipher string) string {
	enigmaMachine := enigma.New()

	enigmaMachine.InstallRotors([]enigma.Rotor{enigma.Rotors[config.wheels.R1], enigma.Rotors[config.wheels.R2], enigma.Rotors[config.wheels.R3]})
	enigmaMachine.SetRingSettings([]byte(config.Rings))
	enigmaMachine.SetRotorPositions([]byte(config.Position))
	enigmaMachine.InstallReflector(enigma.Reflectors[config.Reflector])
	enigmaMachine.SetPlugboard(enigma.Plugboard{Mapping: createPlugboard(config.Plugboard)})
	return enigma.Type(enigmaMachine, cipher)
}
