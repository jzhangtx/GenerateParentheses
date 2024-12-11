GenerateParentheses: GenerateParentheses.o
	g++ -g -o GenerateParentheses.exe GenerateParentheses.o -pthread    

GenerateParentheses.o: GenerateParentheses/GenerateParentheses.cpp
	g++ -g  -c -pthread GenerateParentheses/GenerateParentheses.cpp
