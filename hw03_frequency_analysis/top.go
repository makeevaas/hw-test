// package hw03frequencyanalysis
package main

import "fmt"

func Top10(s string) []string {
	// Place your code here.
	var sLst []string
	sLst = append(sLst, s)
	return sLst
}
func main() {
	res := Top10("привет")
	fmt.Println(res)
}
