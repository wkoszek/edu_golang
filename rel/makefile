# (c) 2015 Wojciech A. Koszek <wojciech@koszek.com>

SRCS:=$(wildcard *.go)
PROGS:=$(SRCS:.go=.prog)

objs: $(PROGS)

%.prog: %.go
	go build -i -o $@ $<

all:
	env GOOS=linux   GOARCH=386   go build -o hello.linux.386 hello.go
	env GOOS=linux   GOARCH=amd64 go build -o hello.linux.amd64 hello.go
	env GOOS=freebsd GOARCH=386   go build -o hello.freebsd.386 hello.go
	env GOOS=freebsd GOARCH=amd64 go build -o hello.freebsd.amd64 hello.go
	env GOOS=darwin  GOARCH=386   go build -o hello.darwin.386 hello.go
	env GOOS=darwin  GOARCH=amd64 go build -o hello.darwin.amd64 hello.go

#const goosList = "android darwin dragonfly freebsd linux nacl netbsd openbsd plan9 solaris windows zos "
#const goarchList = "386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 ppc64le mips mipsle mips64 mips64le mips64p32 mips64p32le ppc s390 s390x sparc sparc64 "
#
clean:
	rm -rf *.prog
	rm -rf *.linux.* *.freebsd.* *.darwin.*
	rm -rf hello hello-2.10 hello-2.10.tar.gz hello.prog hello_2.10.orig.tar.gz
