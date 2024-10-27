package main
import (
	"fmt"
)
func quadA(x int , y int){
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if (i == 0 || i == y-1) && (j == 0 || j == x-1) {
				fmt.Print("o")
			} else if i == 0 || i == y-1 {
				fmt.Print("-")
			} else if j == 0 || j == x-1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}


func quadB(x, y int) {
	// Eğer x veya y 0 ya da negatifse fonksiyon çalışmaz
	if x <= 0 || y <= 0 {
		return
	}
	// Üst kenar
	fmt.Print("/")
	for i := 0; i < x; i++ {
		fmt.Print("*")
	}
	fmt.Println("\\")
	// Yan kenarlar
	for i := 0; i < y-2; i++ {
		fmt.Print("*")
		for j := 0; j < x; j++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}
	// Alt kenar
	fmt.Print("\\")
	for i := 0; i < x; i++ {
		fmt.Print("*")
	}
	fmt.Println("/")
}


func quadC(x, y int) {
	
	if x <= 0 || y <= 0 {
		return
	}
	
	fmt.Print("A")
	for i := 0; i < x-2; i++ {
		fmt.Print("B")
	}
	if x > 1 {
		fmt.Println("A")
	} else {
		fmt.Println()
	}
	
	for i := 0; i < y-2; i++ {
		fmt.Print("B")
		for j := 0; j < x-2; j++ {
			fmt.Print(" ")
		}
		if x > 1 {
			fmt.Println("B")
		} else {
			fmt.Println()
		}
	}
	if y > 1 {
		fmt.Print("C")
		for i := 0; i < x-2; i++ {
			fmt.Print("B")
		}
		if x > 1 {
			fmt.Println("C")
		} else {
			fmt.Println()
		}
	}

func quadD(x, y int) {
	
	if x <= 0 || y <= 0 {
		return
	}
	
	fmt.Print("A")
	for i := 0; i < x-2; i++ {
		fmt.Print("B")
	}
	if x > 1 {
		fmt.Println("C")
	} else {
		fmt.Println()
	}
	
	for i := 0; i < y-2; i++ {
		fmt.Print("B")
		for j := 0; j < x-2; j++ {
			fmt.Print(" ")
		}
		if x > 1 {
			fmt.Println("B")
		} else {
			fmt.Println()
		}
	}
	
	if y > 1 {
		fmt.Print("A")
		for i := 0; i < x-2; i++ {
			fmt.Print("B")
		}
		if x > 1 {
			fmt.Println("C")
		} else {
			fmt.Println()
		}
	}
}
func quadE(x, y int) {
	
	if x <= 0 || y <= 0 {
		return
	}
	
	fmt.Print("A")
	for i := 0; i < x-2; i++ {
		fmt.Print("B")
	}
	if x > 1 {
		fmt.Println("C")
	} else {
		fmt.Println()
	}
	
	for i := 0; i < y-2; i++ {
		fmt.Print("B")
		for j := 0; j < x-2; j++ {
			fmt.Print(" ")
		}
		if x > 1 {
			fmt.Println("B")
		} else {
			fmt.Println()
		}
	}
	
	if y > 1 {
		fmt.Print("C")
		for i := 0; i < x-2; i++ {
			fmt.Print("B")
		}
		if x > 1 {
			fmt.Println("A")
		} else {
			fmt.Println()
		}
	}
}
