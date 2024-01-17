package service

import "net/http"

func (s *Service) auth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Bearer token from authorization header.
		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" {
			http.Error(w, "missing authorization", http.StatusUnauthorized)
			return
		}

		// Strip the Bearer prefix.
		bearerToken = bearerToken[len("Bearer "):]

		// Verify the token.
		_, err := s.verifier.Verify(r.Context(), bearerToken)
		if err != nil {
			http.Error(w, "failed to verify", http.StatusInternalServerError)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func (s *Service) authCallback(w http.ResponseWriter, r *http.Request) {
	// Exchange the code for a token.
	oauth2Token, err := s.oauth2.Exchange(r.Context(), r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Verify the token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "missing id_token", http.StatusInternalServerError)
		return
	}

	_, err = s.verifier.Verify(r.Context(), rawIDToken)
	if err != nil {
		http.Error(w, "failed to verify", http.StatusInternalServerError)
		return
	}

	// Redirect to the main page with the token in the query string.
	http.Redirect(w, r, "/?token="+rawIDToken, http.StatusFound)
}

func (s *Service) authLogin(w http.ResponseWriter, r *http.Request) {
	// Redirect to the OAuth2 provider for authorization.
	http.Redirect(w, r, s.oauth2.AuthCodeURL("state"), http.StatusFound)
}
