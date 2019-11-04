package logstash

type Hook func(massage Massage) (err error)

var DefaultHook = Handle

func Handle(massage Massage) (err error) {
    return
}
