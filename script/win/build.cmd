:: Script for building an application for Windows

cd ../..

:: Del old data
IF EXIST bin RMDIR bin /Q /S
IF EXIST app.exe DEL app.exe /Q /S
IF EXIST debug.out DEL debug.out /Q /S
IF EXIST test.out DEL test.out /Q /S

IF EXIST go.mod DEL go.mod /Q /S
IF EXIST go.sum DEL go.sum /Q /S

:: Set env variable
set GOROOT=c:\go
set GOOS=windows

:: Code format
go fmt

go mod init github.com/karpovdl/app

:: Code build x64
set GOARCH=amd64
go build -ldflags "-s -w" -i -v -o bin/release/64/app.exe
go build -race -i -v -o bin/debug/64/app.exe

:: Code build x86
set GOARCH=386
go build -ldflags "-s -w" -i -v -o bin/release/32/app.exe
go build -i -v -o bin/debug/32/app.exe