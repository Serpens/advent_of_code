#include <iostream>
#include <string>

int main() {
    int depth = 0;
    int forward = 0;
    int aim = 0;
    std::string direction;
    int amount;
    
    while (std::cin >> direction >> amount) {
        if (direction == "forward") {
            forward += amount;
            depth += aim * amount;
        } else if (direction == "up") {
            aim -= amount;
        } else {
            aim += amount;
        };
    };

    std::cout << depth * forward << std::endl;
};

