#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void
get_flag()
{
        char c;
        FILE * f = fopen("flag.txt", "r");
        if(!f)
                exit(1);
        while((c = fgetc(f)) != EOF)
                putchar(c);
        fclose(f);
        exit(0);
}

int
main()
{
	char hackme[16];
	char buff[256];

	printf("Hello everyone, say your name: ");
	gets(buff);
	if(strcmp("CyberCastors", hackme) == 0)
		get_flag();
	printf("I caught an old tire! I sure am tired of that!\n");
	return 0;
}

