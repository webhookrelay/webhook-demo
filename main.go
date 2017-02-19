package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// Port - default port to start application on
var Port = ":8080"

// ReceivedWebhook - keeps info about received webhook
type ReceivedWebhook struct {
	Payload    string
	ReceivedAt time.Time
}

func main() {
	// preparing HTTP server
	srv := &http.Server{Addr: Port, Handler: http.DefaultServeMux}

	// all received webhooks will be stored in this variable
	var webhooks []*ReceivedWebhook
	mu := &sync.RWMutex{}

	go func() {
		fmt.Println("Listening on ", Port)
		fmt.Println("Press enter to shutdown server")
		fmt.Scanln()
		log.Println("Shutting down server...")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatalf("could not shutdown: %v", err)
		}
	}()

	// handler to display all received webhooks
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Received webhooks:")
		mu.RLock()
		for _, v := range webhooks {
			fmt.Fprintln(w, fmt.Sprintf("%s: %s", v.ReceivedAt.String(), v.Payload))
		}
		mu.RUnlock()
	})

	// incomming webhook handler
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		bd, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		// appending webhook to received list
		mu.Lock()
		webhooks = append(webhooks, &ReceivedWebhook{
			ReceivedAt: time.Now(),
			Payload:    string(bd),
		})
		mu.Unlock()
		w.WriteHeader(http.StatusOK)
	})

	// starting server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
