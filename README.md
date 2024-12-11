Generate Parentheses


Given a number n denoting n pairs of parentheses, return all valid expressions using the n pairs of parentheses.

Examples
n: 2
Answer: [
    "(())",
    "()()",
]
n: 3
Answer: [
    "((()))",
    "(()())",
    "(())()",
    "()(())",
    "()()()"
]
Testing
Input Format
The first line contains an integer ‘T’, denoting the number of test cases.

For each test case, the input has an integer n.

Output Format
For each test case, the output has multiple lines each denoting a valid parenthesis expression.

Sample Input
3
1
2
3
Expected Output
()
(())
()()
((()))
(()())
(())()
()(())
()()()