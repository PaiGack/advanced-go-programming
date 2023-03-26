#include <iostream>

extern "C"
{
#include "hello.h"
}

void say_hello(const char *s)
{
    std::cout << s;
}