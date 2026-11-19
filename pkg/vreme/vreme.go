//Package vreme omogoča upravljanje vremenskih postaj,
//dodajanje temperaturnih meritev in izpis vremenskih poročil.
package vreme

import "fmt"

//Postaja predstavlja vremensko postajo z imenom, krajem in meritvami.
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

//DodajMeritev doda temperaturno meritev izbrani postaji,
//če postaja obstaja in je temperatura znotraj dovoljenih meja.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp float64, maxTemp float64) {
	if minTemp == 0 {
		minTemp = -50
	}
	if maxTemp == 0 {
		maxTemp = 50
	}

	e, ok := postaje[oznaka]

	if ok == true {
		if temperatura <= maxTemp && temperatura >= minTemp {
			e.meritve = append(e.meritve, temperatura)
			postaje[oznaka] = e
		} else {
			fmt.Printf("Meritev ni veljavna\n")
		}
	} else {
		fmt.Printf("Postaja %s ne obstaja\n", oznaka)
	}
}

//IzpisMeritev izpiše vse vremenske postaje in njihove meritve.
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Printf("MERITVE\n")

	for i, v := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", i, v.ime, v.kraj, v.meritve)
	}
}

//IzpisPorocila izračuna in izpiše vremensko poročilo za vse postaje.
//Za izračun statistike mora imeti postaja vsaj stMeritev meritev.
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	for i, v := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(postaje, i, stMeritev)

		if ok == false {
			fmt.Printf("%s: premalo meritev\n", v.ime)
			continue
		}

		fmt.Printf("%s: min %.1f°C, max %.1f°C, povprečje %.1f°C -> ", v.ime, najnizja, najvisja, povprecna)

		if povprecna >= 25 {
			fmt.Printf("Vroče\n")
		} else if povprecna < 25 && povprecna >= 15 {
			fmt.Printf("Toplo\n")
		} else if povprecna < 15 && povprecna >= 5 {
			fmt.Printf("Hladno\n")
		} else {
			fmt.Printf("Mrzlo\n")
		}
	}
}

func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	if stMeritev == 0 {
		stMeritev = 3
	}

	e, ok := postaje[oznaka]

	if !ok || len(e.meritve) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja = e.meritve[0]
	najvisja = e.meritve[0]
	var vsota float64

	for _, temperatura := range e.meritve {
		if temperatura < najnizja {
			najnizja = temperatura
		}

		if temperatura > najvisja {
			najvisja = temperatura
		}

		vsota += temperatura
	}

	povprecna = vsota / float64(len(e.meritve))

	return najnizja, najvisja, povprecna, true
}

//DodajPostajo doda novo vremensko postajo v slovar postaj.
func DodajPostajo(postaje map[string]Postaja, kljuc string, ime string, kraj string) {
	postaje[kljuc] = Postaja{ime, kraj, []float64{}}
}