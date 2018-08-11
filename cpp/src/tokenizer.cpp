#include <iostream>
#include <vector>

#include "./token.cpp"
#include "./enums/tokenType.cpp"

using namespace std;

#ifndef TOKENIZER
#define TOKENIZER

class Tokenizer
{
public:
  static vector<Token *> Tokenize(const string line)
  {
    vector<Token *> list;

    for (auto ch : line)
    {
      auto token = new Token(TokenType::Character, string(1, ch));
      list.push_back(token);
    }

    return list;
  }
};

#endif