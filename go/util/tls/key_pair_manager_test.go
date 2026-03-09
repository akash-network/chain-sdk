package tls

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func generateTestCertAndKey(t *testing.T) (*ecdsa.PrivateKey, []byte) {
	t.Helper()

	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "akash1testaddr"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour * 24 * 365),
		KeyUsage:     x509.KeyUsageDataEncipherment | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}

	certDer, err := x509.CreateCertificate(rand.Reader, &template, &template, priv.Public(), priv)
	if err != nil {
		t.Fatal(err)
	}

	return priv, certDer
}

func TestReadImpl_UnencryptedPKCS8(t *testing.T) {
	priv, certDer := generateTestCertAndKey(t)

	keyDer, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	_ = pem.Encode(&buf, &pem.Block{Type: "CERTIFICATE", Bytes: certDer})
	_ = pem.Encode(&buf, &pem.Block{Type: "PRIVATE KEY", Bytes: keyDer})

	kpm := &keyPairManager{
		addr: sdk.AccAddress("testaddr"),
	}

	cert, privKeyData, pubKey, err := kpm.readImpl(&buf)
	if err != nil {
		t.Fatalf("readImpl failed for unencrypted PKCS#8 key: %v", err)
	}
	if cert == nil {
		t.Fatal("expected non-nil cert")
	}
	if privKeyData == nil {
		t.Fatal("expected non-nil private key data")
	}
	if pubKey == nil {
		t.Fatal("expected non-nil public key")
	}
}

func TestReadImpl_SEC1ECPrivateKey(t *testing.T) {
	priv, certDer := generateTestCertAndKey(t)

	ecDer, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	_ = pem.Encode(&buf, &pem.Block{Type: "CERTIFICATE", Bytes: certDer})
	_ = pem.Encode(&buf, &pem.Block{Type: "EC PRIVATE KEY", Bytes: ecDer})

	kpm := &keyPairManager{
		addr: sdk.AccAddress("testaddr"),
	}

	cert, privKeyData, pubKey, err := kpm.readImpl(&buf)
	if err != nil {
		t.Fatalf("readImpl failed for SEC 1 EC private key: %v", err)
	}
	if cert == nil {
		t.Fatal("expected non-nil cert")
	}
	if privKeyData == nil {
		t.Fatal("expected non-nil private key data")
	}
	if pubKey == nil {
		t.Fatal("expected non-nil public key")
	}
}

func TestReadImpl_UnknownPEMType(t *testing.T) {
	_, certDer := generateTestCertAndKey(t)

	var buf bytes.Buffer
	_ = pem.Encode(&buf, &pem.Block{Type: "CERTIFICATE", Bytes: certDer})
	_ = pem.Encode(&buf, &pem.Block{Type: "SOME UNKNOWN KEY", Bytes: []byte("garbage")})

	kpm := &keyPairManager{
		addr: sdk.AccAddress("testaddr"),
	}

	_, _, _, err := kpm.readImpl(&buf)
	if err == nil {
		t.Fatal("expected error for unknown PEM type, got nil")
	}
}
