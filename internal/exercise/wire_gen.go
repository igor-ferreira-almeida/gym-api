// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package exercise

// Injectors from wire.go:

func HandlerInject() (Handler, error) {
	exerciseRedisCache := newRedisCache()
	dynamoDB := newDynamoDB()
	exerciseServiceImpl := newService(exerciseRedisCache, dynamoDB)
	handler := NewHandler(exerciseServiceImpl)
	return handler, nil
}
