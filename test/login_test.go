package test

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/ezrantn/ranko/web"
	"io"
	"testing"
)

func TestLoginView(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		_ = web.Login().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}

	// Expect the component to have a form
	if doc.Find(`form`).Length() == 0 {
		t.Error("expected form to be rendered, but it wasn't")
	}
}
