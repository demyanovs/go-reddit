package web

import (
	goreddit "github.com/demyanovs/go-reddit"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
	"net/http"
)

func NewHandler(store goreddit.Store) * Handler {
	h := &Handler{
		Mux: chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)
	h.Route("/threads", func(r chi.Router) {
		r.Get("/", h.ThreadsList())
	})

	return h
}

type Handler struct {
	*chi.Mux

	store goreddit.Store 
}

const threadsListHtml = `
<h1>Threads</h1>
<dl>
{{range .Threads}}
	<dt><strong>{{.Title}}</strong></dt>
	<dd>{{.Description}}</dd>
{{end}}
</dl>
`

func (h *Handler) ThreadsList() http.HandlerFunc {
	type data struct {
		Threads []goreddit.Thread
	}

	tmpl := template.Must(template.New("").Parse(threadsListHtml))
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data{Threads: tt})
	}
}