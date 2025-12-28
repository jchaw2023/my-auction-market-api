@echo off
cd /d %~dp0
echo Compiling MyXAuctionV2.sol with viaIR enabled...
echo Output: ABI and BIN files
echo.

REM 使用 PowerShell 执行标准 JSON 编译并提取 ABI 和 BIN
powershell -Command ^
"$json = Get-Content 'compile-with-viair.json' -Raw; ^
$output = $json | solcjs --base-path . --include-path node_modules --standard-json 2>&1; ^
$result = $output | ConvertFrom-Json; ^
if ($result.contracts.'MyXAuctionV2.sol'.MyXAuctionV2) { ^
  $contract = $result.contracts.'MyXAuctionV2.sol'.MyXAuctionV2; ^
  if ($contract.abi) { ^
    $contract.abi | ConvertTo-Json -Depth 10 | Out-File -Encoding utf8 'MyXAuctionV2.abi'; ^
    Write-Host 'ABI saved to: MyXAuctionV2.abi' -ForegroundColor Green; ^
  } ^
  if ($contract.evm.bytecode.object) { ^
    $contract.evm.bytecode.object | Out-File -Encoding ascii -NoNewline 'MyXAuctionV2.bin'; ^
    Write-Host 'Bytecode (bin) saved to: MyXAuctionV2.bin' -ForegroundColor Green; ^
  } ^
  $output | Out-File -Encoding utf8 'output.json'; ^
  Write-Host 'Full output saved to: output.json' -ForegroundColor Green; ^
} else { ^
  Write-Host 'Compilation failed!' -ForegroundColor Red; ^
  $output | Out-File -Encoding utf8 'output.json'; ^
  $result.errors | Where-Object { $_.severity -eq 'error' } | ForEach-Object { Write-Host $_.message -ForegroundColor Red }; ^
  exit 1; ^
}"

if %ERRORLEVEL% EQU 0 (
    echo.
    echo ✅ Compilation completed successfully!
) else (
    echo.
    echo ❌ Compilation failed! Check output.json for details.
)

