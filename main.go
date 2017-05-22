package main

import (
	"fmt"
	"net/http"

	"github.com/williamhgough/hydra/defense/shields"
	"github.com/williamhgough/hydra/logger"
)

func main() {
	loggy := logger.GetInstance()
	loggy.Rec("Starting Hydra Web Service", nil)

	builder := shields.NewShieldBuilder()
	shield := builder.RaiseLeft().RaiseRight().Build()
	loggy.Rec("Sheild status:", *shield)

	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	loggy := logger.GetInstance()
	fmt.Fprintf(w, "Welcome to hydra bby")

	loggy.Rec("Request ::/", nil)
}
