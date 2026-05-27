package sdl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	manifest "pkg.akt.dev/go/manifest/v2beta3"
	dtypes "pkg.akt.dev/go/node/deployment/v1beta4"
	atypes "pkg.akt.dev/go/node/types/attributes/v1"
	rtypes "pkg.akt.dev/go/node/types/resources/v1beta4"
	"pkg.akt.dev/go/node/types/unit"
	"pkg.akt.dev/go/testutil"
)

func TestV2_1_ParseSimpleGPU(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-simple-gpu.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	assert.Len(t, groups, 1)

	group := groups[0]
	assert.Len(t, group.GetResourceUnits(), 1)
	assert.Len(t, group.Requirements.Attributes, 2)

	assert.Equal(t, atypes.Attribute{
		Key:   "region",
		Value: "us-west",
	}, group.Requirements.Attributes[1])

	assert.Len(t, group.GetResourceUnits(), 1)

	assert.Equal(t, dtypes.ResourceUnit{
		Count: 2,
		Resources: rtypes.Resources{
			ID: 1,
			CPU: &rtypes.CPU{
				Units: rtypes.NewResourceValue(randCPU),
			},
			GPU: &rtypes.GPU{
				Units: rtypes.NewResourceValue(randGPU),
				Attributes: atypes.Attributes{
					{
						Key:   "vendor/nvidia/model/a100",
						Value: "true",
					},
				},
			},
			Memory: &rtypes.Memory{
				Quantity: rtypes.NewResourceValue(randMemory),
			},
			Storage: rtypes.Volumes{
				{
					Name:     "default",
					Quantity: rtypes.NewResourceValue(randStorage),
				},
			},
			Endpoints: rtypes.Endpoints{
				{
					Kind: rtypes.Endpoint_SHARED_HTTP,
				},
				{
					Kind: rtypes.Endpoint_RANDOM_PORT,
				},
			},
		},
		Price: testutil.ACTDecCoin(t, 50),
	}, group.GetResourceUnits()[0])

	mani, err := sdl.Manifest()
	require.NoError(t, err)

	assert.Len(t, mani.GetGroups(), 1)

	expectedHosts := make([]string, 1)
	expectedHosts[0] = "ahostname.com" // nolint: goconst
	assert.Equal(t, manifest.Group{
		Name: "westcoast",
		Services: []manifest.Service{
			{
				Name:  "web",
				Image: "nginx",
				Resources: rtypes.Resources{
					ID: 1,
					CPU: &rtypes.CPU{
						Units: rtypes.NewResourceValue(100),
					},
					GPU: &rtypes.GPU{
						Units: rtypes.NewResourceValue(1),
						Attributes: atypes.Attributes{
							{
								Key:   "vendor/nvidia/model/a100",
								Value: "true",
							},
						},
					},
					Memory: &rtypes.Memory{
						Quantity: rtypes.NewResourceValue(128 * unit.Mi),
					},
					Storage: rtypes.Volumes{
						{
							Name:     "default",
							Quantity: rtypes.NewResourceValue(1 * unit.Gi),
						},
					},
					Endpoints: rtypes.Endpoints{
						{
							Kind: rtypes.Endpoint_SHARED_HTTP,
						},
						{
							Kind: rtypes.Endpoint_RANDOM_PORT,
						},
					},
				},
				Count: 2,
				Expose: []manifest.ServiceExpose{
					{Port: 80, Global: true, Proto: manifest.TCP, Hosts: expectedHosts,
						HTTPOptions: manifest.ServiceExposeHTTPOptions{
							MaxBodySize: 1048576,
							ReadTimeout: 60000,
							SendTimeout: 60000,
							NextTries:   3,
							NextTimeout: 0,
							NextCases:   []string{"error", "timeout"},
						}},
					{Port: 12345, Global: true, Proto: manifest.UDP,
						HTTPOptions: manifest.ServiceExposeHTTPOptions{
							MaxBodySize: 1048576,
							ReadTimeout: 60000,
							SendTimeout: 60000,
							NextTries:   3,
							NextTimeout: 0,
							NextCases:   []string{"error", "timeout"},
						}},
				},
			},
		},
	}, mani.GetGroups()[0])
}

func TestV2_1_Parse_Deployments(t *testing.T) {
	sdl1, err := ReadFile("_testdata/legacy/deployment-v2.1.yaml")
	require.NoError(t, err)
	_, err = sdl1.DeploymentGroups()
	require.NoError(t, err)

	_, err = sdl1.Manifest()
	require.NoError(t, err)

	sha1, err := sdl1.Version()
	require.NoError(t, err)
	assert.Len(t, sha1, 32)

	sdl2, err := ReadFile("_testdata/legacy/deployment-v2.yaml")
	require.NoError(t, err)
	sha2, err := sdl2.Version()

	require.NoError(t, err)
	assert.Len(t, sha2, 32)
	require.NotEqual(t, sha1, sha2)
}

func Test_V2_1_Cross_Validates(t *testing.T) {
	sdl2, err := ReadFile("_testdata/legacy/deployment-v2.yaml")
	require.NoError(t, err)
	dgroups, err := sdl2.DeploymentGroups()
	require.NoError(t, err)
	m, err := sdl2.Manifest()
	require.NoError(t, err)

	// This is a single document producing both the manifest & deployment groups
	// These should always agree with each other. If this test fails at least one of the
	// following is ture
	// 1. Cross validation logic is wrong
	// 2. The DeploymentGroups() & Manifest() code do not agree with one another
	err = m.CheckAgainstGSpecs(dgroups)
	require.NoError(t, err)

	// Repeat the same test with another file
	sdl2, err = ReadFile("./_testdata/v2.1-simple.yaml")
	require.NoError(t, err)
	dgroups, err = sdl2.DeploymentGroups()
	require.NoError(t, err)
	m, err = sdl2.Manifest()
	require.NoError(t, err)

	// This is a single document producing both the manifest & deployment groups
	// These should always agree with each other
	err = m.CheckAgainstGSpecs(dgroups)
	require.NoError(t, err)

	// Repeat the same test with another file
	sdl2, err = ReadFile("./_testdata/v2.1-simple3.yaml")
	require.NoError(t, err)
	dgroups, err = sdl2.DeploymentGroups()
	require.NoError(t, err)
	m, err = sdl2.Manifest()
	require.NoError(t, err)

	// This is a single document producing both the manifest & deployment groups
	// These should always agree with each other
	err = m.CheckAgainstGSpecs(dgroups)
	require.NoError(t, err)

	// Repeat the same test with another file
	sdl2, err = ReadFile("./_testdata/v2.1-private_service.yaml")
	require.NoError(t, err)
	dgroups, err = sdl2.DeploymentGroups()
	require.NoError(t, err)
	m, err = sdl2.Manifest()
	require.NoError(t, err)

	// This is a single document producing both the manifest & deployment groups
	// These should always agree with each other
	err = m.CheckAgainstGSpecs(dgroups)
	require.NoError(t, err)

}

func Test_V2_1_Parse_simple(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-simple.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	assert.Len(t, groups, 1)

	group := groups[0]
	assert.Len(t, group.GetResourceUnits(), 1)

	assert.Equal(t, atypes.Attribute{
		Key:   "region",
		Value: "us-west",
	}, group.Requirements.Attributes[0])

	assert.Len(t, group.GetResourceUnits(), 1)

	assert.Equal(t, dtypes.ResourceUnit{
		Count: 2,
		Resources: rtypes.Resources{
			ID: 1,
			CPU: &rtypes.CPU{
				Units: rtypes.NewResourceValue(randCPU),
			},
			GPU: &rtypes.GPU{
				Units: rtypes.NewResourceValue(0),
			},
			Memory: &rtypes.Memory{
				Quantity: rtypes.NewResourceValue(randMemory),
			},
			Storage: rtypes.Volumes{
				{
					Name:     "default",
					Quantity: rtypes.NewResourceValue(randStorage),
				},
			},
			Endpoints: rtypes.Endpoints{
				{
					Kind: rtypes.Endpoint_SHARED_HTTP,
				},
				{
					Kind: rtypes.Endpoint_RANDOM_PORT,
				},
			},
		},
		Price: testutil.ACTDecCoin(t, 50),
	}, group.GetResourceUnits()[0])

	mani, err := sdl.Manifest()
	require.NoError(t, err)

	assert.Len(t, mani.GetGroups(), 1)

	expectedHosts := make([]string, 1)
	expectedHosts[0] = "ahostname.com"
	assert.Equal(t, manifest.Group{
		Name: "westcoast",
		Services: []manifest.Service{
			{
				Name:  "web",
				Image: "nginx",
				Resources: rtypes.Resources{
					ID: 1,
					CPU: &rtypes.CPU{
						Units: rtypes.NewResourceValue(100),
					},
					GPU: &rtypes.GPU{
						Units: rtypes.NewResourceValue(0),
					},
					Memory: &rtypes.Memory{
						Quantity: rtypes.NewResourceValue(128 * unit.Mi),
					},
					Storage: rtypes.Volumes{
						{
							Name:     "default",
							Quantity: rtypes.NewResourceValue(1 * unit.Gi),
						},
					},
					Endpoints: rtypes.Endpoints{
						{
							Kind: rtypes.Endpoint_SHARED_HTTP,
						},
						{
							Kind: rtypes.Endpoint_RANDOM_PORT,
						},
					},
				},
				Count: 2,
				Expose: []manifest.ServiceExpose{
					{Port: 80, Global: true, Proto: manifest.TCP, Hosts: expectedHosts,
						HTTPOptions: manifest.ServiceExposeHTTPOptions{
							MaxBodySize: 1048576,
							ReadTimeout: 60000,
							SendTimeout: 60000,
							NextTries:   3,
							NextTimeout: 0,
							NextCases:   []string{"error", "timeout"},
						}},
					{Port: 12345, Global: true, Proto: manifest.UDP,
						HTTPOptions: manifest.ServiceExposeHTTPOptions{
							MaxBodySize: 1048576,
							ReadTimeout: 60000,
							SendTimeout: 60000,
							NextTries:   3,
							NextTimeout: 0,
							NextCases:   []string{"error", "timeout"},
						}},
				},
			},
		},
	}, mani.GetGroups()[0])

	assert.Nil(t, mani.GetGroups()[0].Services[0].Credentials)

}

func Test_V2_1_Parse_credentials(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-credentials.yaml")
	require.NoError(t, err)

	mani, err := sdl.Manifest()
	require.NoError(t, err)

	assert.Len(t, mani.GetGroups(), 1)

	grp := mani.GetGroups()[0]
	assert.Len(t, grp.Services, 1)

	svc := grp.Services[0]

	assert.NotNil(t, svc)

	creds := svc.Credentials
	assert.NotNil(t, creds)

	assert.Equal(t, "https://test.com/v1", creds.Host)
	assert.Equal(t, "foo", creds.Username)
	assert.Equal(t, "foo", creds.Password)
}

func Test_V2_1_Parse_credentials_error(t *testing.T) {
	_, err := ReadFile("./_testdata/v2.1-credentials-error.yaml")
	require.Error(t, err)
}

func Test_v2_1_Parse_ProfileNameNotServiceName(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-profile-svc-name-mismatch.yaml")
	require.NoError(t, err)

	dgroups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	assert.Len(t, dgroups, 1)

	mani, err := sdl.Manifest()
	require.NoError(t, err)
	assert.Len(t, mani.GetGroups(), 1)
}

func Test_v2_1_Parse_DeploymentNameServiceNameMismatch(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-deployment-svc-mismatch.yaml")
	require.Error(t, err)
	require.Nil(t, sdl)
	require.Contains(t, err.Error(), "no service profile named")

	sdl, err = ReadFile("./_testdata/v2.1-simple2.yaml")
	require.NoError(t, err)
	require.NotNil(t, sdl)

	dgroups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	assert.Len(t, dgroups, 1)

	mani, err := sdl.Manifest()
	require.NoError(t, err)
	assert.Len(t, mani.GetGroups(), 1)

	require.Equal(t, dgroups[0].Name, mani.GetGroups()[0].Name)
	// SDL lists 2 services, but particular deployment specifies only one
	require.Len(t, mani.GetGroups()[0].Services, 1)

	// make sure deployment maps to the right service
	require.Len(t, mani.GetGroups()[0].Services[0].Expose, 2)
	require.Len(t, mani.GetGroups()[0].Services[0].Expose[0].Hosts, 1)
	require.Equal(t, mani.GetGroups()[0].Services[0].Expose[0].Hosts[0], "ahostname.com")
}

func TestV2_1_ParseServiceMix(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-service-mix.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	assert.Len(t, groups, 1)

	group := groups[0]
	assert.Len(t, group.GetResourceUnits(), 2)
	assert.Len(t, group.Requirements.Attributes, 2)

	assert.Equal(t, atypes.Attribute{
		Key:   "region",
		Value: "us-west",
	}, group.Requirements.Attributes[1])

	assert.Equal(t, dtypes.ResourceUnits{
		{
			Count: 1,
			Resources: rtypes.Resources{
				ID: 1,
				CPU: &rtypes.CPU{
					Units: rtypes.NewResourceValue(randCPU),
				},
				GPU: &rtypes.GPU{
					Units: rtypes.NewResourceValue(randGPU),
					Attributes: atypes.Attributes{
						{
							Key:   "vendor/nvidia/model/*",
							Value: "true",
						},
					},
				},
				Memory: &rtypes.Memory{
					Quantity: rtypes.NewResourceValue(randMemory),
				},
				Storage: rtypes.Volumes{
					{
						Name:     "default",
						Quantity: rtypes.NewResourceValue(randStorage),
					},
				},
				Endpoints: rtypes.Endpoints{
					{
						Kind: rtypes.Endpoint_SHARED_HTTP,
					},
					{
						Kind: rtypes.Endpoint_RANDOM_PORT,
					},
				},
			},
			Price: testutil.ACTDecCoin(t, 50),
		},
		{
			Count: 1,
			Resources: rtypes.Resources{
				ID: 2,
				CPU: &rtypes.CPU{
					Units: rtypes.NewResourceValue(randCPU),
				},
				GPU: &rtypes.GPU{
					Units: rtypes.NewResourceValue(0),
				},
				Memory: &rtypes.Memory{
					Quantity: rtypes.NewResourceValue(randMemory),
				},
				Storage: rtypes.Volumes{
					{
						Name:     "default",
						Quantity: rtypes.NewResourceValue(randStorage),
					},
				},
				Endpoints: rtypes.Endpoints{
					{
						Kind: rtypes.Endpoint_SHARED_HTTP,
					},
					{
						Kind: rtypes.Endpoint_RANDOM_PORT,
					},
				},
			},
			Price: testutil.ACTDecCoin(t, 50),
		},
	}, group.GetResourceUnits())

	mani, err := sdl.Manifest()
	require.NoError(t, err)

	assert.Len(t, mani.GetGroups(), 1)

	assert.Equal(t, manifest.Group{
		Name: "westcoast",
		Services: []manifest.Service{
			{
				Name:  "svca",
				Image: "nginx",
				Resources: rtypes.Resources{
					ID: 1,
					CPU: &rtypes.CPU{
						Units: rtypes.NewResourceValue(100),
					},
					GPU: &rtypes.GPU{
						Units: rtypes.NewResourceValue(1),
						Attributes: atypes.Attributes{
							{
								Key:   "vendor/nvidia/model/*",
								Value: "true",
							},
						},
					},
					Memory: &rtypes.Memory{
						Quantity: rtypes.NewResourceValue(128 * unit.Mi),
					},
					Storage: rtypes.Volumes{
						{
							Name:     "default",
							Quantity: rtypes.NewResourceValue(1 * unit.Gi),
						},
					},
					Endpoints: rtypes.Endpoints{
						{
							Kind: rtypes.Endpoint_SHARED_HTTP,
						},
						{
							Kind: rtypes.Endpoint_RANDOM_PORT,
						},
					},
				},
				Count: 1,
				Expose: []manifest.ServiceExpose{
					{
						Port: 80, Global: true, Proto: manifest.TCP, Hosts: []string{"ahostname.com"},
						HTTPOptions: defaultHTTPOptions,
					},
					{
						Port: 12345, Global: true, Proto: manifest.UDP,
						HTTPOptions: defaultHTTPOptions,
					},
				},
			},
			{
				Name:  "svcb",
				Image: "nginx",
				Resources: rtypes.Resources{
					ID: 2,
					CPU: &rtypes.CPU{
						Units: rtypes.NewResourceValue(100),
					},
					GPU: &rtypes.GPU{
						Units: rtypes.NewResourceValue(0),
					},
					Memory: &rtypes.Memory{
						Quantity: rtypes.NewResourceValue(128 * unit.Mi),
					},
					Storage: rtypes.Volumes{
						{
							Name:     "default",
							Quantity: rtypes.NewResourceValue(1 * unit.Gi),
						},
					},
					Endpoints: rtypes.Endpoints{
						{
							Kind: rtypes.Endpoint_SHARED_HTTP,
						},
						{
							Kind: rtypes.Endpoint_RANDOM_PORT,
						},
					},
				},
				Count: 1,
				Expose: []manifest.ServiceExpose{
					{
						Port: 80, Global: true, Proto: manifest.TCP, Hosts: []string{"bhostname.com"},
						HTTPOptions: defaultHTTPOptions,
					},
					{
						Port: 12346, Global: true, Proto: manifest.UDP,
						HTTPOptions: defaultHTTPOptions,
					},
				},
			},
		},
	}, mani.GetGroups()[0])
}

func TestV2_1_ParseServiceMix2(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-service-mix2.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	assert.Len(t, groups, 1)

	group := groups[0]
	assert.Len(t, group.GetResourceUnits(), 1)
	assert.Len(t, group.Requirements.Attributes, 2)

	assert.Equal(t, atypes.Attribute{
		Key:   "region",
		Value: "us-west",
	}, group.Requirements.Attributes[1])

	assert.Equal(t, dtypes.ResourceUnits{
		{
			Count: 2,
			Resources: rtypes.Resources{
				ID: 1,
				CPU: &rtypes.CPU{
					Units: rtypes.NewResourceValue(randCPU),
				},
				GPU: &rtypes.GPU{
					Units: rtypes.NewResourceValue(randGPU),
					Attributes: atypes.Attributes{
						{
							Key:   "vendor/nvidia/model/*",
							Value: "true",
						},
					},
				},
				Memory: &rtypes.Memory{
					Quantity: rtypes.NewResourceValue(randMemory),
				},
				Storage: rtypes.Volumes{
					{
						Name:     "default",
						Quantity: rtypes.NewResourceValue(randStorage),
					},
				},
				Endpoints: rtypes.Endpoints{
					{
						Kind: rtypes.Endpoint_SHARED_HTTP,
					},
					{
						Kind: rtypes.Endpoint_RANDOM_PORT,
					},
					{
						Kind: rtypes.Endpoint_SHARED_HTTP,
					},
					{
						Kind: rtypes.Endpoint_RANDOM_PORT,
					},
				},
			},
			Price: testutil.ACTDecCoin(t, 50),
		},
	}, group.GetResourceUnits())

	mani, err := sdl.Manifest()
	require.NoError(t, err)

	assert.Len(t, mani.GetGroups(), 1)

	assert.Equal(t, manifest.Group{
		Name: "westcoast",
		Services: []manifest.Service{
			{
				Name:  "svca",
				Image: "nginx",
				Resources: rtypes.Resources{
					ID: 1,
					CPU: &rtypes.CPU{
						Units: rtypes.NewResourceValue(100),
					},
					GPU: &rtypes.GPU{
						Units: rtypes.NewResourceValue(1),
						Attributes: atypes.Attributes{
							{
								Key:   "vendor/nvidia/model/*",
								Value: "true",
							},
						},
					},
					Memory: &rtypes.Memory{
						Quantity: rtypes.NewResourceValue(128 * unit.Mi),
					},
					Storage: rtypes.Volumes{
						{
							Name:     "default",
							Quantity: rtypes.NewResourceValue(1 * unit.Gi),
						},
					},
					Endpoints: rtypes.Endpoints{
						{
							Kind: rtypes.Endpoint_SHARED_HTTP,
						},
						{
							Kind: rtypes.Endpoint_RANDOM_PORT,
						},
					},
				},
				Count: 1,
				Expose: []manifest.ServiceExpose{
					{
						Port: 80, Global: true, Proto: manifest.TCP, Hosts: []string{"ahostname.com"},
						HTTPOptions: defaultHTTPOptions,
					},
					{
						Port: 12345, Global: true, Proto: manifest.UDP,
						HTTPOptions: defaultHTTPOptions,
					},
				},
			},
			{
				Name:  "svcb",
				Image: "nginx",
				Resources: rtypes.Resources{
					ID: 1,
					CPU: &rtypes.CPU{
						Units: rtypes.NewResourceValue(100),
					},
					GPU: &rtypes.GPU{
						Units: rtypes.NewResourceValue(1),
						Attributes: atypes.Attributes{
							{
								Key:   "vendor/nvidia/model/*",
								Value: "true",
							},
						},
					},
					Memory: &rtypes.Memory{
						Quantity: rtypes.NewResourceValue(128 * unit.Mi),
					},
					Storage: rtypes.Volumes{
						{
							Name:     "default",
							Quantity: rtypes.NewResourceValue(1 * unit.Gi),
						},
					},
					Endpoints: rtypes.Endpoints{
						{
							Kind: rtypes.Endpoint_SHARED_HTTP,
						},
						{
							Kind: rtypes.Endpoint_RANDOM_PORT,
						},
					},
				},
				Count: 1,
				Expose: []manifest.ServiceExpose{
					{
						Port: 80, Global: true, Proto: manifest.TCP, Hosts: []string{"bhostname.com"},
						HTTPOptions: defaultHTTPOptions,
					},
					{
						Port: 12346, Global: true, Proto: manifest.UDP,
						HTTPOptions: defaultHTTPOptions,
					},
				},
			},
		},
	}, mani.GetGroups()[0])
}

func TestV2_1_ConfidentialComputeGPU(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-confidential-compute-gpu.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	require.Len(t, groups, 1)

	group := groups[0]

	// Verify CC attributes survive SDL → GroupSpec round-trip
	ccFound := false
	teeFound := false
	for _, attr := range group.Requirements.Attributes {
		if attr.Key == "confidential-compute" && attr.Value == "true" {
			ccFound = true
		}
		if attr.Key == "confidential-compute-tee" && attr.Value == "amd-sev-snp" {
			teeFound = true
		}
	}
	assert.True(t, ccFound, "confidential-compute attribute must survive round-trip")
	assert.True(t, teeFound, "confidential-compute-tee attribute must survive round-trip")

	// Verify GPU resources are present
	require.Len(t, group.GetResourceUnits(), 1)
	ru := group.GetResourceUnits()[0]
	require.NotNil(t, ru.GPU)
	assert.Equal(t, uint64(1), ru.GPU.Units.Value())
	assert.Equal(t, atypes.Attributes{
		{Key: "vendor/nvidia/model/h100", Value: "true"},
	}, ru.GPU.Attributes)

	// Verify manifest also round-trips correctly
	mani, err := sdl.Manifest()
	require.NoError(t, err)
	require.Len(t, mani.GetGroups(), 1)

	svc := mani.GetGroups()[0].Services[0]
	assert.Equal(t, "web", svc.Name)
	assert.Equal(t, "nginx", svc.Image)
	require.NotNil(t, svc.Resources.GPU)
	assert.Equal(t, uint64(1), svc.Resources.GPU.Units.Value())
}

func TestV2_1_ConfidentialComputeCPUOnly(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-confidential-compute-cpu-only.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	require.Len(t, groups, 1)

	group := groups[0]

	// Verify CC attribute survives SDL → GroupSpec round-trip
	ccFound := false
	for _, attr := range group.Requirements.Attributes {
		if attr.Key == "confidential-compute" && attr.Value == "true" {
			ccFound = true
		}
	}
	assert.True(t, ccFound, "confidential-compute attribute must survive round-trip")

	// Verify no GPU resources (CPU-only CC)
	require.Len(t, group.GetResourceUnits(), 1)
	ru := group.GetResourceUnits()[0]
	require.NotNil(t, ru.GPU)
	assert.Equal(t, uint64(0), ru.GPU.Units.Value())

	// Verify manifest round-trip
	mani, err := sdl.Manifest()
	require.NoError(t, err)
	require.Len(t, mani.GetGroups(), 1)
	assert.Equal(t, "web", mani.GetGroups()[0].Services[0].Name)
}

func TestV2_1_ConfidentialComputeAttributeInline(t *testing.T) {
	stream := `
version: "2.1"
services:
  app:
    image: myapp:latest
    expose:
      - port: 8080
        to:
          - global: true
profiles:
  compute:
    app:
      resources:
        cpu:
          units: "500m"
        gpu:
          units: 1
          attributes:
            vendor:
              nvidia:
              - model: h100
                ram: 80Gi
        memory:
          size: "512Mi"
        storage:
        - size: "2Gi"
  placement:
    dc1:
      attributes:
        confidential-compute: "true"
        confidential-compute-tee: amd-sev-snp
      pricing:
        app:
          denom: uakt
          amount: 100
deployment:
  app:
    dc1:
      profile: app
      count: 1
`
	sdl, err := Read([]byte(stream))
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	require.Len(t, groups, 1)

	group := groups[0]

	// Both CC attributes must be present
	attrMap := make(map[string]string)
	for _, attr := range group.Requirements.Attributes {
		attrMap[attr.Key] = attr.Value
	}
	assert.Equal(t, "true", attrMap["confidential-compute"])
	assert.Equal(t, "amd-sev-snp", attrMap["confidential-compute-tee"])

	// GPU with RAM attribute
	ru := group.GetResourceUnits()[0]
	require.NotNil(t, ru.GPU)
	assert.Equal(t, uint64(1), ru.GPU.Units.Value())
	assert.Contains(t, ru.GPU.Attributes[0].Key, "vendor/nvidia/model/h100")
}

func TestV2_1_ConfidentialComputeAttestationOptOut(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-confidential-compute-no-attestation.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	require.Len(t, groups, 1)

	// Verify CC attribute is present
	ccFound := false
	for _, attr := range groups[0].Requirements.Attributes {
		if attr.Key == "confidential-compute" && attr.Value == "true" {
			ccFound = true
		}
	}
	assert.True(t, ccFound, "confidential-compute attribute must survive round-trip")

	// Verify manifest parses — the attestation opt-out is in service params
	mani, err := sdl.Manifest()
	require.NoError(t, err)
	require.Len(t, mani.GetGroups(), 1)

	svc := mani.GetGroups()[0].Services[0]
	assert.Equal(t, "web", svc.Name)

	// Verify attestation opt-out flows through to manifest ServiceParams
	require.NotNil(t, svc.Params, "service params must not be nil")
	require.NotNil(t, svc.Params.Attestation, "attestation params must not be nil")
	assert.False(t, svc.Params.Attestation.Enabled, "attestation must be disabled")
}

func TestV2_1_ConfidentialComputeTDXCPUOnly(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-confidential-compute-tdx.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	require.Len(t, groups, 1)

	group := groups[0]

	attrMap := make(map[string]string)
	for _, attr := range group.Requirements.Attributes {
		attrMap[attr.Key] = attr.Value
	}
	assert.Equal(t, "true", attrMap["confidential-compute"])
	assert.Equal(t, "intel-tdx", attrMap["confidential-compute-tee"])

	// CPU-only: no GPU
	ru := group.GetResourceUnits()[0]
	require.NotNil(t, ru.GPU)
	assert.Equal(t, uint64(0), ru.GPU.Units.Value())

	mani, err := sdl.Manifest()
	require.NoError(t, err)
	require.Len(t, mani.GetGroups(), 1)
	assert.Equal(t, "web", mani.GetGroups()[0].Services[0].Name)
}

func TestV2_1_ConfidentialComputeTDXGPU(t *testing.T) {
	sdl, err := ReadFile("./_testdata/v2.1-confidential-compute-tdx-gpu.yaml")
	require.NoError(t, err)

	groups, err := sdl.DeploymentGroups()
	require.NoError(t, err)
	require.Len(t, groups, 1)

	group := groups[0]

	attrMap := make(map[string]string)
	for _, attr := range group.Requirements.Attributes {
		attrMap[attr.Key] = attr.Value
	}
	assert.Equal(t, "true", attrMap["confidential-compute"])
	assert.Equal(t, "intel-tdx", attrMap["confidential-compute-tee"])

	// GPU present
	ru := group.GetResourceUnits()[0]
	require.NotNil(t, ru.GPU)
	assert.Equal(t, uint64(1), ru.GPU.Units.Value())
	assert.Equal(t, atypes.Attributes{
		{Key: "vendor/nvidia/model/h100", Value: "true"},
	}, ru.GPU.Attributes)

	mani, err := sdl.Manifest()
	require.NoError(t, err)
	require.Len(t, mani.GetGroups(), 1)
	svc := mani.GetGroups()[0].Services[0]
	assert.Equal(t, "nvidia-app", svc.Image)
	require.NotNil(t, svc.Resources.GPU)
	assert.Equal(t, uint64(1), svc.Resources.GPU.Units.Value())
}
