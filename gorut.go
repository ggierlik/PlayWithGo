//based on http://golang.org.pl/getting_started/02_my_tutorial.html#krotki-tutorial
package main

import "fmt"

const N = 0xfffffffff

func decr(n uint64, c chan int, id int) {
    for n > 0 {
        n--
    }
    c <- id //koniec wykonywania goroutyny
}

func main() {
    c := make(chan int, 4)
    go decr(N/2, c, 1) // wszystkie
    go decr(N/3, c, 2) // goroutyny piszą
    go decr(N, c, 3) // do jednego
    go decr(N/4, c, 4) // kanału

	fmt.Println(<-c);
	fmt.Println(<-c);
	fmt.Println(<-c);
	fmt.Println(<-c);
}

