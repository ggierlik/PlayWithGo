package main //pakiet uruchomieniowy nazywa się main

import "fmt"
import funcs "PlayWithGo/funcs"
import objs "PlayWithGo/objs"

func main() { //i mieć zdefiniowaną funkcję main()
	fmt.Printf("hello, world\n")

	var s string //deklaracja zmiennej
	s = "Hello, world 2\n"
	fmt.Printf(s)

	var s1 = "Hello, world 3" //deklaracja i inicjalizacja zmiennej
	fmt.Println(s1)

	s2 := "Hello, world 4" //deklaracja i inicjalizacja z inferencją typu
	fmt.Println(s2)

	//Funkcje
	var a, b int
	a, b = funcs.DivMod(13, 5)       //zwraca wiele wartości
	fmt.Println("13 div 5 = ", a, b) //Print obsługuje wiele parametrów

	_, x := funcs.DivMod(13, 4) //ignorowanie wyniku _
	fmt.Println("13 mod 4 = ", x)

	funcs.Sum(1, 2, 3, 10, 7, 8)

	ex := funcs.ExtraInt(3)
	fmt.Println("3.isOdd = ", ex.IsOdd())

	r := objs.Rect{W: 10, H: 6}
	fmt.Println("%d x %d = %d", r.W, r.H, r.Area())

}
