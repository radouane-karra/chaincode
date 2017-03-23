# Energy Trading Example App
This repository contains an example energy trader application that talks to a blockchain backend.

## Get Started
First clone this repo and install the dependencies.

1. Clone the repo and go into web directory

```
$ git clone https://github.com/predix/chaincode_example.git && cd energy_trading/web
```

2. Install dependencies (npm and bower)

```
$ npm install && bower install
```

3. Start application locally

```
$ npm start
```

> Open web browser to http://localhost:9001

4. Deploy application to Cloud Foundry

```
$ npm run deploy
```


## Customizing

5. To change port and blockchain endpoint, simply sent the following env vars.

```
$ export PORT=9049;
$ export BLOCKCHAIN_ENDPOINT=https://endpoint
```

2. To change the chaincodeID open the `app/scripts/main.js` file and update the values.

```
//app/scripts/main.js
scope.config = {
	endpoint: '/api/v1',
	secureContext: null,
	chaincodeID: {
		report: '30268bf2818712b14161bd47db875bd5786b357641c2e09a218ff120dc2b072a15edc2e05a87bf5664debefab25880e91fa10ad0f62dde9ffb9ac47f91c8f73e',
		settle: '30268bf2818712b14161bd47db875bd5786b357641c2e09a218ff120dc2b072a15edc2e05a87bf5664debefab25880e91fa10ad0f62dde9ffb9ac47f91c8f73e'
	}
};
```

> Note: The endpoint property is the local express server that proxies requests to `BLOCKCHAIN_ENDPOINT`


## Usage
The web application demonstrates using the blockchain REST api to invoke and query nodes.

Here are some steps to get you started.

1. Once the server is running open the URL in the browser.
2. The `start/stop clock` button will simulate reporting random data and settling every other second.
3. The `random report` button will report random data for each meter, (+ val is a producer / - val is a consumer).
4. The `settle` button will settle the meters and refresh the meters thus transfering the account_balance from meters.
5. The `refresh` icon button will refresh all meters.

> Note: If values are not updating correctly that is because of the lag on the blockchain code
