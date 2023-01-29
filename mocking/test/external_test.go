package test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mocking/external"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

var (
	server *httptest.Server
	ext    external.External
)

func mockFetchDataEndpoint(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]

	sc := http.StatusOK
	m := make(map[string]interface{})

	if !ok || len(ids[0]) == 0 {
		sc = http.StatusBadRequest
	} else {
		m["id"] = "mock"
		m["name"] = "mock"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(sc)
	json.NewEncoder(w).Encode(m)
}

func TestMain(m *testing.M) {
	fmt.Println("mocking server")
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch strings.TrimSpace(r.URL.Path) {
		case "/":
			mockFetchDataEndpoint(w, r)
		default:
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	}))

	fmt.Println("mocking external")
	ext = external.New(server.URL, http.DefaultClient, time.Second)

	fmt.Println("run tests")
	m.Run()
}

func fatal(t *testing.T, want, got interface{}) {
	t.Helper()
	t.Fatalf(`want: %v, got: %v`, want, got)
}

func Test_External_FetchData(t *testing.T) {
	tt := []struct {
		name     string
		id       string
		wantData *external.Data
		wantErr  error
	}{
		{
			name:     "response not ok",
			id:       "",
			wantData: nil,
			wantErr:  external.ErrResponseNotOK,
		},
		{
			name: "data found",
			id:   "mock",
			wantData: &external.Data{
				ID:   "mock",
				Name: "mock",
			},
			wantErr: nil,
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			gotData, gotErr := ext.FetchData(context.Background(), tc.id)

			if !errors.Is(gotErr, tc.wantErr) {
				fatal(t, tc.wantErr, gotErr)
			}

			if !reflect.DeepEqual(gotData, tc.wantData) {
				fatal(t, tc.wantData, gotData)
			}
		})
	}
}
