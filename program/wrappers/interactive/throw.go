package interactive

import "fmt"

func Throw(tag string, message string, err error, params map[string]any) {
	fmt.Printf("[%s] %s %s %s\n", tag, message, err, params)
}
