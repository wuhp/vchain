package vchain

import (
    "os"
    "log"
    "fmt"
    "bytes"
    "net/http"
)

type client struct {
    Server string
    Secret string
}

var conn *client

type message struct {
    Event   string      `json:"event"`
    Payload interface{} `json:"payload"`
}

func send(body []byte) {
    cli := &http.Client{}

    req, err := http.NewRequest(
        "POST",
        fmt.Sprintf("http://%s/v1/vchain/%s", conn.Server, conn.Secret),
        bytes.NewBuffer(body),
    )

    if err != nil {
        log.Printf("Failed to create http request, %s\n", err.Error())
        log.Printf("Failed to upload message, %s\n", string(body))
        return
    }

    req.Header.Set("Content-Type", "application/json")
    res, err := cli.Do(req)
    if err != nil {
        log.Printf("Failed to do http request, %s\n", err.Error())
        log.Printf("Failed to upload message, %s\n", string(body))
        return
    }

    if res.StatusCode != http.StatusOK {
        log.Printf("Bad http response, status = %d\n", res.StatusCode)
        log.Printf("Failed to upload message, %s\n", string(body))
        return
    }
}

func getHostname() string {
    name, err := os.Hostname()
    if err != nil {
        name = "unknown"
    }

    return name
}
