package hpc_test

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/rpc/v2"
	"github.com/tj/go-hpc"
)

type Users struct {
	users []string
}

type ListInput struct {
	Prefix string `json:"prefix"`
}

type ListOutput struct {
	Names []string `json:"names"`
}

func (u *Users) List(r *http.Request, in *ListInput, out *ListOutput) error {
	out.Names = []string{}
	for _, name := range u.users {
		if strings.HasPrefix(name, in.Prefix) {
			out.Names = append(out.Names, name)
		}
	}
	return nil
}

func Example() {
	users := []string{"Tobi", "Loki", "Jane"}

	r := rpc.NewServer()
	r.RegisterCodec(hpc.NewCodec(), "application/json")
	r.RegisterService(&Users{users}, "")

	http.Handle("/", r)
	log.Fatalln(http.ListenAndServe(":3000", nil))
}
