EXECUTE="java Main"

compile: Main.class

%.class: %.java
	@javac $<

.PHONY: clean
clean:
	@find . -iname \*.class -exec rm -f {} \;
