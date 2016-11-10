package noblocklog

const (
    Info = iota
    Error
    Warning
)

var dft = New(&Config{
    Info,
    Error,
    Warning,
})





