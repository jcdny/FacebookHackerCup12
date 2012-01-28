include $(GOROOT)/src/Make.inc

TARG=sq
GOFILES=\
	sq.go\

include $(GOROOT)/src/Make.cmd

test: sq
	./sq < testdata/test.txt
