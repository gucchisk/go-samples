package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v79/github"
)

func main() {
	client, err := github.NewClient(nil).WithEnterpriseURLs("https://api.github.com/", "https://uploads.github.com/")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	releases, _, err := client.Repositories.ListReleases(ctx, "gucchisk", "int64_t", nil)
	if err != nil {
		panic(err)
	}
	for _, release := range releases {
		fmt.Printf("Release: %s Tag: %s\n", release.GetName(), release.GetTagName())
	}
}
