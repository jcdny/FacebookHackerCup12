include $(GOROOT)/src/Make.inc

TARG=checkpoint
GOFILES=\
	checkpoint.go\

include $(GOROOT)/src/Make.cmd

test: checkpoint
	./checkpoint < testdata/test.txt
