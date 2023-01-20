package main

import (
	"fmt"
)

func sayHello(name string) string {
	if name == "" {
		return "Hello Anonymous"
	}

	return fmt.Sprintf("Hello %s", name)
}

func sayGoodBye(name string) string {
	if name == "" {
		return "Bye Bye Anonymous!"
	}

	return fmt.Sprintf("Bye Bye %s", name)
}
