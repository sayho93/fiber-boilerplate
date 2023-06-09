env-up:
	oci os object put -bn environments --file .env --name fiber/.env --no-multipart --force

env-down:
	oci os object get -bn environments --file .env --name fiber/.env

run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
	go build -o out/main main.go
	GOOS=linux GOARCH=arm go build -o out/main-linux-arm main.go
    GOOS=linux GOARCH=arm64 go build -o out/main-linux-arm64 main.go
    GOOS=freebsd GOARCH=386 go build -o out/main-freebsd-386 main.go
	GOOS=windows GOARCH=386 go build -o out/main-windows-386 main.go

BINARY_NAME=fiber

build:
	go build -o out/main main.go
	GOARCH=amd64 GOOS=darwin go build -o out/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o out/${BINARY_NAME}-linux main.go

clean:
	go clean
	rm -r out

deps:
	go mod download

vet:
	go vet