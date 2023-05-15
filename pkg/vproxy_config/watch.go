package vproxy_config

import (
	"encoding/json"
	"fmt"
	"github.com/vproxy-tools/vpctl/pkg/http_parser"
	"net"
	"strconv"
)

type HealthCheckEventChannelMessage struct {
	Evt HealthCheckEvent
	Err error
}

func WatchHealthCheck(chnl chan *HealthCheckEventChannelMessage, stop chan bool) error {
	hostport := fmt.Sprintf("%s:%d", GetHost(), GetHttpPort())
	host := GetHost()
	if GetHttpPort() != 80 {
		host += strconv.Itoa(GetHttpPort())
	}
	conn, err := net.Dial("tcp", hostport)
	if err != nil {
		return err
	}
	req := []byte("" +
		"GET /api/v1/watch/server-group/-/server/-/health-check HTTP/1.1\r\n" +
		"Host: " + host + "\r\n" +
		"\r\n")
	n, err := conn.Write(req)
	if err != nil {
		return err
	}
	if n != len(req) {
		return fmt.Errorf("request not sent completely")
	}
	statusCode, err := http_parser.ReadStatusAndSkipHttpHeaders(conn)
	if err != nil {
		return err
	}
	if statusCode != "200" {
		return fmt.Errorf("response is not 200, %v", statusCode)
	}
	stopped := false
	go func() {
		for {
			b := <-stop
			if b {
				// do stop
				stopped = true
				_ = conn.Close()
				return
			}
		}
	}()
	go func() {
		for {
			bytes, err := http_parser.ReadChunk(conn)
			if stopped {
				return
			}
			msg := HealthCheckEventChannelMessage{}
			if err != nil {
				msg.Err = err
				chnl <- &msg
				return
			}
			wire := HealthCheckEventOnWire{}
			err = json.Unmarshal(bytes, &wire)
			if err != nil {
				msg.Err = err
				chnl <- &msg
				return
			}
			msg.Evt.from(&wire)
			chnl <- &msg
		}
	}()
	return nil
}
