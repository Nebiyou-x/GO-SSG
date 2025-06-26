# ⛓️ Blockchain – Project Board (in Go)

## 📝 To Do

### 🧱 Core Blockchain Engine
- [ ] Define Block struct (PrevHash, Timestamp, Nonce, TxRoot)
- [ ] Create transaction struct (sender, receiver, amount, sig)
- [ ] Implement Merkle Tree builder (recursive SHA-256)
- [ ] Serialize/deserialize chain to disk

### ⚙️ Proof of Work
- [ ] Write PoW miner loop (adjustable difficulty, nonce grind)
- [ ] Add block reward (coinbase transaction)
- [ ] Implement difficulty adjustment algorithm

### 🌐 P2P Networking
- [ ] Basic TCP peer-to-peer messaging (`net`, `encoding/gob`)
- [ ] Gossip protocol: broadcast transactions and new blocks
- [ ] Implement sync logic (compare chains, request missing blocks)
- [ ] Handle forks (longest chain rule)

### 💳 Wallet & Signing
- [ ] Generate ECDSA keypairs (`crypto/ecdsa`, `elliptic.P256`)
- [ ] Sign and verify transactions
- [ ] Track balances using UTXO or account model

### 🧪 Testing & Interface
- [ ] CLI: mine, send, check balance
- [ ] Add unit tests for block validation and consensus
- [ ] Log node activity (new blocks, peers, forks)
- [ ] Write `README.md` with architecture and usage

---

## 🔧 In Progress
- [ ] Block and transaction struct implementation
- [ ] Merkle root hashing and serialization

---

## ✅ Done
<!-- Add completed tasks here -->
