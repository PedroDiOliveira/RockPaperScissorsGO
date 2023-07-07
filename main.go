package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	jogar        bool
	respostaMenu int
)

func main() { //                                                          Main
	Introducao()
	jogar = true

	for jogar {
		MenuDoJogo()
		respostaMenu = VerificaRespostaMenu()
		fmt.Println("")

		switch respostaMenu {
		case 1:
			Jogando()
			fmt.Println("")
			fmt.Println("")
		case 2:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Resposta invalida")
		}

	}
}

func Introducao() {
	fmt.Println("|----------------------------|")
	fmt.Println("|----Pedra Papel e Tesoura---|")
	fmt.Println("|----------------------------|")
}

func MenuDoJogo() {
	fmt.Println("1- Novo Jogo")
	fmt.Println("2- Fechar o jogo")
}

func VerificaRespostaMenu() int {
	var x int
	fmt.Println("")
	fmt.Scan(&x)
	if x != 1 && x != 2 && x != 3 {
		return 4
	} else {
		return x
	}
}

var ( //Variaveis relacionadas ao jogo
	vidas              int
	vidasDoBot         int
	resposta           string
	respostaAdversario string
	numeroDeRodadas    int
)

func Jogando() {
	vidas = 3
	vidasDoBot = 3

	for vidas > 0 && vidasDoBot > 0 {

		fmt.Println("Rodada", numeroDeRodadas+1)
		fmt.Println("")

		DisplayDeEscolha()
		resposta = EscolhaDoDisplay()                 //sua resposta
		respostaAdversario = SorteioRespostaInimigo() //resposta do Bot

		fmt.Println("")

		switch resposta {
		case "Pedra":
			if respostaAdversario == "Pedra" {
				fmt.Println("Empate")
				estatisticas(resposta, respostaAdversario, vidas, vidasDoBot)
			} else if respostaAdversario == "Papel" {
				fmt.Println("Voce perdeu!")
				vidas--
				estatisticas(resposta, respostaAdversario, vidas, vidasDoBot)
			} else if respostaAdversario == "Tesoura" {
				fmt.Println("Voce ganhou!")
				vidasDoBot--
				estatisticas(resposta, respostaAdversario, vidas, vidasDoBot)
			}
		case "Papel":
			if respostaAdversario == "Pedra" {
				fmt.Println("Voce ganhou!")
				vidasDoBot--
				estatisticas(resposta, respostaAdversario, vidas, vidasDoBot)
			} else if respostaAdversario == "Papel" {
				fmt.Println("Empate")
				estatisticas(resposta, respostaAdversario, vidas, vidasDoBot)
			} else if respostaAdversario == "Tesoura" {
				fmt.Println("Voce perdeu!")
				vidas--
				estatisticas(resposta, respostaAdversario, vidas, vidasDoBot)
			}
		case "Tesoura":
			if respostaAdversario == "Pedra" {
				fmt.Println("Voce perdeu!")
				vidas--
				estatisticas(resposta, respostaAdversario, vidas, vidasDoBot)
			} else if respostaAdversario == "Papel" {
				fmt.Println("Voce ganhou!")
				vidasDoBot--
				estatisticas(resposta, respostaAdversario, vidas, vidasDoBot)
			} else if respostaAdversario == "Tesoura" {
				fmt.Println("Empate")
				estatisticas(resposta, respostaAdversario, vidas, vidasDoBot)
			}
		}
		if vidas == 0 {
			fmt.Println("Voce Perdeu!")
		} else if vidasDoBot == 0 {
			fmt.Println("Voce Ganhou!")
		}
		numeroDeRodadas++
	}
}

func DisplayDeEscolha() {
	fmt.Println("(1)Pedra")
	fmt.Println("(2)Papel")
	fmt.Println("(3)Tesoura")
}

func EscolhaDoDisplay() string {
	var x int
	for x != 1 && x != 2 && x != 3 {
		var x int
		fmt.Println("")
		fmt.Scan(&x)
		if x != 1 && x != 2 && x != 3 {
			fmt.Println("Valor invalido")
		}
		switch x {
		case 1:
			return "Pedra"
		case 2:
			return "Papel"
		case 3:
			return "Tesoura"
		}
	}
	return "Erro"
}

func SorteioRespostaInimigo() string {
	rand.Seed(time.Now().UnixNano())
	numeroAleatorio := rand.Intn(3) + 1
	switch numeroAleatorio {
	case 1:
		return "Pedra"
	case 2:
		return "Papel"
	case 3:
		return "Tesoura"
	}
	return "erro"
}
func estatisticas(x string, y string, z int, w int) {
	fmt.Println("PLAYER:", x, "     BOT:", y)
	fmt.Println("Suas vida:", z)
	fmt.Println("Suas vida:", w)
	fmt.Println("")
}
