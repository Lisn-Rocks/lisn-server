package index

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "database/sql"
    "path"
    "html/template"

    "github.com/sharpvik/lisn-backend/config"
)


// TemplateData is a struct used for template generation.
type TemplateData struct {
    CssData            template.HTML
    BodyContents    template.HTML
    JavaScript        template.HTML
}


// MainHTMLTemplate is a string read from /templates/base.html -- it is the most
// basic HTML template used in the project.
var CSSData, _ = ioutil.ReadFile(
    path.Join(config.AppsFolder, "index", "static", "style.html"),
)


// MainHTMLTemplate is a string read from /templates/base.html -- it is the most
// basic HTML template used in the project.
var MainHTMLTemplate, _ = ioutil.ReadFile(
    path.Join(config.TemplatesFolder, "base.html"),
)


// Serve function is used to serve the main index of songs to the user.
func Serve(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    data := TemplateData{
        template.HTML( string(CSSData) ),
        template.HTML( FormBodyContents(db) ),
        template.HTML(""),
    }

    tmpl, err := template.New("index").Parse( string(MainHTMLTemplate) )

    if err != nil {
        fmt.Println(err)
        return
    }

    err = tmpl.Execute(w, data)

    if err != nil {
        fmt.Println(err)
        return
    }
}


// FormBodyContents returns an HTML list of songs we have in the database.
func FormBodyContents(db *sql.DB) (cont string) {
    rows, _ := db.Query("SELECT title, author FROM songs")

    var title, author string

    cont += "<ol>"

    for rows.Next() {
        rows.Scan(&title, &author)
        cont += fmt.Sprintf("<li><i>%s</i> by %s</li>", title, author)
    }

    cont += "</ol>"
    return
}
