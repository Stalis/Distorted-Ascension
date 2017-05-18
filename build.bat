@echo off
cd %~dp0
echo Generating...
go generate ./Engine/Map ./Engine
echo Building...
go build -o bin/game.exe
echo Done!