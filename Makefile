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

test:
	./2021/day-01/main < ./2021/day-01/test
	./2021/day-02/main < ./2021/day-02/test
	./2021/day-03/main < ./2021/day-03/test
	./2021/day-04/main < ./2021/day-04/test
	./2021/day-05/main < ./2021/day-05/test
	./2021/day-06/main < ./2021/day-06/test
	./2021/day-07/main < ./2021/day-07/test

clean:
	rm */*/main
