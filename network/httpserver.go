package network

import (
    "net/http"
    "log"
    "html/template"
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
    wh := &WrappedHandler{
        routes: h.routes,
    }
    err := http.ListenAndServe(":8081", wh)
    if err != nil {
        log.Fatal(err)
    }
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

type Person struct {
    Name string
}

func SendHttpResponse(response *HttpResponse, w http.ResponseWriter, r *http.Request) {
    t := template.New("newtemplate")

    t, _ = t.Parse(response.Payload.String())

    t.Execute(w, response.GetTemplateModel())
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
    TemplateModel   interface{}
    Payload         MessagePayload
}

func (r *HttpResponse) GetError() error {
    return r.err
}

func (r *HttpResponse) GetPayload() ([]byte) {
    return r.Payload.GetBytes()
}

func (r *HttpResponse) GetContentType() (string) {
    return r.contentType
}

func (r *HttpResponse) GetTemplateModel() (interface{}) {
    return r.TemplateModel
}