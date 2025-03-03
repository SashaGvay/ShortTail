package root

import (
	"context"
	"net/http"
	"strings"
)

func (r *Root) initHttp(_ context.Context) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			http.ServeFile(w, req, "static/index.html")
			return
		}

		alias := strings.TrimPrefix(req.URL.Path, "/")
		dto, err := r.Entity.Url.Service.UnShort(req.Context(), alias)
		if err != nil {
			http.Error(w, "oops", http.StatusNotFound)
			return
		}

		http.Redirect(w, req, "//"+dto.Original, http.StatusFound)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
