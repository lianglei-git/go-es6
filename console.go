package es6

import "fmt"

type Console struct{}

func (c *Console) Log(msg string) {
	fmt.Println(msg)
}
