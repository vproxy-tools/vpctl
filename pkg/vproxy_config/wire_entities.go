package vproxy_config

import "vproxy_client_model"

type HealthCheckEventOnWire struct {
	Server      vproxy_client_model.Server      `json:"server"`
	ServerGroup vproxy_client_model.ServerGroup `json:"serverGroup"`
}
