#include <iostream>
#include <string>

int main() {
    int depth = 0;
    int forward = 0;
    std::string direction;
    int amount;
    
    while (std::cin >> direction >> amount) {
        if (direction == "forward") {
            forward += amount;
        } else if (direction == "up") {
            depth -= amount;
        } else {
            depth += amount;
        };
    };

    std::cout << depth * forward << std::endl;
};

