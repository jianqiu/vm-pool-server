package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/jianqiu/vm-pool-server/models"
	"github.com/jianqiu/vm-pool-server/restapi/operations"
	"github.com/jianqiu/vm-pool-server/restapi/operations/pool"
	"github.com/jianqiu/vm-pool-server/restapi/operations/vms"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name vm-pool --spec ../docs/vm_pool_server_api.json

func configureFlags(api *operations.VMPoolServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.VMPoolServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.VmsGetVmsHandler = vms.GetVmsHandlerFunc(func(params vms.GetVmsParams) middleware.Responder {
		return middleware.NotImplemented("operation vms.GetVms has not yet been implemented")
	})
	api.PoolRequestVMHandler = pool.RequestVMHandlerFunc(func(params pool.RequestVMParams) middleware.Responder {

		requestVM := pool.NewRequestVMOK()
		vm := models.VM{
			CPU:         params.Body.CPU,
			Deployment:  "",
			Hostname:    "",
			Memory:      params.Body.Memory,
			PrivateIP:   "",
			PrivateVlan: params.Body.PrivateVlan,
			PublicVlan:  params.Body.PublicVlan,
			Status:      "",
			VMID:        0,
		}
		requestVM.SetPayload(&vm)
		return requestVM
	})
	api.PoolReturnVMHandler = pool.ReturnVMHandlerFunc(func(params pool.ReturnVMParams) middleware.Responder {
		returnVM := pool.NewReturnVMOK()

		//To-do: Get all properties through SL API

		vm := models.VM{
			CPU:         4,
			Deployment:  "",
			Hostname:    "test-hostname-1.softlayer.com",
			Memory:      32768,
			PrivateIP:   "10.0.0.1",
			PrivateVlan: 123456,
			PublicVlan:  123457,
			Status:      "free",
			VMID:        params.Body.VMID,
		}

		//To-do: Insert to DB
		vm.CPU = 4

		return returnVM

	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}