/*
This package provides an HTTP handler for serving a Single Page Application (SPA)
from embedded files using Go's embed package. It serves the SPA's static files
and ensures that all non-existent routes fall back to the main index.html file,
enabling client-side routing.

NOTE: Embedding requires the files to be present at compile time and must be in a subdirectory.
For this reason, this package is located in the 'web' directory, and the SPA files should be in 'web/dist'.
*/
package frontend

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/cooperlutz/go-full/pkg/hteeteepee"
)

//go:embed dist/*
var spaFiles embed.FS

// https://github.com/go-chi/chi/issues/611#issuecomment-1804702959
func spaHandler() http.HandlerFunc {
	spaFS, err := fs.Sub(spaFiles, "dist")
	if err != nil {
		panic(fmt.Errorf("failed getting the sub tree for the site files: %w", err))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; "+ // Allow content only from self
				"script-src 'self' 'unsafe-inline'; "+ // Allow scripts from self and inline scripts
				"style-src 'self' 'unsafe-inline' https://fonts.googleapis.com; "+ // Allow inline styles and loading from external stylesheets
				"img-src 'self' data:; "+ // Allow images from self and data URLs
				"connect-src 'self'; "+ // Allow connections to self
				"font-src 'self' https://fonts.gstatic.com; "+ // Allow fonts from self and Google Fonts
				"object-src 'none'; "+ // Disallow all object embeddings
				"frame-ancestors 'none';", // Disallow framing of the content
		)
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")

		f, err := spaFS.Open(strings.TrimPrefix(path.Clean(r.URL.Path), "/"))
		if err == nil {
			defer f.Close()
		}

		if os.IsNotExist(err) {
			r.URL.Path = "/"
		}

		http.FileServer(http.FS(spaFS)).ServeHTTP(w, r)
	}
}

func SpaRouter() http.Handler {
	spaRouter := hteeteepee.NewRouter("web")
	spaRouter.Handle("/*", spaHandler())
	return spaRouter
}
