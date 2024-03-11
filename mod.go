package test_telelog_lib

import "github.com/telenornorway/slf4go"

var log = slf4go.GetLogger()

func Test() {
	log.Info("Hello, Lib!")
}
