package main

import "fmt"

func requestPrice(url string) {
	fmt.Println(url)
}

func main() {
	requestPrice("https://www.avito.ru/vladimir/rasteniya/palma_2294526731")
}
