package config_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/federicodosantos/hello-world-app/config"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	conf := &config.Config{
		AppPort: "8000",
	}

	app := config.NewApp(conf)

	app.SetupRoutes()

	req:= httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Fiber.Test(req)
	
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	got := string(body)
	want := "Hallo boskuhh semuanyaa....."
	assert.Equal(t, want, got)

}