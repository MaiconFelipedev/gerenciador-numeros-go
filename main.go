package main

// ============================================================
// ESQUELETO - remova todos os comentários antes de entregar.
// O enunciado exige que o código final não tenha comentários.
// ============================================================

func main() {
	// TODO: criar o slice de inteiros que guarda os números
	// TODO: laço principal -> exibirMenu(), ler a opção, chamar a função correspondente
	// TODO: opção 0 encerra o laço
}

func exibirMenu() {
	// TODO: imprimir as 7 opções do menu (1..6 e 0)
}

func lerInteiro(rotulo string) (int, error) {
	// TODO: mostrar o rótulo, ler a entrada do usuário e converter para int
	// TODO: devolver o erro de conversão em vez de tratá-lo aqui
	return 0, nil
}

func adicionarNumero(numeros []int, n int) []int {
	// TODO: acrescentar n ao slice e devolver o slice novo
	return numeros
}

func listarNumeros(numeros []int) {
	// TODO: se estiver vazio, avisar; senão imprimir índice + valor de cada elemento
}

func removerPorIndice(numeros []int, indice int) ([]int, error) {
	// TODO: validar se o índice está dentro dos limites do slice
	// TODO: remover o elemento e devolver o slice resultante
	return numeros, nil
}

func estatisticas(numeros []int) (int, int, float64, error) {
	// TODO: mínimo, máximo e média - erro se o slice estiver vazio
	return 0, 0, 0, nil
}

func divisaoSegura(a int, b int) (int, error) {
	// TODO: erro se b for zero; senão devolver a divisão
	return 0, nil
}

func limparLista() []int {
	// TODO: devolver um slice vazio
	return nil
}

// ---------- bônus (opcionais) ----------

func ordenar(numeros []int, crescente bool) []int {
	// TODO
	return numeros
}

func apenasPares(numeros []int) []int {
	// TODO
	return nil
}

func exportarParaArquivo(numeros []int, caminho string) error {
	// TODO
	return nil
}
