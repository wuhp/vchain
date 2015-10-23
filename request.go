package vchain

import (
    "time"
    "fmt"
    "strconv"
    "net/http"
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

func NewRequest(ch *ChainHeader, service, category string) *Request {
    r := new(Request)
    r.Uuid = uuid.NewV4().String()
    r.Service = service
    r.Category = category
    r.BeginTs = time.Now().UTC().UnixNano()
    if ch == nil {
        r.ParentUuid = ""
        r.Sync = true
    } else {
        r.ParentUuid = ch.Uuid
        r.Sync = ch.Sync
    }

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

// Below two methods give a simple solution for http solution
func NewRequestFromHttp(req *http.Request, service, category string) *Request {
    ch := fetchChainHeader(req)
    return NewRequest(ch, service, category)
}

func WrapHttpRequest(req *http.Request, ch *ChainHeader) {
    if ch != nil {
        req.Header.Set("Vchain-Uuid", ch.Uuid)
        req.Header.Set("Vchain-Sync", fmt.Sprintf("%v", ch.Sync))
    }
}

////////////////////////////////////////////////////////////////////////////////

type ChainHeader struct {
    Uuid string
    Sync bool
}

func fetchChainHeader(req *http.Request) *ChainHeader {
    uuid := req.Header.Get("Vchain-Uuid")
    sync := req.Header.Get("Vchain-Sync")
    if uuid == "" || sync == "" {
        return nil
    }

    ch := new(ChainHeader)
    ch.Uuid = uuid
    ch.Sync, _ = strconv.ParseBool(sync)

    return ch
}

func NewChainHeader(uuid string, sync bool) *ChainHeader {
    ch := new(ChainHeader)
    ch.Uuid = uuid
    ch.Sync = sync
    return ch
}
