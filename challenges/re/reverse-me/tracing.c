#include <stdio.h>
#include <string.h>

void rot(char *input);
void add(char *input);
void check(char *input, char *flag);
void remove_newline(char *string);

char password[30];

void remove_newline(char *string)
{
    // Replace newline with null character
    for (int i = 0; i < strlen(string); i++)
    {
        if (string[i] == '\n')
        {
            string[i] = '\0';
        }
    }
    
}

void rot(char *input) 
{
    int len = strlen(input);
    for (int i = 0; i < len; i++)
    {
        if (input[i] >= 'a' && input[i] <= 'z') 
        {
            input[i] = (((input[i] - 'a') + 0xA) % 26) + 'a'; 
        }
    }
}

void add(char *input)
{
    int len = strlen(input);
    for (int i = 0; i < len; i++)
    {
        input[i] = input[i] + 0x2;
    }
}

void check(char *input, char *flag)
{   
    int result = strcmp(input, flag);

    if (result == 0) {
        printf("Correct!\n");
        printf("castorsCTF{%s}", password);
        fflush(stdout);
    } else {
        printf("Wrong!\n");
        fflush(stdout);
    }
}

int main(int argc, char const *argv[])
{
    FILE *fptr;
    fptr = fopen("flag.txt", "r");

    char encrypted[30];
    char usr_input[30];

    if (fptr != NULL) {
        fgets(password, 30, fptr);
        fclose(fptr);
    } else {
        printf("flag.txt not found.");
        fflush(stdout);
        return 1;
    }

    remove_newline(password);
    
    strcpy(encrypted, password);

    add(encrypted);
    rot(encrypted);
    
    printf("System Error...\nDumping memory...\n");

    for (int i = 0; i < strlen(encrypted); i++)
    {
        printf("%x ", encrypted[i]);
    }
    
    printf("\nEnter password: ");
    fflush(stdout);
    fgets(usr_input, 50, stdin);

    remove_newline(usr_input);

    add(usr_input);
    rot(usr_input);

    check(usr_input, encrypted);
    
    return 0;
}
