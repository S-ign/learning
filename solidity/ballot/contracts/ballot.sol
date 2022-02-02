// SPDX-License-Identifier: MIT
pragma solidity ^0.8.11;

contract Ballot {

  address private owner;
  address[] private players;
  uint private entryFee;
  uint private prizePoolWei;
  uint private maxParticipants;
  mapping(string=>uint) public prizePool;

  // save initial caller of contract's address
  constructor(uint fee, uint participants) {
    owner = msg.sender;

    // set fee in Wei
    entryFee = fee;

    // set how many people can enter the game
    maxParticipants = participants;
  }

  // create modifier for allowing only the initial caller of this contract to
  // run certain functions
  modifier onlyOwner {
    require(msg.sender == owner);
    _;
  }

  // only allow user to call function one time
  modifier onlyOnce {
    for (uint16 i = 0; i < players.length; i++) {
      require(msg.sender != players[i], "you can only enter the game once");
    }
    _;
  }

  // changes who owns contract
  function transferOwner(address newOwner) public onlyOwner {
    owner = newOwner;
  }

  // allows owner to change how many people can join the game after the game
  // has started
  function setMaxParticipants(uint participants) public onlyOwner{
    maxParticipants = participants;
  }

  // get entrants address and verify the correct amount of wei was sent
  function enter() public onlyOnce payable {

    // error out and revert transaction if function called without sending the
    // right amount of wei
    require(msg.value == entryFee, string(abi.encodePacked("must send exactly: ",  entryFee, " wei to enter")));

    // do not allow the same person to enter twice
    //require(newPlayer.Registered != true, "you can not enter twice from the same address");

    // prohibit owner of contract from entering the game
    require(msg.sender != owner, "owner of contract can not enter the game, enter from another address");

    // add new player to list
    players.push(msg.sender);

    // add to prize pool variable so people can check how large the pool is
    prizePoolWei += entryFee;
    prizePool['wei']     = prizePoolWei;
    prizePool['kwei']    = prizePoolWei / 1000; // factor: 3
    prizePool['mwei']    = prizePoolWei / 1000000; // factor: 6
    prizePool['gwei']    = prizePoolWei / 1000000000; // factor: 9
    prizePool['szabo']   = prizePoolWei / 1000000000000; // factor: 12
    prizePool['finney']  = prizePoolWei / 1000000000000000; // factor: 15
    prizePool['ether']   = prizePoolWei / 1000000000000000000; // factor: 18

    if (players.length == maxParticipants) {
      pickWinner();
    }
  }

  // get full list of players
  function getPlayers() public view onlyOwner returns(address[] memory) {
    return players;
  }

  function random() private view returns (uint) {
    return uint(keccak256(abi.encodePacked(block.difficulty, block.timestamp, players.length)));
  }

  // allow owner to end entry phase and have the contract randomly select a
  // winner
  function pickWinner() private {
    require(players.length == maxParticipants);

    // send winner all aquired eth, also causing balance to be 0
    payable(players[random()%players.length]).transfer(address(this).balance);

    // reset player count
    delete players;

    // reset prizePoolWei
    prizePoolWei = 0;
  }
}
