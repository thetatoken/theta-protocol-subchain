## Usage

For subchain registration:

```
subchain_e2e_test_tools  RegisterSubchain (it will register 360777)
```

To stake to an validator from an account:

```
subchain_e2e_test_tools  AccountStake --id=(the id of accountList)
```

For testing transfer cases:

```
subchain_e2e_test_tools     MainchainTNT20Burn   --amount=(the amount or tokenid you want to operate)
                            MainchainTNT20Lock  
                            MainchainTNT721Burn 
                            MainchainTNT721Lock 
                            MainchainTFuelLock   
                            SubchainTNT20Burn   
                            SubchainTNT20Lock   
                            SubchainTNT721Burn  
                            SubchainTNT721Lock
                            SubchainTFuelBurn  

```

For help 

```
subchain_e2e_test_tools --help
```
