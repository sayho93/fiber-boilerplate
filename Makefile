BINARY_NAME=fiber

env-up:
	oci os object put -bn environments --file .env --name ${BINARY_NAME}/.env --no-multipart --force

env-down:
	oci os object get -bn environments --file .env --name ${BINARY_NAME}/.env

run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
	go build -o out/${BINARY_NAME} main.go
	GOOS=linux GOARCH=arm go build -o out/${BINARY_NAME}-linux-arm main.go
    GOOS=linux GOARCH=arm64 go build -o out/${BINARY_NAME}-linux-arm64 main.go
    GOOS=freebsd GOARCH=386 go build -o out/${BINARY_NAME}-freebsd-386 main.go
	GOOS=windows GOARCH=386 go build -o out/${BINARY_NAME}-windows-386 main.go

build:
	go build -o out/${BINARY_NAME} main.go
	GOARCH=amd64 GOOS=darwin go build -o out/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o out/${BINARY_NAME}-linux main.go

clean:
	go clean
	rm -r out

run-prod:
	./out/${BINARY_NAME}

deps:
	go mod download

vet:
	go vet