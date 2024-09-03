package mytypes

import "strings"

// Twice multiplies its receiver by 2 and returns the result
type MyInt int
type MyString string
type Mybuilder struct {
	Contents strings.Builder
}

type StringUppercaser struct {
	Contents strings.Builder
}

func (i MyInt) Twice() MyInt {
	return 2 * i
}

func (s MyString) Len() int {
	return len(s)
}

func (mb Mybuilder) Hello() string {
	return "Hello, Gophers!"
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}

func (input *MyInt) Double() {
	*input *= 2
}
