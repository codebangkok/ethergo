//SPDX-License-Identifier: MIT

pragma solidity ^0.6.0;

contract BondCoin {
    
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    
    uint256 _totalSupply;
    
    mapping(address => uint256) _balances;
    
    function name() public pure returns (string memory) {
        return "Bond Coin";
    }
    
    function symbol() public pure returns (string memory) {
        return "BON";
    }
    
    function decimals() public pure returns (uint8) {
        return 0;
    }
    
    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }
    
    function balanceOf(address _owner) public view returns (uint256 balance) {
        return _balances[_owner];
    }
    
    function transfer(address _to, uint256 _value) public returns (bool success) {
        require(_balances[msg.sender] >= _value, "Not enough coin");
        
        _balances[msg.sender] -= _value;
        _balances[_to] += _value;
        emit Transfer(msg.sender, _to, _value);
        return true;
    }
    
    function mint(address _owner, uint256 _value) public {
        _totalSupply += _value;
        _balances[_owner] += _value;
        emit Transfer(address(0), _owner, _value);
    }
}