#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <cerrno>

#include "./tokenizer.cpp"

std::string get_file_contents(const char *filename)
{
  std::FILE *fp = std::fopen(filename, "rb");
  if (fp)
  {
    std::string contents;
    std::fseek(fp, 0, SEEK_END);
    contents.resize(std::ftell(fp));
    std::rewind(fp);
    std::fread(&contents[0], 1, contents.size(), fp);
    std::fclose(fp);
    return (contents);
  }
  throw(errno);
}

int main(int argc, char *argv[])
{
  if (argc <= 1)
  {
    std::cout << "Invalid arguments" << std::endl;
    return -1;
  }

  auto text = get_file_contents(argv[1]);
  auto list = Tokenizer::Tokenize(text);

  for (auto a : list)
  {
    std::cout << a->data << std::endl;

    delete a;
  }

  return 0;
}