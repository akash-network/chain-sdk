# Akash Migration CLI Guide: v0.38.x → v1.0.0

This guide helps you migrate from Akash v0.38.x to v1.0.0. The main changes involve CLI command updates and new API structures.

## 🚀 Quick Overview

The v1.0.0 release introduces several breaking changes to improve the user experience and security:
- **Deposit**: Replaced `--depositor` flag with `--deposit-sources`
- **Authorization**: Moved from deployment-specific authz to standard Cosmos authz
- **Escrow operations**: Consolidated under the `escrow` module

## 📋 CLI Changes

### Transaction Commands

#### 1. Deployment Creation

**What Changed**: The `--depositor` flag has been replaced with `--deposit-sources`

**Before (v0.38.x)**:
```bash
akash tx deployment create --depositor=<granter_address>
```

**After (v1.0.0)**:
```bash
akash tx deployment create --deposit-sources=grant
```

**Options for `--deposit-sources`**:
- `grant` - Use funds from authorization grants
- `balance` - Use funds from your account balance
- `grant,balance` - Try grants first, then take balance (default)

---

#### 2. Market Bid Creation

**What Changed**: Same as deployment creation - `--depositor` replaced with `--deposit-sources`

**Before (v0.38.x)**:
```bash
akash tx market bid create --depositor=<granter_address>
```

**After (v1.0.0)**:
```bash
akash tx market bid create --deposit-sources=grant
```

---

#### 3. Authorization Grants

**What Changed**: Deployment-specific authz commands moved to standard Cosmos authz

**Before (v0.38.x)**:
```bash
akash tx deployment authz grant <grantee_address> 50akt
```

**After (v1.0.0)**:
```bash
akash tx authz grant <grantee_address> deposit \
  --spend-limit=50akt \
  --scope=deployment
```

**Key Differences**:
- Use `tx authz grant` instead of `tx deployment authz grant`
- Specify `deposit` as the message type
- Add `--scope=deployment` to restrict the grant to deployment operations only

---

#### 4. Authorization Revocation

**What Changed**: Updated to use the new escrow message path

**Before (v0.38.x)**:
```bash
akash tx deployment authz revoke <grantee_address>
```

**After (v1.0.0)**:
```bash
akash tx authz revoke <akash1…> /akash.escrow.v1.MsgAccountDeposit
```

**Note**: The new path `/akash.escrow.v1.MsgAccountDeposit` specifically targets deposit-related permissions.

---

#### 5. Deployment Deposits

**What Changed**: Moved from `deployment` module to `escrow` module

**Before (v0.38.x)**:
```bash
akash tx deployment deposit 5000000uakt \
  --dseq=<deployment_sequence> \
  --from=<owner_address>
```

**After (v1.0.0)**:
```bash
akash tx escrow deposit deployment 5000000uakt \
  --dseq=<deployment_sequence> \
  --from=<owner_address>
```

**Key Change**: Use `tx escrow deposit deployment` instead of `tx deployment deposit`

---

### New Escrow Commands

#### Transaction Commands

**`tx escrow deposit deployment [amount]`**
- **Purpose**: Deposit funds to an escrow account for a deployment
- **Usage**: `akash tx escrow deposit deployment 5000000uakt --dseq=<dseq> --from=<owner>`
- **Flags**:
  - `--dseq` - Deployment sequence number
  - `--gseq` - Group sequence number (optional)
  - `--oseq` - Order sequence number (optional)
  - `--owner` - Owner address (optional, defaults to signer)

#### Query Commands

**`query escrow accounts [state] [xid]`**
- **Purpose**: Query escrow accounts with optional filtering
- **Usage Examples**:
  ```bash
  # Query all accounts
  akash query escrow accounts

  # Query accounts in specific state
  akash query escrow accounts open

  # Query accounts for deployment scope
  akash query escrow accounts open deployment

  # Query specific deployment account
  akash query escrow accounts open deployment/akash1.../123
  ```
- **States**: `open`, `closed`, `overdrawn`
- **Scopes**: `deployment`, `bid`

**`query escrow payments [state] [xid]`**
- **Purpose**: Query escrow payments with optional filtering
- **Usage**: Similar to accounts command but for payment records
- **States**: `open`, `closed`, `overdrawn`

**`query escrow blocks-remaining`**
- **Purpose**: Calculate remaining blocks for an escrow account
- **Usage**: `akash query escrow blocks-remaining --owner=<owner> --dseq=<dseq>`
- **Output**: Shows balance remaining, blocks remaining, and estimated time remaining
- **Required Flags**: `--dseq` (deployment sequence), `--owner` (deployment owner)

---

## 🏗️ Genesis Command Changes

### What Changed

In v1.0.0, all initialization and genesis-related commands have been consolidated under the `genesis` command group. This provides better organization and follows Cosmos SDK conventions.

### Command Structure Changes

**Before (v0.38.x)**:
```bash
# Individual init commands
akash init <moniker>
akash gentx <key_name> <amount>
akash add-genesis-account <address> <coins>
akash collect-gentxs
akash validate-genesis
```

**After (v1.0.0)**:
```bash
# All commands now under genesis
akash genesis init <moniker>
akash genesis gentx <key_name> <amount>
akash genesis add-account <address> <coins>
akash genesis collect
akash genesis validate [file]
```

### Detailed Command Changes

#### 1. Node Initialization

**Before (v0.38.x)**:
```bash
akash init <moniker> --chain-id=<chain_id>
```

**After (v1.0.0)**:
```bash
akash genesis init <moniker> --chain-id=<chain_id>
```

**New Flags Available**:
- `--init-height` - Specify initial block height (default: 1)
- `--consensus-key-algo` - Algorithm for consensus key (default: ed25519)
- `--recover` - Recover existing key using mnemonic

---

#### 2. Genesis Transaction Generation

**Before (v0.38.x)**:
```bash
akash gentx <key_name> <amount> --chain-id=<chain_id>
```

**After (v1.0.0)**:
```bash
akash genesis gentx <key_name> <amount> --chain-id=<chain_id>
```

**Key Differences**:
- Command moved under `genesis` subcommand
- Same functionality and flags maintained
- Better integration with genesis workflow

---

#### 3. Genesis Account Addition

**Before (v0.38.x)**:
```bash
akash add-genesis-account <address> <coins>
```

**After (v1.0.0)**:
```bash
akash genesis add-account <address> <coins>
```

**Enhanced Features**:
- Support for vesting accounts with `--vesting-amt`, `--vesting-start`, `--vesting-end`
- Module account support with `--module-name`
- Append mode with `--append` flag

---

#### 4. Genesis Transaction Collection

**Before (v0.38.x)**:
```bash
akash collect-gentxs
```

**After (v1.0.0)**:
```bash
akash genesis collect
```

**Improvements**:
- Better error handling and validation
- Enhanced output formatting
- Integration with genesis validation

---

#### 5. Genesis Validation

**Before (v0.38.x)**:
```bash
akash validate-genesis
```

**After (v1.0.0)**:
```bash
akash genesis validate [file]
```

**Enhanced Validation**:
- Optional file path argument for custom genesis files
- Better error messages with upgrade guidance
- CometBFT consensus parameter validation

---

### Migration Steps for Genesis Commands

1. **Update Scripts**: Replace all standalone init commands with `genesis` prefixed versions
2. **Update Documentation**: Update any documentation or scripts that reference old command paths
3. **Test Workflows**: Verify that genesis creation and validation workflows still function correctly
4. **Update CI/CD**: Update any automated deployment scripts that use genesis commands

### Example Migration Workflow

**Before (v0.38.x)**:
```bash
# Initialize node
akash init mynode --chain-id=testnet-1

# Add genesis account
akash add-genesis-account akash1... 1000000uakt

# Generate genesis transaction
akash gentx mykey 1000000uakt --chain-id=testnet-1

# Collect genesis transactions
akash collect-gentxs

# Validate genesis
akash validate-genesis
```

**After (v1.0.0)**:
```bash
# Initialize node
akash genesis init mynode --chain-id=testnet-1

# Add genesis account
akash genesis add-account akash1... 1000000uakt

# Generate genesis transaction
akash genesis gentx mykey 1000000uakt --chain-id=testnet-1

# Collect genesis transactions
akash genesis collect

# Validate genesis
akash genesis validate
```

---

## ❓ Common Issues & Solutions

### Issue: "Unknown flag --depositor"
**Solution**: Replace with `--deposit-sources=grant` or `--deposit-sources=balance`

### Issue: "Command not found: tx deployment authz"
**Solution**: Use `tx authz grant` and `tx authz revoke` instead

### Issue: "Module not found: deployment deposit"
**Solution**: Use `tx escrow deposit deployment` instead

### Issue: "Invalid account scope" in escrow commands
**Solution**: Use `deployment` as the scope for deployment-related escrow operations

### Issue: "Command not found: init"
**Solution**: Use `genesis init` instead

### Issue: "Command not found: gentx"
**Solution**: Use `genesis gentx` instead

### Issue: "Command not found: add-genesis-account"
**Solution**: Use `genesis add-account` instead

---

## 📚 Additional Resources

- [Akash v1.0.0 Release Notes](https://github.com/akash-network/akash/releases)
- [Cosmos SDK Authz Module Documentation](https://docs.cosmos.network/v0.50/modules/authz)
- [Akash Provider Documentation](https://docs.akash.network/)

---

## 🆘 Need Help?

If you encounter issues during migration:
1. Check the [Akash Discord](https://discord.gg/akash)
2. Review [GitHub Issues](https://github.com/akash-network/akash/issues)
3. Consult the [Akash Documentation](https://docs.akash.network/)
