package routebuilder

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouteBuilder_Get(t *testing.T) {
	builder := NewBuilder()
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("102"))//w.Write([]byte(r.Context().Value("{id}").(string)))
	})
	builder.Get("/user/{id}", testHandler)
	router := builder.Build()

	request, _ := http.NewRequest("GET", "/user/102", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	expected := "102"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body got %v, but expected %v",
			recorder.Body.String(), expected)
	}
}
