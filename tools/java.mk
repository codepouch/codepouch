EXECUTE="java Main"

compile: prepare Main.class

%.class: %.java
	javac $<

.PHONY: clean
clean:
	find . -iname \*.class -exec rm -f {} \;
