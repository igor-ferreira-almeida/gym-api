//go:build wireinject
// +build wireinject

package exercise

import "github.com/google/wire"

func HandlerInject() (Handler, error) {
	panic(wire.Build(
		providerService,
		NewHandler,

		wire.Bind(new(service), new(serviceImpl)),
	))
}
