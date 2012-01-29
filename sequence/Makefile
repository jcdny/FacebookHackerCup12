include $(GOROOT)/src/Make.inc

TARG=ms
GOFILES=\
	ms.go\

include $(GOROOT)/src/Make.cmd

test: ms
	./ms < testdata/test.txt
