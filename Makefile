all: day-01 day-02 day-03

day-01: 2021/day-01/main.go
	go build -o 2021/day-01/main 2021/day-01/main.go

day-02: 2021/day-02/main.go
	go build -o 2021/day-02/main 2021/day-02/main.go

day-03: 2021/day-03/main.go
	go build -o 2021/day-03/main 2021/day-03/main.go

run:
	./2021/day-01/main < ./2021/day-01/input
	./2021/day-02/main < ./2021/day-02/input
	./2021/day-03/main < ./2021/day-03/input

test:
	./2021/day-01/main < ./2021/day-01/test
	./2021/day-02/main < ./2021/day-02/test
	./2021/day-03/main < ./2021/day-03/test

clean:
	rm */*/main
