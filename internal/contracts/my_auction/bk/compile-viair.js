const fs = require('fs');
const { spawn } = require('child_process');
const path = require('path');

// 读取合约源代码
const contractPath = path.join(__dirname, 'MyXAuctionV2.sol');
const contractContent = fs.readFileSync(contractPath, 'utf8');

// 创建标准 JSON 输入（使用 content 字段）
const input = {
  language: "Solidity",
  sources: {
    "MyXAuctionV2.sol": {
      content: contractContent
    }
  },
  settings: {
    optimizer: {
      enabled: true,
      runs: 200
    },
    viaIR: true,  // 启用 viaIR 解决 Stack too deep 错误
    outputSelection: {
      "*": {
        "*": [
          "abi",
          "evm.bytecode",
          "evm.deployedBytecode"
        ]
      }
    }
  }
};

// 将输入写入临时文件
const inputFile = path.join(__dirname, 'input.json');
fs.writeFileSync(inputFile, JSON.stringify(input, null, 2));

console.log('Compiling MyXAuctionV2.sol with viaIR enabled...');
console.log('Using base-path and include-path for imports...');
console.log('');

// 执行编译（使用标准 JSON 输入，通过 stdin 传递）
const inputContent = fs.readFileSync(inputFile, 'utf8');

const child = spawn('solcjs', [
  '--base-path', '.',
  '--include-path', 'node_modules',
  '--standard-json'
], {
  cwd: __dirname,
  shell: true,
  stdio: ['pipe', 'pipe', 'pipe']
});

let output = '';
let errorOutput = '';

child.stdout.on('data', (data) => {
  output += data.toString();
});

child.stderr.on('data', (data) => {
  errorOutput += data.toString();
});

child.on('close', (code) => {
  // 保存输出
  const outputFile = path.join(__dirname, 'output.json');
  const finalOutput = output || errorOutput;
  fs.writeFileSync(outputFile, finalOutput);
  
  // 检查是否有错误输出
  if (errorOutput && errorOutput.trim().length > 0 && !output) {
    console.error('❌ Compilation failed with errors:');
    console.error(errorOutput);
    process.exit(1);
    return;
  }
  
  // 检查输出是否为空
  if (!finalOutput || finalOutput.trim().length === 0) {
    console.error('❌ Compilation failed: No output received');
    process.exit(1);
    return;
  }
  
  // 尝试提取 JSON（可能输出中有其他内容）
  let jsonOutput = finalOutput.trim();
  // 如果输出不是以 { 开头，尝试找到 JSON 部分
  if (!jsonOutput.startsWith('{')) {
    const jsonStart = jsonOutput.indexOf('{');
    if (jsonStart !== -1) {
      jsonOutput = jsonOutput.substring(jsonStart);
    }
    // 如果输出不是以 } 结尾，尝试找到 JSON 结束
    if (!jsonOutput.endsWith('}')) {
      const jsonEnd = jsonOutput.lastIndexOf('}');
      if (jsonEnd !== -1) {
        jsonOutput = jsonOutput.substring(0, jsonEnd + 1);
      }
    }
  }
  
  try {
    // 解析输出
    const result = JSON.parse(jsonOutput);
    
    if (result.errors) {
      const errors = result.errors.filter(e => e.severity === 'error');
      if (errors.length > 0) {
        console.error('❌ Compilation errors:');
        errors.forEach(err => {
          console.error(`  ${err.message}`);
          if (err.formattedMessage) {
            console.error(`  ${err.formattedMessage}`);
          }
        });
        process.exit(1);
      }
    }
    
    if (result.contracts && result.contracts['MyXAuctionV2.sol'] && result.contracts['MyXAuctionV2.sol']['MyXAuctionV2']) {
      console.log('✅ Compilation successful!');
      console.log('Output saved to:', outputFile);
      
      // 提取 ABI 和字节码
      const contract = result.contracts['MyXAuctionV2.sol']['MyXAuctionV2'];
      
      // 输出 ABI 文件（使用脚本所在目录，确保文件在同一目录）
      const outputDir = __dirname;
      const abiPath = path.join(outputDir, 'MyXAuctionV2.abi');
      const binPath = path.join(outputDir, 'MyXAuctionV2.bin');
      const deployedBinPath = path.join(outputDir, 'MyXAuctionV2.deployed.bin');
      
      console.log('Output directory:', outputDir);
      
      if (contract.abi) {
        fs.writeFileSync(abiPath, JSON.stringify(contract.abi, null, 2));
        console.log('✅ ABI saved to:', abiPath);
      } else {
        console.warn('⚠️  Warning: ABI not found in compilation output');
      }
      
      // 输出 bin 文件（字节码）
      if (contract.evm && contract.evm.bytecode && contract.evm.bytecode.object) {
        fs.writeFileSync(binPath, contract.evm.bytecode.object);
        console.log('✅ Bytecode (bin) saved to:', binPath);
      } else {
        console.warn('⚠️  Warning: Bytecode not found in compilation output');
      }
      
      // 同时输出部署字节码（可选）
      if (contract.evm && contract.evm.deployedBytecode && contract.evm.deployedBytecode.object) {
        fs.writeFileSync(deployedBinPath, contract.evm.deployedBytecode.object);
        console.log('✅ Deployed bytecode saved to:', deployedBinPath);
      }
    } else {
      console.error('❌ Compilation failed: Contract not found in output');
      if (result.errors && result.errors.length > 0) {
        result.errors.forEach(err => {
          console.error(`  ${err.message}`);
        });
      }
      process.exit(1);
    }
  } catch (parseError) {
    console.error('❌ Failed to parse compilation output:', parseError.message);
    console.log('Raw output:', finalOutput);
    process.exit(1);
  }
});

// 通过 stdin 传递 JSON
child.stdin.write(inputContent);
child.stdin.end();
