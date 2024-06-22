# blockchain
this repo to demo how blockchain working


# Run the program using:
```go run main.go blockchain.go block.go node.go transaction.go```

## Creating a more complex example of a blockchain with multiple instances involves setting up a peer-to-peer network where each node (instance) can add blocks and synchronize with others. This example will include:

- A `Block` structure to represent individual blocks.
- A `Blockchain` structure to manage the chain of blocks.
- A `Node` structure to represent each instance in the network.
- Networking code to allow nodes to communicate, synchronize, and broadcast new blocks.

Step-by-Step Guide: 
- Define the Block Structure.
- Define the Blockchain Structure.
- Define the Node Structure.
- Implement Networking for Synchronization.
- Run Multiple Instances.


To ensure that the data in Node 1 and Node 2 are synchronized, we'll need to add a mechanism to request the blockchain data from one node and update the other node with this data. This process is typically referred to as blockchain synchronization.

Step-by-Step Guide to Synchronize Nodes
- Update the Node Struct to Handle Synchronization Requests
- Implement a Synchronization Mechanism
- Modify the Main Function to Synchronize Data

Explanation
- Node Structure:

    - Nodes handle synchronization requests (sync message) by sending their entire blockchain to the requester.
    - Nodes can request synchronization from their peers using the RequestSync method.
    - Nodes can update their blockchain using the Update method, which ensures the new blockchain is valid and longer than the current one.

- Blockchain Structure:

    - The blockchain can be serialized and deserialized for transmission between nodes.
    - The Update method ensures the new blockchain is valid before accepting it.

- Main Function:

    - Node 1 adds a block.
    - Node 2 requests synchronization from Node 1.
    - Node 2 updates its blockchain with the data received from Node 1 and validates it.

## Send Btc to others
### Explanation
    Transaction Structure:

    - Transactions represent the transfer of BTC from one user to another.
    - Blocks contain multiple transactions.
    
    Node Structure:

    - Nodes handle synchronization requests (sync message) by sending their entire blockchain to the requester.
    - Nodes can request synchronization from their peers using the RequestSync method.
    - Nodes can update their blockchain using the Update method, which ensures the new blockchain is valid and longer than the current one.
    
    Blockchain Structure:

    - The blockchain can be serialized and deserialized for transmission between nodes.
    - The Update method ensures the new blockchain is valid before accepting it.
    
    Main Function:

    - Node 1 adds a block with multiple transactions.
    - Node 2 requests synchronization from Node 1.
    - Node 2 updates its blockchain with the data received from Node 1 and validates it.

```
Transaction: Alice -> Bob: 1.000000 BTC
Transaction: Bob -> Charlie: 2.000000 BTC
Transaction: Charlie -> Dave: 3.000000 BTC
```

To extend the example where Charlie sends BTC to Mike using Node 2, we'll follow these steps:

- Node 1 creates a block with multiple transactions.
- Node 2 synchronizes with Node 1.
- Node 2 creates a new transaction from Charlie to Mike and adds it to its blockchain.
- Node 1 synchronizes with Node 2 to get the new transaction.

```
    Explanation:
    - Initial Setup: Nodes are created, started, and peers are added.
    - Node 1 Adds a Block: Node 1 adds a block with multiple transactions.
    - Node 2 Synchronizes: Node 2 synchronizes with Node 1 to get the latest blockchain state.
    - Node 2 Adds a New Transaction: Node 2 adds a new transaction from Charlie to Mike.
    - Node 1 Synchronizes: Node 1 synchronizes with Node 2 to get the new transaction.
    - Validation: Both Node 1 and Node 2 validate their blockchain.
```

Output in terminal : 
```
    Node 1 adds a block with multiple transactions
    Node 2 requests synchronization
    Node 2 Blockchain after synchronization:
    Prev. hash:
    Transaction: genesis -> genesis: 0.000000 BTC
    Hash: 84f6f96975c97a1f1362561be21d7138ce3d2edf61bc7382a136abf2fe0e42a2

    Prev. hash: 84f6f96975c97a1f1362561be21d7138ce3d2edf61bc7382a136abf2fe0e42a2
    Transaction: Alice -> Bob: 1.000000 BTC
    Transaction: Bob -> Charlie: 2.000000 BTC
    Transaction: Charlie -> Dave: 3.000000 BTC
    Hash: 0be159e141874dc40ed36d662d6af2929b347f80e16f489290449ec576640636

    Node 2 Blockchain is valid
    Node 2 adds a new block with a transaction from Charlie to Mike
    Node 1 requests synchronization
    Node 1 Blockchain after synchronization:
    Prev. hash:
    Transaction: genesis -> genesis: 0.000000 BTC
    Hash: 84f6f96975c97a1f1362561be21d7138ce3d2edf61bc7382a136abf2fe0e42a2

    Prev. hash: 84f6f96975c97a1f1362561be21d7138ce3d2edf61bc7382a136abf2fe0e42a2
    Transaction: Alice -> Bob: 1.000000 BTC
    Transaction: Bob -> Charlie: 2.000000 BTC
    Transaction: Charlie -> Dave: 3.000000 BTC
    Hash: 0be159e141874dc40ed36d662d6af2929b347f80e16f489290449ec576640636

    Prev. hash: 0be159e141874dc40ed36d662d6af2929b347f80e16f489290449ec576640636
    Transaction: Charlie -> Mike: 1.500000 BTC
    Hash: e8688ee730b8c12a30e568b460b58cdc34c94263708f6e7ce9dad3ceaf2a6677

    Node 1 Blockchain is valid
```

## Update Balance

To handle balance updates when transactions occur, we'll need to add account balances to our blockchain nodes and ensure they are updated correctly when transactions are processed. Here's how we can achieve this:

- Add a Balance field to track account balances.
- Update balances during transactions.
- Ensure synchronization reflects updated balances.

```
Explanation: 
    1. Initialization:
        - Nodes are created, started, and peers are added.
        - Balances are initialized for each user.

    2. Node 1 Adds a Block:
        - Node 1 adds a block with multiple transactions and logs all transactions in the blockchain.

    3. Node 2 Synchronizes:
        - Node 2 synchronizes with Node 1 and logs all transactions and balances.
    
    4. Node 2 Adds a New Transaction:
        - Node 2 adds a new transaction from Charlie to Mike.

    5. Node 1 Synchronizes:
        - Node 1 synchronizes with Node 2 and logs the new transaction added along with the previous and current hash.

    6. Balances Logging:
        - After each synchronization, balances are printed to show the updated state.
```

```
Expected Output
You should see the following output, demonstrating that Node 2 synchronized its blockchain with Node 1, added a new transaction, and Node 1 synchronized back to include the new transaction while printing the new incoming transaction along with the previous and current hashes and updated balances:

    Node 1 adds a block with multiple transactions
    Node 2 requests synchronization
    Node 2 Blockchain after synchronization:
    Prev. hash: 
    Transaction: genesis -> genesis: 0.000000 BTC
    Hash: 79054025255fb1a26e4bc422aef54eb4

    Prev. hash: 79054025255fb1a26e4bc422aef54eb4
    Transaction: Alice -> Bob: 1.000000 BTC
    Transaction: Bob -> Charlie: 2.000000 BTC
    Transaction: Charlie -> Dave: 3.000000 BTC
    Hash: 04b7fbae83eb0a66d2d3f9470d0de16f7a4b2e7f2d41bdb99a76b7a8a450d24b

    Balances after first synchronization:
    Alice: 9.000000 BTC
    Bob: 6.000000 BTC
    Charlie: 5.000000 BTC
    Dave: 5.000000 BTC

    Node 2 Blockchain is valid
    Node 2 adds a new block with a transaction from Charlie to Mike
    Node 1 requests synchronization
    Incoming next transaction in Node 1:
    Prev. hash: 04b7fbae83eb0a66d2d3f9470d0de16f7a4b2e7f2d41bdb99a76b7a8a450d24b
    Transaction: Charlie -> Mike: 1.500000 BTC
    Hash: <new block hash>

    Balances after second synchronization:
    Alice: 9.000000 BTC
    Bob: 6.000000 BTC
    Charlie: 3.500000 BTC
    Dave: 5.000000 BTC
    Mike: 1.500000 BTC
```