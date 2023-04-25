package octocat

import (
	"fmt"
	"strings"

	"github.com/promacanthus/put/pkg/sparrow"
)

const (
	aboutSlug = "about"
	README    = "README.md"
)

type Document struct {
	Slug       string `json:"slug,omitempty"`
	Name       string `json:"name,omitempty"`
	Repository string `json:"repository,omitempty"`
	Content    string `json:"content,omitempty"`
	Message    string `json:"message,omitempty"`
	SHA        string `json:"sha,omitempty"`
}

func NewDocument(data sparrow.Data) *Document {
	doc := &Document{
		Slug:       data.Slug,
		Name:       data.Title,
		Repository: data.Book.Name,
		Content:    data.BodyDraft,
	}
	doc.updateContent()
	doc.updateMessage(data)
	return doc
}

// GetPath returns the file path of a document in a markdown format.
// The path is composed of the document name and the ".md" extension.
func (d *Document) GetPath() string {
	if d.Slug == aboutSlug {
		return README
	}
	return fmt.Sprintf("%s.md", d.Name)
}

// updateContent updates the content of the Document object.
// It sets the content to the result of joining the string elements.
func (d *Document) updateContent() {
	d.Content = strings.Join([]string{"# ", d.Name, "\n", d.Content}, "")
}

// updateMessage updates the message stored in the Document struct.
// The message is a string generated from the data received from the sparrow.Data object.
func (d *Document) updateMessage(data sparrow.Data) {
	d.Message = fmt.Sprintf("%s %s %s", data.ContentUpdatedAt.String(), data.WebhookSubjectType, data.Title)
}
