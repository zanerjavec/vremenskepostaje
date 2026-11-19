package main

import (
	"fmt"
	"context"
	"log"
	"os"
	"github.com/urfave/cli/v3"
	"vremenskepostaje/pkg/vreme"
)

func main() {
    cmd := &cli.Command{
        Name:  "vremenske-postaje",
        Usage: "Program za upravljanje vremenskih postaj",

        // Definicija stikal
        Flags: []cli.Flag{
            &cli.Float64Flag{
                Name:  "minTemp",
                Value: -50,
                Usage: "Najnižja dovoljena temperatura",
            },
            &cli.Float64Flag{
                Name:  "maxTemp",
                Value: 50,
                Usage: "Najvišja dovoljena temperatura",
            },
            &cli.IntFlag{
                Name:  "stMeritev",
                Value: 3,
                Usage: "Najmanjše število meritev za statistiko",
            },
        },

        // Funkcija, ki se izvede ob zagonu programa
        Action: func(ctx context.Context, cmd *cli.Command) error {

            // Pridobivanje vrednosti stikal
            minTemp := cmd.Float64("minTemp")
            maxTemp := cmd.Float64("maxTemp")
            stMeritev := cmd.Int("stMeritev")

            // Preverjanje nastavitev
            if minTemp > maxTemp {
                return fmt.Errorf("minTemp ne sme biti večji od maxTemp")
            }

            if stMeritev < 1 {
                return fmt.Errorf("stMeritev mora biti vsaj 1")
            }

            // Ustvarjanje slovarja
            postaje := make(map[string]vreme.Postaja)

            // Dodajanje postaj
            vreme.DodajPostajo(postaje, "KP01", "Koper Luka", "Koper")
            vreme.DodajPostajo(postaje, "KR01", "Kredarica", "Triglav")
            vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Bežigrad", "Ljubljana")
            vreme.DodajPostajo(postaje, "MB01", "Maribor Tabor", "Maribor")
            vreme.DodajPostajo(postaje, "LJ02", "Ljubljana ZOO", "Ljubljana")

            // Dodajanje meritev
            vreme.DodajMeritev(postaje, "KP01", 22.5, minTemp, maxTemp)
            vreme.DodajMeritev(postaje, "KP01", 25.3, minTemp, maxTemp)
            vreme.DodajMeritev(postaje, "KP01", 27.1, minTemp, maxTemp)

            vreme.DodajMeritev(postaje, "KR01", -5.2, minTemp, maxTemp)
            vreme.DodajMeritev(postaje, "KR01", -3.8, minTemp, maxTemp)
            vreme.DodajMeritev(postaje, "KR01", -1.5, minTemp, maxTemp)

            vreme.DodajMeritev(postaje, "LJ01", 18.2, minTemp, maxTemp)
            vreme.DodajMeritev(postaje, "LJ01", 20.5, minTemp, maxTemp)
            vreme.DodajMeritev(postaje, "LJ01", 21.7, minTemp, maxTemp)

            vreme.DodajMeritev(postaje, "MB01", 12.4, minTemp, maxTemp)
            vreme.DodajMeritev(postaje, "MB01", 14.1, minTemp, maxTemp)
            vreme.DodajMeritev(postaje, "MB01", 16.3, minTemp, maxTemp)

            // Izpis
            fmt.Println("Izpis meritev")
            vreme.IzpisMeritev(postaje)

            fmt.Println("\nIzpis poročila")
            vreme.IzpisPorocila(postaje, stMeritev)

            return nil
        },
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}