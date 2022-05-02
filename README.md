# theta-protocol-subchain-poc

## Compilation

```shell
git clone https://github.com/thetatoken/theta-protocol-ledger.git $GOPATH/src/github.com/thetatoken/theta
export THETA_HOME=$GOPATH/src/github.com/thetatoken/theta
cd $THETA_HOME
git checkout sc-privatenet
git pull origin sc-privatenet

git clone https://github.com/thetatoken/theta-protocol-subchain-poc $GOPATH/src/github.com/thetatoken/thetasubchain
export SUBCHAIN_HOME=$GOPATH/src/github.com/thetatoken/thetasubchain
cd $SUBCHAIN_HOME

make install
```

## Start a single validator testnet

### Setup (run once)

```shell
cd $SUBCHAIN_HOME
mkdir -p ../data/subchain/privatenet/node
rm -rf ../data/subchain/privatenet/node/*
cp -r integration/privatenet/node/* ../data/subchain/privatenet/node/
```

### Start the validator

```shell
cd $SUBCHAIN_HOME
thetasubchain start --config=../data/subchain/privatenet/node --password=qwertyuiop
```

## Query the validator

```shell
# query status
thetasubcli query status

# query validator set at a given block height
thetasubcli query vs --height=20

# query block(s) at a given height
thetasubcli query block --height=122
```

## Misc

If you need to generate a new genesis snapshot for the single node testnet, please use the following command. The json file `init_validator_set.json` specifies the initial validator set.

```shell
cd $SUBCHAIN_HOME/integration/privatenet/node
subchain_generate_genesis -chainID=private_subchain -initValidatorSet=./data/init_validator_set.json -genesis=./snapshot
```

The above commands should print the genesis block hash (see below). Please replace the `genesis.hash` parameter in your `config.yaml` file with the genesis block hash printed.

```shell
-----------------------------------------------------------------------------------------
Genesis block hash: <GENESIS_BLOCK_HASH>
-----------------------------------------------------------------------------------------
```


