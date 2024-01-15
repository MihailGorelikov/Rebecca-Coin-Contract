// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./ERC20.sol";

/// @custom:security-contact Mihail.Gorelikov.Dev@outlook.com
contract RebeccaCoin is ERC20 {
    constructor(uint256 _totalSupply) ERC20("RebeccaCoin", "RBC", _totalSupply) {}
}