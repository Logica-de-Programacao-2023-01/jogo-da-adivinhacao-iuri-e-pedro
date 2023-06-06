package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var Matriztentativas [][]int
	jogar := true

	for jogar {
		num := rand.Intn(100)
		fmt.Println("Qual número entre 0 e 100 você escolhera? ")

		var tentativasFeitas []int
		tentativa := 0
		contagemTentativas := 0

		for tentativa != num {
			_, err := fmt.Scan(&tentativa)
			if tentativa <= -1 || tentativa >= 101 {
				fmt.Println("O númeor está fora do limite de 0 e 100")
			} else if err != nil {
				fmt.Println("Entrada invalida, Por favor digite um valor inteiro entre 0 e 100")
			}

			contagemTentativas++
			tentativasFeitas = append(tentativasFeitas, tentativa)

			if tentativa < num {
				fmt.Println("O num é maior que", tentativa)
			} else if tentativa > num {
				fmt.Println("O numero é menor que ", tentativa)
			}
		}
		fmt.Printf("Você completou em %d tentativas. \n", contagemTentativas)
		Matriztentativas = append(Matriztentativas, tentativasFeitas)

		fmt.Println("deseja jogar novamente? (s/n):")
		var inputjogar string
		fmt.Scan(&inputjogar)
		jogar = inputjogar == "s"

		fmt.Println("Numero de tentativas feitas em todas as jogadas: ")
		for i, tentativas := range Matriztentativas {
			fmt.Printf("Jogada %d: %v\n", i+1, tentativas)
		}
	}
}
