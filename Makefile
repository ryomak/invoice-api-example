install: install-def install-sqlboiler

.PHONY: install-def
install-def:
	curl -L -O https://github.com/k0kubun/sqldef/releases/download/v0.11.58/mysqldef_darwin_amd64.zip
	unzip mysqldef_darwin_amd64.zip
	rm mysqldef_darwin_amd64.zip
	mv mysqldef /usr/local/bin/mysqldef

.PHONY: install-sqlboiler
install-sqlboiler:
	go install github.com/volatiletech/sqlboiler/v4@v4.14.2
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.14.2

.PHONY: generate-mysql
generate-mysql:
	mysqldef -uapp -ppassword db < schema.sql
	sqlboiler mysql

.PHONY: generate-di
generate-di:
	wire ./di/wire.go

.PHONY: run
run:
	go run cmd/app/main.go

.PHONY: test
test:
	go test -v ./...