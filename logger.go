package noblocklog

import (
    "io"
    "time"
)

type Config struct {
    writer io.Writer
}

type Logger interface {
    Log(args ...interface{})
    Logf(format string, args ...interface{})
}

func New(conf *Config) Logger {
    return l{}
}

type messages struct {
    time time.Time
    args []interface{}
}

type l struct {
    messages string
}

func (l) Log(args ...interface{}) {
    
}

func (l) Logf(format string, args ...interface{}) {

}





