#include <iostream>
#include <vector>
#include <string>
#include <bitset>
#include <algorithm>
#include <numeric>

int main() {
    std::vector<int> numbers;
    std::string current;
    int gamma = 0 , epsilon = 0;
    
    while (std::cin >> current) {
        numbers.push_back((int) std::bitset<64>(current).to_ulong());
    };

    auto getGammaDigit  = [](std::vector<int> nums, int position) {
        std::vector<int> digits;
        std::transform(nums.begin(), nums.end(), 
            std::back_inserter(digits), 
            [&position](int n) {return (n & ( 1 << position )) >> position;});
        return std::accumulate(digits.begin(), digits.end(), 0 ) > nums.size()/2;
    };
    
    for (int i=0; i<current.size(); i++){
        auto digit = getGammaDigit(numbers, i);
        gamma += digit << i;
        epsilon += (digit ? 0 : 1)  << i;
    };

    std::cout << gamma * epsilon << std::endl;
};

