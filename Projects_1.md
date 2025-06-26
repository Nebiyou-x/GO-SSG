🛡️ Password Manager with Zero-Knowledge-Inspired Proof (Go)
markdown
Copy
Edit
# 🔐 Password Manager – Project Board (in Go)

## 📝 To Do

### 🔑 Encryption & Key Derivation
- [ ] Use Argon2 to derive key from master password (`x/crypto/argon2`)
- [ ] Encrypt secrets using AES-GCM (`crypto/aes`, `crypto/cipher`)
- [ ] Generate secure random salt & IV (`crypto/rand`)
- [ ] Implement secure challenge–response proof for password check

### 💾 Local Vault Storage
- [ ] Create vault format (e.g., encrypted `vault.json`)
- [ ] Implement add/edit/delete secret CLI commands
- [ ] Implement vault lock + timeout

### ☁️ Encrypted Cloud Sync
- [ ] Serialize and encrypt entire vault
- [ ] Upload to cloud storage API (e.g., Dropbox, or REST backend)
- [ ] Implement auto-sync and conflict resolution

### 🧪 Testing & Docs
- [ ] Write tests for encryption/decryption
- [ ] Test Argon2 behavior with different params
- [ ] Document file format and security model
- [ ] Write `README.md` with setup, usage, and dev guide

---

## 🔧 In Progress
- [ ] Vault format + encryption pipeline
- [ ] CLI: add/edit/list secrets

---

## ✅ Done
<!-- Add completed tasks here -->
