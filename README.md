# Log Stash
Push json formatted log to local file system

## Usage
_1_. go get github.com/vonlippmann/logstash<br>
_2_. import "github.com/vonlippmann/logstash" in your code<br>
_3_. make a new instance of Logger:<br>
```go
logger := NewLogStash(&Config{
        LogPath:     "<your specified path>",
        LogKeepDays: 0,
        FileName:    "<your specified file name>",
        CleanLog:    false,
})
```

_4_. then you can use logger in anywhere you need to sink log to the specified path:
```go
logger.Sink(Massage{
    "auth": "pep-pig",
    "age":  "18",
    })
``` 
### Hooks
* you can use hook to postprocess the massage you sinked, such as add some other fields. You can do this simply by register a hook function:
```go
logger.RegisterHook(func(msg Massage)(err error){
    mas["email"]="###@example.com"
    return
})
```

### Example

```go
package main

var logger *Logger

func main(){
    for{
         logger.Sink(Massage{
            "auth": "pep-pig",
            "age":  "18",
        })
    }
}

func addField(massage Massage) (err error) {
    massage["email"] = "###@example.com"
    return
}

func init() {
    logger = NewLogStash(&Config{
        LogPath:     "/var/log/logstash",
        LogKeepDays: 0,
        FileName:    "elk",
        CleanLog:    false,
    })
    logger.RegisterHook(addField)
}

```
