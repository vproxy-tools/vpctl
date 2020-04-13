package vproxy_controller

import (
	"sync"
	"time"
)

const FailRetryDelaySeconds = 1
const FailRetryDelay = FailRetryDelaySeconds * time.Second
const MaxRetry = 5

type ResourcePool struct {
	tl     []*TcpLb
	socks5 []*Socks5Server
	dns    []*DnsServer
	ups    []*Upstream
	sg     []*ServerGroup
	secg   []*SecurityGroup
	ep     []*Endpoints
	ck     []*Secret
}

type NamespacedResourcePool struct {
	pools map[string]*ResourcePool
}

type PoolEventHandler interface {
	Name() string
	Handle(*NamespacedResourcePool) error
}

type Pool interface {
	SetAndTriggerEventHandler(PoolEventHandler)

	GetNamespaces() []string

	UpdateTcpLb(string, func([]*TcpLb) ([]*TcpLb, bool))
	UpdateSocks5Server(string, func([]*Socks5Server) ([]*Socks5Server, bool))
	UpdateDnsServer(string, func([]*DnsServer) ([]*DnsServer, bool))
	UpdateUpstream(string, func([]*Upstream) ([]*Upstream, bool))
	UpdateServerGroup(string, func([]*ServerGroup) ([]*ServerGroup, bool))
	UpdateSecurityGroup(string, func([]*SecurityGroup) ([]*SecurityGroup, bool))
	UpdateEndpoints(string, func([]*Endpoints) ([]*Endpoints, bool))
	UpdateSecret(string, func([]*Secret) ([]*Secret, bool))

	ClearPendingTcpLb(string)
	ClearPendingSocks5Server(string)
	ClearPendingDnsServer(string)
	ClearPendingUpstream(string)
	ClearPendingServerGroup(string)
	ClearPendingSecurityGroup(string)
	ClearPendingEndpoints(string)
	ClearPendingSecret(string)
}

func NewPool() Pool {
	return &_pool{
		pool:    &NamespacedResourcePool{pools: map[string]*ResourcePool{}},
		handler: nil,
		lock:    sync.RWMutex{},
	}
}

type _pool struct {
	pool    *NamespacedResourcePool
	handler PoolEventHandler
	lock    sync.RWMutex
}

func (p *_pool) SetAndTriggerEventHandler(handler PoolEventHandler) {
	p.lock.Lock()
	p.handler = handler
	p.lock.Unlock()

	p.trigger()
}

func (p *_pool) GetNamespaces() []string {
	ret := make([]string, len(p.pool.pools))
	index := 0
	for k := range p.pool.pools {
		ret[index] = k
		index += 1
	}
	return ret
}

func (p *_pool) ensureAndGetNamespace(namespace string) *ResourcePool {
	pool := p.pool.pools[namespace]
	if pool == nil {
		pool = &ResourcePool{
			tl:     make([]*TcpLb, 0),
			socks5: make([]*Socks5Server, 0),
			dns:    make([]*DnsServer, 0),
			ups:    make([]*Upstream, 0),
			sg:     make([]*ServerGroup, 0),
			secg:   make([]*SecurityGroup, 0),
			ep:     make([]*Endpoints, 0),
			ck:     make([]*Secret, 0),
		}
		p.pool.pools[namespace] = pool
	}
	return pool
}

func (p *_pool) trimNamespace(namespace string) {
	pool := p.pool.pools[namespace]
	if pool == nil {
		return
	}
	if len(pool.tl) == 0 &&
		len(pool.socks5) == 0 &&
		len(pool.dns) == 0 &&
		len(pool.ups) == 0 &&
		len(pool.sg) == 0 &&
		len(pool.secg) == 0 &&
		len(pool.ep) == 0 &&
		len(pool.ck) == 0 {
		delete(p.pool.pools, namespace)
	}
}

func (p *_pool) trigger() {
	p.lock.RLock()
	defer p.lock.RUnlock()

	if p.handler == nil {
		return
	}

	failTimes := 0
	for {
		err := p.handler.Handle(p.pool)
		if err == nil {
			break
		}
		if failTimes >= MaxRetry {
			Log("%s handles resource pool failed, will not retry: %v", p.handler.Name(), err)
			break
		} else if failTimes == 0 {
			Log("%s handles resource pool failed, will retry later: %v", p.handler.Name(), err)
		} else {
			unit := "time"
			if failTimes > 1 {
				unit = unit + "s"
			}
			Log("%s handles resource pool failed, will retry later (already retried %v %v): %v", p.handler.Name(), failTimes, unit, err)
		}
		failTimes += 1
		time.Sleep(FailRetryDelay)
	}
}
