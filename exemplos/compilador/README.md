# Exemplo de compilação de código

Go é uma linguagem compilada, ou seja, a partir do código fonte é gerado (compilado) um binário, executável pela CPU.

O Go conta com um interpretador nativo (`go run`), usado para desenvolvimento e/ou execução de scripts.

Nesse diretório vamos exemplificar como rodar um *hello world* utilizando um binário e utilizando o interpretador do Go.

## Compilando e executando o binário

**Os comandos a seguir são válidos para execução nesse diretório.**

```bash
go build main.go
# irá compilar o arquivo main.go, do diretório atual, para um binário
# nomeado <main>. Para alterar o nome do binário a ser criado, acrescente
# <-o NOVO_NOME> depois do comando <build>.
./main
```

```bash
go run main.go
```
