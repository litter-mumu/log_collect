package tailf

import (
	"github.com/hpcloud/tail"
)

var tailObj *tail.Tail

func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tailObj, err = tail.TailFile(fileName, config)
	return err
}

func ReadLog() <-chan *tail.Line{
	return tailObj.Lines
}
