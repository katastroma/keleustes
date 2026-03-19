package keleustes

import (
	"context"
	"io/fs"

	katastromaorgv0 "github.com/katastroma/tropis/apis/katastroma.org/v0"
)

// Retriever fetches an Application's source into a readable filesystem.
// The credential may be nil for public repositories.
type Retriever interface {
	Retrieve(
		ctx context.Context,
		source katastromaorgv0.ApplicationSource,
		cred *katastromaorgv0.RepoCredential,
	) (fs.FS, error)
}
