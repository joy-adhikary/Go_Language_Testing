package Http_Unit_testing

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

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

	//get the v query value
	text := r.FormValue("v")

	//if v is empty
	if text == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}

	query, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number: "+text, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, query*2)
}
