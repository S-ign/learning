const assert = require('assert');
const ganache = require('ganache-cli');
const Web3 = require('web3');
const web3 = new Web3(ganache.provider());
const {abi, bytecode} = require('../compile')

let accounts;
let ballot;
const MAX_PLAYERS = 5;

beforeEach(async () => {
  // Get a list of all accounts
  accounts = await web3.eth.getAccounts();

  // Use one of those accounts to deploy the contract
  ballot = await new web3.eth.Contract(abi)
    .deploy({ data: bytecode, arguments: [100000000000, MAX_PLAYERS] })
    .send({ from: accounts[0], gas: 5373352}) // gasLimit 6721975
  console.log(ballot)
});

describe('Ballot', () => {
  it('deploys a contract', () => {
    assert.ok(ballot.options.address);
  });

  it('adds a player to the game', async () => {
    await ballot.methods.enter().send({
      from: accounts[1],
      value: 100000000000,
      gas: 205781
    });

    const players = await ballot.methods.getPlayers().call({
      from: accounts[0],
    });

    assert.equal(accounts[1], players[0]);
    assert.equal(1, players.length);

  });

  it('can not enter as contract owner', async () => {
    assert.rejects(
      async () => {
        await ballot.methods.enter().send({
          from: accounts[0],
          value: 100000000000,
          gas: 205781
        });
      }
    )

    const players = await ballot.methods.getPlayers().call({
      from: accounts[0],
    });

    assert.equal(0, players.length);

  });

  it('can not add the same player twice', async () => {
    await ballot.methods.enter().send({
      from: accounts[1],
      value: 100000000000,
      gas: 205781
    });

    assert.rejects(
      async () => {
        await ballot.methods.enter().send({
          from: accounts[1],
          value: 100000000000,
          gas: 205781
        });
      }
    )

    const players = await ballot.methods.getPlayers().call({
      from: accounts[0]
    });

    assert.equal(1, players.length);

  });

  it('can call prizePool', async () => {
    await ballot.methods.enter().send({
      from: accounts[1],
      value: 100000000000,
      gas: 205781
    });

    const prizePool = await ballot.methods.prizePool('gwei').call({
      from: accounts[1]
    });

    assert.equal(100, prizePool);

  });

  it('resets after a winner has been chosen', async () => {
    for ( let i = 0; i < MAX_PLAYERS; i++) {
      await ballot.methods.enter().send({
        from: accounts[i+1],
        value: 100000000000,
        gas: 205781
      });
    }

    const players = await ballot.methods.getPlayers().call({
      from: accounts[0],
    });

    assert.equal(0, players.length)
  })


});
