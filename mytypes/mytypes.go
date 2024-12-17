package mytypes

type MyInt int

func (i MyInt) Twice() MyInt {
	return i * 2
}
