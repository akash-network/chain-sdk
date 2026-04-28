import type { NetworkId } from "./types.ts";

/**
 * The network ID for the mainnet environment.
 */
export const MAINNET_ID = "mainnet" satisfies NetworkId;

/**
 * The network ID for the sandbox environment.
 */
export const SANDBOX_ID = "sandbox" satisfies NetworkId;

/**
 * The network ID for the testnet environment.
 */
export const TESTNET_ID = "testnet" satisfies NetworkId;

/**
 * The denomination for the AKT token.
 */
export const AKT_DENOM = "uakt";

/**
 * The denomination for ACT token.
 */
export const ACT_DENOM = "uact";
