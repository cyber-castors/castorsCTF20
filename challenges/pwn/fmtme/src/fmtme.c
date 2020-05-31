#include <stdio.h>
#include <stdlib.h>
int
main()
{
  char buffer[256];
  char buffer2[256];
  system("echo Tell me something cool :B");
  fgets(buffer, 256, stdin);
  printf("That's so cool that I need to repeat: ");
  printf(buffer);
  do {
    printf("Love what you said, can you tell me more? ");
    fgets(buffer2, 256, stdin);
    puts(buffer2);
  } while(1);
}
