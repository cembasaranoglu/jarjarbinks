package main

import (
	"jarjarbinks/pkg/api"
	"jarjarbinks/pkg/application"
	"jarjarbinks/pkg/domain"
	"jarjarbinks/pkg/domain/commands"
	"jarjarbinks/pkg/domain/queries"
	"jarjarbinks/pkg/infrastructure/assembler"
	http_server "jarjarbinks/pkg/infrastructure/http"
	"jarjarbinks/pkg/infrastructure/http/options"
	lg "jarjarbinks/pkg/infrastructure/logging"
	options2 "jarjarbinks/pkg/infrastructure/logging/options"
	"jarjarbinks/pkg/infrastructure/mediator"
	_ "jarjarbinks/pkg/infrastructure/policies/fifo"
	"jarjarbinks/pkg/infrastructure/ports/http"
)

// @BasePath /
func main() {
	// *** Logging Implementations *** ///
	const serviceName = "cache_store_service"
	logger, _ := lg.New(&options2.LoggerOptions{
		DefaultParameters: map[string]interface{}{
			"service.name": serviceName,
		},
		Development: false,
	})
	// *** Cache Store *** ///
	cacheStore := domain.FIFO.New(0)

	// *** Assembler *** ///
	cacheAssembler := assembler.NewCacheAssembler()

	// *** HTTP Port Initialization *** ///
	httpPort := http.NewHttpPort()
	/// *** Mediator Initialization *** ///
	mediator, _ := mediator.New().
		RegisterHandler(&queries.FindCacheEntryByKeyQuery{}, application.NewFindCacheEntryByKeyQueryHandler(logger, cacheStore, cacheAssembler)).
		RegisterHandler(&commands.CreateCacheEntryCommand{}, application.NewCreateCacheEntryCommandHandler(logger, cacheStore, cacheAssembler)).
		Build()
	// *** HTTP Server *** ///
	httpServer := http_server.New(serviceName).
		WithLogger(logger).
		UseLogger().
		UseCompress().
		UseRecover().
		UseRequestID().
		UseSwagger().
		UseHealthCheck("/health", "OK!").
		WithController(api.NewCacheController(mediator, httpPort))
	defer httpServer.Shutdown()
	httpServer.Start(options.HttpServerStartOption{Port: 8080, GracefullyShutdown: false})
}

