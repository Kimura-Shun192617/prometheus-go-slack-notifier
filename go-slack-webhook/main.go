package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "bytes"
)

type Alert struct {
    Status string `json:"status"`
    Alerts []struct {
        Labels struct {
            AlertName string `json:"alertname"`
        } `json:"labels"`
    } `json:"alerts"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    var alert Alert
    err := json.NewDecoder(r.Body).Decode(&alert)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
    if webhookURL == "" {
        log.Println("SLACK_WEBHOOK_URL not set")
        return
    }

    for _, a := range alert.Alerts {
        msg := map[string]string{"text": "Alert: " + a.Labels.AlertName + " (" + alert.Status + ")"}
        jsonMsg, _ := json.Marshal(msg)

        http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonMsg))
    }
    w.WriteHeader(http.StatusOK)
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
