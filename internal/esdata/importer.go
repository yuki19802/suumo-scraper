package esdata

import (
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/evertras/suumo-scraper/internal/suumo"
)

type Importer struct {
	esClient *elasticsearch.Client
}

func NewImporter(esClient *elasticsearch.Client) *Importer {
	return &Importer{
		esClient: esClient,
	}
}

func (i *Importer) DeleteAllDataIndices() error {
	toDelete := []string{
		IndexNameListing,
	}

	for _, index := range toDelete {
		// Do these separately since they may not all exist
		res, err := i.esClient.Indices.Delete([]string{index})

		if err != nil {
			return fmt.Errorf("i.esClient.Indices.Delete %q: %w", index, err)
		}

		if res.StatusCode/200 != 1 && res.StatusCode != 404 {
			return fmt.Errorf("unexpected status code for %q: %d", index, res.StatusCode)
		}
	}

	return nil
}

func (i *Importer) ImportListings(ctx context.Context, listings []suumo.Listing) error {
	bulk, err := startBulkAdder(ctx, i.esClient, IndexNameListing)
	if err != nil {
		return fmt.Errorf("startBulkAdder: %w", err)
	}
	defer bulk.closeWithLoggedError(ctx)

	for i, entry := range listings {
		err = bulk.add(ctx, entry)
		if err != nil {
			return fmt.Errorf("bulk.add #%d: %w", i, err)
		}
	}

	return nil
}
