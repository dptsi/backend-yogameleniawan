package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/layout.gohtml"))

	t.ExecuteTemplate(writer, "layout.gohtml", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Yoga Meleniawan Pamungkas",
	})
}

func TestTemplateLayout(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Printf(string(body))

}
