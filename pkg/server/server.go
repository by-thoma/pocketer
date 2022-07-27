package server

import (
	"net/http"

	"github.com/zhashkevych/go-pocket-sdk"
)

type AuthorizationServer struct {
	server       *http.Server
	pocketClient *pocket.Client
}
