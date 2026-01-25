 

GOCMD=go
MYSQL_PASS= --db-pass
MYSQL_HOST= --db-host
MYSQL_USER= --db-user
MYSQL_PORT= --db-port
MYSQL_DB_NAME= --db-name
EMAIL_API_KEY= --email-api-key
HTTP_SERVER_PORT= --http-port
GOBUILD=$(GOCMD) build


BINARY_NAME=iamgenii

BUILD_PATH=./bin/$(BINARY_NAME)

remove-build: 
	rm -rf bin/

build:
	rm -rf bin/
	$(GOBUILD) -o bin/$(BINARY_NAME) cmd/main.go
   
run:
	$(BUILD_PATH) $(MYSQL_HOST) localhost $(MYSQL_USER) root $(MYSQL_PASS) password $(MYSQL_PORT) 3306 $(MYSQL_DB_NAME) iamgenii $(EMAIL_API_KEY) 3aLd0KyItbpCJxN4 $(HTTP_SERVER_PORT) 8081

upload: 
	scp bin/iamgenii root@31.220.62.185:/root/work/


runs:
	./iamgenii --db-host 185.201.11.212 --db-user u156722531_rahuls --db-pass password --db-port 3306 --db-name u156722531_iamgenii --email-api-key 3aLd0KyItbpCJxN4

ssh: 
	ssh root@31.220.62.185

local_start: build  run
