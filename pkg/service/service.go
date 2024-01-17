package service

import (
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/gorilla/mux"
	"github.com/mtaylor91/desktopctl/pkg/kubevirt/client/versioned"
	"golang.org/x/oauth2"
)

type Service struct {
	client    *versioned.Clientset
	namespace string
	oauth2    *oauth2.Config
	router    *mux.Router
	server    *http.Server
	verifier  *oidc.IDTokenVerifier
}

func New(
	addr, namespace, oidcClientID, oidcClientSecret, oidcRedirectURL string,
	oidcProvider *oidc.Provider,
	client *versioned.Clientset,
) *Service {
	s := &Service{client: client, namespace: namespace}

	s.router = router(s)
	s.server = &http.Server{Addr: addr, Handler: s.router}

	s.oauth2 = &oauth2.Config{
		ClientID:     oidcClientID,
		ClientSecret: oidcClientSecret,
		Endpoint:     oidcProvider.Endpoint(),
		RedirectURL:  oidcRedirectURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	s.verifier = oidcProvider.Verifier(&oidc.Config{ClientID: oidcClientID})

	return s
}
