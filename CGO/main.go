package main

/*
#include <stdio.h>
// Khai báo hàm trong ngôn ngữ C
static void SayHello(const char* s) {
    puts(s);
}
*/
import "C"

func main() {
	C.SayHello(C.CString("Hello World\n"))
}
