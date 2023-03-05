# Settings on chain

## Setting examples for admin

### set protocol fee receiver

```bash
sojahubd tx ledger set-protocol-fee-receiver did:fury:1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --from admin --chain-id local-sojahub --keyring-backend file

sojahubd query ledger protocol-fee-receiver 
```

### add new rtoken

```bash
# set rtoken metadata
sojahubd tx rbank add-denom cosmos cosmosvaloper ./metadata/metadata_ratom.json --chain-id local-sojahub --from admin --keyring-backend file

sojahubd query bank denom-metadata

sojahubd query rbank address-prefix uratom

# set relay fee receiver
sojahubd tx ledger set-relay-fee-receiver uratom did:fury:1mgjkpyfm00mxk0nmhvfvwhlr65067d538l6cxa --from admin --chain-id local-sojahub --keyring-backend file

sojahubd query ledger relay-fee-receiver uratom

# this will init bonded pool, exchange rate, pipeline
sojahubd tx ledger init-pool uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 --from admin --chain-id local-sojahub --keyring-backend file

sojahubd query ledger bonded-pools uratom

sojahubd query ledger exchange-rate uratom

sojahubd query ledger bond-pipeline uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



# add relayers
sojahubd tx relayers add-relayers ledger uratom did:fury:1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx:did:fury:1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --from admin --chain-id local-sojahub

sojahubd query relayers relayers ledger uratom

# set threshold
sojahubd tx relayers set-threshold ledger uratom 1 --from admin --keyring-backend file --chain-id local-sojahub

sojahubd query relayers threshold ledger uratom

# set params used by relay
sojahubd tx ledger set-r-params uratom 0.00001stake 600 0 2 0stake --from admin --keyring-backend file --chain-id local-sojahub

sojahubd query ledger r-params uratom

# set pool detail for multisig/ica pool
sojahubd tx ledger set-pool-detail uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmos1cad0efr25faywnjp8qp36l8zlqa2sgz0jwn0hl:cosmos13mwxtgrljf9d5r72sc28496ua4lsga0jvmqz8x 1 --from admin --chain-id local-sojahub --keyring-backend file

sojahubd query ledger pool-detail uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



# default 0.1
sojahubd tx ledger set-staking-reward-commission uratom 0.15 --from admin --chain-id local-sojahub --keyring-backend file

sojahubd q ledger staking-reward-commission uratom

# default 0.002
sojahubd tx ledger set-unbond-commission uratom 0.0025 --from admin --chain-id local-sojahub --keyring-backend file

sojahubd q ledger unbond-commission uratom

# default 1000000ufury
sojahubd tx ledger set-unbond-relay-fee uratom 1000005ufury --from admin --chain-id local-sojahub --keyring-backend file

sojahubd q ledger unbond-relay-fee uratom

```

### register ica pool
```
# register ica pool (need set rtoken metadata before this)
sojahubd tx ledger register-ica-pool uratom connection-0 --keyring-backend file --from admin --chain-id local-sojahub --gas 410000

sojahubd q ledger ica-pool-list uratom

# set withdrawal address
sojahubd tx ledger set-withdrawal-addr cosmos1gsth46z50w256p4kq36xquh4q90mfjq0t4lm9scln6zucg64epyqudzqzm --keyring-backend file --from admin --chain-id local-sojahub --gas 410000

```

### rvalidator

```bash
# add relayers
sojahubd tx relayers add-relayers rvalidator uratom did:fury:14z467aut40mcrt2ddyxf7e74fq99udul7kaf9g:did:fury:15lne70yk254s0pm2da6g59r82cjymzjqvvqxz7 --keyring-backend file --from admin --chain-id local-sojahub

sojahubd q relayers relayers rvalidator uratom

# set threshold
sojahubd tx relayers set-threshold rvalidator uratom 1 --from admin --keyring-backend file --chain-id local-sojahub

sojahubd q relayers threshold rvalidator uratom

# init rvalidator (should init target validators of pool before rtoken relay start)
sojahubd tx rvalidator init-r-validator uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9npre --from admin --chain-id local-sojahub --keyring-backend file  

sojahubd q rvalidator r-validator-list uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75

# add rvalidator
sojahubd tx rvalidator add-r-validator uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper1cad0efr25faywnjp8qp36l8zlqa2sgz0h686mv  --chain-id local-sojahub --keyring-backend file --from admin

sojahubd q rvalidator r-validator-list uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75

# rm rvalidator
sojahubd tx rvalidator rm-r-validator uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper1cad0efr25faywnjp8qp36l8zlqa2sgz0h686mv cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9npre --from admin --chain-id local-sojahub --keyring-backend file
```



### bridge

```bash
sojahubd tx bridge add-chain-id 1 --from admin --keyring-backend file --chain-id local-sojahub

sojahubd query bridge chain-ids



sojahubd tx relayers add-relayers bridge 1 did:fury:1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-sojahub

sojahubd query relayers relayers bridge 1



sojahubd tx relayers set-threshold bridge 1 1 --from admin --keyring-backend file --chain-id local-sojahub

sojahubd query relayers threshold bridge 1



sojahubd tx bridge set-resourceid-to-denom  000000000000000000000000000000a9e0095b8965c01e6a09c97938f3860901 uratom NATIVE --from admin --keyring-backend file --chain-id local-sojahub

sojahubd query bridge resourceid-to-denoms



sojahubd tx bridge set-relay-fee-receiver did:fury:1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-sojahub

sojahubd query bridge relay-fee-receiver



sojahubd tx bridge set-relay-fee 1 1000000ufury --from admin --keyring-backend file --chain-id local-sojahub

sojahubd query bridge  relay-fee 1


sojahubd tx bridge add-banned-denom 1 uratom --from admin --keyring-backend file --chain-id local-sojahub

sojahubd q bridge banned-denom-list
```

### migrate rtoken (after adding new rtoken step)

```bash
sojahubd tx ledger migrate-init uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100000000 150000000 200000000 300000000 1.23 --from admin --keyring-backend file --chain-id local-sojahub

sojahubd query bank  total 

sojahubd query ledger exchange-rate uratom

sojahubd query ledger bond-pipeline uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



sojahubd tx ledger migrate-unbondings uratom --unbondings ./unbondings_example.json --from admin --keyring-backend file --chain-id local-sojahub

sojahubd query ledger pool-unbond uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 3



sojahubd tx bridge set-denom-type uratom  1 --from admin --keyring-backend file --chain-id local-sojahub

sojahubd query bridge denom-types
```

### rdex

```bash
sojahubd tx rdex create-pool 10ufury 20uratom --from admin --chain-id local-sojahub --keyring-backend file

sojahubd tx rdex add-provider did:fury:1qzt0qajzr9df3en5sk06xlk26n30003c8uhdkg --from admin --chain-id local-sojahub --keyring-backend file

sojahubd tx rdex add-liquidity  100ufury 200uratom --from admin --chain-id local-sojahub --keyring-backend file

sojahubd tx rdex remove-liquidity 10 5 1uratom 1ufury ufury --from admin --chain-id local-sojahub --keyring-backend file

sojahubd tx rdex swap 2ufury 1uratom  --from admin --chain-id local-sojahub --keyring-backend file
```

### mining

```bash
sojahubd tx mining add-mining-provider did:fury:1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx  --from admin --chain-id local-sojahub --keyring-backend file

sojahubd tx mining add-reward-token ufury 200 --from admin --chain-id local-sojahub --keyring-backend file


sojahubd tx mining add-stake-pool ufury ./add_stake_pool_example.json  --from relay1 --chain-id local-sojahub --keyring-backend file

sojahubd tx mining stake 0 10ufury 0 --from my-account --chain-id local-sojahub --keyring-backend file 

sojahubd tx mining claim-reward 0 0 --from my-account --chain-id local-sojahub --keyring-backend file

sojahubd tx mining add-reward 1 0 300 0 0 --from relay1 --chain-id local-sojahub --keyring-backend file

sojahubd tx mining withdraw 1 10ufury 0 --from test --chain-id local-sojahub --keyring-backend file
```



## Operate examples for user

### liquidity bond (gaiad example)

```bash
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1000000stake --note 1:did:fury:1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --chain-id local-cosmos
```

### recover (gaiad example)

```bash
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1stake --note 2:did:fury:1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc:9A80F3E6A007E1144BE34F4A0AC35B9288C19641BCAD3464277168000AF5FC66 --keyring-backend file --chain-id local-cosmos
```

### liquidity unbond

```bash
sojahubd tx ledger liquidity-unbond cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100uratom cosmos1j9dues7ey2a39nes4ewfvyma96d3f5zrdhnfan --keyring-backend file --from user --home /Users/tpkeeper/gowork/soja/rtoken-relay-core/keys/sojahub --chain-id local-sojahub
```

### deposit (transfer token to external chain)

```bash
sojahubd tx bridge deposit 1 uratom 800 dccf954570063847d73746afa0b0878f2c779d42089c5d9a107f2aca176e985f --from my-account --chain-id local-sojahub --keyring-backend file
```
