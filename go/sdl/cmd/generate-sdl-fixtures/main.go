package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"pkg.akt.dev/go/sdl"
)

func main() {
	fixturesRoot := "../../testdata/sdl"
	versions := []string{"v2.0", "v2.1"}

	for _, version := range versions {
		versionDir := filepath.Join(fixturesRoot, version)

		if _, err := os.Stat(versionDir); os.IsNotExist(err) {
			continue
		}

		entries, err := os.ReadDir(versionDir)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", versionDir, err)
			os.Exit(1)
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
				fmt.Printf("  Error: %v\n", err)
				os.Exit(1)
			}

			if err := generateManifest(obj, fixtureDir); err != nil {
				fmt.Printf("  %v\n", err)
				os.Exit(1)
			}

			if err := generateGroups(obj, fixtureDir); err != nil {
				fmt.Printf("  %v\n", err)
				os.Exit(1)
			}
		}
	}

	fmt.Println("\nFixture generation complete!")
}

func generateManifest(obj sdl.SDL, fixtureDir string) error {
	manifest, err := obj.Manifest()
	if err != nil {
		return fmt.Errorf("manifest error: %w", err)
	}

	manifestJSON, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON marshal error: %w", err)
	}

	manifestPath := filepath.Join(fixtureDir, "manifest.json")
	if err := os.WriteFile(manifestPath, manifestJSON, 0600); err != nil {
		return fmt.Errorf("write error: %w", err)
	}

	fmt.Printf("  Generated %s\n", manifestPath)
	return nil
}

func generateGroups(obj sdl.SDL, fixtureDir string) error {
	groups, err := obj.DeploymentGroups()
	if err != nil {
		return fmt.Errorf("groups error: %w", err)
	}

	groupsJSON, err := json.MarshalIndent(groups, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON marshal error: %w", err)
	}

	groupsPath := filepath.Join(fixtureDir, "groups.json")
	if err := os.WriteFile(groupsPath, groupsJSON, 0600); err != nil {
		return fmt.Errorf("write error: %w", err)
	}

	fmt.Printf("  Generated %s\n", groupsPath)
	return nil
}
