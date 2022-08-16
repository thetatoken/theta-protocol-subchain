### To launch a local privatenet with 4 validators ###

First follow the steps [here](https://docs.thetatoken.org/docs/setup) to compile the latest code of the `privatenet` branch. Next, run the commands below to launch the privatenet with 4 validators:

```
cd $THETA_HOME
cp -r ./integration/privatenet_multinode ../privatenet_multinode
mkdir ../privatenet_multinode/logs

# In terminal 1
thetasubchain start --config=../privatenet_multinode/node1 --password=qwertyuiop |& tee ../privatenet_multinode/logs/node1.log

# In terminal 2
thetasubchain start --config=../privatenet_multinode/node2 --password=qwertyuiop |& tee ../privatenet_multinode/logs/node2.log

# In terminal 3
thetasubchain start --config=../privatenet_multinode/node3 --password=qwertyuiop |& tee ../privatenet_multinode/logs/node3.log

# In terminal 4
thetasubchain start --config=../privatenet_multinode/node4 --password=qwertyuiop |& tee ../privatenet_multinode/logs/node4.log
```
