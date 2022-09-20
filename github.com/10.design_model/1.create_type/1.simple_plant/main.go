package main

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	}
	if t == 2 {
		return &helloAPI{}
	}

	return nil
}

type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	api := NewAPI(1)
	s := api.Say("tom")
	if s == "Hi, tom" {
		fmt.Println(true)
	}
}
