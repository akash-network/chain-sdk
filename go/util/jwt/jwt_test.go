package jwt

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"text/template"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"

	jwttests "pkg.akt.dev/testdata/jwt"
)

type JWTTestSuite struct {
	IntegrationTestSuite
}

type jwtTestCase struct {
	Description string `json:"description"`
	Claims      Claims `json:"claims"`
	TokenString string `json:"tokenString"`
	Expected    struct {
		SignFail               bool   `json:"signFail"`
		VerifyFail             bool   `json:"verifyFail"`
		BypassSchemaValidation bool   `json:"bypassSchemaValidation"`
		Err                    string `json:"error"`
	} `json:"expected"`
}

type testTemplate struct {
	Issuer   string
	Provider string
	IatCurr  int64
	Iat24h   int64
	Nbf24h   int64
	NbfCurr  int64
	Exp48h   int64
}

func (s *JWTTestSuite) initClaims(tc jwtTestCase) jwtTestCase {
	var err error

	if tc.Claims.Issuer != "" {
		tc.Claims.iss, err = sdk.AccAddressFromBech32(tc.Claims.Issuer)
		require.NoError(s.T(), err)
	}

	if tc.TokenString != "" {
		return tc
	}

	ehdr := encodeSegment([]byte(`{"alg":"ES256K","typ":"JWT"}`))

	claims, err := json.Marshal(tc.Claims)
	require.NoError(s.T(), err)

	eclaims := encodeSegment(claims)
	data := ehdr + "." + eclaims

	method := jwt.GetSigningMethod("ES256K")
	sig, err := method.Sign(data, Signer{
		Signer: s.kr,
		addr:   s.addr,
	})

	require.NoError(s.T(), err)

	tc.TokenString = data + "." + encodeSegment(sig)

	return tc
}

func (s *JWTTestSuite) TestSigning() {
	testCases := s.prepareTestCases(s.T())

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			token := jwt.NewWithClaims(jwt.GetSigningMethod("ES256K"), tc.Claims)
			tokenString, err := token.SignedString(Signer{
				Signer: s.kr,
				addr:   s.addr,
			})

			if tc.Expected.SignFail {
				require.Error(s.T(), err)
				require.Empty(s.T(), tokenString)
			} else {
				require.NoError(s.T(), err)
				require.Equal(s.T(), tc.TokenString, tokenString)
			}

			if !tc.Expected.VerifyFail {
				claims := &Claims{}
				_, err := jwt.ParseWithClaims(tokenString, claims, func(_ *jwt.Token) (interface{}, error) {
					return s.pubKey, nil
				}, jwt.WithValidMethods([]string{"ES256K"}))

				require.Equal(s.T(), &tc.Claims, claims)
				require.NoError(s.T(), err)
			} else {
				claims := &Claims{}
				_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
					if token.Header["alg"] != "ES256K" {
						return nil, jwt.ErrInvalidKeyType
					}
					return s.pubKey, nil
				}, jwt.WithValidMethods([]string{"ES256K"}))

				require.ErrorContains(s.T(), err, tc.Expected.Err)
			}
		})
	}
}

// TestSchema ensure JSON schema is
func (s *JWTTestSuite) TestSchema() {
	testCases := s.prepareTestCases(s.T())

	for _, tc := range testCases {
		s.Run(tc.Description, func() {
			parts := strings.Split(tc.TokenString, ".")
			claims := decodeSegment(s.T(), parts[1])

			res, err := schemaLoader.Validate(gojsonschema.NewBytesLoader(claims))
			require.NotNil(s.T(), res)
			require.NoError(s.T(), err)
			if tc.Expected.VerifyFail && !tc.Expected.BypassSchemaValidation {
				require.False(s.T(), res.Valid())
				require.Greater(s.T(), len(res.Errors()), 0)
			} else {
				require.True(s.T(), res.Valid(), res.Errors())
				require.Len(s.T(), res.Errors(), 0)
			}
		})
	}
}

func (s *JWTTestSuite) prepareTestCases(t *testing.T) []jwtTestCase {
	t.Helper()
	var testCases []jwtTestCase

	data, err := jwttests.GetTestsFile("cases_jwt.json.tmpl")
	if err != nil {
		s.T().Fatalf("could not read test data file: %v", err)
	}

	now := time.Now()

	testTemplate := testTemplate{
		Issuer:   s.addr.String(),
		Provider: s.addr.String(),
		IatCurr:  jwt.NewNumericDate(now).Unix(),
		NbfCurr:  jwt.NewNumericDate(now).Unix(),
		Iat24h:   jwt.NewNumericDate(now.Add(24 * time.Hour)).Unix(),
		Nbf24h:   jwt.NewNumericDate(now.Add(24 * time.Hour)).Unix(),
		Exp48h:   jwt.NewNumericDate(now.Add(48 * time.Hour)).Unix(),
	}

	tmpl, err := template.New("tests").Parse(string(data))
	if err != nil {
		s.T().Fatalf("could not parse test data template: %v", err)
	}

	parsedTmpl := &bytes.Buffer{}
	err = tmpl.Execute(parsedTmpl, testTemplate)
	if err != nil {
		s.T().Fatalf("could not execute test data template: %v", err)
	}

	err = json.Unmarshal(parsedTmpl.Bytes(), &testCases)
	if err != nil {
		s.T().Fatalf("could not unmarshal test data: %v", err)
	}

	for i := range testCases {
		testCases[i] = s.initClaims(testCases[i])
	}

	return testCases
}

func TestGetSupportedScopes(t *testing.T) {
	scopes := GetSupportedScopes()

	expectedScopes := []PermissionScope{
		PermissionScopeSendManifest,
		PermissionScopeGetManifest,
		PermissionScopeLogs,
		PermissionScopeShell,
		PermissionScopeEvents,
		PermissionScopeStatus,
		PermissionScopeRestart,
		PermissionScopeHostnameMigrate,
		PermissionScopeIPMigrate,
	}

	require.Len(t, scopes, len(expectedScopes), "should return all expected scopes")

	for _, expectedScope := range expectedScopes {
		found := false
		for _, scope := range scopes {
			if scope == expectedScope {
				found = true
				break
			}
		}
		require.True(t, found, "expected scope %s should be present", expectedScope)
	}

	scopeMap := make(map[PermissionScope]bool)
	for _, scope := range scopes {
		require.False(t, scopeMap[scope], "duplicate scope found: %s", scope)
		scopeMap[scope] = true
	}
}
