package server

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

type Route interface {
	Serve(http.ResponseWriter, *http.Request)
}

type route struct {
	routes []routePack
}

func newRoutePack(method, pattern string, handler http.HandlerFunc) routePack {
	return routePack{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type routePack struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func (ro *route) Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range ro.routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}

type ctxKey struct{}

func NewRoute(routes []routePack) Route {

	return &route{routes}
}
