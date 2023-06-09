env-up:
	oci os object put -bn environments --file .env --name fiber/.env --no-multipart --force
env-down:
	oci os object get -bn environments --file .env --name fiber/.env
build:
	go build -o out/main main.go
run:
	go run main.go
