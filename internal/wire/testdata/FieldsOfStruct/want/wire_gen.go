// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/zeriontech/google-wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectedMessage() string {
	s := provideS()
	string2 := s.Foo
	return string2
}
