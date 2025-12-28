@echo off
cd /d %~dp0
echo Compiling MyXAuctionV2.sol with viaIR enabled...
echo.

REM 使用标准 JSON 输入，通过 PowerShell 处理
powershell -Command "$json = Get-Content 'compile-with-viair.json' -Raw | ConvertFrom-Json; $json.settings.viaIR = $true; $json | ConvertTo-Json -Depth 10 | solcjs --standard-json > output.json"

if %ERRORLEVEL% EQU 0 (
    echo.
    echo Compilation completed! Check output.json for results.
) else (
    echo.
    echo Compilation failed! Check output.json for errors.
    type output.json
)

