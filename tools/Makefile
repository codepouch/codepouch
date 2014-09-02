.PHONY: all prepare compile test clean
all: prepare compile test

prepare:
	@for t in *.template; do \
		[ -e "$${t%.template}" ] || cp $$t $${t%.template}; \
	done

include ../../../tools/$(shell basename `pwd`).mk

