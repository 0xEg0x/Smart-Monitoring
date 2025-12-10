package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	reqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_requests_total",
			Help: "Total number of requests",
		},
		[]string{"path"},
	)

	gauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "app_load_value",
			Help: "Simulated load value",
		},
	)
)

func init() {
	prometheus.MustRegister(reqs)
	prometheus.MustRegister(gauge)
}

func main() {
	n8nWebhook := os.Getenv("N8N_WEBHOOK") // ex: http://n8n:5678/webhook/monitor
	threshold := 70.0

	// Rota principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqs.WithLabelValues("/").Inc()
		fmt.Fprintf(w, "Hello - visit /simulate to change metric\n")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		msg := `{"alert":"test from go container"}`
		resp, err := http.Post(os.Getenv("N8N_WEBHOOK"), "application/json", bytes.NewBuffer([]byte(msg)))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(w, "N8N response: %s", string(body))
	})

	// Endpoint para simular valores de carga
	http.HandleFunc("/simulate", func(w http.ResponseWriter, r *http.Request) {
		reqs.WithLabelValues("/simulate").Inc()

		v := r.URL.Query().Get("v")

		var val float64
		if v == "" {
			val = float64(time.Now().Second() % 100)
		} else {
			_, err := fmt.Sscanf(v, "%f", &val)
			if err != nil {
				http.Error(w, "invalid v", 400)
				return
			}
		}

		gauge.Set(val)

		fmt.Fprintf(w, "set load to %.2f\n", val)

		// Se passar do limite, envia alerta ao n8n
		if val > threshold {
			go notifyN8N(n8nWebhook, val)
		}
	})

	// Endpoint de métricas do Prometheus
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func notifyN8N(webhook string, value float64) {
	if webhook == "" {
		log.Println("N8N_WEBHOOK not set, skipping alert")
		return
	}

	msg := fmt.Sprintf(`{"alert": "high load", "value": %.2f}`, value)
	_, err := http.Post(webhook, "application/json", bytes.NewBuffer([]byte(msg)))

	if err != nil {
		log.Printf("Failed to send alert to n8n: %v\n", err)
	} else {
		log.Println("Alert sent to n8n!")
	}
}
