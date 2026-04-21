package main

import "fmt"

func square(i int) int {
	// if i > 9989 && i < 9999 {
	// 	return i
	// }
	return i * i
}

func dayOfWeek(i int) string {
	switch i {
	case 1:
		return "شنبه"
	case 2:
		return "یکشنبه"
	case 3:
		return "دوشنبه"
	case 4:
		return "سه شنبه"
	case 5:
		return "چهارشنبه"
	case 6:
		return "پنجشنبه"
	case 7:
		return "جمعه"
	default:
		return ""
	}
}

func main() {

	fmt.Println(square(1))
	fmt.Println(square(2))
	fmt.Println(square(3))

	// implement for loop with label and goto keyword
	// 	i := 0

	// start:
	// 	fmt.Println("i", i)
	// 	i++
	// 	if i < 10 {
	// 		goto start
	// 	}
	// 	fmt.Println("Finish")

	// 	sum := 10

	// 	if sum > 0 {
	// 		sum++

	// 		if sum > 10 {
	// 			goto label
	// 		}

	// 		fmt.Println("sum", sum)
	// 	}

	// label:
	// 	fmt.Println("Here")
}
