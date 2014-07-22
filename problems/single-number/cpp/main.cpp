#include <iostream>

int singleNumber(int *A, int n);

int main() {
    int cases = 0;
    std::cin >> cases;

    for (int i = 0; i < cases; i++) {
        int n = 0;
        std::cin >> n;

        int *A = new int[n];
        for (int k = 0; k < n; k++) {
            std::cin >> A[k];
        }

        int output = singleNumber(A, n);
        std::cout << output << std::endl;

        delete[] A;
    }

    return 0;
}

