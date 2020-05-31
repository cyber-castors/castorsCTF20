#include "stdlib.h"
#include "unistd.h"
#include "stdio.h"
#include <sys/personality.h>
#include <string.h>

personality(ADDR_NO_RANDOMIZE);

int winnersLevel(int idNumber){
    if (idNumber == 386 || idNumber == 258) {
        printf("Wow! Please excuse me sir I had no idea...here are your chips\n");
        system("cat ./flag.txt");
	return 1;
    }
    else {
        printf("You guessed right but it seems your badge number isn't on our list.\n");
	return 0;
    }
}
char *gets(char *);

void start() {
    char buffer[64];
    gets(buffer);
}

int main(int argv, char **argc){

    char buffer[64];
    printf("Do you really think you can get to the winners table?\n");
    printf("I'll give you one shot at it, what floor is the table at: \n");

    start();

    puts("Yeah that's what I thougt, LOL.\n");


    return 0;
}
