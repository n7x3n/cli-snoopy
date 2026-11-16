package docs

import (
	"context"
	_ "embed"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//go:embed docs.html
var html []byte

func DocsServer() error {
	ip, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	defer ip.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(html)
	})
	url := fmt.Sprintf("http://%s", ip.Addr().String())
	fmt.Printf("\nDocumentation is live at: %s\n", url)
	fmt.Println("To exit, press Ctrl+C")
	server := &http.Server{Handler: mux}
	go func() {
		server.Serve(ip)
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return server.Shutdown(ctx)
}
