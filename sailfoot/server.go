package sailfoot

import (
	"fmt"
	"net/http"
)

type Response struct {
}

func (c *Case) handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GET: ", r.URL.Query())

	keyword := r.URL.Query().Get("keyword")

	if keyword != "" {
		fmt.Println(keyword)
	}

	args := r.URL.Query()["arg"]

	if args != nil {
		fmt.Println(args)
	}

	command := append([]string{keyword}, args...)
	commands := [][]string{command}
	c.RootKeyword.Commands = commands

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": \"ok\"}"))
}

func (c *Case) Listen(port int) {
	http.HandleFunc("/", c.handler)
	http.ListenAndServe(":"+fmt.Sprint(port), nil)
}
