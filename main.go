package main

import "fmt"

func main() {
	// fmt.Println("\nHello there! I'm a Go program!")
	// fmt.Print("I like the green color\n")
	// fmt.Println("Today is the 11 of August 2026")
	// fmt.Println("\n")

	// fmt.Print("Happy New 2027 Year!\n")
	// fmt.Println("     *   ")
	// fmt.Println("    ***   ")
	// fmt.Println("   *****   ")
	// fmt.Println("  *******   ")
	// fmt.Println(" *********   ")
	// fmt.Println("***********   ")
	// fmt.Println("    |||\n   ")

	// var a int
	// var b float32
	// var c float64
	// var d string
	// var e bool
	// fmt.Println(a, b, c, d, e)

	// count := 9
	// price := 12.5
	// total := float64(count) * price
	// fmt.Println("Count: ", count)
	// fmt.Println("Price: ", price)
	// fmt.Println("Total: ", total)

	// const pi = 3.14
	// fmt.Println("Pi constant: ", pi)

	// // checking seconds in the week
	// fmt.Println("Seconds in a week: ", 7*24*60*60)

	// fmt.Println("My friends kandy: ", 137/8)
	// fmt.Println("Only my kandy: ", 137%8)

	// f_temp := 98.6
	// fmt.Println("Celsium temperature: ", ((f_temp-32)*5)/9)

	var score int;
	score = 7;
	fmt.Println("My score is: ", score);
	
	name := "Ann";
	avg := 8.667;
	fmt.Printf("%s gets %.3f points\n", name, avg);
	fmt.Printf("Avg degree: %.2f\n", avg);
	fmt.Printf("Percentage: %d%%\n", int(avg*10));

	msg := fmt.Sprintf("Player %s, level %d", name, score);
	// msg := fmt.Sprintf("Игрок %s, уровень %d", "Аня", 7)

	fmt.Println(msg);
	fmt.Println(len(msg), "bytes long");
}
н