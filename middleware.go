package tornote

import (
	"fmt"
	"net/http"
	"strings"
)

func RedirectToHTTPSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		proto := req.Header.Get("x-forwarded-proto")
		if !strings.EqualFold("https", proto) {
			http.Redirect(res, req, fmt.Sprintf("https://%s%s", req.Host, req.URL), http.StatusPermanentRedirect)
			return
		}
		next.ServeHTTP(res, req)
	})
}
