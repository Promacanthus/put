package put

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Promacanthus/put/pkg/octocat"
	"github.com/Promacanthus/put/pkg/pointer"
	"github.com/Promacanthus/put/pkg/sparrow"
	"github.com/google/go-github/v52/github"
)

const (
	// No organization provided.
	emptyOrg = ""
)

type Server struct {
	client  *github.Client
	owner   string
	private bool
	branch  string
}

func NewServer() *Server {
	owner := os.Getenv("GITHUB_OWNER")
	if len(owner) == 0 {
		log.Fatal("Set github owner into GITHUB_OWNER")
	}
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if len(token) == 0 {
		log.Fatal("Unauthorized: No token present")
	}
	client := github.NewTokenClient(context.Background(), token)
	return &Server{
		client:  client,
		owner:   owner,
		private: false,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	payload := &sparrow.Payload{}
	if err := json.Unmarshal(b, payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	doc := octocat.NewDocument(payload.Data)
	_, resp, err := s.client.Repositories.Get(ctx, s.owner, doc.Repository)
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			if _, err = s.CreateRepository(ctx, &doc.Repository, &s.private); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	opts := &github.RepositoryContentGetOptions{Ref: "heads/main"}
	content, _, resp, err := s.client.Repositories.GetContents(ctx, s.owner, doc.Repository, doc.GetPath(), opts)
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			if err := s.CreateFile(ctx, doc); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		doc.SHA = *content.SHA
		if err := s.UpdateFile(ctx, doc); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

// CreateRepository creates a GitHub repository with the given name and private flag.
// autoInit is set to true. It returns the created repository or an error if something went wrong.
func (s *Server) CreateRepository(ctx context.Context, name *string, private *bool) (*github.Repository, error) {
	repo := &github.Repository{
		Name:     name,
		Private:  private,
		AutoInit: pointer.Bool(false),
	}
	r, _, err := s.client.Repositories.Create(ctx, emptyOrg, repo)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// CreateFile creates a new file in the specified repository.
// It takes in a context and Document struct which consists of the repository, path and other needed information.
// It then uses the CreateFile function from the Github client to create the file and returns any errors.
func (s *Server) CreateFile(ctx context.Context, doc *octocat.Document) error {
	opts := &github.RepositoryContentFileOptions{
		Message: &doc.Message,
		Content: []byte(doc.Content),
	}
	_, _, err := s.client.Repositories.CreateFile(ctx, s.owner, doc.Repository, doc.GetPath(), opts)
	return err
}

// UpdateFile updates an existing Document in the specified Repository with the given content.
// It returns an error in case of any failure.
func (s *Server) UpdateFile(ctx context.Context, doc *octocat.Document) error {
	opts := &github.RepositoryContentFileOptions{
		Message: &doc.Message,
		Content: []byte(doc.Content),
		SHA:     &doc.SHA,
	}
	_, _, err := s.client.Repositories.UpdateFile(ctx, s.owner, doc.Repository, doc.GetPath(), opts)
	return err
}
