install: install-def install-sqlboiler

install-def:
	curl -L -O https://github.com/k0kubun/sqldef/releases/download/v0.11.58/mysqldef_darwin_amd64.zip
	unzip mysqldef_darwin_amd64.zip
	rm mysqldef_darwin_amd64.zip
	mv mysqldef /usr/local/bin/mysqldef
install-sqlboiler:
	go install github.com/volatiletech/sqlboiler/v4@v4.14.2
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.14.2

generate-mysql:
	mysqldef -uapp -ppassword db < schema.sql
	sqlboiler mysql

generate-di:
	wire ./di/wire.go

run:
	go run cmd/app/main.go