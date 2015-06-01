package network

import (
    "net/http"
    "log"
)

func NewDefaultHttpServer() (*HttpServer) {
    return &HttpServer{}
}

type HttpServer struct {
    routes     []*Route
}

func (h *HttpServer) Start() {
    h.serveServer()
}

func (h *HttpServer) serveServer() {
    log.Println("Started HTTP Server @ Port 8080")

    wh := &WrappedHandler{
        routes: h.routes,
    }
    http.ListenAndServe(":8081", wh)
}

func (h *HttpServer) NewRoute(path string, method string, fn RouteHandler) *Route {
    route := CreateNewRoute(path, method, fn)
    h.routes = append(h.routes, route)

    return route
}

type WrappedHandler struct {
    routes     []*Route
}

func (wh *WrappedHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    route, _, err := MatchingRoute(r.URL.Path, r.Method, nil, wh.routes)

    if err != nil {
        if err == ERR_NO_MATCHING_ROUTE {
            http.NotFound(w, r)
            return
        }

        if err == ERR_NO_MATCHING_METHOD {
            // TODO
            return
        }

        if err == ERR_UNSUPPORTED_CONTENT_FORMAT {
            // TODO
            return
        }
    } else {

        req := NewRequestFromHttp()
        resp := route.Handler(req).(*HttpResponse)

        SendHttpResponse(resp, w, r)
        return
    }
}

func SendHttpResponse(response *HttpResponse, w http.ResponseWriter, r *http.Request) {

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func NewRequestFromHttp() (*HttpRequest) {
    return &HttpRequest{}
}

type HttpRequest struct {

}

type HttpResponse struct {
    contentType     string
    err             error
    Payload         MessagePayload
}

func (r *HttpResponse) GetError() error {
    return r.err
}

func (r *HttpResponse) GetPayload() ([]byte) {
    return r.Payload.GetBytes()
}
