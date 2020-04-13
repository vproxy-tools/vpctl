package vproxy_controller

import (
	"../http_parser"
	"crypto/tls"
	"encoding/json"
	"net"
	"strconv"
	"strings"
	"time"
)

const FailRetryIntervalSecondInt = 2
const FailRetryInterval = FailRetryIntervalSecondInt * time.Second

func buildConnection(watchConf *WatchConfig) (net.Conn, error) {
	if watchConf.UseHttps {
		conf := tls.Config{}
		if watchConf.CaRoots == nil {
			conf.InsecureSkipVerify = true
		} else {
			conf.RootCAs = watchConf.CaRoots
		}
		conn, err := tls.Dial("tcp", watchConf.HostPort, &conf)
		if err != nil {
			return nil, err
		}
		return conn, nil
	} else {
		conn, err := net.Dial("tcp", watchConf.HostPort)
		if err != nil {
			return nil, err
		}
		return conn, nil
	}
}

func watch(resourceType string, watchConf *WatchConfig, reqUri string,
	resourceHandler func([]byte, *meta),
	removeHandler func(string, string),
	cleanHandler func(),
) {
	reqUri = reqUri + "?watch=1&allowWatchBookmarks=true"
	lastVer := 0
	for {
		curVer := doWatch(resourceType, watchConf, reqUri, lastVer, resourceHandler, removeHandler)
		if curVer == -1 {
			cleanHandler()
			lastVer = 0
		} else if curVer != 0 {
			lastVer = curVer
		}
		Debug("wait for %v seconds before reconnecting for %v, ver='%v'", FailRetryIntervalSecondInt, resourceType, lastVer)
		time.Sleep(FailRetryInterval)
	}
}

func doWatch(resourceType string, watchConf *WatchConfig, reqUri string, lastVer int,
	resourceHandler func([]byte, *meta),
	removeHandler func(string, string),
) (curVer int) {
	conn, err := buildConnection(watchConf)
	if err != nil {
		Log("failed to establish connection for watching %v: %v", resourceType, err)
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	if lastVer != 0 {
		reqUri = reqUri + "&resourceVersion=" + strconv.Itoa(lastVer)
	}
	req := "" +
		"GET " + reqUri + " HTTP/1.1\r\n" +
		"Host: " + watchConf.HostPort + "\r\n" +
		"User-Agent: vproxy-k8s-controller\r\n"
	if watchConf.Token != "" {
		req += "Authorization: Bearer " + watchConf.Token + "\r\n"
	}
	req += "\r\n"
	reqBytes := []byte(req)
	n, err := conn.Write(reqBytes)
	if err != nil {
		Log("failed to send watch request for %v: %v", resourceType, err)
		return
	}
	if n != len(reqBytes) {
		Log("not all bytes of the request sent for watching %v", resourceType)
		return
	}
	status, err := http_parser.ReadStatusAndSkipHttpHeaders(conn)
	if err != nil {
		Log("failed when reading status and headers for watching %v: %v", resourceType, err)
		return
	}
	if status == "410" {
		Log("The API server responded with 410 Gone, need to clear all cache of %v and re-fetch them.", resourceType)
		return -1
	} else if status != "200" {
		Log("The API server responded with "+status+", need to clear all cache of %v and re-fetch them.", resourceType)
		return -1
	}

	lastLine := ""
	for {
		bytes, err := http_parser.ReadChunk(conn)
		if err != nil {
			Log("failed when reading chunk for watching %v: %v", resourceType, err)
			return -1
		}
		if len(bytes) == 0 {
			Debug("response bytes len=0 for %v", resourceType)
			return
		}
		strChunk := string(bytes)
		lines := strings.Split(strChunk, "\n")
		if len(lines) != 0 {
			lines[0] = lastLine + lines[0]
			last := lines[len(lines)-1]
			if strings.TrimSpace(last) == "" {
				lastLine = ""
			} else {
				lastLine = last
			}
			lines = lines[:len(lines)-1]
		}
	forEachLine:
		for _, line := range lines {
			line = strings.TrimSpace(line)
			chunk := Chunk{}
			err = json.Unmarshal([]byte(line), &chunk)
			if err != nil {
				Log("received chunk %v is invalid for watching %v: %v", line, resourceType, err)
				return
			}
			if chunk.Type == "ERROR" {
				Debug("got error %v, should clear cache and read again", line)
				return -1
			}

			m := map[string]interface{}{}
			err = json.Unmarshal([]byte(line), &m)
			if err != nil {
				Log("received chunk %v for watching %v is invalid when deserialize into map: %v", line, resourceType, err)
				return
			}
			o, ok := m["object"].(map[string]interface{})
			if !ok {
				Log("received chunk %v for watching %v is invalid in entry 'object': %v", line, resourceType, err)
				return
			}
			oBytes, err := json.Marshal(o)
			if err != nil {
				Log("serialize 'object' for watching %v failed: %v", resourceType, err)
				return
			}

			ver := chunk.Object.Metadata.ResourceVersion
			if ver == "" {
				Log("received chunk %v for watching %v is invalid, no resourceVersion present", line, resourceType)
				return
			}
			verInt, err := strconv.ParseInt(ver, 10, 64)
			if err != nil {
				Log("received chunk %v for watching %v is invalid, resourceVersion is not a number", line, resourceType)
				return
			}
			if chunk.Type == "BOOKMARK" {
				if curVer < int(verInt) {
					curVer = int(verInt)
				}
				continue
			}
			if chunk.Object.Metadata.Namespace == "" {
				Log("received chunk %v for watching %v is invalid, no namespace present", line, resourceType)
				return
			}
			if chunk.Object.Metadata.Name == "" {
				Log("received chunk %v for watching %v is invalid, no name present", line, resourceType)
				return
			}
			if curVer < int(verInt) {
				curVer = int(verInt)
			}

			for _, ns := range watchConf.IgnoreNs {
				if chunk.Object.Metadata.Namespace == ns {
					continue forEachLine
				}
			}
			if len(watchConf.WatchNs) != 0 {
				notWatched := true
				for _, ns := range watchConf.WatchNs {
					if chunk.Object.Metadata.Namespace == ns {
						notWatched = false
						break
					}
				}
				if notWatched {
					continue forEachLine
				}
			}

			Debug("received record: %v", line)

			if chunk.Type == "DELETED" {
				removeHandler(chunk.Object.Metadata.Namespace, chunk.Object.Metadata.Name)
			} else {
				meta := meta{
					ns:   chunk.Object.Metadata.Namespace,
					name: chunk.Object.Metadata.Name,
					ver:  int(verInt),
				}
				resourceHandler(oBytes, &meta)
			}
		}
	}
}
