package workers

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

type apiServer struct {
	lock     sync.Mutex
	managers map[string]*Manager
}

func (s *apiServer) registerManager(m *Manager) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.managers[m.uuid] = m
}

func (s *apiServer) deregisterManager(m *Manager) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.managers, m.uuid)
}

var globalHTTPServer *http.Server

var globalAPIServer = &apiServer{
	managers: map[string]*Manager{},
}

// RegisterAPIEndpoints sets up API server endpoints
func RegisterAPIEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("/stats", globalAPIServer.Stats)
	mux.HandleFunc("/retries", globalAPIServer.Retries)
}

// StartAPIServer starts the API server
func StartAPIServer(port int) {
	mux := http.NewServeMux()
	RegisterAPIEndpoints(mux)

	Logger.Println("APIs are available at", fmt.Sprintf("http://localhost:%v/", port))

	globalHTTPServer = &http.Server{Addr: fmt.Sprint(":", port), Handler: mux}
	if err := globalHTTPServer.ListenAndServe(); err != nil {
		Logger.Println(err)
	}
}

// StopAPIServer stops the API server
func StopAPIServer() {
	if globalHTTPServer != nil {
		globalHTTPServer.Shutdown(context.Background())
	}
}
