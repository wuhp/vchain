package vchain

import (
    "log"
    "os"
    "fmt"
    "errors"
)

var logger *log.Logger

func SetOutput(path string) error {
    f, e := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if e != nil {
        msg := fmt.Sprintf("Can not open file `%s`, with err `%s`", path, e.Error())
        return errors.New(msg)
    }

    logger = log.New(f, "", log.Ldate | log.Ltime | log.Lmicroseconds)
    return nil
}

func print(category, message string) {
    if logger == nil {
        return
    }

    logger.Println(category, message)
}
