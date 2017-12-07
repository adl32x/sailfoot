package sailfoot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Result bool
	Reason string `json:"reason,omitempty"`
}

func (c *Case) handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("GET: ", r.URL.Query())

	keyword := r.URL.Query().Get("keyword")

	if keyword == "" {
		result := &Response{false, "no-keyword"}
		json.NewEncoder(w).Encode(result)
		return
	}

	args := r.URL.Query()["arg"]

	if args != nil {
		fmt.Println(args)
	}

	command := append([]string{keyword}, args...)
	commands := [][]string{command}
	c.RootKeyword.Commands = commands

	c.RootKeyword.Run(c.Driver, c.KnownKeywords, args)

	result := &Response{c.RootKeyword.LastResult, ""}

	json.NewEncoder(w).Encode(result)
}

func (c *Case) Listen(port int) {
	http.HandleFunc("/", c.handler)
	http.ListenAndServe(":"+fmt.Sprint(port), nil)
}
