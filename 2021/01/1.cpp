#include <iostream>

int main() {
    int depth, oldDepth = 100000000;
    int result = 0;
    while (std::cin >> depth) {
        if (depth > oldDepth) {
            result++;
        };
        oldDepth = depth;
    };
    std::cout << result << std::endl;
};

