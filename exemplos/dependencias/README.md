# Básico de dependências em Go

## Módulo

Cada aplicação em Go contém um mais módulos. Um módulo contém um arquivo `go.mod`, onde definimos o nome do módulo (usualmente o endereço do repositório), a versão do go utilizada, e as dependências diretas e indiretas.

## Import

Em Go, quando queremos acessar um recurso (tipo, interface, variável, função etc.) de outro pacote, seja da lib padrão, de terceiros ou da própria aplicação, importamos o pacote todo. Apesar disso, o compilador analisa quais partes do código não estão sendo utilizadas e as exclui do binário.

Para importar um pacote, basta utilizar a sintaxe

```go
import "nomeDoPacote"
```

podendo o `nomeDoPacote` ser:

1. Apenas o nome do pacote (pacotes da lib padrão)
2. O nome pacote importado no arquivo `go.mod` (pacotes de terceiros)
3. O nome declarado em `module` no arquivo `go.mod` + caminho para o pacote (pacotes declarados no mesmo projeto)

Quando importando mais de um pacote, a sintaxe e estilo mais utilizado é:

```go
import (
    //pacotes da lib padrão

    //pacotes de terceiros

    //pacotes do mesmo módulo
)
```

## Instalando dependências

Quando clonar um projeto existente, para baixar as dependências basta executar, no mesmo diretório onde estiver o `go.mod` (geralmente a raiz do projeto), o comando `go mod download`.

Durante o desenvolvimento, para baixar uma dependência específica utilize o comando `go get <nome do pacote>`.
