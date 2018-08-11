#include <iostream>

#ifndef TOKEN
#define TOKEN
#include "./enums/tokenType.cpp"
#endif

using namespace std;

class Token
{
public:
  TokenType type = TokenType::Unknown;
  string data = "";

  Token(TokenType type, string data)
  {
    this->type = type;
    this->data = data;
  }
};