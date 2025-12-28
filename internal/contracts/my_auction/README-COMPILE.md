# 编译 MyXAuctionV2.sol（启用 viaIR）

## 问题
直接使用 `solcjs` 编译时会出现 "Stack too deep" 错误，需要启用 `viaIR` 选项。

## 输出文件
编译成功后会生成以下文件：
- `MyXAuctionV2.abi` - 合约 ABI（JSON 格式）
- `MyXAuctionV2.bin` - 合约字节码（十六进制字符串）
- `MyXAuctionV2.deployed.bin` - 部署后的字节码（可选）
- `output.json` - 完整的编译输出（JSON 格式）

## 解决方案

### 方法 1: 使用 Node.js 脚本（推荐）⭐

```bash
npm run compile
```

或者直接运行：
```bash
node compile-viair.js
```

**输出：**
- ✅ 自动生成 `MyXAuctionV2.abi`
- ✅ 自动生成 `MyXAuctionV2.bin`
- ✅ 自动生成 `MyXAuctionV2.deployed.bin`
- ✅ 保存完整输出到 `output.json`

### 方法 2: 使用批处理脚本（Windows）

```cmd
compile-direct.bat
```

**输出：**
- ✅ 自动生成 `MyXAuctionV2.abi`
- ✅ 自动生成 `MyXAuctionV2.bin`
- ✅ 保存完整输出到 `output.json`

### 方法 3: 使用标准 JSON 输入文件（手动）

1. 使用 PowerShell 执行：
```powershell
Get-Content compile-with-viair.json | solcjs --base-path . --include-path node_modules --standard-json > output.json
```

2. 然后从 `output.json` 中提取 ABI 和 bin：
```powershell
# 提取 ABI
$result = Get-Content output.json | ConvertFrom-Json
$result.contracts.'MyXAuctionV2.sol'.MyXAuctionV2.abi | ConvertTo-Json -Depth 10 | Out-File MyXAuctionV2.abi

# 提取 bin
$result.contracts.'MyXAuctionV2.sol'.MyXAuctionV2.evm.bytecode.object | Out-File -NoNewline MyXAuctionV2.bin
```

### 方法 4: 使用 Hardhat（如果可用）

如果项目中有 Hardhat 配置，可以在 `hardhat.config.js` 中启用 `viaIR: true`，然后使用：
```bash
npx hardhat compile
```

## 注意事项

- `viaIR: true` 会增加编译时间，但能解决 "Stack too deep" 错误
- 确保已安装所有依赖：`npm install`
- 确保 `node_modules` 目录存在且包含所需的 OpenZeppelin 和 Chainlink 合约
- ABI 文件是 JSON 格式，可以直接用于 Web3 库
- BIN 文件是十六进制字符串，用于合约部署

