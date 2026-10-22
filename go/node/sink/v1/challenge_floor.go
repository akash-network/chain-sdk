//go:build !sinkdev

package v1

// MinChallengePeriodSeconds is the shortest void challenge window parameters
// may set. AEP-92 sizes the defense against a false void quorum on 72 hours.
const MinChallengePeriodSeconds uint64 = 72 * 60 * 60
