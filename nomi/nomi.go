package nomi

import (
	"context"
	"fmt"
	"iter"
	"net/url"

	"github.com/atomotic/iccu/client"
	"github.com/carlmjohnson/requests"
)

const baseURL = "https://api.iccu.sbn.it/nomi/1.0.0/search"

// SearchOptions configures the search behavior
type SearchOptions struct {
	PageSize int // Number of results per page (default: 500)
}

// Search returns an iterator over search results for names in the ICCU names database.
// The iterator automatically handles pagination and can be used with range loops.
//
// Example:
//
//	for doc := range nomi.Search(ctx, client, "alei* crowley", nil) {
//	    fmt.Println(doc.Name())
//	}
func Search(ctx context.Context, c *client.Client, query string, opts *SearchOptions) iter.Seq[*Doc] {
	pageSize := 500
	if opts != nil && opts.PageSize > 0 {
		pageSize = opts.PageSize
	}

	return func(yield func(*Doc) bool) {
		startAt := 0
		total := 0

		for {
			// Check if we've fetched all results
			if total > 0 && startAt >= total {
				return
			}

			// Build URL with pagination parameters
			q := url.QueryEscape(query)
			apiURL := fmt.Sprintf("%s?format=json&detail=full&page-size=%d&start-at=%d&monocampo=%s",
				// apiURL := fmt.Sprintf("%s?format=json&detail=full&page-size=%d&start-at=%d&nome=%s",
				baseURL, pageSize, startAt, q)

			// Fetch page
			var r MainResponse
			err := requests.
				URL(apiURL).
				Client(c.HTTP()).
				ToJSON(&r).
				Fetch(ctx)

			if err != nil {
				// Errors are swallowed in iterator; consider logging or alternative handling
				return
			}

			// Update total on first fetch
			if total == 0 {
				total = r.Response.NumFound
			}

			// No more results
			if len(r.Response.Docs) == 0 {
				return
			}

			// Yield each document
			for i := range r.Response.Docs {
				if !yield(&r.Response.Docs[i]) {
					return // Consumer stopped iteration
				}
			}

			// Move to next page
			startAt += len(r.Response.Docs)
		}
	}
}

// SearchWithTotal returns an iterator that yields both the document and metadata about the search.
// The total count is fetched immediately and available before iteration begins.
func SearchWithTotal(ctx context.Context, c *client.Client, query string, opts *SearchOptions) (iter.Seq[*Doc], *int, error) {
	pageSize := 500
	if opts != nil && opts.PageSize > 0 {
		pageSize = opts.PageSize
	}

	// Fetch first page immediately to get total count
	q := url.QueryEscape(query)
	apiURL := fmt.Sprintf("%s?format=json&detail=full&page-size=%d&start-at=%d&monocampo=%s",
		baseURL, pageSize, 0, q)

	var firstPage MainResponse
	err := requests.
		URL(apiURL).
		Client(c.HTTP()).
		ToJSON(&firstPage).
		Fetch(ctx)

	if err != nil {
		return nil, nil, err
	}

	total := new(int)
	*total = firstPage.Response.NumFound

	seq := func(yield func(*Doc) bool) {
		// Yield documents from the first page
		for i := range firstPage.Response.Docs {
			if !yield(&firstPage.Response.Docs[i]) {
				return
			}
		}

		// If first page had all results, we're done
		startAt := len(firstPage.Response.Docs)
		if startAt >= *total {
			return
		}

		// Continue fetching remaining pages
		for {
			// Check if we've fetched all results
			if startAt >= *total {
				return
			}

			// Build URL with pagination parameters
			apiURL := fmt.Sprintf("%s?format=json&detail=full&page-size=%d&start-at=%d&monocampo=%s",
				baseURL, pageSize, startAt, q)

			// Fetch page
			var r MainResponse
			err := requests.
				URL(apiURL).
				Client(c.HTTP()).
				ToJSON(&r).
				Fetch(ctx)

			if err != nil {
				return
			}

			// No more results
			if len(r.Response.Docs) == 0 {
				return
			}

			// Yield each document
			for i := range r.Response.Docs {
				if !yield(&r.Response.Docs[i]) {
					return
				}
			}

			// Move to next page
			startAt += len(r.Response.Docs)
		}
	}

	return seq, total, nil
}
