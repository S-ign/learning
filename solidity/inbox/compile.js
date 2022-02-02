var path = require('path');
var fs = require('fs');
var solc = require('solc');
var inboxPath = path.resolve(__dirname, 'contracts', 'Inbox.sol');
var source = fs.readFileSync(inboxPath, 'UTF-8');
var input = {
    language: 'Solidity',
    sources: {
        'Inbox.sol': {
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
module.exports.abi = output.contracts['Inbox.sol']['Inbox']['abi']
module.exports.bytecode = output.contracts['Inbox.sol']['Inbox'].evm.bytecode.object
