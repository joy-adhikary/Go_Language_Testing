package testtda

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type testdata struct {
	Name         string
	server       *httptest.Server
	informations *info
	err          error
}

func TestInfo(t *testing.T) {
	tests := []testdata{
		{
			Name: "joya",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"name":"joya", "age":23,"sex":"F"}`))
			})),
			informations: &info{
				Name: "joya",
				Age:  23,
				Sex:  "F",
			},
			err: nil,
		},
		{
			Name: "joy",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"name":"joy", "age":23,"sex":"M"}`))
			})),
			informations: &info{
				Name: "joy",
				Age:  23,
				Sex:  "M",
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			defer test.server.Close()
			resp, err := Getinfo(test.server.URL)

			if !reflect.DeepEqual(resp, test.informations) {
				t.Errorf("FAILED: expected %v, got %v\n", test.informations, resp)
			}
			if !errors.Is(err, test.err) {
				t.Errorf("Expected error FAILED: expected %v got %v\n", test.err, err)
			}
		})
	}

}
