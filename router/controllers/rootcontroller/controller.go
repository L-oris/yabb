package rootcontroller

import (
	"flag"
	"net/http"

	"github.com/L-oris/yabb/logger"
	"github.com/L-oris/yabb/models/tpl"
	"github.com/gorilla/mux"
)

type Config struct {
	Tpl tpl.Template
	Serve
}

type Controller struct {
	Router *mux.Router
	serve  Serve
	tpl    tpl.Template
}

// New creates a new controller and registers the routes
func New(config *Config) Controller {
	c := Controller{
		serve: config.Serve,
		tpl:   config.Tpl,
	}

	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(c.static())
	router.HandleFunc("/", c.home).Methods("GET")
	router.HandleFunc("/ping", c.ping).Methods("GET")
	router.HandleFunc("/favicon.ico", c.favicon).Methods("GET")

	c.Router = router
	return c
}

// static serves static files
func (c Controller) static() http.Handler {
	var dir string
	flag.StringVar(&dir, "dir", "public/", "the directory to serve files from /public")
	flag.Parse()

	return http.StripPrefix("/static/", http.FileServer(http.Dir(dir)))
}

// home serves the Home page
func (c Controller) home(w http.ResponseWriter, req *http.Request) {
	c.tpl.Render(w, "home.gohtml", nil)
}

// ping is used for health check
func (c Controller) ping(w http.ResponseWriter, req *http.Request) {
	logger.Log.Debug("ping pong request")
	w.Write([]byte("pong"))
}

func (c Controller) favicon(w http.ResponseWriter, req *http.Request) {
	c.serve(w, req, "public/favicon.ico")
}
