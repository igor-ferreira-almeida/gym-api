//go:build wireinject
// +build wireinject

package workout

import "github.com/google/wire"

func HandlerInject() (Handler, error) {
	panic(wire.Build(
		providerService,
		NewHandler,

		wire.Bind(new(service), new(serviceImpl)),
	))
}
