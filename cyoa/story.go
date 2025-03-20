package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"
)

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Choose your own adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
        <p>{{.}}</p>
    {{end}}
    <ul>
        {{range .Options}}
            <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
</html>`

func NewHandler(s Story) http.Handler {
	return handler{s: s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

// try running below coomand in terminal
// godoc -http :3030
// then open
// localhost:3030/pkg

func JsonReader(r io.Reader) (Story, error) {
	var story Story
	if err := json.NewDecoder(r).Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

// this struct was created using this
// https://mholt.github.io/json-to-go/
// just give it any json and it will create a go struct for you
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// to add automatic tags to your struct
// select the complete struct
// press cmd+shift+p
// type gotags
// select the option add tags to struct field
// you can uncomment and experiment on the below struct
// type Demo struct {
// 	Name string
// 	Age  int
// }
