@echo off
setlocal

if "%1"=="" (
    echo Please specify package version, e.g.  0.1.0
    exit /b 1
)
set PACKAGE_VERSION=%1

rd /s /q ..\docs 2>nul
rd /s /q ..\hwmux_client 2>nul

curl -k http://localhost:80/schema/download -o hwmux.yaml
docker compose run --rm openapi-generator-cli generate -i /local/tools/hwmux.yaml -g go --additional-properties=projectName=hwmux-client-golang,packageVersion=!PACKAGE_VERSION!,packageName=hwmux -o /local/
