#include <iostream>
#include <vector>

int main() {
    int depth;
    int result = 0;
    
    std::vector<int> depths;
    while (std::cin >> depth) {
        depths.push_back(depth);
    };

    for (int i=0; i<depths.size()-2; i++) {
        if (depths[i] < depths[i + 3]) {
            result++;
        };
    };  

    std::cout << result << std::endl;
};

