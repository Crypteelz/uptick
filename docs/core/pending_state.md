<!--
order: 2
-->

# Pending State

Learn how Uptick handles pending state queries. {synopsis}

## Pre-requisite Readings

- [Tendermint Mempool](https://docs.tendermint.com/master/tendermint-core/mempool.htm) {prereq}

## Uptick vs Ethereum

In Ethereum, pending blocks are generated as they are queued for production by miners. These pending
blocks include pending transactions that are picked out by miners, based on the highest reward paid
in gas. This mechanism exists as block finality is not possible on the Ethereum network. Blocks are
committed with probabilistic finality, which means that transactions and blocks become less likely
to become reverted as more time (and blocks) passes.

Uptick is designed quite differently on this front as there is no concept of a "pending state".
Uptick uses [Tendermint Core](https://docs.tendermint.com/) BFT consensus which provides instant
finality for transaction. For this reason, Etheremint does not require a pending state mechanism, as
all (if not most) of the transactions will be committed to the next block (avg. block time on Cosmos chains is ~8s). However, this causes a
few hiccups in terms of the Ethereum Web3-compatible queries that can be made to pending state.

Another significant difference with Ethereum, is that blocks are produced by validators or block producers, who include transactions from their local mempool into blocks in a
first-in-first-out (FIFO) fashion. Transactions on Uptick cannot be ordered or cherry picked out from the Tendermint node [mempool](https://docs.tendermint.com/master/tendermint-core/mempool.html#transaction-ordering).

## Pending State Queries

Uptick will make queries which will account for any unconfirmed transactions present in a node's
transaction mempool. A pending state query made will be subjective and the query will be made on the
target node's mempool. Thus, the pending state will not be the same for the same query to two
different nodes.

### JSON-RPC Calls on Pending Transactions

- [`eth_getBalance`](./../api/json-rpc/endpoints.md#eth_getbalance)
- [`eth_getTransactionCount`](./../api/json-rpc/endpoints.md#eth-gettransactioncount)
- [`eth_getBlockTransactionCountByNumber`](./../api/json-rpc/endpoints.md#eth-getblocktransactioncountbynumber)
- [`eth_getBlockByNumber`](./../api/json-rpc/endpoints.md#eth-getblockbynumber)
- [`eth_getTransactionByHash`](./../api/json-rpc/endpoints.md#eth-gettransactionbyhash)
- [`eth_getTransactionByBlockNumberAndIndex`](./../api/json-rpc/endpoints.html#eth-gettransactionbyblockhashandindex)
- [`eth_sendTransaction`](./../api/json-rpc/endpoints.md#eth-sendtransaction)

## Next {hide}

Learn how to deploy a Solidity smart contract on Uptick using [Truffle](./../guides/truffle.md) {hide}