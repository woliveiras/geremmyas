package dashboard

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Server serves the generated dashboard and SSE reload events.
type Server struct {
	Dir      string
	Port     int
	mu       sync.Mutex
	clients  map[chan string]struct{}
}

// NewServer creates a dashboard HTTP server.
func NewServer(dir string, port int) *Server {
	return &Server{Dir: dir, Port: port, clients: map[chan string]struct{}{}}
}

// Run starts the server until interrupted.
func (s *Server) Run() error {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(s.Dir)))
	mux.HandleFunc("/events", s.handleEvents)

	addr := fmt.Sprintf("127.0.0.1:%d", s.Port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("port %d in use: %w (try --port)", s.Port, err)
	}

	srv := &http.Server{Handler: mux}
	fmt.Printf("Dashboard: http://localhost:%d\n", s.Port)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		shutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdown)
	}()

	if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) handleEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}
	ch := make(chan string, 1)
	s.mu.Lock()
	s.clients[ch] = struct{}{}
	s.mu.Unlock()
	defer func() {
		s.mu.Lock()
		delete(s.clients, ch)
		s.mu.Unlock()
	}()
	fmt.Fprintf(w, "data: connected\n\n")
	flusher.Flush()
	for {
		select {
		case <-r.Context().Done():
			return
		case msg := <-ch:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		}
	}
}

// NotifyReload broadcasts reload to SSE clients.
func (s *Server) NotifyReload() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for ch := range s.clients {
		select {
		case ch <- "reload":
		default:
		}
	}
}
