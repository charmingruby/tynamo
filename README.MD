# Tynamo

Local DynamoDB with LocalStack, made using Go and Terraform.

# Requisitos

Algumas ferramentas são necessárias para utilizar a aplicação:

- Docker
- Terraform(tflocal): mesma coisa que Terraform, mas já adaptado para o ambiente local.

```
pip install terraform-local
```

- Go
- Make

# Execução

## 1 - Iniciar LocalStack

```
$ docker compose up
```

## 2 - Iniciar configurações feitas no Terraform

```
$ cd ./terraform
$ make run-tf
```

## 3 - Executar Go app (raiz do projeto)

```
$ make run
```
