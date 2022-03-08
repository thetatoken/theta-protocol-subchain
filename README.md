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