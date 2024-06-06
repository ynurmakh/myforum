package http

import (
	"flag"
	"forum/internal/business"
	businessrealiz "forum/internal/business/businessRealiz"
	"forum/internal/transport"
	"forum/ui"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

type Transport struct {
	service       businessrealiz.Service
	port          string
	templateCache map[string]*template.Template
	Configs       conf
}

type conf struct {
	CookiesMaxAge int
}

func InitTransport(b business.Business) (transport.Transport, error) {
	var t Transport
	flag.StringVar(&t.port, "p", "8080", "port")
	flag.Parse()
	templateCache, err := newTemplateCache()
	if err != nil {
		return nil, err
	}
	t.templateCache = templateCache

	// todo kek
	log.Println("server started on http://localhost:" + t.port)
	err = http.ListenAndServe(":"+t.port, t.routes())
	if err != nil {
		return nil, err
	}

	return t, nil
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.html",
			"html/partials/*.html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}
