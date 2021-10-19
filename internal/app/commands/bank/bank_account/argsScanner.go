package bank_account

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// todo Наверное необходимо переместить куда-то.

type argsScanner struct {
	reader *bufio.Reader
}

func newArgsScanner(args string) *argsScanner {

	return &argsScanner{
		reader: bufio.NewReader(strings.NewReader(args)),
	}
}

func (s *argsScanner) nextUInt64() (uint64, error) {
	var i uint64
	n, err := fmt.Fscanf(s.reader, "%d", &i)
	if err != nil {
		return 0, err
	}
	if n <= 0 {
		return 0, fmt.Errorf("wrong args")
	}

	return i, nil
}

func (s *argsScanner) stringToEnd() (string, error) {
	var str string
	_, err := fmt.Fscanln(s.reader, &str)
	if err != nil {
		return "", err
	}

	return str, nil
}

func (s *argsScanner) bytesToEnd() ([]byte, error) {
	str, err := s.reader.ReadBytes(0)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return str, nil
}
