# Energy trading smart contract
Smart contract to demonstrate peer to peer energy trading. We are using [this](https://github.com/olegabu/decentralized-energy-utility) as our starting point. It was a project done for hackathon on hyperledger and we are changing it to demonstrate peer to peer energy trading. Following changes have been made:

1. Allow chain code deployer to specify the commission that will be charged by the exchange smart contract from producers.
1. Exchange account (grid maintainer) is initialized at time of deploy.
1. A method to register a new meter is provided with the rate at which it is ready to sell/buy energy.
1. A method to delete existing meter is provided.
1. Data on meter id, meter name, reported kwh and account balance is stored in a table.
1. During settlement all the rows in the table are considered unlike hardcoded meter ids from 1 to 10 in original chain code implementation. Also, it matches buyers with sellers based on the rate and transfers account balance accordingly.
1. Additional query methods are provided to give meter information and exchange account balance.

## Steps to deploy and use this smart contract
1. Deploy chaincode

    ```
    curl -k -XPOST -d @scripts/energy_chaincode.txt https://<blockchain ip>/chaincode
    ```
1. Enroll new meters

    ```
    curl -k -XPOST -d @scripts/enroll.txt https://<blockchain ip>/chaincode
    ```
1. Fund new meter accounts with some coins

    ```
    curl -k -XPOST -d @scripts/change_account_balance.txt https://<blockchain ip>/chaincode
    ```
1. Report power consumed or produced by each meter (+ve is produced and -ve is consumed)

    ```
    curl -k -XPOST -d @scripts/report_kwh.txt https://<blockchain ip>/chaincode
    ```
1. Query reported kwh

    ```
    curl -k -XPOST -d @scripts/reportkwh_query.txt https://<blockchain ip>/chaincode
    ```
1. Query account balance for each account

    ```
    curl -k -XPOST -d @scripts/balance_query.txt https://<blockchain ip>/chaincode
    ```
1. Query for exchange rate/comission

    ```
    curl -k -XPOST -d @scripts/exchangerate_query.txt https://<blockchain ip>/chaincode
    ```
1. Settle accounts by transferring money from consumers to producers

    ```
    curl -k -XPOST -d @scripts/settle.txt https://<blockchain ip>/chaincode
    ```
1. Query exchange account balance

    ```
    curl -k -XPOST -d @scripts/exchangeaccountbalance_query.txt https://<blockchain ip>/chaincode
    ```
1. Query meter information

    ```
    curl -k -XPOST -d @scripts/meter_query.txt https://<blockchain ip>/chaincode
    ```
1. Query all meters

    ```
    curl -k -XPOST -d @scripts/meters.txt https://<blockchain ip>/chaincode
    ```
1. Delete a meter

    ```
    curl -k -XPOST -d @scripts/delete_meter.txt https://<blockchain ip>/chaincode
    ```