package http

import (
	"flag"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"forum/internal/business"
	"forum/internal/models"
	"forum/internal/transport"
	"forum/ui"

	businessrealiz "forum/internal/business/businessRealiz"

	"gopkg.in/yaml.v2"
)

type Transport struct {
	service       *businessrealiz.Service
	port          string
	templateCache map[string]*template.Template
	configs       *configType
	UserId        int
	User          *models.User
}

type configType struct {
	CookiesMaxAge int `yaml:"CookieMaxAge"`
}

func InitTransport(b business.Business) (transport.Transport, error) {
	var t Transport
	flag.StringVar(&t.port, "p", "8080", "port")
	flag.Parse()
	t.service = b.(*businessrealiz.Service)

	templateCache, err := newTemplateCache()
	if err != nil {
		return nil, err
	}
	t.templateCache = templateCache

	conf, err := configParce()
	if err != nil {
		return nil, err
	}
	t.configs = conf

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

func configParce() (*configType, error) {
	c := &configType{}
	file, err := os.ReadFile("./internal/transport/http/config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
