package vchain

import (
    "time"
    "encoding/json"

    "github.com/satori/go.uuid"
)

////////////////////////////////////////////////////////////////////////////////

func Connect(server, secret string) error {
    // http solution should not create a connection
    // TBD consider tcp connection
    conn = new(client)
    conn.Server = server
    conn.Secret = secret

    return nil
} 

////////////////////////////////////////////////////////////////////////////////

type Service struct {
    Uuid       string `json:"uuid"`
    Category   string `json:"category"`
    Instance   string `json:"instance"`
    Hostname   string `json:"hostname"`
    StartTs    int64  `json:"start_ts"`
    StopTs     int64  `json:"stop_ts"`
}

type Request struct {
    Uuid           string            `json:"uuid"`
    Service        *Service          `json:"service"`
    ParentUuid     string            `json:"parent_uuid"`
    Category       string            `json:"category"`
    BeginTs        int64             `json:"begin_ts"`
    EndTs          int64             `json:"end_ts"`
    BeginMetadata  map[string]string `json:"begin_metadata,omitempty"`
    EndMetadata    map[string]string `json:"end_metadata,omitempty"`
}

func StartService(category, instance string) *Service {
    s := new(Service)
    s.Uuid = uuid.NewV4().String()
    s.Category = category
    s.Instance = instance
    s.Hostname = getHostname()
    s.StartTs = time.Now().UTC().UnixNano()

    m := new(message)
    m.Event = "service.start"
    m.Payload = s

    ms := make([]*message, 0)
    ms = append(ms, m)
    body, _ := json.Marshal(ms)
    send(body)

    return s
}

func (s *Service) Stop() {
    s.StopTs = time.Now().UTC().UnixNano()

    m := new(message)
    m.Event = "service.stop"
    m.Payload = s

    ms := make([]*message, 0)
    ms = append(ms, m)
    body, _ := json.Marshal(ms) 
    send(body)
}

func (s *Service) NewRequest(parentUuid, category string, metadata map[string]string) *Request {
    r := new(Request)
    r.Uuid = uuid.NewV4().String()
    r.ParentUuid = parentUuid
    r.Service = s
    r.Category = category
    r.BeginTs = time.Now().UTC().UnixNano()
    r.BeginMetadata = metadata

    m := new(message)
    m.Event = "request.begin"
    m.Payload = r

    ms := make([]*message, 0)
    ms = append(ms, m)
    body, _ := json.Marshal(ms)
    send(body) 
    
    return r
}

func (r *Request) Done(metadata map[string]string) {
    r.EndTs = time.Now().UTC().UnixNano()
    r.EndMetadata = metadata

    m := new(message)
    m.Event = "request.end"
    m.Payload = r

    ms := make([]*message, 0)
    ms = append(ms, m)
    body, _ := json.Marshal(ms)
    send(body)
}
