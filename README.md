# pass

[![License][1]][2] [![golang][10]][11]

[1]: https://img.shields.io/badge/license-MIT-blue.svg?label=License&maxAge=86400 "License"
[2]: ./LICENSE

[10]: https://img.shields.io/badge/golang-1.21.3-blue.svg?style=flat "Golang"
[11]: https://golang.org

:green_book: [Samples](./SAMPLES.md)
:green_book: [Tests](./TESTS.md)

* app --version, -v
* app --help, -h

To build an application for windows make script/build.cmd

## app

Application app - application for running processes
* app -h - help command

### app exec

* app exec, app e - run aplication
* app exec --help, -h

Configuration exec.json file

```json
{
    "port": {APP_PORT},
    "auth_user": "{APP_AUTH_USER}",
    "auth_pass": "{APP_AUTH_PASS}"
}
```

### sample run

Sample run server without parameters

```bash
app.exe exec
```

Sample run with parameters

```bash
app.exe ^
  "exec" ^
  "--port=9000" ^
  "--auth_user=user" ^
  "--auth_pass=****"
```
