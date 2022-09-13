package HTTP

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//
//// Ints return sum of a list of integers
//func Ints(vs ...int) int {
//	return ints(vs)
//
//}
//func ints(vs []int) int {
//	if len(vs) == 0 {
//		return 0
//	}
//	return ints(vs[1:]) + vs[0]
//}

func main() {
	err := http.ListenAndServe(":8080", handler())
	if err != nil {
		log.Fatal(err)
	}
}

// a Handler responds to an HTTP request.
func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/joy", joy)
	return r
}

func joy(w http.ResponseWriter, r *http.Request) {

	text := r.FormValue("v")
	if text == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}

	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number: "+text, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, v*2)
}
