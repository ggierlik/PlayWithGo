package funcs

import "fmt"

//deklaracja funkcji, która zwraca dwie wartości
//nazwa funkcji zaczyna się wielką literą jeżeli jest publiczna
func DivMod(a, b int) (int, int) {
	return a / b, a % b
}

//variadic function -- funkcja ze zmienną liczbą parametrów
func Sum(nums ...int) {
	fmt.Print(nums, " = ")
	total := 0
	for _, num := range nums { //range zwraca index i wartość elementu -- tutaj index jest ignorowany
		total += num
	}
	fmt.Println(total)
}

/*
func main() {
	var a, b int
	a, b = divMod(13, 5)             //zwraca wiele wartości
	fmt.Println("13 div 5 = ", a, " reszty ", b) //Print obsługuje wiele parametrów

	_, x := divMod(13, 4) //ignorowanie wyniku _
	fmt.Println("13 mod 4 = ", x)
}
*/
