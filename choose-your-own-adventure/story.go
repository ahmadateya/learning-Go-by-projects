package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`


/* 
  if you are not exporting a type (starts with Capital letter) => its better to not return it (in a func)
  Reason => it will not apear appropriately
*/

// breaks "return structs accept interfaces "
func NewHandler(s Story, t *template.Template) http.Handler {
  if t == nil {
    t = tpl
  }
  return handler{s, t}
}

type handler struct {
  s Story
  t template.Template
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  path := strings.TrimSpace(r.URL.Path)
  if path == "" || path == "/" {
    path = "/intro"
  }
  path = path[1:]

  if chapter, ok := h.s[path]; ok {
    err := tpl.Execute(w, chapter)
    if err != nil {
      log.Printf("%v", err)
      http.Error(w, "Something went wrong..!", http.StatusInternalServerError)
    }
    return
  }
  http.Error(w, "Chapter not found.", http.StatusNotFound)
}


func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story

	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

// remember in maps order is not guaranteed
type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

