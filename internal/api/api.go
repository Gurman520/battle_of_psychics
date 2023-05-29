package api

import (
	"strconv"

	"battle_of_psychics/internal/app"
	"battle_of_psychics/internal/def"
	"battle_of_psychics/openapi/models"
	"battle_of_psychics/openapi/restapi"
	"battle_of_psychics/openapi/restapi/operations"
	"battle_of_psychics/openapi/restapi/operations/standard"

	"github.com/go-openapi/loads"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"go.uber.org/zap"
)

type service struct {
	l             *zap.SugaredLogger
	battleServers app.BattleServerService
}

func NewServer(cfg *def.Config, l *zap.SugaredLogger, battleSvc app.BattleServerService) (*restapi.Server, error) {
	svc := &service{
		l:             l,
		battleServers: battleSvc,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load embedded swagger spec")
	}
	if cfg.BasePath == "" {
		cfg.BasePath = swaggerSpec.BasePath()
	}
	swaggerSpec.Spec().BasePath = cfg.BasePath

	api := operations.NewBattleAPI(swaggerSpec)

	api.KeyAuth = ValidateHeader

	api.Logger = structlog.New(structlog.KeyUnit, "swagger").Printf

	api.StandardHealthCheckHandler = standard.HealthCheckHandlerFunc(healthCheck)

	svc.initBattleServerHandlers(api)

	server := restapi.NewServer(api)
	server.Host = string(cfg.Host)
	port, err := strconv.ParseInt(cfg.HTTPPort, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse port")
	}
	server.Port = int(port)

	return server, nil
}

func healthCheck(params standard.HealthCheckParams, principal *models.Principal) standard.HealthCheckResponder {
	return standard.NewHealthCheckOK().WithPayload(&standard.HealthCheckOKBody{Ok: true})
}
