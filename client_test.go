package statuscake_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/StatusCakeDev/statuscake-go"
)

const apiToken = "abcdefg123456789"

func TestClient(t *testing.T) {
	client := statuscake.NewAPIClient(apiToken)
	if client == nil {
		t.Error("expected a client, got nil")
	}
}

func createTestEndpoint(h http.HandlerFunc) (*httptest.Server, *statuscake.APIClient) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	}))

	// Create a StatusCake API client using the above server as the host.
	client := statuscake.NewAPIClient(apiToken)

	// Configure the client to use the test server.
	cfg := client.GetConfig()
	cfg.Scheme = "http"
	cfg.Host = strings.TrimPrefix(server.URL, "http://")
	cfg.HTTPClient = server.Client()

	return server, client
}

func mustParse(t *testing.T, r *http.Request) url.Values {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal("failed to read request body")
	}

	v, err := url.ParseQuery(string(body))
	if err != nil {
		t.Fatal("failed to parse request body")
	}

	return v
}

func mustRead(t *testing.T, f string) []byte {
	j, err := ioutil.ReadFile(f)
	if err != nil {
		t.Fatal("failed to read JSON file")
	}
	return j
}
