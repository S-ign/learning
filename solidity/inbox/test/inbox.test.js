const assert = require('assert');
const ganache = require('ganache-cli');
const Web3 = require('web3');
const web3 = new Web3(ganache.provider());
const {abi, bytecode} = require('../compile')

let accounts;
let inbox;

beforeEach(async () => {
  // Get a list of all accounts
  accounts = await web3.eth.getAccounts();

  // Use one of those accounts to deploy the contract
  inbox = await new web3.eth.Contract(abi)
    .deploy({ data: bytecode, arguments: ['Hello World'] })
    .send({ from: accounts[0], gas: 1000000 })
  console.log(inbox)
});

describe('Inbox', () => {
  it('deploys a contract', () => {
    assert.ok(inbox.options.address);
  });

  it('has a default message', async () => {
    const message = await inbox.methods.message().call();
    console.log(message);

    assert.equal(message, 'Hello World');
  });

  it('can change message', async () => {
    await inbox.methods.setMessage('The World Hears You').send({ from:
      accounts[0] });

    const message = await inbox.methods.message().call();
    console.log(message);
    assert.equal(message, 'The World Hears You')
  });

});
