#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int
main()
{
	char hackme[256];
	char buff[256];
        char c;
        FILE * f = fopen("flag.txt", "r");
        if(!f)
                exit(1);
	fscanf(f, "%s", hackme);
        fclose(f);

	printf("Hello everyone, this is babyfmt! say something: ");
	fgets(buff, 255, stdin);
	printf(buff);
	return 0;
}

