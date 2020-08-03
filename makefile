# (c) 2015 Wojciech A. Koszek <wojciech@koszek.com>

SRCS:=$(wildcard *.go)
PROGS:=$(SRCS:.go=.prog)

objs: $(PROGS)

%.prog: %.go
	gofmt -w *.go
	goimports -w *.go
	go build -i -o $@ $<

clean:
	rm -rf *.prog
