package ratwebserver

import (
	"fmt"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, ErrorString string) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(fmt.Sprintf("<html><body><b>%s<b></body></html>", ErrorString)))
}
