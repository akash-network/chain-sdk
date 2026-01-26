package sdl

type sdlTestBuilder struct {
	version         string
	endpoints       string
	hasTwoServices  bool
	serviceBlock    string
	exposeBlock     string
	resourcesBlock  string
	placementBlock  string
	deploymentBlock string
}

func (b sdlTestBuilder) build() string {
	version := `version: "2.0"`
	if b.version != "" {
		version = b.version
	}

	endpoints := ""
	if b.endpoints != "" {
		endpoints = b.endpoints + "\n"
	}

	service := `  web:`
	hasImage := b.serviceBlock != "" && (len(b.serviceBlock) > 10 && b.serviceBlock[:10] == "    image:")
	if !hasImage {
		service += `
    image: nginx`
	}
	if b.serviceBlock != "" {
		service += "\n" + b.serviceBlock
	}
	if b.exposeBlock != "" {
		service += "\n" + b.exposeBlock
	}

	services := "services:\n" + service
	if b.hasTwoServices {
		services += `
  db:
    image: postgres`
	}

	resources := `        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi`
	if b.resourcesBlock != "" {
		resources = b.resourcesBlock
	}

	compute := `  compute:
    web:
      resources:
` + resources
	if b.hasTwoServices {
		compute += `
    db:
      resources:
        cpu:
          units: 1
        memory:
          size: 1Gi
        storage:
          - size: 1Gi`
	}

	pricing := `        web:
          denom: uakt
          amount: 1`
	if b.hasTwoServices {
		pricing += `
        db:
          denom: uakt
          amount: 1`
	}

	placement := ""
	if b.placementBlock != "" {
		if len(b.placementBlock) >= 6 && b.placementBlock[:6] == "    dc" {
			placement = b.placementBlock
		} else {
			placement = `    dc:
` + b.placementBlock
		}
	} else {
		placement = `    dc:
      pricing:
` + pricing
	}

	deployment := ""
	if b.deploymentBlock != "" && len(b.deploymentBlock) >= 3 && b.deploymentBlock[:2] == "  " && b.deploymentBlock[2] != ' ' {
		deployment = b.deploymentBlock
	} else {
		deployment = `  web:
    dc:
      profile: web
      count: 1`
		if b.deploymentBlock != "" {
			deployment += "\n" + b.deploymentBlock
		}
		if b.hasTwoServices {
			deployment += `
  db:
    dc:
      profile: db
      count: 1`
		}
	}

	return version + `
` + endpoints + services + `
profiles:
` + compute + `
  placement:
` + placement + `
deployment:
` + deployment + `
`
}
