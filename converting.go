package main

import (
	"bytes"
	"fmt"
	"strings"
)

func Composed(str string) string {
	i := strings.LastIndex(str, "/")
	str = str[i+1:]
	if dot := strings.LastIndex(str, "."); dot >= 0 {
		str = str[:dot]
	}
	return str
}
func IntToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
func main() {

	fmt.Println(IntToString([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(Composed("a/b/c.d.go"))
}
