all: day-01

day-01: 2021/day-01/main.go
	go build -o 2021/day-01/main 2021/day-01/main.go

run:
	./2021/day-01/main

clean:
	rm */*/main
