package Parser

import (
	"github.com/pkg/errors"
	"strings"
	"fmt"
)

type Parser struct {
	HttpPakStruct HttpPakStruct
}

func GetNewParser() (*Parser) {
	return new(Parser)
}

type HttpPakStruct struct {
	RequestType     string
	Url             string
	ProtocolVersion string
	Headers         []map[string]string
	BodyBytes       []byte
}

func (*Parser) getHttpStructFromRaw(httpRaw string) (err error) {
	rawLines := strings.Split(httpRaw, "\r\n")
	for i, line := range rawLines {
		fmt.Printf("Line %d = %s\n", i, line)
	}

	return errors.New("nil");
}
