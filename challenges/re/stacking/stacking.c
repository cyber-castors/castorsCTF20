#include <stdio.h>

void func();

const char* fake = "Strings_will_not_work";

int main(int argc, char const *argv[])
{
    func();
    return 0;
}


void func()
{
    char flag[42] = {'c', 'a', 's', 't', 'o', 'r', 's', 'C', 'T', 'F', '{', 'w', '3', 'l', 'c', '0', 'm', '3', '_', '7', '0', '_', 'r', '3', 'v', '3', 'r', '5', '3', '_', '3', 'n', '6', '1', 'n', '3', '3', 'r', '1', 'n', '6', '}'};
    int P = 20;

    for (int i = 0; i < P; i++)
    {
        printf("=");
    }

    printf("\nWhere's the flag?\n");

    for (int i = 0; i < P; i++)
    {
        printf("=");
    }

    printf("\n");
}
