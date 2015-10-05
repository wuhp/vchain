package vchain

import (
    "time"
    "encoding/json"

    "github.com/satori/go.uuid"
)

type Request struct {
    Uuid       string `json:"uuid"`
    ParentUuid string `json:"parent_uuid"`
    Service    string `json:"service"`
    Category   string `json:"category"`
    Sync       bool   `json:"sync"`
    BeginTs    int64  `json:"begin_ts"`
    EndTs      int64  `json:"end_ts"`
}

func NewRequest(puuid, service, category string) *Request {
    r := new(Request)
    r.Uuid = uuid.NewV4().String()
    r.ParentUuid = puuid
    r.Service = service
    r.Category = category
    r.Sync = true
    r.BeginTs = time.Now().UTC().UnixNano()
    
    return r
}

func (r *Request) End() {
    r.EndTs = time.Now().UTC().UnixNano()
}

func (r *Request) EndWithCommit() {
    r.End()
    r.Commit()
}

func (r *Request) Commit() {
    message, _ := json.Marshal(r)
    print("request", string(message))
}
