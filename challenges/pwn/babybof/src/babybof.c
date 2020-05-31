#include <stdio.h>
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

void
main()
{
	char buff[256];
	printf("Welcome to the cybercastors Babybof\n");
	printf("Say your name: ");
	gets(buff);
}
