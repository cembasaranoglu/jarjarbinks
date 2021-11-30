package http_server

import (
	"fmt"
	sentryfiber "github.com/aldy505/sentry-fiber"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/nobuyo/nrfiber"
	controller2 "jarjarbinks/pkg/infrastructure/http/controller"
	"jarjarbinks/pkg/infrastructure/http/controller/interfaces"
	"jarjarbinks/pkg/infrastructure/http/middlewares"
	"jarjarbinks/pkg/infrastructure/http/options"
	"jarjarbinks/pkg/infrastructure/json"
	lg "jarjarbinks/pkg/infrastructure/logging/interfaces"
	"time"
)

type HttpServer struct {
	serverOptions *options.HttpServerOption
	http          *fiber.App
	controller    interfaces.ControllerBase
	logger        lg.Logger
}

const idleTimeout = 5 * time.Second
func New(name string) *HttpServer {
	marshaller := json.New()
	server := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
		AppName: name,
		JSONEncoder: func(v interface{}) ([]byte, error) {
			return marshaller.Marshall(v)
		},
		JSONDecoder: func(data []byte, v interface{}) error {
			return marshaller.Unmarshall(data, v)
		},
	})
	return &HttpServer{
		http:          server,
		serverOptions: &options.HttpServerOption{Name: name},
	}
}

func (k *HttpServer) WithLogger(logger lg.Logger) *HttpServer {
	if logger == nil {
		panic("logger must be specified")
	}
	if k.logger != nil {
		panic("you can not register logger more than once")
	}
	k.logger = logger
	return k
}

func (k *HttpServer) UseCompress() *HttpServer{
	k.http.Use(compress.New())
	return k
}

func (k *HttpServer) UseProfiler() *HttpServer{
	k.http.Use(pprof.New())
	return k
}

func (k *HttpServer) UseRequestID() *HttpServer{
	k.http.Use(requestid.New())
	return k
}

func (k *HttpServer) UseLogger() *HttpServer{
	k.http.Use(middlewares.Logging(k.logger))
	return k
}

func (k *HttpServer) UseRecover() *HttpServer{
	k.http.Use(recover.New())
	return k
}

func (k *HttpServer) UseHealthCheck(path string, response string) *HttpServer{
	k.http.Get(path, func(ctx *fiber.Ctx) error {
		return ctx.JSON(response)
	})
	return k
}

func (k *HttpServer) UseSwagger() *HttpServer{
	k.http.Get("/swagger/*", swagger.Handler) // default
	return k
}
func (k *HttpServer) UseNewRelic(application *newrelic.Application) *HttpServer{
	k.http.Use(nrfiber.New(nrfiber.Config{
		NewRelicApp: application,
	}))
	return k
}

func (k *HttpServer) UseMiddlewares(args ...interface{}) *HttpServer{
	k.http.Use(args...)
	return k
}


func (k *HttpServer) UseSentry() *HttpServer{
	k.http.Use(sentryfiber.New(sentryfiber.Options{
		Timeout: time.Second * 2,
	}))
	return k
}
func (k *HttpServer) WithControllers(controllers []interfaces.ControllerBase) *HttpServer {
	for _, controller := range controllers{
		k.WithController(controller)
	}
	return k
}
func (k *HttpServer) WithController(controller interfaces.ControllerBase) *HttpServer {
	if controller == nil {
		panic("controller must be specified")
	}
	k.controller = controller

	var (
		httpController controller2.HttpController
	)
	switch svc := controller.(type) {
	case controller2.HttpController:
		httpController = svc
	default:
		panic("controller must implement to Service interface")
	}

	if httpController != nil {
		for path, endpointMethod := range *httpController.Endpoints() {
			for method, endpointHandler := range endpointMethod {
				var endpoint string
				if len(httpController.Prefix()) > 0 {
					endpoint += fmt.Sprintf("/%s", httpController.Prefix())
				}
				if len(httpController.Name()) > 0 {
					endpoint += fmt.Sprintf("/%s", httpController.Name())
				}
				if len(httpController.Version()) > 0 {
					endpoint += fmt.Sprintf("/%s", httpController.Version())
				}
				endpoint += fmt.Sprintf("%s", path)
				httpController.Version()
				k.http.Add(method, endpoint, endpointHandler)
			}
		}
	}
	return k
}


func (k *HttpServer) Start(option options.HttpServerStartOption) *HttpServer {
	err := k.http.Listen(fmt.Sprintf(":%d", option.Port))
	if err != nil{
		panic(err)
	}
	return k
}
func (k *HttpServer) Shutdown() *HttpServer {
	k.http.Shutdown()
	return k
}