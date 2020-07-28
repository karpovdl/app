package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	afero "github.com/spf13/afero"
	cli "github.com/urfave/cli/v2"
)

// appConfiguration type
type appConfiguration struct {
	Port     int    `json:"port"`
	AuthUser string `json:"auth_user"`
	AuthPass string `json:"auth_pass"`
}

// appRequest type
type appRequest struct {
	Name string `json:"name"`
}

var (
	appPort         int
	appAuthUser     string
	appAuthPassword string
)

const (
	app                      = "app"
	appAlias                 = "a"
	appPortFlagName          = "port"
	appPortFlagAlias         = "p"
	appAuthUserFlagName      = "auth_user"
	appAuthUserFlagAlias     = "au"
	appAuthPasswordFlagName  = "auth_pass"
	appAuthPasswordFlagAlias = "ap"
)

func appCommand() *cli.Command {
	return &cli.Command{
		Name:        "exec",
		Aliases:     []string{"e"},
		Usage:       "",
		Description: "",
		Flags: []cli.Flag{
			appPortFlag(),
			appAuthUserFlag(),
			appAuthPassFlag(),
		},
		Action: appAction,
	}
}

func appPortFlag() *cli.IntFlag {
	return &cli.IntFlag{
		Name:    appPortFlagName,
		Aliases: []string{appPortFlagAlias},
		Value:   9000,
		Usage:   "`port` to use listen server",
	}
}

func appAuthUserFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    appAuthUserFlagName,
		Aliases: []string{appAuthUserFlagAlias},
		Value:   "",
		Usage:   "auth `user`",
	}
}

func appAuthPassFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    appAuthPasswordFlagName,
		Aliases: []string{appAuthPasswordFlagAlias},
		Value:   "",
		Usage:   "auth `password`",
	}
}

func (conf appConfiguration) read(path string) (interface{}, error) {
	var (
		file    afero.File
		decoder *json.Decoder
		err     error
	)

	if ok, err := isExist(path); path == "" || !ok || err != nil {
		return appConfiguration{}, err
	}

	if file, err = appFs.Open(path); err != nil {
		return appConfiguration{}, err
	}

	defer file.Close()

	decoder = json.NewDecoder(file)
	conf = appConfiguration{}

	if err = decoder.Decode(&conf); err != nil {
		return appConfiguration{}, err
	}

	return conf, nil
}

func appAction(c *cli.Context) error {
	cli.ShowVersion(c)

	var bufout = bufio.NewWriter(os.Stdout)

	appPort = c.Int(appPortFlagName)
	appAuthUser = c.String(appAuthUserFlagName)
	appAuthPassword = c.String(appAuthPasswordFlagName)

	if conf, err := confRead(appConfiguration{}, "conf/exec.json"); err == nil {
		c := conf.(appConfiguration)

		if appPort == 0 {
			appPort = c.Port
		}

		if appAuthUser == empty {
			appAuthUser = c.AuthUser
		}

		if appAuthPassword == empty {
			appAuthPassword = c.AuthPass
		}
	}

	if appPort == 0 {
		appPort, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	}

	if appAuthUser == empty {
		appAuthUser = os.Getenv("APP_AUTH_BASIC_USER")
	}

	if appAuthPassword == empty {
		appAuthPassword = os.Getenv("APP_AUTH_BASIC_PASSWORD")
	}

	bufout.WriteString("Port '" + strconv.Itoa(appPort) + "'" + endl)
	bufout.WriteString("Auth user '" + appAuthUser + "'" + endl)
	bufout.WriteString("Auth password '" + appAuthPassword + "'" + endl)

	bufout.Flush()

	appExec()

	return nil
}

func appExec() {
	r := mux.NewRouter()

	r.HandleFunc("/app/start", authorized(appStart)).Methods("POST")

	http.ListenAndServe(":"+strconv.Itoa(appPort), handlers.LoggingHandler(os.Stdout, r))
}

func appStart(w http.ResponseWriter, r *http.Request) {
	var (
		body []byte
		err  error
	)

	if r.Header.Get("Content-Type") != "application/json" {
		jsonResponse(w, http.StatusBadRequest, "unknown payload")
		return
	}

	if body, err = ioutil.ReadAll(r.Body); err != nil {
		jsonResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer r.Body.Close()

	ar := &appRequest{}
	if err = json.Unmarshal(body, ar); err != nil {
		jsonResponse(w, http.StatusBadRequest, "cant unpack payload")
		return
	}

	cmd := exec.Command(ar.Name)
	if err := cmd.Start(); err != nil {
		jsonResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	jsonResponse(w, http.StatusOK, "success")
}
