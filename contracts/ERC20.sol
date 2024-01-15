// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./IERC20.sol";

abstract contract ERC20 is IERC20 {
    string private _name;
    string private _symbol;
    uint256 private _totalSupply;

    mapping(address account => uint256) private _balances;
    mapping(address account => mapping(address spender => uint256)) private _allowances;

    constructor(string memory name_, string memory symbol_, uint256 totalSupply_) {
        _name = name_;
        _symbol = symbol_;
        _totalSupply = totalSupply_;
    }

    function name() override public view returns (string memory) {
        return _name;
    }

    function symbol() override public view returns (string memory) {
        return _symbol;
    }

    function decimals() override public pure returns (uint8) {
        return 8;
    }

    function totalSupply() override public view returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address _owner) override public view returns (uint256) {
        return _balances[_owner];
    }

    function transfer(address _to, uint256 _value) public returns(bool) {
    }

    function transferFrom(address _from, address _to, uint256 _value) public returns (bool success) {}

    function approve(address _spender, uint256 _value) public returns (bool success) {}

    function allowance(address _owner, address _spender) public view returns (uint256 remaining) {}
}