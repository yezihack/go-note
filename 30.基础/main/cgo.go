package main

/*
#include <stdio.h>
#include <stdlib.h>
void hello() {
printf("Hello, World!\n");
}
*/
import "C"
func main() {
	C.hello()
}
