package routebuilder

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouteBuilder_Get(t *testing.T) {
	builder := NewBuilder()
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PASSED"))
	})
	builder.Get("/test/path", testHandler)
	router := builder.Build()

	request, _ := http.NewRequest("GET", "/test/path", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	expected := "PASSED"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body got %v, but expected %v",
			recorder.Body.String(), expected)
	}
}
