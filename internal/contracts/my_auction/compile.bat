@echo off
cd /d %~dp0
echo Compiling MyXAuctionV2.sol with viaIR enabled...
solcjs --base-path . --include-path node_modules --standard-json compile-with-viair.json > output.json
if %ERRORLEVEL% EQU 0 (
    echo Compilation successful!
    echo Output saved to output.json
) else (
    echo Compilation failed!
    type output.json
)

