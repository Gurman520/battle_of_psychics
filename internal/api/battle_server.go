package api

import (
	"battle_of_psychics/internal/app"
	"battle_of_psychics/openapi/models"
	"battle_of_psychics/openapi/restapi/operations"
	"battle_of_psychics/openapi/restapi/operations/server"
	"io/ioutil"
)

func (svc *service) initBattleServerHandlers(api *operations.BattleAPI) {
	api.ServerGetSessionHandler = server.GetSessionHandlerFunc(svc.session)
	api.ServerConceiveHandler = server.ConceiveHandlerFunc(svc.conceive)
	api.ServerResultHandler = server.ResultHandlerFunc(svc.result)
	api.ServerGameHandler = server.GameHandlerFunc(svc.start)
}

// Handler for the request to open an html page
func (svc *service) start(params server.GameParams, principal *models.Principal) server.GameResponder {
	fileBytes, err := ioutil.ReadFile("Front/Home.html")
	if err != nil {
		svc.l.Error("Upload file is fail. Err: ", err)
	}
	f := string(fileBytes)
	svc.l.Info("Upload file to game is good")
	return server.NewGameOK().WithPayload(f)
}

// Session Request Handler
func (svc *service) session(params server.GetSessionParams, principal *models.Principal) server.GetSessionResponder {
	token, err := svc.battleServers.CreateSession(params.HTTPRequest.Context())
	switch err {
	case nil:
		svc.l.Info("Session create is success")
		return server.NewGetSessionOK().WithPayload(&models.ResponseSession{Jwt: token})
	case app.ErrIncorrectApiKey:
		svc.l.Error("Session create is fail. Err: ", app.ErrIncorrectApiKey)
		return server.NewGetSessionBadRequest()
	default:
		return server.NewGetSessionInternalServerError()
	}
}

// The handler of the request to receive predictions
func (svc *service) conceive(params server.ConceiveParams, principal *models.Principal) server.ConceiveResponder {
	battle, err := svc.battleServers.GetHypothesis(params.HTTPRequest.Context(), string(*principal))

	switch err {
	case nil:
		svc.l.Info("Conceive success")
		hypo := ConceiveToRest(battle)
		return server.NewConceiveOK().WithPayload(hypo)
	case app.ErrIncorrectApiKey:
		svc.l.Error("Conceive is fail. Err: ", app.ErrIncorrectApiKey)
		return server.NewConceiveUnauthorized()
	case app.ErrNotFoundSession:
		svc.l.Info("Not found session")
		return server.NewConceiveNotFound()
	default:
		return server.NewConceiveInternalServerError()
	}
}

// The handler of the request for obtaining the reliability of predictions
func (svc *service) result(params server.ResultParams, principal *models.Principal) server.ResultResponder {
	battle, err := svc.battleServers.GetReliability(params.HTTPRequest.Context(), int(params.Body.Number), string(*principal))

	switch err {
	case nil:
		svc.l.Info("Result success")
		res := ResultToRest(battle)
		return server.NewResultOK().WithPayload(res)
	case app.ErrIncorrectApiKey:
		svc.l.Error("Result is fail. Err: ", app.ErrIncorrectApiKey)
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
