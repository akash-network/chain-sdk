package v1

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protowire"
)

func TestAttestationGPUReportBinaryRoundTrip(t *testing.T) {
	expected := AttestationGPUReport{
		DeviceIndex:       7,
		Report:            "legacy-aggregate",
		AttestationReport: "gpu-attestation-report",
		CecReport:         "cec-attestation-report",
		CertificateChain:  "device-certificate-chain",
	}

	data, err := proto.Marshal(&expected)
	require.NoError(t, err)

	fields := protobufFieldNumbers(t, data)
	require.ElementsMatch(t, []protowire.Number{1, 2, 3, 4, 5}, fields)

	var actual AttestationGPUReport
	require.NoError(t, proto.Unmarshal(data, &actual))
	require.Equal(t, expected, actual)
}

func TestAttestationGPUReportJSONRoundTrip(t *testing.T) {
	expected := AttestationGPUReport{
		DeviceIndex:       7,
		Report:            "legacy-aggregate",
		AttestationReport: "gpu-attestation-report",
		CecReport:         "cec-attestation-report",
		CertificateChain:  "device-certificate-chain",
	}

	data, err := json.Marshal(&expected)
	require.NoError(t, err)

	var fields map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(data, &fields))
	require.Contains(t, fields, "report")
	require.Contains(t, fields, "attestation_report")
	require.Contains(t, fields, "cec_report")
	require.Contains(t, fields, "certificate_chain")
	require.NotContains(t, fields, "attestationReport")
	require.NotContains(t, fields, "cecReport")
	require.NotContains(t, fields, "certificateChain")

	var actual AttestationGPUReport
	require.NoError(t, json.Unmarshal(data, &actual))
	require.Equal(t, expected, actual)
}

func protobufFieldNumbers(t *testing.T, data []byte) []protowire.Number {
	t.Helper()

	var fields []protowire.Number
	for len(data) > 0 {
		number, wireType, count := protowire.ConsumeTag(data)
		require.GreaterOrEqual(t, count, 0)
		data = data[count:]
		fields = append(fields, number)

		count = protowire.ConsumeFieldValue(number, wireType, data)
		require.GreaterOrEqual(t, count, 0)
		data = data[count:]
	}

	return fields
}
