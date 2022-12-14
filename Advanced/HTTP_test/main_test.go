package HTTP

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

//type table struct {
//	numberOfTest string
//	number       []int
//	sum          int
//}

//func TestInts(t *testing.T) {
//
//	//c := Ints(1, 2, 3, 4)
//	//if c != 10 {
//	//	t.Errorf("expected but gain %v", c)
//	//}
//
//	test := []table{
//		{"1", []int{1, 2, 3, 4}, 10},
//		{"2", []int{1, 2, 3}, 6},
//		{"3", []int{}, 0},
//		{"4", []int{1100, 500, 500, 500}, 2600},
//	}
//
//	for _, tc := range test {
//		t.Run(tc.numberOfTest, func(t *testing.T) { // this is the syntax of subtest //  go test -v -run Ints/2 using this command we can run only the 2nd case
//			c := Ints(tc.number...)
//			if c != tc.sum {
//				t.Errorf("expected %v", c)
//			}
//		})
//	}
//
//}

func TestDoubleHandler(t *testing.T) {
	tt := []struct {
		name   string
		value  string
		double int
		err    string
	}{
		{name: "double of two", value: "2", double: 4, err: ""},
		{name: "missing value", value: "", err: "missing value"},
		{name: "not a number", value: "x", err: "not a number: x"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:8080/double?v="+tc.value, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()
			joy(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tc.err != "" {
				// do something
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status Bad Request; got %v", res.StatusCode)
				}
				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %q; got %q", tc.err, msg)
				}
				return
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", res.Status)
			}

			d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
			if err != nil {
				t.Fatalf("expected an integer; got %s", b)
			}

			if d != tc.double {
				t.Fatalf("expected double to be %v; got %v", tc.double, d)
			}
		})
	}
}

//func TestRouting(t *testing.T) {
//	srv := httptest.NewServer(handler())
//	defer srv.Close()
//
//	res, err := http.Get(fmt.Sprintf("%s/double?v=2", srv.URL))
//	if err != nil {
//		t.Fatalf("could not send GET request: %v", err)
//	}
//	defer res.Body.Close()
//
//	if res.StatusCode != http.StatusOK {
//		t.Errorf("expected status OK; got %v", res.Status)
//	}
//
//	b, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		t.Fatalf("could not read response: %v", err)
//	}
//
//	d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
//	if err != nil {
//		t.Fatalf("expected an integer; got %s", b)
//	}
//	if d != 4 {
//		t.Fatalf("expected double to be 4; got %v", d)
//	}
//}
