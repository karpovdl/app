package main

import (
	"encoding/json"
	"os"
	"runtime"
	"time"

	"net/http"
	_ "net/http/pprof"

	cli "github.com/urfave/cli/v2"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	initApp()
}

func initApp() bool {
	app := cli.App{
		Name:     "app",
		Version:  "v1.0.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Denis Karpov",
				Email: "KarpovDL@hotmail.com",
			},
		},
		Copyright: "(c) 2020 Denis Karpov",
		HelpName:  "run",
		Usage:     "application for running processes",

		EnableBashCompletion:   true,
		HideHelp:               false,
		HideVersion:            false,
		UseShortOptionHandling: true,

		Commands: []*cli.Command{
			appCommand(),
		},

		Action: func(c *cli.Context) error {
			cli.ShowVersion(c)

			return nil
		},
	}

	app.Run(os.Args)

	return true
}

func jsonResponse(w http.ResponseWriter, status int, msg string) {
	var (
		resp []byte
		err  error
	)

	if resp, err = json.Marshal(response{Message: msg, Status: status}); err != nil {
		resp, _ = json.Marshal(map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(resp)
}
