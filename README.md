# Bank Project

Um projeto de sistema bancário desenvolvido em Go, com funcionalidades de gestão de contas e clientes.

## 📋 Descrição

Este projeto implementa um sistema bancário básico com suporte a:
- Contas Correntes
- Contas Poupança
- Operações de saque, depósito e transferência
- Gestão de titular/cliente

## 🏗️ Estrutura do Projeto

```
bank-project/
├── main.go                 # Arquivo principal da aplicação
├── clientes/
│   └── cliente.go         # Definição do tipo Titular
└── contas/
    ├── contaCorrente.go   # Implementação de Conta Corrente
    └── contaPoupanca.go   # Implementação de Conta Poupança
```

## 🚀 Como Executar

### Pré-requisitos
- Go 1.21 ou superior instalado

### Instalação e Execução

1. Clone o repositório:
```bash
git clone <URL-do-repositorio>
cd bank-project
```

2. Execute o programa:
```bash
go run ./main.go
```

## 💳 Funcionalidades

### Conta Corrente
- **Sacar**: Realiza saque da conta com validação de saldo
- **Depositar**: Adiciona valor à conta
- **Transferir**: Transfere dinheiro entre contas
- **Obter Saldo**: Retorna o saldo atual da conta

### Conta Poupança
- **Sacar**: Realiza saque com validação de saldo
- **Depositar**: Adiciona valor à conta
- **Obter Saldo**: Retorna o saldo atual

## 📝 Exemplo de Uso

```go
contaDoDenis := contas.ContaPoupanca{}
contaDoDenis.Depositar(100)
```

## 🤝 Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

## 📄 Licença

Este projeto está sob a licença MIT.

## 👨‍💻 Autor

Desenvolvido como projeto educacional da Alura.
