package ratwebserver

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func LogHTTPRequest(r *http.Request) {
	log.Printf("%s request for %s Origin: %s", r.Method, r.URL, r.RemoteAddr)
}

func getFile(p string) string {
	f := fmt.Sprintf("%s/index.html", p)
	return filepath.Join(usedConfiguration.GetHTMLLocation(), f)
}

func GetURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	LogHTTPRequest(r)
	w.Header().Set("Content-Type", "text/html")
	filename := getFile(r.URL.String())
	index, _ := os.Open(filename)
	defer index.Close()
	indexString, _ := ioutil.ReadAll(index)
	w.Write(indexString)
}

func addBodyToHTML(htmlTags string) string {
	f := getFile("/templates")
	htmlTemplate, _ := ioutil.ReadFile(f)
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
