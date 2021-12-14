GO = go build
SRCS := $(wildcard 2021/*/main.go)
BINS := $(SRCS:%.go=%)

all: ${BINS}

%: %.go
	${GO} -o $@ $<

run:
	./2021/day-01/main < ./2021/day-01/input
	./2021/day-02/main < ./2021/day-02/input
	./2021/day-03/main < ./2021/day-03/input
	./2021/day-04/main < ./2021/day-04/input
	./2021/day-05/main < ./2021/day-05/input
	./2021/day-06/main < ./2021/day-06/input
	./2021/day-07/main < ./2021/day-07/input
	./2021/day-08/main < ./2021/day-08/input
	./2021/day-09/main < ./2021/day-09/input
	./2021/day-10/main < ./2021/day-10/input
	./2021/day-11/main < ./2021/day-11/input
	./2021/day-12/main < ./2021/day-12/input
	./2021/day-13/main < ./2021/day-13/input
	./2021/day-14/main < ./2021/day-14/input

test:
	./2021/day-01/main < ./2021/day-01/test
	./2021/day-02/main < ./2021/day-02/test
	./2021/day-03/main < ./2021/day-03/test
	./2021/day-04/main < ./2021/day-04/test
	./2021/day-05/main < ./2021/day-05/test
	./2021/day-06/main < ./2021/day-06/test
	./2021/day-07/main < ./2021/day-07/test
	./2021/day-08/main < ./2021/day-08/test
	./2021/day-09/main < ./2021/day-09/test
	./2021/day-10/main < ./2021/day-10/test
	./2021/day-11/main < ./2021/day-11/test
	./2021/day-12/main < ./2021/day-12/test
	./2021/day-13/main < ./2021/day-13/test
	./2021/day-14/main < ./2021/day-14/test

aux:
	./2021/day-08/main < ./2021/day-08/aux
	./2021/day-12/main < ./2021/day-12/aux1
	./2021/day-12/main < ./2021/day-12/aux2

plot:
	python ./2021/day-09/plot.py ./2021/day-09/test
	python ./2021/day-09/plot.py ./2021/day-09/input

clean:
	rm */*/main
