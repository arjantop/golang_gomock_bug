package foo

//go:generate mockgen -source=foo.go -destination=foo_mock.go

type Arg struct {}

type Foo interface {
	Bar(a Arg) error
}
