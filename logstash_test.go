package logstash

import (
    "testing"
)

var (
    logger *Logger
)

func TestLogStash(t *testing.T) {
    for{
        logger.Sink(Massage{
            "auth": "fengjiabin",
            "cop":  "xiaomi",
        })
    }
}

func AddField(massage Massage) (err error) {
    massage["name"] = "fengjb"
    massage["test"] = 1
    massage["program"] = "hello world"
    return
}

func init() {
    logger = NewLogStash(&Config{
        LogPath:     "/Users/fengjb/GoProjects/gopath/src/logstash",
        LogKeepDays: 0,
        FileName:    "falcon",
        CleanLog:    false,
    })
    logger.RegisterHook(AddField)
}
