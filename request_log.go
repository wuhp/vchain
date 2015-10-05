package vchain

import (
    "fmt"
    "time"
    "encoding/json"
)

type RequestLog struct {
    Uuid      string `json:"uuid"`
    Timestamp int64  `json:"timestamp"`
    Log       string `json:"log"`
}

func Log(r *Request, format string, v ...interface{}) {
    rlog := new(RequestLog)
    rlog.Uuid = r.Uuid
    rlog.Timestamp = time.Now().UTC().UnixNano()
    rlog.Log = fmt.Sprintf(format, v...)

    message, _ := json.Marshal(rlog)
    print("request-log", string(message))
}
