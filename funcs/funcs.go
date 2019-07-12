package funcs

import "fmt"

//własny unsinged int
type ExtraInt int //własny int

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

func (i ExtraInt) IsOdd() bool {
	_, x := DivMod(int(i), 2)
	return (x == 1)
}
