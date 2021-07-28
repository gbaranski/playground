#include <bits/stdc++.h>
#include <iostream>
#include <vector>
#include <iterator>

int main() {
    std::vector<int> vec;
    for (int i = 0; i < 10; i++) {
        vec.push_back(i);
    }

    for (auto it = begin(vec); it != end(vec); ++it) {
        if (*it == 5) {
            vec.erase(it);
        }
    }

    for (auto it = begin(vec); it != end(vec); ++it) {
        if (*it == 4) {
            vec.erase(it);
        }
    }

    for (auto it = begin(vec); it != end(vec); ++it) {
        std::cout << *it << std::endl;
    }
}

