package main

import (
	"compress/gzip"
	"fmt"
	"net/http"
	neturl "net/url"
	"sort"
	"strings"

	"github.com/ceshihao/go-yum"
)

type Repo struct {
	BaseURL string
}

func main() {
	repomds := []string{
		"https://dl.rockylinux.org/pub/rocky/9.6/BaseOS/x86_64/os/repodata/repomd.xml",
		"https://dl.rockylinux.org/pub/rocky/9.6/AppStream/x86_64/os/repodata/repomd.xml",
	}

	for _, repomdUrl := range repomds {
		url, err := neturl.Parse(repomdUrl)
		if err != nil {
			panic(err)
		}
		// baseUrl := "https://dl.rockylinux.org/pub/rocky/9.6/BaseOS/x86_64/os"
		// repomdUrl := fmt.Sprintf("%s/repodata/repomd.xml", baseUrl)
		repoPath := strings.TrimSuffix(url.Path, "/repodata/repomd.xml")
		respMeta, err := http.Get(url.String())
		if err != nil {
			panic(err)
		}
		defer respMeta.Body.Close()
		meta, err := yum.ReadRepoMetadata(respMeta.Body)
		if err != nil {
			panic(err)
		}
		var primary string
		for _, db := range meta.Databases {
			fmt.Printf("%s: %s\n", db.Type, db.Location.Href)
			if db.Type == "primary" {
				primary = db.Location.Href
			}
		}
		primaryURL, err := neturl.Parse(url.String())
		if err != nil {
			panic(err)
		}
		primaryURL.Path = fmt.Sprintf("%s/%s", repoPath, primary)
		respPrimary, err := http.Get(primaryURL.String())
		if err != nil {
			panic(err)
		}
		defer respPrimary.Body.Close()
		gzReader, err := gzip.NewReader(respPrimary.Body)
		if err != nil {
			panic(err)
		}
		primaryMeta, err := yum.ReadPrimaryMetadata(gzReader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Total packages: %d\n", len(primaryMeta.Packages))
		pkgs := []yum.PackageEntry{}
		for _, pkg := range primaryMeta.Packages {
			if strings.Contains(pkg.Arch, "x86_64") {
				pkgs = append(pkgs, pkg)
				// fmt.Printf("%s %s %d %s %s\n", pkg.PackageName, pkg.Version(), pkg.Epoch(), pkg.Release(), pkg.Name())
			}
		}
		sort.Slice(pkgs, func(i, j int) bool {
			return strings.Compare(pkgs[i].Name(), pkgs[j].Name()) < 0
		})
		for _, pkg := range pkgs {
			fmt.Printf("name:%s epoch: %d version:%s release%s\n", pkg.Name(), pkg.Epoch(), pkg.Version(), pkg.Release())
		}
	}
}
