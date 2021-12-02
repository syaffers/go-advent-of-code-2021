all: day-01 day-02

day-01: 2021/day-01/main.go
	go build -o 2021/day-01/main 2021/day-01/main.go

day-02: 2021/day-02/main.go
	go build -o 2021/day-02/main 2021/day-02/main.go

run:
	./2021/day-01/main < ./2021/day-01/input
	./2021/day-02/main < ./2021/day-02/input

test:
	./2021/day-01/main < ./2021/day-01/test
	./2021/day-02/main < ./2021/day-02/test

clean:
	rm */*/main
