package main

import (
	"fmt"
	"net/http"
	"time"
)

func healthCheck1() bool {

	url := "http://db-svc:9090/isReady"

	client := http.Client{
		Timeout: 2 * time.Second,
	}
	res, err := client.Get(url)

	if err == nil && res.StatusCode == 200 {
		return true
	}
	return false
}

var ready = true

func healthCheck2() bool {
	//Write any code to determine the health of this app service
	return true
}

func checkReady(w http.ResponseWriter, r *http.Request) {
	//custom logic, I am using a mocking service to validate the concept
	if ready {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	} else {
		w.WriteHeader(500)
		w.Write([]byte("failed"))
	}
}

func switchReady(w http.ResponseWriter, r *http.Request) {
	ready = !ready
}

func readyness(w http.ResponseWriter, r *http.Request) {
	if getStatusMonitor().isReady() {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	} else {
		w.WriteHeader(500)
		w.Write([]byte("healthz failed"))
	}
}

func main() {

	mon := getStatusMonitor()
	defer mon.stop()

	//chain all health checks
	mon.addCheck(healthCheck1)
	mon.addCheck(healthCheck2)
	mon.start()

	http.HandleFunc("/healthz", readyness)

	//mock services to make the service not ready
	http.HandleFunc("/isReady", checkReady)
	http.HandleFunc("/switchReady", switchReady)

	fmt.Println("Listening to port 3000 !")
	http.ListenAndServe(":3000", nil)
}
