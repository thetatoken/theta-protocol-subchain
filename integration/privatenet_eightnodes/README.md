### To launch a local privatenet with 4 validators ###

First follow the steps [here](https://docs.thetatoken.org/docs/setup) to compile the latest code of the `privatenet` branch. Next, run the commands below to launch the privatenet with 4 validators:

```
cd $THETA_HOME
cp -r ./integration/privatenet_eightnodes ../privatenet_eightnodes
mkdir ../privatenet_eightnodes/logs

# In terminal 1
thetasubchain start --config=../privatenet_eightnodes/node1 --password=qwertyuiop |& tee ../privatenet_eightnodes/logs/node1.log

# In terminal 2
thetasubchain start --config=../privatenet_eightnodes/node2 --password=qwertyuiop |& tee ../privatenet_eightnodes/logs/node2.log

# In terminal 3
thetasubchain start --config=../privatenet_eightnodes/node3 --password=qwertyuiop |& tee ../privatenet_eightnodes/logs/node3.log

# In terminal 4
thetasubchain start --config=../privatenet_eightnodes/node4 --password=qwertyuiop |& tee ../privatenet_eightnodes/logs/node4.log

# In terminal 5
thetasubchain start --config=../privatenet_eightnodes/node5 --password=qwertyuiop |& tee ../privatenet_eightnodes/logs/node5.log

# In terminal 6
thetasubchain start --config=../privatenet_eightnodes/node6 --password=qwertyuiop |& tee ../privatenet_eightnodes/logs/node6.log

# In terminal 7
thetasubchain start --config=../privatenet_eightnodes/node7 --password=qwertyuiop |& tee ../privatenet_eightnodes/logs/node7.log

# In terminal 8
thetasubchain start --config=../privatenet_eightnodes/node8 --password=qwertyuiop |& tee ../privatenet_eightnodes/logs/node8.log
```

# How to generate plain private key?
use the web3 Cli tool
web3 account extract --keyfile ./node1/key/encrypted/9F1233798E905E173560071255140b4A8aBd3Ec6 --password qwertyuiop
