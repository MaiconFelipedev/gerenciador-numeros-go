# Gerenciador de Números em Go

Miniprojeto de nivelamento em Go — gerenciador de números inteiros em linha de comando.

**Instituto Federal da Paraíba — Campus João Pessoa**
Unidade Acadêmica de Informática · Programação Distribuída
Prof. Ruan Delgado Gomes

## Sobre

Aplicação CLI que permite adicionar, listar, remover e calcular estatísticas sobre
números inteiros armazenados em um slice. Exercita variáveis, condicionais, laços,
funções, slices e funções com múltiplos retornos.

## Como executar

```
go run main.go
```

## Menu

| Opção | Funcionalidade |
| --- | --- |
| 1 | Adicionar número |
| 2 | Listar números |
| 3 | Remover por índice |
| 4 | Estatísticas (mínimo, máximo e média) |
| 5 | Divisão segura |
| 6 | Limpar lista |
| 0 | Sair |

## Requisitos do enunciado

- [ ] Implementado em um único arquivo `main.go`
- [ ] Compila e executa com `go run main.go`
- [ ] Slice de inteiros como estrutura de armazenamento
- [ ] Código modularizado em funções
- [ ] Estatísticas com múltiplos retornos
- [ ] Divisão retorna erro quando o divisor é zero
- [ ] Tratamento de erro no padrão Go (`if err != nil`)
- [ ] Código sem comentários

## Bônus

- [ ] Impedir números negativos
- [ ] Ordenação crescente e decrescente
- [ ] Exibir apenas números pares
- [ ] Exportar a lista para arquivo texto

## Pacotes úteis

`fmt` · `bufio` · `os` · `strconv` · `strings` · `errors` · `sort`
