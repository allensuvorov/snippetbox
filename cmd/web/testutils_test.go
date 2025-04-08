package main

import (
	"log/slog"
	"net/http/httptest"
	"testing"
)

func newTestApplication(t *testing.T) *application {
	return &application{
		logger: slog.New(slog.DiscardHandler),
	}

	type testServer struct {
		*httptest.Server
	}
}
