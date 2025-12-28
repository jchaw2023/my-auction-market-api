@echo off
cd /d %~dp0
echo Compiling MyXAuctionV2.sol with viaIR enabled...
echo.

REM 使用 PowerShell 执行标准 JSON 编译
powershell -Command "$json = Get-Content 'compile-with-viair.json' -Raw; $json | solcjs --base-path . --include-path node_modules --standard-json > output.json 2>&1"

if %ERRORLEVEL% EQU 0 (
    echo.
    echo Compilation completed! Check output.json for results.
) else (
    echo.
    echo Compilation failed! Check output.json for errors.
    type output.json
)

