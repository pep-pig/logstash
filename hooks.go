package logstash

type Hook func(massage Massage) (shouldDrop bool)

var DefaultHook = Handle

func Handle(massage Massage) (shouldDrop bool) {
    return false
}
