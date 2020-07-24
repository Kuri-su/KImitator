package main

import (
	"reflect"

	"github.com/go-ozzo/ozzo-di"
)

func main() {
	c := di.NewContainer()

	// Step 1: register values, types, providers

	// register the value Foo{"abc"} as type Foo
	c.Register(Foo{"abc"})
	// register the value Foo{"abc"} as the interface Bar
	c.RegisterAs(Foo{"abc"}, di.InterfaceOf((*Bar)(nil)))
	// register the struct Foo as the interface Bar
	c.RegisterAs(reflect.TypeOf(Foo{}), di.InterfaceOf((*Bar)(nil)))
	// register a provider that returns a shared value as the interface Bar

	provider := func(container di.Container) reflect.Value {
		return reflect.ValueOf(&Foo{"xyz"})
	}
	c.RegisterProvider(provider, di.InterfaceOf((*Bar)(nil)), true)

	// Step 2: dependency injection

	// use `inject` tag to indicate which fields can be injected
	type Tee struct {
		Foo `inject`
		bar Bar `inject`
	}
	// inject the fields of a struct
	t := Tee{}
	c.Inject(&t)
	// inject function parameters
	c.Call(func(bar Bar, foo Foo) {})
	// build a value of the specified type with injection
	t2 := c.Make(reflect.TypeOf(&Tee{})).(*Tee)
}
