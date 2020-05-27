package main

import "strings"

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func formatMessage(message *string) {
	*message = strings.ToLower(*message)
	*message = strings.ReplaceAll(*message, "é", "e")
	*message = strings.ReplaceAll(*message, "ê", "e")
	*message = strings.ReplaceAll(*message, "è", "e")
	*message = strings.ReplaceAll(*message, "à", "a")
	*message = strings.ReplaceAll(*message, "œ", "oe")
	*message = strings.ReplaceAll(*message, "ô", "o")
	*message = strings.ReplaceAll(*message, "î", "i")
	*message = strings.ReplaceAll(*message, " ", "")
	*message = strings.ReplaceAll(*message, ",", "")
	*message = strings.ReplaceAll(*message, ".", "")
	*message = strings.ReplaceAll(*message, "'", "")
	*message = strings.ReplaceAll(*message, "-", "")
}

//Ic Return the Index of coincidence IC = Σ[q=a->q=z] nq (nq-1) / n(n-1)
//where nq = number of time that letter appear in text
//and n the size of the text
func Ic(message string) (indice float64) {
	formatMessage(&message)
	msg, n := message, len(message)
	var ic float64 = 0

	for _, alpha := range alphabet {
		numberOfAlpha := strings.Count(msg, string(alpha))

		num := float64(numberOfAlpha * (numberOfAlpha - 1))
		den := float64(n * (n - 1))
		ic += num / den
	}

	return ic
}
