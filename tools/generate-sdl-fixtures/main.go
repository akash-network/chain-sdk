package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"pkg.akt.dev/go/sdl"
)

func main() {
	inputRoot := filepath.Join("testdata", "sdl", "input")
	outputRoot := filepath.Join("testdata", "sdl", "output-fixtures")
	versions := []string{"v2.0", "v2.1"}

	for _, version := range versions {
		inputVersionDir := filepath.Join(inputRoot, version)

		if _, err := os.Stat(inputVersionDir); os.IsNotExist(err) {
			continue
		}

		entries, err := os.ReadDir(inputVersionDir)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", inputVersionDir, err)
			os.Exit(1)
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			inputPath := filepath.Join(inputVersionDir, entry.Name(), "input.yaml")

			if _, err := os.Stat(inputPath); os.IsNotExist(err) {
				fmt.Printf("Missing input.yaml for fixture %s (%s)\n", entry.Name(), inputPath)
				os.Exit(1)
			} else if err != nil {
				fmt.Printf("Error accessing %s: %v\n", inputPath, err)
				os.Exit(1)
			}

			outputDir := filepath.Join(outputRoot, version, entry.Name())
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				fmt.Printf("Error creating output dir %s: %v\n", outputDir, err)
				os.Exit(1)
			}

			fmt.Printf("Processing %s...\n", inputPath)

			obj, err := sdl.ReadFile(inputPath)
			if err != nil {
				fmt.Printf("  Error: %v\n", err)
				os.Exit(1)
			}

			if err := generateManifest(obj, outputDir); err != nil {
				fmt.Printf("  %v\n", err)
				os.Exit(1)
			}

			if err := generateGroups(obj, outputDir); err != nil {
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
