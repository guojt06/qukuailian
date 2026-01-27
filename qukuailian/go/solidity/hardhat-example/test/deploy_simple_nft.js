const { ethers } = require("hardhat");

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log("部署者地址:", deployer.address);

  // 合约参数
  const name = "SimpleNFT";
  const symbol = "SNFT";

  // 部署合约
  const SimpleNFT = await ethers.getContractFactory("SimpleNFT");
  const nft = await SimpleNFT.deploy(name, symbol);

  await nft.waitForDeployment();
  
  const contractAddress = await nft.getAddress();
  console.log("✅ NFT合约部署成功!");
  console.log("合约地址:", contractAddress);
  console.log("集合名称:", name);
  console.log("集合符号:", symbol);
  console.log("初始铸造者:", deployer.address);
  
  return { nft, contractAddress };
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});