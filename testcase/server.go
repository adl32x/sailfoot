package testcase

import (
	"net/http"
	"fmt"
)

type Response struct {

}

func (t *Testcase) handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GET: ", r.URL.Query())

	keyword := r.URL.Query().Get("keyword")

	if keyword != "" {
		fmt.Println(keyword)
	}

	args := r.URL.Query()["arg"]

	if args != nil {
		fmt.Println(args)
	}

	commands := [][]string{[]string{keyword}}
	t.Command.Commands = commands


	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": \"ok\"}"))
}

func (t *Testcase) Listen(port int) {
	http.HandleFunc("/", t.handler)
	http.ListenAndServe(":"+fmt.Sprint(port), nil)
}