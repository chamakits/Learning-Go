package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main(){
	/*
	 var1,var2 := swap("a","b")
	 fmt.Println(var1,var2)
	 */
	fmt.Println(swap("a","b"))
}