compile: prepare main

main: *.cpp
	g++ -o main -Wall *.cpp

clean:
	rm -f main
	find . -iname \*.o -exec rm -f {} \;
