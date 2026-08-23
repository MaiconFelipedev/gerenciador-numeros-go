package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

var ErrFimDaEntrada = errors.New("fim da entrada")

func main() {
	var numeros []int
	for {
		exibirMenu()

		opcao, err := lerInteiro("Escolha uma opção: ")

		if errors.Is(err, ErrFimDaEntrada) {
			fmt.Println("\nEntrada encerrada.")
			return
		}

		if err != nil {
			fmt.Println("Entrada inválida. Tente novamente.")
			continue
		}

		switch opcao {
		case 1:
			n, err := lerInteiro("Digite o número: ")
			if errors.Is(err, ErrFimDaEntrada) {
				fmt.Println("\nEntrada encerrada.")
				return
			}
			if err != nil {
				fmt.Println("Número inválido. Tente novamente.")
				continue
			}
			if n < 0 {
				fmt.Println("Números negativos não são permitidos. Tente novamente.")
				continue
			}
			numeros = adicionarNumero(numeros, n)
			fmt.Println("Número adicionado com sucesso.")
		case 2:
			listarNumeros(numeros)
		case 3:
			listarNumeros(numeros)
			i, err := lerInteiro("Digite o índice: ")
			if errors.Is(err, ErrFimDaEntrada) {
				fmt.Println("\nEntrada encerrada.")
				return
			}
			if err != nil {
				fmt.Println("Índice inválido. Tente novamente.")
				continue
			}
			numeros, err = removerPorIndice(numeros, i)
			if err != nil {
				fmt.Println("Erro:", err)
				continue
			}
			fmt.Println("Número removido com sucesso.")
		case 4:
			menor, maior, media, err := estatisticas(numeros)
			if err != nil {
				fmt.Println("Erro:", err)
				continue
			}
			fmt.Printf("Mínimo: %d\n", menor)
			fmt.Printf("Máximo: %d\n", maior)
			fmt.Printf("Média: %.2f\n", media)
		case 5:
			a, err := lerInteiro("Digite o dividendo: ")
			if errors.Is(err, ErrFimDaEntrada) {
				fmt.Println("\nEntrada encerrada.")
				return
			}
			if err != nil {
				fmt.Println("Número inválido. Tente novamente.")
				continue
			}
			b, err := lerInteiro("Digite o divisor: ")
			if errors.Is(err, ErrFimDaEntrada) {
				fmt.Println("\nEntrada encerrada.")
				return
			}
			if err != nil {
				fmt.Println("Número inválido. Tente novamente.")
				continue
			}
			resultado, err := divisaoSegura(a, b)
			if err != nil {
				fmt.Println("Erro:", err)
				continue
			}
			fmt.Printf("%d / %d = %d\n", a, b, resultado)
		case 6:
			numeros = limparLista()
			fmt.Println("Lista limpa com sucesso.")
		case 7:
			listarNumeros(apenasPares(numeros))
		case 8:
			listarNumeros(apenasImpares(numeros))
		case 9:
			ordem, err := lerInteiro("1) Crescente\n2) Decrescente\nEscolha a ordem: ")
			if errors.Is(err, ErrFimDaEntrada) {
				fmt.Println("\nEntrada encerrada.")
				return
			}
			if err != nil || (ordem != 1 && ordem != 2) {
				fmt.Println("Opção de ordem inválida. Tente novamente.")
				continue
			}
			numeros = ordenar(numeros, ordem == 1)
			fmt.Println("Lista ordenada com sucesso.")
			listarNumeros(numeros)
		case 10:
			if err := exportarParaArquivo(numeros, "numeros.txt"); err != nil {
				fmt.Println("Erro:", err)
				continue
			}
			fmt.Println("Lista exportada para numeros.txt")
		case 0:
			fmt.Println("Saindo do programa...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}

func exibirMenu() {
	fmt.Println("=== Gerenciador de Numeros ===")
	fmt.Println("1) Adicionar numero")
	fmt.Println("2) Listar numeros")
	fmt.Println("3) Remover por indice")
	fmt.Println("4) Estatisticas")
	fmt.Println("5) Divisao segura")
	fmt.Println("6) Limpar lista")
	fmt.Println("7) Listar apenas pares")
	fmt.Println("8) Listar apenas impares")
	fmt.Println("9) Ordenar lista")
	fmt.Println("10) Exportar para arquivo")
	fmt.Println("0) Sair")
}

func lerInteiro(rotulo string) (int, error) {
	fmt.Print(rotulo)
	if !scanner.Scan() {
		return 0, ErrFimDaEntrada
	}
	linha := strings.TrimSpace(scanner.Text())
	return strconv.Atoi(linha)
}

func adicionarNumero(numeros []int, n int) []int {
	numeros = append(numeros, n)
	return numeros
}

func listarNumeros(numeros []int) {
	if len(numeros) == 0 {
		fmt.Println("A lista está vazia.")
		return
	}
	for i, v := range numeros {
		fmt.Printf("[%d] %d\n", i, v)
	}
}

func removerPorIndice(numeros []int, indice int) ([]int, error) {
	if indice < 0 || indice >= len(numeros) {
		return numeros, fmt.Errorf("índice %d inválido", indice)
	}
	return append(numeros[:indice], numeros[indice+1:]...), nil
}

func estatisticas(numeros []int) (int, int, float64, error) {
	if len(numeros) == 0 {
		return 0, 0, 0, errors.New("não é possível calcular estatísticas de uma lista vazia")
	}
	menor := numeros[0]
	maior := numeros[0]
	soma := 0

	for _, v := range numeros {
		if v < menor {
			menor = v
		}
		if v > maior {
			maior = v
		}
		soma += v
	}
	media := float64(soma) / float64(len(numeros))

	return menor, maior, media, nil
}

func divisaoSegura(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não é permitida")
	}
	return a / b, nil
}

func limparLista() []int {
	return nil
}

func ordenar(numeros []int, crescente bool) []int {
	sort.Slice(numeros, func(i, j int) bool {
		if crescente {
			return numeros[i] < numeros[j]
		}
		return numeros[i] > numeros[j]
	})
	return numeros
}

func apenasPares(numeros []int) []int {
	var pares []int
	for _, v := range numeros {
		if v%2 == 0 {
			pares = append(pares, v)
		}
	}
	return pares
}

func apenasImpares(numeros []int) []int {
	var impares []int
	for _, v := range numeros {
		if v%2 != 0 {
			impares = append(impares, v)
		}
	}
	return impares
}

func exportarParaArquivo(numeros []int, caminho string) error {
	if len(numeros) == 0 {
		return errors.New("lista vazia, nada para exportar")
	}
	arquivo, err := os.Create(caminho)
	if err != nil {
		return err
	}
	defer arquivo.Close()
	for _, v := range numeros {
		if _, err := fmt.Fprintln(arquivo, v); err != nil {
			return err
		}
	}
	return nil
}
