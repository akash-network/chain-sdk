#!/usr/bin/env bash
set -euo pipefail

# Script to setup and fund test account for local testnet functional tests

AKASH_NODE="${AKASH_NODE:-http://localhost:26657}"
REST_API_URL="${REST_API_URL:-http://localhost:1317}"
GENESIS_ACCOUNT="${GENESIS_ACCOUNT:-main}"
MIN_BALANCE="${MIN_BALANCE:-100000000}"  # 100 AKT in uakt
CHAIN_ID="${CHAIN_ID:-local}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if node is running
check_node_running() {
    log_info "Checking if node is running at ${AKASH_NODE}..."
    if curl -s "${AKASH_NODE}/status" > /dev/null 2>&1; then
        log_info "Node is running"
        return 0
    else
        log_error "Node is not running at ${AKASH_NODE}"
        log_error "Please start the node first:"
        log_error "  make local-node-run"
        return 1
    fi
}

# Get test account address from mnemonic
get_test_account_address() {
    TEST_MNEMONIC="${TEST_MNEMONIC:-abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about}"
    
    # Use Node.js to derive address from mnemonic (using cosmjs)
    # This requires @cosmjs/proto-signing to be installed
    local script_dir
    script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    # Script is at ts/test/functional/setup-local-testnet.sh
    # So ts directory is two levels up: ts/test/functional/../.. = ts
    local ts_root
    ts_root="$(cd "${script_dir}/../.." && pwd)"
    
    # Change to ts directory where node_modules are located
    cd "${ts_root}" || return 1
    
    # Check if node_modules exists
    if [ ! -d "node_modules" ]; then
        log_error "node_modules not found in ${ts_root}. Please run 'npm install' first."
        return 1
    fi
    
    TEST_MNEMONIC="${TEST_MNEMONIC}" node -e "
    const { DirectSecp256k1HdWallet } = require('@cosmjs/proto-signing');
    (async () => {
        const wallet = await DirectSecp256k1HdWallet.fromMnemonic(
            process.env.TEST_MNEMONIC,
            { prefix: 'akash' }
        );
        const [account] = await wallet.getAccounts();
        console.log(account.address);
    })().catch(err => {
        console.error('Error:', err.message);
        process.exit(1);
    });
    " || return 1
}

# Check balance via REST API
get_balance() {
    local address="$1"
    local balance=0
    
    if response=$(curl -s "${REST_API_URL}/cosmos/bank/v1beta1/balances/${address}" 2>/dev/null); then
        if uakt_balance=$(echo "$response" | grep -o '"denom":"uakt"[^}]*"amount":"[^"]*"' | grep -o '"amount":"[^"]*"' | cut -d'"' -f4); then
            balance="${uakt_balance:-0}"
        fi
    fi
    
    echo "${balance}"
}

# Fund account from genesis account
fund_account() {
    local to_address="$1"
    local amount="${MIN_BALANCE}uakt"
    
    log_info "Funding test account ${to_address} with ${amount} from ${GENESIS_ACCOUNT}..."
    
    # Determine which akash command to use
    local AKASH_TO_USE
    if command -v akash >/dev/null 2>&1; then
        AKASH_TO_USE="akash"
    else
        log_error "akash command not found. Please ensure akash is in PATH or installed by 'make local-node-install-akash'."
        return 1
    fi
    
    # Get genesis account address
    local from_address
    if [ -n "${AKASH_HOME}" ]; then
        from_address=$("${AKASH_TO_USE}" keys show "${GENESIS_ACCOUNT}" -a --keyring-backend test --home "${AKASH_HOME}" 2>/dev/null || echo "")
    else
        from_address=$("${AKASH_TO_USE}" keys show "${GENESIS_ACCOUNT}" -a --keyring-backend test 2>/dev/null || echo "")
    fi
    
    if [ -z "${from_address}" ]; then
        log_error "Genesis account '${GENESIS_ACCOUNT}' not found. Make sure you've run 'make local-node-init'"
        return 1
    fi
    
    # Send tokens
    if [ -n "${AKASH_HOME}" ]; then
        if "${AKASH_TO_USE}" tx bank send "${from_address}" "${to_address}" "${amount}" \
            --chain-id "${CHAIN_ID}" \
            --node "${AKASH_NODE}" \
            --keyring-backend test \
            --from "${GENESIS_ACCOUNT}" \
            --yes \
            --gas auto \
            --gas-adjustment 2.0 \
            --gas-prices 0.025uakt \
            --home "${AKASH_HOME}" > /dev/null 2>&1; then
            log_info "Successfully funded account"
            log_info "Waiting for transaction to be included in a block..."
            sleep 3
            return 0
        else
            log_error "Failed to fund account"
            return 1
        fi
    else
        if "${AKASH_TO_USE}" tx bank send "${from_address}" "${to_address}" "${amount}" \
            --chain-id "${CHAIN_ID}" \
            --node "${AKASH_NODE}" \
            --keyring-backend test \
            --from "${GENESIS_ACCOUNT}" \
            --yes \
            --gas auto \
            --gas-adjustment 2.0 \
            --gas-prices 0.025uakt > /dev/null 2>&1; then
            log_info "Successfully funded account"
            log_info "Waiting for transaction to be included in a block..."
            sleep 3
            return 0
        else
            log_error "Failed to fund account"
            return 1
        fi
    fi
}

# Main function
main() {
    log_info "Setting up local testnet for functional tests..."
    
    # Check node is running
    if ! check_node_running; then
        exit 1
    fi
    
    # Get test account address
    log_info "Deriving test account address from TEST_MNEMONIC..."
    test_address=$(get_test_account_address)
    if [ -z "${test_address}" ]; then
        log_error "Failed to derive test account address"
        exit 1
    fi
    log_info "Test account address: ${test_address}"
    
    # Check balance
    log_info "Checking balance..."
    balance=$(get_balance "${test_address}")
    if [ -z "${balance}" ] || [ "${balance}" = "0" ]; then
        balance=0
    fi
    if command -v bc &> /dev/null && [ "${balance:-0}" -gt 0 ]; then
        balance_akt=$(echo "scale=2; ${balance}/1000000" | bc)
        log_info "Current balance: ${balance} uakt (${balance_akt} AKT)"
    else
        log_info "Current balance: ${balance} uakt"
    fi
    
    # Fund if needed
    if [ "${balance:-0}" -lt "${MIN_BALANCE}" ]; then
        log_warn "Balance is below minimum required (${MIN_BALANCE} uakt)"
        if ! fund_account "${test_address}"; then
            log_error "Failed to fund account. Please fund manually:"
            log_error "  akash tx bank send <genesis-account> ${test_address} ${MIN_BALANCE}uakt --chain-id ${CHAIN_ID} --node ${AKASH_NODE} --keyring-backend test"
            exit 1
        fi
        
        # Re-check balance
        sleep 2
        balance=$(get_balance "${test_address}")
        if [ -z "${balance}" ]; then
            balance=0
        fi
        if command -v bc &> /dev/null && [ "${balance}" -gt 0 ]; then
            balance_akt=$(echo "scale=2; ${balance}/1000000" | bc)
            log_info "New balance: ${balance} uakt (${balance_akt} AKT)"
        else
            log_info "New balance: ${balance} uakt"
        fi
    else
        log_info "Account has sufficient balance"
    fi
    
    log_info "Setup complete! Test account is ready."
}

main "$@"

