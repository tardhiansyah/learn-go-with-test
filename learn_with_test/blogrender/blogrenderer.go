package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

type PostRenderer struct {
	templ *template.Template
}

type PostViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags                                     []string
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	// extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	// parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
