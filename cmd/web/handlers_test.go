package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"testing"

	"github.com/allensuvorov/snippetbox.git/internal/assert"
)

func TestPing(t *testing.T) {
	app := &application{
		logger: slog.New(slog.DiscardHandler),
	}

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	body = bytes.TrimSpace(body)

	assert.Equal(t, string(body), "OK")
}
