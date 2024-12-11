// GenerateParentheses.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <vector>

void GenerateParentheses(int n, int m, std::string s, std::vector<std::string>& result)
{
    if (n == 0 && m == 0)
    {
        result.push_back(s);
        return;
    }

    if (n > 0)
        GenerateParentheses(n - 1, m, s + "(", result);
    if (m > n)
        GenerateParentheses(n, m - 1, s + ")", result);
}

std::vector<std::string> GenerateParentheses(int n)
{
    std::vector<std::string> result;
    std::string s;
    int m = n;

    GenerateParentheses(n, m, s, result);
    return result;
}

int main()
{
    for (int i = 1; i < 5; ++i)
    {
        std::vector<std::string> r = GenerateParentheses(i);
        for (auto s : r)
            std::cout << s << std::endl;
        std::cout << std::endl;
    }
}
