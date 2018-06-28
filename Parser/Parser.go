package Parser

import (
	"github.com/pkg/errors"
	"strings"
)

type HttpPakStruct struct {
	RequestType     string
	Url             string
	ProtocolVersion string
	Headers         map[string]string
	BodyBytes       []byte
}

type parserStruct struct {
	HttpPakStruct
}

func (Parser *parserStruct) GetHttpStructFromRaw(httpRaw string) (err error) {
	rawLines := strings.Split(httpRaw, "\r\n")
	rawcount := len(rawLines)
	loopContinue := true;
	Parser.Headers = make(map[string]string)
	for i, line := range rawLines {
		if loopContinue == false {
			break;
		}

		switch i {
		case 0: //First line
			parts := strings.Split(line, " ")
			if (len(parts) == 3) {
				Parser.RequestType = parts[0]
				Parser.Url = parts[1]
				Parser.ProtocolVersion = parts[2]
			} else {
				err = errors.New("http package struct is error")
				loopContinue = false
			}
		case rawcount - 1: //Last line
			if line != "" {
				Parser.BodyBytes = []byte(line)
			}
		}
		if i > 0 && i < rawcount-1 {
			if line != "" {
				flagIndex := strings.Index(line, ":")
				if (flagIndex == -1) {
					err = errors.New("http header struct is error")
					loopContinue = false
				} else {
					key := string(line[0:flagIndex])
					val := string(line[flagIndex+1:])
					Parser.Headers[key] = val
				}
			}
		}
	}
	return;
}

func GetNewParser() (*parserStruct) {
	return new(parserStruct)
}
