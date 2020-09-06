package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

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

func run(ctx context.Context) error {
	db, err := sql.Open("sqlite3", "ghost.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Query("select * from posts")
	if err != nil {
		return err
	}
	defer stmt.Close()

	var i int
	for stmt.Next() {
		var gpost post.GhostPost
		err := stmt.Scan(
			&gpost.Id,
			&gpost.Uuid,
			&gpost.Title,
			&gpost.Slug,
			&gpost.ContentsMarkdown,
			&gpost.ContentsHtml,
			&gpost.Image,
			&gpost.Featured,
			&gpost.Page,
			&gpost.Status,
			&gpost.Language,
			&gpost.MetaTitle,
			&gpost.MetaDescription,
			&gpost.AuthorId,
			&gpost.CreatedAt,
			&gpost.CreatedBy,
			&gpost.UpdatedAt,
			&gpost.UpdatedBy,
			&gpost.PublishedAt,
			&gpost.PublishedBy,
		)
		if err != nil {
			return err
		}
		filename := fmt.Sprintf("%d-%s.json", i, gpost.Slug)
		i++
		if err := gpost.WriteToFile(filename); err != nil {
			return err
		}
		fmt.Printf("wrote %s\n", filename)
	}
	return nil
}
