package foo

import "github.com/netbrain/importbug/bar"

func Foo() string {
	return "Foo"
}

func FooBar() string {
	return Foo() + bar.Bar()
}
