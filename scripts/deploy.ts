import {
  BaseContract,
  ContractFactory,
  ContractTransactionResponse,
} from "ethers";
import { ethers } from "hardhat";
import { RebeccaCoin__factory } from "../typechain-types/factories/contracts/RebeccaCoin__factory";

async function main() {
  const factory: ContractFactory = (await ethers.getContractFactory(
    "RebeccaCoin"
  )) as RebeccaCoin__factory;

  const initialAuthority: String = process.env.INITIAL_AUTHORITY as string;
  if (!initialAuthority) {
    throw new Error("No initial authority provided");
  }

  const token: BaseContract & {
    deploymentTransaction(): ContractTransactionResponse;
  } & Omit<BaseContract, keyof BaseContract> = await factory.deploy(
    initialAuthority
  );

  const deployedCode: string | null = await token.getDeployedCode();
  if (!deployedCode) {
    throw new Error("No deployed code found");
  }

  console.log(`Deployed to ${await token.getAddress()}`);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
