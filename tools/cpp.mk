compile: main

main: *.cpp
	@g++ -o main -Wall *.cpp

clean:
	@find . -iname \*.o -exec rm -f {} \;
