# theta-protocol-subchain

In April 2022, Theta Labs [announced the Theta Metachain concept](https://twitter.com/theta_network/status/1512555910568292355). Targeted to be launched on Dec 1, 2022, the Theta Metachain is an interconnected network of blockchains, a “chain of chains”. The goal is to allow permissionless horizontal scaling of the Theta blockchain network in order to achieve potentially unlimited transactional throughput, and 1-2 seconds, or even subsecond block finalization time.

The [Theta Metachain](https://assets.thetatoken.org/theta-mainnet-4-whitepaper.pdf) consists of one “main chain” and an unlimited number of “subchains”. Just as “meta-” as a prefix refers to something that transcends or is more comprehensive than the subject, ex. metaphysics describing what exists beyond physics, Theta Metachain refers to an overarching “main chain” above many purpose-specific “subchains”. The “main chain” in this case refers to the existing Theta mainnet. Theta will offer an easy-to-use SDK that developers can quickly use to launch a subchain and plug it into the main chain. Since each subchain can execute transactions independently, this provides a viable path to infinitely scale the processing capacity of the Metachain. The subchain SDK will implement a built-in interchain messaging channel which connects the subchains and the main chain, and thus allows crypto assets like TNT20/721 tokens to flow freely across the chains. The process of creating a subchain is permissionless, meaning that anyone can register and launch a subchain. No approval from Theta Labs is required. This repository contain the reference implementation of the Theta subchain.

Initially implemented as a multi-chain solution, the Theta Metachain can be extended into a zk-rollup by adding a few gadgets. Such an extension can achieve a higher level of security guarantees. To learn more about Theta Metachain, please check out the [Theta Mainnet 4.0 whitepaper](https://assets.thetatoken.org/theta-mainnet-4-whitepaper.pdf).

## Compilation

```shell
git clone https://github.com/thetatoken/theta-protocol-ledger.git $GOPATH/src/github.com/thetatoken/theta
export THETA_HOME=$GOPATH/src/github.com/thetatoken/theta
cd $THETA_HOME
git checkout sc-privatenet
git pull origin sc-privatenet

git clone https://github.com/thetatoken/theta-protocol-subchain $GOPATH/src/github.com/thetatoken/thetasubchain
export SUBCHAIN_HOME=$GOPATH/src/github.com/thetatoken/thetasubchain
cd $SUBCHAIN_HOME

make install
```

## Misc

If you need to generate a new genesis snapshot for the single node testnet, please use the following command. The json file `init_validator_set.json` specifies the initial validator set.

```shell
cd $SUBCHAIN_HOME/integration/privatenet/node
subchain_generate_genesis -chainID=tsub_360777 -initValidatorSet=./data/init_validator_set.json -genesis=./snapshot
```

The above commands should print the genesis block hash (see below). Please replace the `genesis.hash` parameter in your `config.yaml` file with the genesis block hash printed.

```shell
-----------------------------------------------------------------------------------------
Genesis block hash: <GENESIS_BLOCK_HASH>
-----------------------------------------------------------------------------------------
```
