# Custom Blockchain Build

## Project Description

This project aims to create a custom blockchain tailored for metadata use cases. The focus is on building a flexible and modular architecture that allows for easy customization and integration with various network protocols. The blockchain will support efficient metadata storage, retrieval, and management, making it ideal for applications that require robust and fast metadata tracing proof.

## Features

- Customizable server container modules
- Configurable transport layer for diverse network protocols
- Efficient metadata storage and retrieval mechanisms
- Modular design for easy integration and scalability
- Comprehensive testing suite for transport and network connectivity
- ...(in progress)

## Process

- Create server container with customisable modules
- Configurable transport layer to allow any type of network protocol
- Create block builder 
- Create tranaction object to append to a block
- Create private/public key generator for sign and verification
- ...(in progress)

## Tests

```shell
make test
```

## Flows

### Flowchart of Adding a New Block
```mermaid
flowchart TD
    Start --> CreateTransaction
    CreateTransaction --> ValidateTransaction
    ValidateTransaction --> |Valid| AddTransactionToBlock
    ValidateTransaction --> |Invalid| DiscardTransaction
    AddTransactionToBlock --> MineBlock
    MineBlock --> VerifyProofOfWork
    VerifyProofOfWork --> |Success| AddBlockToChain
    AddBlockToChain --> BroadcastBlock
    BroadcastBlock --> End
```

### Class Diagram of Blockchain Components
```mermaid
classDiagram
    class Blockchain {
        +[]Block chain
        +addBlock(Block)
        +getLatestBlock() Block
    }
    class Block {
        +int index
        +string timestamp
        +[]Transaction transactions
        +string previousHash
        +string hash
        +int nonce
        +calculateHash() string
    }
    class Transaction {
        +string sender
        +string recipient
        +float amount
        +string timestamp
    }
    Blockchain --> Block : contains
    Block --> Transaction : includes
```

### Sequence Diagram of Transaction Processing
```mermaid
sequenceDiagram
    participant User
    participant Node
    participant Miner
    User->>Node: Submit Transaction
    Node-->>Miner: Broadcast Transaction
    Miner->>Miner: Validate Transaction
    Miner->>Miner: Add to Block
    Miner->>Miner: Mine Block
    Miner-->>Node: New Block
    Node-->>User: Transaction Confirmed
```