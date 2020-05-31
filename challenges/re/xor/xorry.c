#include <stdio.h>
#include <string.h>

int encrypt(char *s1);
void decrypt(char *s1);

char s2[] = {0x67, 0x68, 0x7d, 0x77, 0x5f, 0x7b, 0x61, 0x50, 0x44, 0x53, 0x6d, 0x6b, 0x24, 0x63, 0x68, 0x26, 0x72, 0x2b, 0x41, 0x68, 0x2d, 0x26, 0x46, 0x7c, 0x14, 0x7a, 0x11, 0x50, 0x15, 0x10, 0x1d, 0x52, 0x1e};

int encrypt(char *s1)
{
    int len = strlen(s1);
    int xor;
    for (int i = 0; i < len; i++)
    {
        xor = i + 0xA;
        s1[i] = (s1[i] ^ xor);
        s1[i] -= 2;
    }

    int result = strcmp(s1, s2);

    return result != 0;
}

int main(int argc, char const *argv[])
{
    char buff[44];
    printf("Enter flag: ");
    fgets(buff, 44, stdin);

    int val = encrypt(buff);

    if (val != 0) {
        printf("Wrong flag!\n");
    } else {
        printf("Correct!\n");
    }

    return 0;
}
