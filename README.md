# go-clean-arch

## Changelog

- **v1**: checkout to the [v1 branch](https://github.com/bxcodec/go-clean-arch/tree/v1) <br>
  Proposed on 2017, archived to v1 branch on 2018 <br>
  Desc: Initial proposal by me. The story can be read here: https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047

- **v2**: checkout to the [v2 branch](https://github.com/bxcodec/go-clean-arch/tree/v2) <br>
  Proposed on 2018, archived to v2 branch on 2020 <br>
  Desc: Improvement from v1. The story can be read here: https://medium.com/@imantumorang/trying-clean-architecture-on-golang-2-44d615bf8fdf

- **v3**: master branch <br>
  Proposed on 2019, merged to master on 2020. <br>
  Desc: Introducing Domain package, the details can be seen on this PR [#21](https://github.com/bxcodec/go-clean-arch/pull/21)

## How to run
- go mod init
- go mod tidy
- go mod vendor
- go run app/main.go

## How Push
- git add (pilih file yang mau diadd)
- git commit -m "message commit"
- git push origin development

## Use migration
### Create Migration
- migrate create -ext sql -dir db/migrations -seq create_users_table
### Run Migration
- migrate -database YOUR_DATABASE_URL -path PATH_TO_YOUR_MIGRATIONS up
example up:
- migrate -database "mysql://root:d4esUqz@QpS9XZNv@tcp(localhost:3306)/article" -path db/migrations up
example down:
- migrate -database "mysql://root:#d4esUqzQpS9XZNv@tcp(localhost:3306)/article" -path db/migrations down
### Detail YOUR_DATABASE_URL
["mysql://root:secret@tcp(localhost:3306)/simple_bank"]
- We’re using mysql, so the driver name is mysql.
- Then the username is root
- The password is secret
- The address is localhost, port 3306.
- And the database name is simple_bank.

## Use Mockery
### Install
- go install github.com/vektra/mockery/v2@v2.20.0
### Run Mockery
- cd domain
- mockery --all --case=camel

## Setup Email SMTP Gmail for Sender
### Get Password Aplication
- visiting https://myaccount.google.com/apppasswords should allow you to set up application specific passwords

### Pre-commit Installation
## On Our Local System
- Install PIP on your local machine, you can follow [this link](https://pip.pypa.io/en/stable/installation).
- Before you can run hooks, you need to have the pre-commit package manager installed.
- `pip install pre-commit`
- `pre-commit --version` should show you what version you're using
- Add a pre-commit configuration, create a file named `.pre-commit-config.yaml`, choose one based on your own language
##### Sample file config
```yaml
repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.3.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-added-large-files
  - repo: https://github.com/compilerla/conventional-pre-commit
    rev: v2.3.0
    hooks:
      - id: conventional-pre-commit
        stages: [ commit-msg ]
  - repo: https://github.com/dnephin/pre-commit-golang
    rev: v0.5.0
    hooks:
      - id: golangci-lint
        stages: [ pre-commit ]
      - id: go-unit-tests
        stages: [ pre-commit ]

```