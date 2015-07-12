@echo off

cd %cd%

set GOPATH=%GOPATH%;%cd%


go build -o goWeb.exe main.go

echo build finished
echo start service
goWeb.exe 
pause /s