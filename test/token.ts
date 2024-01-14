import {
  BaseContract,
  ContractFactory,
  ContractTransactionResponse,
} from "ethers";
import { ethers } from "hardhat";
import { RebeccaCoin__factory } from "../typechain-types/factories/contracts/RebeccaCoin__factory";
import { SignerWithAddress } from "@nomicfoundation/hardhat-ethers/signers";
import { RebeccaCoin } from "../typechain-types/contracts/RebeccaCoin";

export default class Token {
  public static readonly TOKEN_NAME: String | any = "RebeccaCoin";
  public static readonly TOKEN_SYMBOL: String = "RBC";
  public static readonly INITIAL_SUPPLY: BigInt = 1000n * 10n ** 18n;

  private _totalSupply: BigInt;
  private _addr1?: SignerWithAddress;
  private _addr2?: SignerWithAddress;
  private _addrs?: SignerWithAddress[];
  private _owner?: SignerWithAddress;
  private _token?: RebeccaCoin;

  public constructor() {
    this._totalSupply = Token.INITIAL_SUPPLY;
    this._owner = undefined;
    this._addr1 = undefined;
    this._addr2 = undefined;
    this._addrs = undefined;
  }

  public async deploy(): Promise<void> {
    const [owner, addr1, addr2, ...addrs] = await ethers.getSigners();

    const factory: ContractFactory = (await ethers.getContractFactory(
      Token.TOKEN_NAME,
      owner
    )) as RebeccaCoin__factory;

    const token: BaseContract & {
      deploymentTransaction(): ContractTransactionResponse;
    } & Omit<BaseContract, keyof BaseContract> = await factory.deploy(
      owner.address.toString()
    );

    this._owner = owner;
    this._addr1 = addr1;
    this._addr2 = addr2;
    this._addrs = addrs;
    this._token = token as RebeccaCoin;
  }

  public async balanceOf(address: string): Promise<bigint | undefined> {
    return await this._token?.balanceOf(address);
  }

  public async transfer(
    to: SignerWithAddress,
    amount: number | bigint
  ): Promise<void> {
    await this._token?.transfer(to.address as string, amount);
  }

  public async transferAs(
    from: SignerWithAddress,
    to: SignerWithAddress,
    amount: number
  ): Promise<void> {
    await this._token?.connect(from).transfer(to.address, amount);
  }

  public get totalSupply(): BigInt {
    return this._totalSupply;
  }

  public get owner(): SignerWithAddress | undefined {
    return this._owner;
  }

  public get addr1(): SignerWithAddress | undefined {
    return this._addr1;
  }

  public get addr2(): SignerWithAddress | undefined {
    return this._addr2;
  }

  public get addrs(): SignerWithAddress[] | undefined {
    return this._addrs;
  }

  public get token(): BaseContract | undefined {
    return this._token;
  }
}
