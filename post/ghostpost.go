package post

import (
	"encoding/json"
	"io/ioutil"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type GhostPost struct {
	// ID Is the post id
	Id int `json:"id"`
	// Uuid is the post unique id
	Uuid string `json:"uuid"`
	// Title is the post title
	Title string `json:"title"`
	// Slug is the post slug
	Slug string `json:"slug"`
	// ContentsMarkdown is the markdown contents
	ContentsMarkdown string `json:"contents_markdown"`
	// ContentsHtml are the html contents
	ContentsHtml string `json:"contents_html"`
	// Image contains the image for the post
	Image *string `json:"image"`
	// Featured indicates if the post is featured.
	Featured bool `json:"featured"`
	// Page is ???
	Page bool `json:"page"`
	// Status is the status of the post.
	Status string `json:"status"`
	// Language is the language string for the post
	Language string `json:"language"`
	// MetaTitle is the metadata title for the post
	MetaTitle *string `json:"meta_title"`
	// MetaDescription is the metadata description for the post
	MetaDescription *string `json:"meta_description"`
	// AuthorId is the author id
	AuthorId int `json:"author_id"`
	// CreatedAt is the created timestamp
	CreatedAt time.Time `json:"created_at"`
	// CreatedBy is the creator id
	CreatedBy int `json:"created_by"`
	// UpdatedAt is the updated timestamp
	UpdatedAt time.Time `json:"updated_at"`
	// UpdatedBy is the updater id
	UpdatedBy int `json:"updated_by"`
	// PublishedAt is the publishing time
	PublishedAt time.Time `json:"published_at"`
	// PublishedBy is the publisher id
	PublishedBy int `json:"published_by"`
}

func (g *GhostPost) WriteToFile(filename string) error {
	dat, err := json.MarshalIndent(g, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, dat, 0644)
}
