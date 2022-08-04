mkdir ../privatenet_multinode/logs

# In terminal 1
thetasubchain start --config=./node1 --password=qwertyuiop |& tee ./logs/node1.log

# In terminal 2
thetasubchain start --config=./node2 --password=qwertyuiop |& tee ./logs/node2.log

# In terminal 3
thetasubchain start --config=./node3 --password=qwertyuiop |& tee ./logs/node3.log

# In terminal 4
thetasubchain start --config=./node4 --password=qwertyuiop |& tee ./logs/node4.log