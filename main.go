package main

import "fmt"

// Hello returns a greeting, the name is optional
func Hello(name ...string) string{
	if len(name) > 0 {
		return fmt.Sprintf("Hello %s!", name[0])
	}
	return "Hello World!"
}

func main (){
	fmt.Println(Hello())
}
