package ratwebserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// LoadIndex sends back the index as the response
func LoadIndex(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")

	index, err := os.Open("index.html")
	if err != nil {
		w.Write([]byte("<html><body><a href=\"/\">Back to Homepage</a></body></html>"))
		return
	}
	defer index.Close()
	indexString, _ := ioutil.ReadAll(index)
	w.Write(indexString)
}

func addBodyToHTML(htmlTags string) string {
	htmlTemplate, err := ioutil.ReadFile("./templates/index.html")
	if err != nil {
		fmt.Println("Couldn't Find file ./templates/index.html")
		return fmt.Sprintf("<html><body>%s</body></html>", htmlTags)
	}

	lines := strings.Split(string(htmlTemplate), "\n")
	lines = append(lines, htmlTags)
	lines = append(lines, "</body>\n</html>")
	return strings.Join(lines, "\n")
}

// SuccessResponse returns the successful response
func SuccessResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	html := addBodyToHTML(fmt.Sprintf("<b>%s</b>", "Command Successful"))
	w.Write([]byte(html))
}

// ErrorResponse creates a page which tells you what your error is
func ErrorResponse(w http.ResponseWriter, ErrorString string) {
	w.Header().Set("Content-Type", "text/html")
	html := addBodyToHTML(fmt.Sprintf("<b>%s</b>", ErrorString))
	w.Write([]byte(html))
}
