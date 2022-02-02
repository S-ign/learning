const HDWalletProvider = require ('@truffle/hdwallet-provider');
const Web3 = require('web3');
const {abi, bytecode} = require('./compile');

const provider = new HDWalletProvider(
  'hip divorce hospital process scrap huge guard enact cash zoo patient poem',
  'https://rinkeby.infura.io/v3/b2602a19ab3747e181bb9cf2c45fa73a'
);

const web3 = new Web3(provider)
const deploy = async () => {
  const accounts = await web3.eth.getAccounts();

  console.log('attempting to deploy from account: ', accounts[0]);

  const contract = await new web3.eth.Contract(abi)
    .deploy({data: bytecode, arguments: ['10000000000000','4']})
    .send({gas: '1000000', from: accounts[0]})

  console.log('contract deployed to', accounts[0])
};
deploy();
