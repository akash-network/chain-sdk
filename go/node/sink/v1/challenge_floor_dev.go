//go:build sinkdev

package v1

// MinChallengePeriodSeconds lets local and team test networks, built with the
// sinkdev tag, run refunds in seconds. Release builds use 72 hours.
const MinChallengePeriodSeconds uint64 = 1
