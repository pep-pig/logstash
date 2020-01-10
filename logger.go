package logstash

import (
    "encoding/json"
    "log"
    "sync"
)

var (
    lock = sync.Mutex{}
)

type Config struct {
    LogPath     string
    LogKeepDays int
    CleanLog    bool
    FileName    string
}

type Massage map[string]interface{}

type Logger struct {
    hook     Hook
    msgQueue chan Massage
    manager  *Manager
}

func (this *Logger) RegisterHook(hook Hook) {
    this.hook = hook
}

func (this *Logger) loop() {
    for msg := range this.msgQueue {
        if shouldDrop := this.hook(msg); shouldDrop {
            log.Printf("will drop massage:%+v\n", msg)
            continue
        }
        jsonMsg, err := json.Marshal(msg)
        if err != nil {
            log.Println("[ERROR]func_loop: failed to marshal request logger:", err)
        } else {
            this.manager.write(jsonMsg)
        }
    }
}

func (this *Logger) Sink(msg Massage) {
    lock.Lock()
    defer lock.Unlock()
    this.msgQueue <- msg
}

func NewLogStash(config ...*Config) *Logger {
    logger := &Logger{
        hook:     DefaultHook,
        msgQueue: make(chan Massage, 1000),
        manager:  &Manager{config: mergeConfig(config...)},
    }
    logger.manager.initLogManager()
    go logger.loop()
    return logger
}

func mergeConfig(config ...*Config) *Config {
    c := &Config{}
    for _, cfg := range config {
        if cfg.LogPath == "" {
            cfg.LogPath = "/var/logger/"
            log.Println("set default logger path: /var/logger")
        }
        if cfg.CleanLog == false {
            log.Println("set logger keep days: ten years")
            cfg.LogKeepDays = 10 * 356
        }
        if cfg.FileName == "" {
            log.Println("set date as file name")
        } else {
            cfg.FileName = "-" + cfg.FileName
        }
        c = cfg
    }
    return c
}
