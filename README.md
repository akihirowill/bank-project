# Bank Project - Alura

Um projeto de sistema bancário simples desenvolvido em Go, demonstrando conceitos de orientação a objetos, interfaces e operações financeiras.

A simple banking system project developed in Go, demonstrating object-oriented concepts, interfaces, and financial operations.

---

## 📚 Português

### Descrição

Este projeto implementa um sistema bancário básico em Go com suporte a dois tipos de contas:

- **Conta Corrente (ContaCorrente)**: Conta com funcionalidades básicas de saque, depósito e transferência
- **Conta Poupança (ContaPoupanca)**: Conta com funcionalidades de saque e depósito

### Funcionalidades

- ✅ Sacar dinheiro de uma conta
- ✅ Depositar dinheiro em uma conta
- ✅ Verificar saldo
- ✅ Transferir dinheiro entre contas (ContaCorrente)
- ✅ Pagar boletos através de uma conta
- ✅ Sistema de titulares com informações de CPF e profissão

### Estrutura do Projeto

```
bank-project/
├── main.go              # Arquivo principal com exemplos de uso
├── go.mod               # Declaração do módulo Go
├── LICENSE              # Licença do projeto
├── clientes/
│   └── cliente.go       # Definição do tipo Titular
└── contas/
    ├── contaCorrente.go    # Implementação da Conta Corrente
    └── contaPoupanca.go    # Implementação da Conta Poupança
```

### Como Usar

#### Pré-requisitos

- Go 1.21 ou superior instalado
- Terminal/Prompt de comando

#### Instalação

1. Clone ou baixe o projeto
2. Navegue até o diretório do projeto:
   ```bash
   cd bank-project
   ```

#### Executar o Projeto

Para executar o programa principal:

```bash
go run .\main.go
```

#### Compilar o Projeto

Para criar um executável:

```bash
go build -o bank.exe
```

Depois execute:

```bash
.\bank.exe
```

### Exemplo de Uso

```go
package main

import (
    "alura/bank-project/contas"
    "fmt"
)

func main() {
    // Criar uma conta poupança
    contaDoDenis := contas.ContaPoupanca{}
    contaDoDenis.Depositar(100)
    fmt.Println(contaDoDenis.ObterSaldo()) // 100
    
    // Criar uma conta corrente
    contaDaLuiza := contas.ContaCorrente{}
    contaDaLuiza.Depositar(200)
    fmt.Println(contaDaLuiza.ObterSaldo()) // 200
}
```

### Operações Disponíveis

#### Depositar
```go
conta.Depositar(valor float64) string
// Retorna mensagem de sucesso ou erro
```

#### Sacar
```go
conta.Sacar(valor float64) string
// Retorna mensagem de sucesso ou erro
```

#### Obter Saldo
```go
conta.ObterSaldo() float64
// Retorna o saldo atual da conta
```

#### Transferir (apenas ContaCorrente)
```go
conta.Transferir(valor float64, contaDestino *ContaCorrente) bool
// Retorna true se bem-sucedido, false caso contrário
```

---

## 📚 English

### Description

This project implements a basic banking system in Go with support for two types of accounts:

- **Checking Account (ContaCorrente)**: Account with basic withdrawal, deposit, and transfer features
- **Savings Account (ContaPoupanca)**: Account with withdrawal and deposit features

### Features

- ✅ Withdraw money from an account
- ✅ Deposit money into an account
- ✅ Check balance
- ✅ Transfer money between accounts (ContaCorrente)
- ✅ Pay bills through an account
- ✅ Account holder system with CPF and profession information

### Project Structure

```
bank-project/
├── main.go              # Main file with usage examples
├── go.mod               # Go module declaration
├── LICENSE              # Project license
├── clientes/
│   └── cliente.go       # Definition of Titular type
└── contas/
    ├── contaCorrente.go    # Checking Account implementation
    └── contaPoupanca.go    # Savings Account implementation
```

### How to Use

#### Prerequisites

- Go 1.21 or higher installed
- Terminal/Command prompt

#### Installation

1. Clone or download the project
2. Navigate to the project directory:
   ```bash
   cd bank-project
   ```

#### Run the Project

To run the main program:

```bash
go run .\main.go
```

#### Build the Project

To create an executable:

```bash
go build -o bank.exe
```

Then execute:

```bash
.\bank.exe
```

### Usage Example

```go
package main

import (
    "alura/bank-project/contas"
    "fmt"
)

func main() {
    // Create a savings account
    savingsAccount := contas.ContaPoupanca{}
    savingsAccount.Depositar(100)
    fmt.Println(savingsAccount.ObterSaldo()) // 100
    
    // Create a checking account
    checkingAccount := contas.ContaCorrente{}
    checkingAccount.Depositar(200)
    fmt.Println(checkingAccount.ObterSaldo()) // 200
}
```

### Available Operations

#### Deposit
```go
account.Depositar(valor float64) string
// Returns success or error message
```

#### Withdraw
```go
account.Sacar(valor float64) string
// Returns success or error message
```

#### Get Balance
```go
account.ObterSaldo() float64
// Returns the current account balance
```

#### Transfer (Checking Account only)
```go
account.Transferir(valor float64, contaDestino *ContaCorrente) bool
// Returns true if successful, false otherwise
```

---

## 📄 License

Este projeto está licenciado sob a licença especificada no arquivo LICENSE.

This project is licensed under the license specified in the LICENSE file.

---

## 👨‍💻 Desenvolvedor

Projeto desenvolvido como parte do curso Alura de Go.

Project developed as part of the Alura Go course.
