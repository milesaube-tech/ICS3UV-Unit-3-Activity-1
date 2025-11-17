/*
 * Author Miles Aube
 * Version 1.0.0
 * Date 2025-11-17
 * This is the This program to display the sum of three Integers; Linked to test case SUM_INT_1
 */

package main

import "fmt"

func main () {
		//set three Integers
	var Integer1 int = 10
	var Integer2 int = -20
	var Integer3 int = 85

	// set answer 
	var answer int = Integer1 + Integer2 + Integer3

	// display sum 
	fmt.Println(" The sum of ",Integer1,"&",Integer2, "&",Integer3,"is:",answer) 
}