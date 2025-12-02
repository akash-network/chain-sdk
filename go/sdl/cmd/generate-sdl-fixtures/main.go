package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"pkg.akt.dev/go/sdl"
)

func main() {
	fixturesRoot := "../../../testdata/sdl"
	versions := []string{"v2.0", "v2.1"}

	for _, version := range versions {
		versionDir := filepath.Join(fixturesRoot, version)

		if _, err := os.Stat(versionDir); os.IsNotExist(err) {
			continue
		}

		entries, err := os.ReadDir(versionDir)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", versionDir, err)
			continue
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			fixtureDir := filepath.Join(versionDir, entry.Name())
			inputPath := filepath.Join(fixtureDir, "input.yaml")

			if _, err := os.Stat(inputPath); os.IsNotExist(err) {
				continue
			}

			fmt.Printf("Processing %s...\n", inputPath)

			obj, err := sdl.ReadFile(inputPath)
			if err != nil {
				fmt.Printf("  ❌ Error: %v\n", err)
				continue
			}

			manifest, err := obj.Manifest()
			if err != nil {
				fmt.Printf("  ❌ Manifest error: %v\n", err)
				continue
			}
			manifestJSON, err := json.MarshalIndent(manifest, "", "  ")
			if err != nil {
				fmt.Printf("  ❌ JSON marshal error: %v\n", err)
				continue
			}
			manifestPath := filepath.Join(fixtureDir, "manifest.json")
			if err := os.WriteFile(manifestPath, manifestJSON, 0644); err != nil {
				fmt.Printf("  ❌ Write error: %v\n", err)
				continue
			}
			fmt.Printf("  ✓ Generated %s\n", manifestPath)

			groups, err := obj.DeploymentGroups()
			if err != nil {
				fmt.Printf("  ❌ Groups error: %v\n", err)
				continue
			}
			groupsJSON, err := json.MarshalIndent(groups, "", "  ")
			if err != nil {
				fmt.Printf("  ❌ JSON marshal error: %v\n", err)
				continue
			}
			groupsPath := filepath.Join(fixtureDir, "groups.json")
			if err := os.WriteFile(groupsPath, groupsJSON, 0644); err != nil {
				fmt.Printf("  ❌ Write error: %v\n", err)
				continue
			}
			fmt.Printf("  ✓ Generated %s\n", groupsPath)
		}
	}

	fmt.Println("\n✅ Fixture generation complete!")
}

