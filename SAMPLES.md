# Samples

## File for run and debug

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "app exec",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "env": {
                "APP_AUTH_BASIC_USER": "basicuser",
                "APP_AUTH_BASIC_PASSWORD": "*basicpassword*",
            },
            "args": [
                "exec",
                "--port=9000",
            ]
        },
    ]
}
```