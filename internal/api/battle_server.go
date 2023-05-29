package api

import (
	"battle_of_psychics/internal/app"
	"battle_of_psychics/openapi/models"
	"battle_of_psychics/openapi/restapi/operations"
	"battle_of_psychics/openapi/restapi/operations/server"
	"fmt"
)

func (svc *service) initBattleServerHandlers(api *operations.BattleAPI) {
	api.ServerGetSessionHandler = server.GetSessionHandlerFunc(svc.session)
	api.ServerConceiveHandler = server.ConceiveHandlerFunc(svc.conceive)
	api.ServerResultHandler = server.ResultHandlerFunc(svc.result)
}

func (svc *service) session(params server.GetSessionParams, principal *models.Principal) server.GetSessionResponder {
	token, err := svc.battleServers.CreateSession(params.HTTPRequest.Context())
	switch err {
	case nil:
		svc.l.Info("Session create is success")
		return server.NewGetSessionOK().WithPayload(&models.ResponseSession{Jwt: token})
	case app.ErrIncorrectApiKey:
		svc.l.Info("Session create is fail. Err: ", app.ErrIncorrectApiKey)
		return server.NewGetSessionBadRequest()
	default:
		return server.NewGetSessionInternalServerError()
	}
}

func (svc *service) conceive(params server.ConceiveParams, principal *models.Principal) server.ConceiveResponder {
	battle, err := svc.battleServers.GetHypothesis(params.HTTPRequest.Context(), string(*principal))

	switch err {
	case nil:
		svc.l.Info("Conceive success")
		hypo := ConceiveToRest(battle)
		return server.NewConceiveOK().WithPayload(hypo)
	case app.ErrIncorrectApiKey:
		svc.l.Info("Conceive is fail. Err: ", app.ErrIncorrectApiKey)
		return server.NewConceiveUnauthorized()
	case app.ErrNotFoundSession:
		svc.l.Info("Not found session")
		return server.NewConceiveNotFound()
	default:
		return server.NewConceiveInternalServerError()
	}
}

func (svc *service) result(params server.ResultParams, principal *models.Principal) server.ResultResponder {
	battle, err := svc.battleServers.GetReliability(params.HTTPRequest.Context(), int(params.Body.Number), string(*principal))

	switch err {
	case nil:
		svc.l.Info("Result success")
		fmt.Println(battle)
		res := ResultToRest(battle)
		fmt.Println(res)
		return server.NewResultOK().WithPayload(res)
	case app.ErrIncorrectApiKey:
		svc.l.Info("Result is fail. Err: ", app.ErrIncorrectApiKey)
		return server.NewResultUnauthorized()
	case app.ErrNotFoundSession:
		svc.l.Info("Not found session")
		return server.NewResultNotFound()
	case app.ErrCallBan:
		svc.l.Info(app.ErrCallBan)
		return server.NewResultNotAcceptable()
	default:
		return server.NewResultInternalServerError()
	}
}
