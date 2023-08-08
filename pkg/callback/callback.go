package callback

// This file handles spin-up and take-down of the Callback server required by
// three-legged authentication.

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"

	"github.com/pkg/browser"
)

type CallbackData struct {
	AuthCode string
	State    string
}

var cbData chan CallbackData
var cbServer *http.Server
var cbState string

// CatchCode() spins up a callback server to catch the AuthCode thrown by an
// Auth0 three-legged authorisation request, and cleans up after itself once the
// code has been received.
func CatchCode(url string, state string, cbURL string) (*CallbackData, error) {
	log.Printf("Request AuthCode from Auth0 server")
	callbackDone := &sync.WaitGroup{}
	callbackDone.Add(1)

	cbServer := initCallbackServer(cbURL, callbackDone, state)

	browser.OpenURL(url)

	callbackData := <-cbData

	log.Printf("Shutting down callback server")
	if err := cbServer.Shutdown(context.Background()); err != nil {
		return nil, err
	}
	callbackDone.Wait()
	log.Printf("Callback server successfully closed")

	if callbackData.AuthCode == "" {
		fmt.Println("*ERROR*: Authentication failed or rejected; cannot continue!")
		return nil, fmt.Errorf("authorisation failed: %s", callbackData.State)
	}

	fmt.Println("Access authorised. Thank you.")
	return &callbackData, nil
}

// initCallbackServer starts a callback server in a goroutine, and returns
// a link to the server, allowing us to close it down when it has caught
// the expected AuthCode value.
func initCallbackServer(cbURL string, wg *sync.WaitGroup, state string) *http.Server {
	cb, err := url.Parse(cbURL)
	if err != nil {
		panic(err)
	}

	log.Printf("Initialising Callback Server on %s", cb.Path)

	cbData = make(chan CallbackData)
	cbState = state

	_, port, _ := net.SplitHostPort(cb.Host)
	cbServer = &http.Server{Addr: ":" + port}
	http.HandleFunc(cb.Path, callbackListener)

	go func() {
		defer wg.Done()

		log.Println("Callback Server started")
		if err := cbServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}

	}()

	return cbServer
}

// AuthCallbackHandler listens at the CallbackEndpoint for the reply from the
// API Auth0 server containing the AuthCode.
func callbackListener(w http.ResponseWriter, r *http.Request) {
	m, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	code := m.Get("code")
	state := m.Get("state")
	errorName := m.Get("error")
	errorDesc := m.Get("error_description")

	if state != cbState {
		log.Printf("Unexpected hit on callback server")
		log.Printf(" - invalid state = '%s' (expected '%s')", state, cbState)
		return
	}

	if len(errorName) > 0 {
		io.WriteString(w, WWW_ERROR)
		log.Printf("ERROR: Authorisation failed on server")
		cbData <- CallbackData{"", errorName + ": " + errorDesc}
		return
	}

	if len(code) > 0 {
		io.WriteString(w, WWW_CALLBACK)
		log.Printf("AuthCode received from server")
		cbData <- CallbackData{code, state}
	}
}
