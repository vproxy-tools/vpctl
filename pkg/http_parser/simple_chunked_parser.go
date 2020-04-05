package http_parser

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func ReadStatusAndSkipHttpHeaders(conn net.Conn) (string, error) {
	bytes := make([]byte, 1)
	state := 0
	// 0 -> initial, space->0, other->1
	// 1 -> reading before statusCode, space->2, other->1
	// 2 -> ready to read statusCode, space->2, other->3
	// 3 -> read statusCode, space->4, other->3
	// 4 -> expecting \r\n\r\n, \r->5, other->4
	// 5 -> \r, \n->6, other->err
	// 6 -> \n, \r->7, other->4
	// 7 -> \r, \n->break, other->err
	var sb strings.Builder
loop:
	for {
		n, err := conn.Read(bytes)
		if err != nil {
			return "", err
		}
		if n == 0 {
			continue
		}
		b := bytes[0]
		switch state {
		case 0:
			if isWhiteSpace(b) {
				state = 0
			} else {
				state = 1
			}
		case 1:
			if isWhiteSpace(b) {
				state = 2
			} else {
				state = 1
			}
		case 2:
			if isWhiteSpace(b) {
				state = 2
			} else {
				state = 3
				sb.Write([]byte{b})
			}
		case 3:
			if isWhiteSpace(b) {
				state = 4
			} else {
				sb.Write([]byte{b})
				state = 3
			}
		case 4:
			if b == '\r' {
				state = 5
			} else {
				state = 4
			}
		case 5:
			if b == '\n' {
				state = 6
			} else {
				return "", fmt.Errorf("invalid protocol \\r should follow \\n")
			}
		case 6:
			if b == '\r' {
				state = 7
			} else {
				state = 4
			}
		case 7:
			if b == '\n' {
				break loop
			} else {
				return "", fmt.Errorf("invalid protocol \\r should follow \\n")
			}
		default:
			return "", fmt.Errorf("BUG: invalid state %d", state)
		}
	}
	return sb.String(), nil
}

func isWhiteSpace(b byte) bool {
	return b == ' ' || b == '\t'
}

func isHex(b byte) bool {
	return b == '0' ||
		b == '1' ||
		b == '2' ||
		b == '3' ||
		b == '4' ||
		b == '5' ||
		b == '6' ||
		b == '7' ||
		b == '8' ||
		b == '9' ||
		b == 'A' || b == 'a' ||
		b == 'B' || b == 'b' ||
		b == 'C' || b == 'c' ||
		b == 'D' || b == 'd' ||
		b == 'E' || b == 'e' ||
		b == 'F' || b == 'f'
}

func ReadChunk(conn net.Conn) ([]byte, error) {
	l, err := readChunkLen(conn)
	if err != nil {
		return nil, err
	}
	if l == 0 {
		return make([]byte, 0), nil
	}
	err = skipUntilLF(conn)
	if err != nil {
		return nil, err
	}
	bytes, err := readChunkContent(conn, l)
	if err != nil {
		return nil, err
	}
	err = skipUntilLF(conn)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func readChunkLen(conn net.Conn) (int, error) {
	bytes := make([]byte, 1)
	state := 0
	// 0 -> initial state, HEX->1, space->0, other->error
	// 1 -> length, ;|\r|space->break, HEX->1, other->error
	var sb strings.Builder
loop:
	for {
		n, err := conn.Read(bytes)
		if err != nil {
			return 0, err
		}
		if n == 0 {
			continue
		}
		b := bytes[0]
		switch state {
		case 0:
			if isHex(b) {
				sb.Write([]byte{b})
				state = 1
			} else if isWhiteSpace(b) {
				state = 0
			} else {
				return 0, fmt.Errorf("invalid protocol, chunk should start with hex")
			}
		case 1:
			if b == ';' || isWhiteSpace(b) || b == '\r' {
				break loop
			} else if isHex(b) {
				sb.Write([]byte{b})
				state = 1
			} else {
				return 0, fmt.Errorf("invalid protocol, chunk length should be followed by extension or \\r, but got %v", string([]byte{b}))
			}
		}
	}
	n, err := strconv.ParseInt(sb.String(), 16, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid protocol, invalid hex %s %v", sb.String(), err)
	}
	return int(n), nil
}

func skipUntilLF(conn net.Conn) error {
	bytes := make([]byte, 1)
	for {
		n, err := conn.Read(bytes)
		if err != nil {
			return err
		}
		if n == 0 {
			continue
		}
		b := bytes[0]
		if b == '\n' {
			break
		}
	}
	return nil
}

func readChunkContent(conn net.Conn, l int) ([]byte, error) {
	bytes := make([]byte, l)
	off := 0
	for {
		n, err := conn.Read(bytes[off:])
		if err != nil {
			return nil, err
		}
		off += n
		if off < l {
			continue
		} else {
			break
		}
	}
	return bytes, nil
}
