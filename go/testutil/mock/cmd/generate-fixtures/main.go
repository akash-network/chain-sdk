package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/types/query"

	dv1beta4 "pkg.akt.dev/go/node/deployment/v1beta4"
	mv1beta5 "pkg.akt.dev/go/node/market/v1beta5"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <output-dir>\n", os.Args[0])
		os.Exit(1)
	}

	outputDir := os.Args[1]

	deploymentsResp := &dv1beta4.QueryDeploymentsResponse{
		Deployments: []dv1beta4.QueryDeploymentResponse{},
		Pagination: &query.PageResponse{
			Total: uint64(0),
		},
	}

	deploymentsData := map[string]interface{}{
		"deployments": deploymentsResp,
	}

	deploymentsJSON, err := json.MarshalIndent(deploymentsData, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal deployments: %v", err)
	}

	deploymentsPath := filepath.Join(outputDir, "deployments.json")
	if err := os.WriteFile(deploymentsPath, deploymentsJSON, 0644); err != nil {
		log.Fatalf("Failed to write %s: %v", deploymentsPath, err)
	}

	fmt.Printf("Generated %s\n", deploymentsPath)

	bidsResp := &mv1beta5.QueryBidsResponse{
		Bids: []mv1beta5.QueryBidResponse{},
		Pagination: &query.PageResponse{
			Total: uint64(0),
		},
	}

	leasesResp := &mv1beta5.QueryLeasesResponse{
		Leases: []mv1beta5.QueryLeaseResponse{},
		Pagination: &query.PageResponse{
			Total: uint64(0),
		},
	}

	marketData := map[string]interface{}{
		"leases": leasesResp,
		"bids":   bidsResp,
	}

	marketJSON, err := json.MarshalIndent(marketData, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal market data: %v", err)
	}

	marketPath := filepath.Join(outputDir, "market.json")
	if err := os.WriteFile(marketPath, marketJSON, 0644); err != nil {
		log.Fatalf("Failed to write %s: %v", marketPath, err)
	}

	fmt.Printf("Generated %s\n", marketPath)
}
