package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/paralin/go-export-ghostdb/post"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
}

var outTemplate = `Title: %s
Date: %s
Slug: %s

%s`

func run(ctx context.Context) error {
	fileList, err := ioutil.ReadDir("../")
	if err != nil {
		return err
	}
	for _, f := range fileList {
		if f.IsDir() || !f.Mode().IsRegular() {
			continue
		}
		fname := f.Name()
		if !strings.HasSuffix(fname, ".json") {
			continue
		}
		fmt.Printf("reading file: %s\n", fname)
		var gpost post.GhostPost
		fdata, err := ioutil.ReadFile(path.Join("../", fname))
		if err != nil {
			return err
		}
		if err := json.Unmarshal(fdata, &gpost); err != nil {
			return err
		}
		of, err := os.Create(fname[:len(fname)-5] + ".md")
		if err != nil {
			return err
		}
		fmt.Fprintf(
			of,
			outTemplate,
			gpost.Title,
			gpost.CreatedAt.String(),
			gpost.Slug,
			gpost.ContentsMarkdown,
		)
		of.Close()
	}

	return nil
}
