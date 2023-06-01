// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"battle_of_psychics/openapi/models"
	"battle_of_psychics/openapi/restapi/operations"
	"battle_of_psychics/openapi/restapi/operations/server"
	serverops "battle_of_psychics/openapi/restapi/operations/server"
	"battle_of_psychics/openapi/restapi/operations/standard"
)

//go:generate swagger generate server --target ..\..\openapi --name Battle --spec ..\swagger.yaml --principal models.Principal --exclude-main --strict-responders

func configureFlags(api *operations.BattleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BattleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.UrlformConsumer = runtime.DiscardConsumer

	api.HTMLProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("html producer has not yet been implemented")
	})
	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-token" header is set
	if api.KeyAuth == nil {
		api.KeyAuth = func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (key) x-token from header param [x-token] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.ServerConceiveHandler == nil {
		api.ServerConceiveHandler = serverops.ConceiveHandlerFunc(func(params serverops.ConceiveParams, principal *models.Principal) server.ConceiveResponder {
			return server.ConceiveNotImplemented()
		})
	}
	if api.ServerGameHandler == nil {
		api.ServerGameHandler = serverops.GameHandlerFunc(func(params serverops.GameParams, principal *models.Principal) server.GameResponder {
			return server.GameNotImplemented()
		})
	}
	if api.ServerGetSessionHandler == nil {
		api.ServerGetSessionHandler = serverops.GetSessionHandlerFunc(func(params serverops.GetSessionParams, principal *models.Principal) server.GetSessionResponder {
			return server.GetSessionNotImplemented()
		})
	}
	if api.StandardHealthCheckHandler == nil {
		api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(func(params standard.HealthCheckParams, principal *models.Principal) standard.HealthCheckResponder {
			return standard.HealthCheckNotImplemented()
		})
	}
	if api.ServerResultHandler == nil {
		api.ServerResultHandler = serverops.ResultHandlerFunc(func(params serverops.ResultParams, principal *models.Principal) server.ResultResponder {
			return server.ResultNotImplemented()
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
