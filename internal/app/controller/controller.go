package controller

import (
	"../../../pkg/vproxy_controller"
	"../../pkg/versions"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func Main(args []string) int {
	if len(args) == 0 {
		fmt.Println("controller: " + versions.CTRL)
		return 0
	}
	if len(args) == 1 {
		cmd := args[0]
		switch cmd {
		case "help":
			fallthrough
		case "--help":
			fallthrough
		case "-help":
			fallthrough
		case "-h":
			fmt.Println(HelpStr)
			return 0
		case "version":
			fallthrough
		case "--version":
			fallthrough
		case "-version":
			fmt.Println(versions.CTRL)
			return 0
		default:
			break
		}
	}

	labels := map[string]string{}
	useHttps := true
	hostport := os.Getenv("$KUBERNETES_SERVICE_HOST") + ":" + os.Getenv("$KUBERNETES_SERVICE_PORT")
	tokenFileName := "/var/run/secrets/kubernetes.io/serviceaccount/token"
	caFileName := "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
	for _, arg := range args {
		if strings.HasPrefix(arg, "--labels=") {
			labelsStr := arg[len("--labels="):]
			err := json.Unmarshal([]byte(labelsStr), &labels)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "invalid labels: %v\n", err)
				return 1
			}
		} else if strings.HasPrefix(arg, "--k8s-api-server=") {
			hp := arg[len("--k8s-api-server="):]
			if strings.HasPrefix(hp, "http://") {
				useHttps = false
				hostport = hp[len("http://"):]
				if !strings.Contains(hostport, ":") {
					hostport = hostport + ":80"
				}
			} else if strings.HasPrefix(hp, "https://") {
				useHttps = true
				hostport = hp[len("https://"):]
				if !strings.Contains(hostport, ":") {
					hostport = hostport + ":443"
				}
			} else {
				_, _ = fmt.Fprintln(os.Stderr, "invalid url for k8s-api-server, should start with http:// or https://")
				return 1
			}
		} else if strings.HasPrefix(arg, "--token=") {
			tokenFileName = arg[len("--token="):]
		} else if strings.HasPrefix(arg, "--ca=") {
			caFileName = arg[len("--ca="):]
		} else {
			_, _ = fmt.Fprintln(os.Stderr, "unknown command '"+arg+"'")
			return 1
		}
	}

	// check labels
	if len(labels) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "invalid labels: empty")
		return 1
	}
	// check tokenFile
	token := ""
	if tokenFileName != "" {
		tokenBytes, err := ioutil.ReadFile(tokenFileName)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "invalid token %s: %v\n", tokenFileName, err)
			return 1
		}
		token = string(tokenBytes)
		token = strings.TrimSpace(token)
	}
	// check caFile
	var tlsConf *tls.Config = nil
	var roots *x509.CertPool = nil
	if useHttps {
		tlsConf = &tls.Config{}
		if caFileName == "" {
			tlsConf.InsecureSkipVerify = true
		} else {
			roots = x509.NewCertPool()
			caBytes, err := ioutil.ReadFile(caFileName)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "invalid ca %s: %v\n", caFileName, err)
				return 1
			}
			ok := roots.AppendCertsFromPEM(caBytes)
			if !ok {
				_, _ = fmt.Fprintf(os.Stderr, "invalid ca %s: invalid format, requires x509 pem\n", caFileName)
				return 1
			}
			tlsConf.RootCAs = roots
		}
	}
	// check host port
	httpClient := &http.Client{}
	if tlsConf != nil {
		tr := &http.Transport{
			IdleConnTimeout: 2 * time.Second,
		}
		tr.TLSClientConfig = tlsConf
		httpClient.Transport = tr
	}
	var protocol string
	if useHttps {
		protocol = "https"
	} else {
		protocol = "http"
	}
	url := protocol + "://" + hostport + "/healthz"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "creating http request failed: GET %s, error: %v\n", url, err)
		return 1
	}
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "requesting k8s failed: %s, error: %v\n", url, err)
		return 1
	}
	if resp.StatusCode != 200 {
		_, _ = fmt.Fprintf(os.Stderr, "requesting k8s failed: %s, response code not 200: %d\n", url, resp.StatusCode)
		return 1
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "requesting k8s failed: %s, reading body failed: %v\n", url, err)
		return 1
	}
	respBody := strings.TrimSpace(string(respBytes))
	if respBody != "ok" {
		_, _ = fmt.Fprintf(os.Stderr, "requesting k8s failed: %s, response body is not 'ok': %s\n", url, respBody)
		return 1
	}

	// launch
	connInfo := vproxy_controller.WatchConfig{
		UseHttps: useHttps,
		HostPort: hostport,
		Token:    token,
		CaRoots:  roots,
		Labels:   labels,
	}
	vproxy_controller.Launch(connInfo)

	for {
		time.Sleep(3600 * time.Second)
	}
	//noinspection GoUnreachableCode
	return 0
}
