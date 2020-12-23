package pkg

type mystruct struct {
	A int
	B int
}

// var decl in pkg
var (
	p []*mystruct // collection of ptr to A
)

//
func Foo() []*mystruct {
	return p
}

func Bar(c mystruct) (mystruct, error) {
	p = append(p, &c)
	return c, nil
}

//
// hang behavior to mystruct
//
func (c mystruct) MethodName() string {
	return "hang behavior to pkg.C struct"
}

//
// no cstor in Go
// cstor func
// avoid unecessary copy
//
func NewC() *mystruct {
	// for struct it is permissive
	// it is a local variable
	// take the address of it
	// can not take the address of numeric type, &42
	return &mystruct{ // if you're comming from C++
		// little bit panic
		A: 0,
		B: 1,
	}
}
