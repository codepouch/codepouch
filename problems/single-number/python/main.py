import sys
from solution import single_number

def main():
    # first line is the number of test cases
    cases = int(sys.stdin.readline().strip())
    for _ in range(cases):
        # first line in each test case is the number of integers given
        count = int(sys.stdin.readline().strip())

        # second line in each test case are the numbers
        A = map(int, sys.stdin.readline().strip().split())
        assert len(A) == count

        # run algorithm
        output = single_number(A)

        # dump output
        sys.stdout.write('%d\n' % output)

if __name__ == '__main__':
    main()

