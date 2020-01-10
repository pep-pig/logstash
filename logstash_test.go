package logstash

import (
    "testing"
    "time"
)

var (
    logger *Logger
)

func TestLogStash(t *testing.T) {
    logger.Sink(Massage{
        "auth": "pep-pig",
    })
    time.Sleep(time.Second*1)
}

func AddField(massage Massage) (shouldDrop bool) {
    if age, exist := massage["age"]; exist {
        if age.(int) >= 18 {
            massage["young"] = true
        } else {
            massage["young"] = false
        }
    } else {
        return true
    }
    return
}

func init() {
    logger = NewLogStash(&Config{
        LogPath:     "/Users/fengjb/GoProjects/gopath/src/github.com/pep-pig/logstash/",
        LogKeepDays: 0,
        FileName:    "test",
        CleanLog:    false,
    })
    logger.RegisterHook(AddField)
}
