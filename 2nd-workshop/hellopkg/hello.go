package hellopkg

import (
	"fmt"
)

func Hello(firstName string) string {
	return fmt.Sprintf("Hello from hellopkg: %s", firstName)
}

func sayHello(firstName string) string {
	return fmt.Sprintf("Say hello %s", firstName)
}
