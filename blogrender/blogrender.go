package blogrender

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/russross/blackfriday"
)

type Post struct {
   Title        string 
   Description  string 
   Body         string 
   Tags         []string
}

func (p Post) SanitisedTitle() string {
    return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

var (
    //go:embed "templates/*"
    postTemplates embed.FS
)

type PostRenderer struct {
    templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
    templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
    if err != nil {
        return nil, err
    }
    return &PostRenderer{templ}, nil

}
// io.Writer is an interface that implements the method Write. 
// Because we are passing a buffer that implements the method Write, we can use it as a parameter.
func (r *PostRenderer) Render(w io.Writer, p Post) error {

    // Convert MARKDOWN to HTML using blackfriday pacakge.
    p.Body = string(blackfriday.MarkdownBasic([]byte(p.Body)))

    return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}
func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error { 
    return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}
