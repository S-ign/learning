var path = require('path');
var fs = require('fs');
var solc = require('solc');
var projectPath = path.resolve(__dirname, 'contracts', 'ballot.sol');
var source = fs.readFileSync(projectPath, 'UTF-8');
var input = {
    language: 'Solidity',
    sources: {
        'ballot.sol': {
            content: source
        }
    },
    settings: {
        outputSelection: {
            '*': {
                '*': ['*']
            }
        }
    }
};
var output = JSON.parse(solc.compile(JSON.stringify(input)));
//console.log(output.contracts['ballot.sol']['Ballot'].evm.bytecode.object)
module.exports.abi = output.contracts['ballot.sol']['Ballot']['abi']
module.exports.bytecode = output.contracts['ballot.sol']['Ballot'].evm.bytecode.object
