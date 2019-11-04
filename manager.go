package logstash

import (
    "io/ioutil"
    "log"
    "os"
    "sort"
    "strings"
    "time"
)

type Manager struct {
    latestDate string
    logFile    *os.File
    config     *Config
}

func (this *Manager) initLogManager() {
    var err error
    _, err = os.Stat(this.config.LogPath)
    if os.IsNotExist(err) {
        err = os.MkdirAll(this.config.LogPath, os.ModePerm)
        if err != nil {
            log.Println("[ERROR]func_initElkLogManager: failed to create logger logFile directory: ", err)
        }
    }
    this.openLogFile()
    this.latestDate = time.Now().Format("2006-01-02")
}

func (this *Manager) openLogFile() {
    file := this.config.LogPath + "/" + time.Now().Format("2006-01-02") + this.config.FileName + ".logger"
    var err error
    this.logFile, err = os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
    if err != nil {
        log.Println("[ERROR]func_openElkLogFile: failed to open logger file: ", err)
    }
}

func (this *Manager) closeLogFile() {
    if err := this.logFile.Close(); err != nil {
        log.Println("[ERROR]func_closeElkLogFile: failed to close logger file:", err)
    }
}

func (this *Manager) cleanLogFile() {
    files := make([]string, 0)
    dir, err := ioutil.ReadDir(this.config.LogPath)
    if err != nil {
        log.Println("[ERROR]func_cleanOldLog: failed to add history logger into logFile queue:", err)
    } else {
        for _, file := range dir {
            if ok := strings.HasSuffix(file.Name(), "-elk.logger"); ok {
                files = append(files, file.Name())
            }
        }
        sort.Sort(sort.Reverse(sort.StringSlice(files)))
        if len(files) > this.config.LogKeepDays {
            for _, file := range files[this.config.LogKeepDays:] {
                if err := os.Remove(this.config.LogPath + file); err != nil {
                    log.Println("[ERROR]func_cleanOldLog: failed to remove file:", err)
                }
            }
        }
    }
}

func (this *Manager) write(msg []byte) {
    if shouldCreateNewFile := time.Now().Format("2006-01-02") != this.latestDate; shouldCreateNewFile {
        this.closeLogFile()
        this.cleanLogFile()
        this.openLogFile()
        this.latestDate = time.Now().Format("2006-01-02")
    }
    if _, err := this.logFile.WriteString(string(msg) + "\n"); err != nil {
        log.Println("[ERROR]func_write failed to write msg: ", err)
    }
}
