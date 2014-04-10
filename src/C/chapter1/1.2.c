#include "stdio.h"
#include "stdlib.h"
#include "string.h"

void reverse(char* str) {
    if (str == NULL) {
        return;
    }

    int count = 0;

    char* p = str;

    while (p != NULL && *p != '\0'){
        count++;
        p ++;
    }

    printf("%d\n", count);
    

    for(int i =0; i < (count >> 1); i++){
        printf("%d\n", count-1-i);
        char a = str[count-1-i];
        char b = str[i];
        str[count-1-i] = b;
        str[i] = a;
    }

}

int main(){

    char* value = (char *) malloc(sizeof(char*) * 10);
    const char* target = "0123456789";
    void * amount = memcpy(value, target, strlen(target));
    reverse(value);
    printf("%s\n", value);
}