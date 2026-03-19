package keleustes

import (
	"context"
	"io/fs"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Renderer produces a resource inventory from a source filesystem.
type Renderer interface {
	Render(ctx context.Context, source fs.FS) ([]unstructured.Unstructured, error)
}
