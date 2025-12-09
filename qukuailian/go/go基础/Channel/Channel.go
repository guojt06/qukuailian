package main

var moneyChan = make(chan int)

func pay(amount int) {
	moneyChan <- amount
}
