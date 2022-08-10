package interactive

import "fmt"

func Throw(tag string, message string, err error, params map[string]any) {
	fmt.Printf("[%s] %s %s %s", tag, message, err, params)
}
