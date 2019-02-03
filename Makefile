include .env
export $(shell sed 's/=.*//' .env)

REPOS:=github.com/hiroykam/goa-sample
DESIGN:=$(REPO)/design

dep:
	@dep ensure

.PHONY: app
app:
	@goagen app -d github.com/hiroykam/goa-sample/design

.PHONY: swagger
swagger:
	@goagen swagger -d github.com/hiroykam/goa-sample/design

.PHONY: controller
controller:
	@goagen controller -d github.com/hiroykam/goa-sample/design --out=./controller

docker-build:
	@docker-compose build

docker-up:
	@docker-compose up

.PHONY: read-env
read-env:
	@. .env

migration-up:
	/bin/sleep 10; DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) DB_HOST=$(DB_HOST) DATABASE=$(DATABASE) sql-migrate up -env=local -config=./configs/dbconfig.yml

compile:
	CompileDaemon -log-prefix=true -directory="." -command="./goa-sample"

docker-cmd: migration-up compile
