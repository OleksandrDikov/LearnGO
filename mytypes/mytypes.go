package mytypes

import (
	"strings"
)

type MyInt int

type MyString string

type MyBuilder struct {
	Contents strings.Builder
}

type StringUppercaser struct {
	Contents strings.Builder
}

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (i *MyInt) Double() {
	*i *= 2
}

func (s MyString) Len() int {
	return len(s)
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}
