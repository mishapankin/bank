bin/bank: *.go
	go build -o bin/bank .

run: bin/bank
	export DB_URL="postgresql://admin:123456@localhost:5432/bank_db" SERVER_URL="localhost:8080";  bin/bank