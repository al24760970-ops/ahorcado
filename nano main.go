package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Dibujos del ahorcado en una variable limpia para no ensuciar el loop principal
var estados = []string{
	"  O",
	"  O\n  |",
	"  O\n /|",
	"  O\n /|\\",
	"  O\n /|\\\n /",
	"  O\n /|\\\n / \\",
}

func main() {
	palabras := []string{"computadora", "guitarra", "profesor", "ite", "ensenada", "musica"}

	rand.Seed(time.Now().UnixNano())
	secreta := palabras[rand.Intn(len(palabras))]

	progreso := make([]string, len(secreta))
	for i := range progreso {
		progreso[i] = "_"
	}

	fallos := 0
	maxVidas := 6

	fmt.Println("--- AHORCADO GO ---")

	for fallos < maxVidas {
		fmt.Printf("\nPalabra: %s  (Vidas: %d)\n", strings.Join(progreso, " "), maxVidas-fallos)
		fmt.Print("Introduce unas letra: ")

		var input string
		fmt.Scan(&input)
		input = strings.ToLower(input)

		if len(input) != 1 {
			fmt.Println("Una sola letra a la vez")
			continue
		}

		// vemos si atino de una formita mas directa
		acierto := false
		for i, letra := range secreta {
			if string(letra) == input {
				progreso[i] = input
				acierto = true
			}
		}

		if !acierto {
			fmt.Println(estados[fallos])
			fallos++
			fmt.Printf("No, la '%s' no está.\n", input)
		}

		// Condición de victoria rápida
		if strings.Join(progreso, "") == secreta {
			fmt.Printf("\nLo lograste La palabra era: %s\n", secreta)
			return
		}
	}

	fmt.Printf("\nTerminaste ahorcado... Era: %s\n", secreta)
}
