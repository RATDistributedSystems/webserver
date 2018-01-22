package ratwebserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

// ErrorResponse creates a page which tells you what your error is
func ErrorResponse(w http.ResponseWriter, ErrorString string) {
	w.Header().Set("Content-Type", "text/html")
	html := addBodyToHTML(fmt.Sprintf("<b>%s</b>", ErrorString))
	w.Write([]byte(html))
}
