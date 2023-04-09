package main

// #include <stdio.h>
// void print() {
// 	printf("Hello, wowo!\n");
// }
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	C.print()
}
