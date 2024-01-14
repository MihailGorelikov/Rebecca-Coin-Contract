// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/manager/AccessManaged.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

/// @custom:security-contact Mihail.Gorelikov.Dev@outlook.com
contract RebeccaCoin is ERC20, AccessManaged, ERC20Permit {
    constructor(address initialAuthority)
        ERC20("RebeccaCoin", "REBECCA")
        AccessManaged(initialAuthority)
        ERC20Permit("RebeccaCoin")
    {
        _mint(msg.sender, 1000 * 10 ** decimals());
    }

    function mint(address to, uint256 amount) public restricted {
        _mint(to, amount);
    }
}