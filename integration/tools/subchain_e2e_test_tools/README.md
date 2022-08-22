For testing transfer cases:
subchain_e2e_test_tools     MainchainTNT20Burn   --amount=(the amount or tokenid you want to operate)
                            MainchainTNT20Lock  
                            MainchainTNT721Burn 
                            MainchainTNT721Lock 
                            MainchainTfuelBurn  
                            MainchainTfuelLock   
                            SubchainTNT20Burn   
                            SubchainTNT20Lock   
                            SubchainTNT721Burn  
                            SubchainTNT721Lock

To register subchain use :
subchain_e2e_test_tools  RegisterSubchain (it will register 360777)
To stake as validator:
subchain_e2e_test_tools  AcoountStake --id=(the id of accountList)

For help 
use --help