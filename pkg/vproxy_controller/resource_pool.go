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
	RegisterEventHandler(PoolEventHandler)

	GetNamespaces() []string

	UpdateTcpLb(string, func([]*TcpLb) ([]*TcpLb, int))
	UpdateSocks5Server(string, func([]*Socks5Server) ([]*Socks5Server, int))
	UpdateDnsServer(string, func([]*DnsServer) ([]*DnsServer, int))
	UpdateUpstream(string, func([]*Upstream) ([]*Upstream, int))
	UpdateServerGroup(string, func([]*ServerGroup) ([]*ServerGroup, int))
	UpdateSecurityGroup(string, func([]*SecurityGroup) ([]*SecurityGroup, int))
	UpdateEndpoints(string, func([]*Endpoints) ([]*Endpoints, int))
	UpdateSecret(string, func([]*Secret) ([]*Secret, int))
}

func NewPool() Pool {
	return &_pool{
		pool:     &NamespacedResourcePool{pools: map[string]*ResourcePool{}},
		handlers: make([]PoolEventHandler, 0),
		lock:     sync.RWMutex{},
	}
}

type _pool struct {
	pool     *NamespacedResourcePool
	handlers []PoolEventHandler
	lock     sync.RWMutex
}

func (p *_pool) RegisterEventHandler(handler PoolEventHandler) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.handlers = append(p.handlers, handler)
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

	if len(p.handlers) == 0 {
		return
	}

	todo := make([]PoolEventHandler, len(p.handlers))
	for i, h := range p.handlers {
		todo[i] = h
	}
	failTimes := 0
	for {
		failed := make([]PoolEventHandler, 0)
		for _, h := range todo {
			err := h.Handle(p.pool)
			if err != nil {
				if failTimes >= MaxRetry {
					Log("%s handles resource pool failed, will not retry: %v", h.Name(), err)
					break
				} else if failTimes == 0 {
					Log("%s handles resource pool failed, will retry later: %v", h.Name(), err)
				} else {
					unit := "time"
					if failTimes > 1 {
						unit = unit + "s"
					}
					Log("%s handles resource pool failed, will retry later (already retried %v %v): %v", h.Name(), failTimes, unit, err)
				}
				failed = append(failed, h)
				failTimes += 1
			}
		}
		todo = failed
		if len(todo) == 0 {
			break
		}
		time.Sleep(FailRetryDelay)
	}
}

func (p *_pool) UpdateTcpLb(namespace string, f func([]*TcpLb) ([]*TcpLb, int)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	tl, ver := f(rp.tl)
	ls := make([]*TcpLb, 0)
	for _, t := range tl {
		if t.M.Pending && t.M.Version < ver {
			continue
		}
		ls = append(ls, t)
	}
	rp.tl = ls
	p.trimNamespace(namespace)
	p.lock.Unlock()
	p.trigger()
}

func (p *_pool) UpdateSocks5Server(namespace string, f func([]*Socks5Server) ([]*Socks5Server, int)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	socks5, ver := f(rp.socks5)
	ls := make([]*Socks5Server, 0)
	for _, s := range socks5 {
		if s.M.Pending && s.M.Version < ver {
			continue
		}
		ls = append(ls, s)
	}
	rp.socks5 = ls
	p.trimNamespace(namespace)
	p.lock.Unlock()
	p.trigger()
}

func (p *_pool) UpdateDnsServer(namespace string, f func([]*DnsServer) ([]*DnsServer, int)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	dns, ver := f(rp.dns)
	ls := make([]*DnsServer, 0)
	for _, d := range dns {
		if d.M.Pending && d.M.Version < ver {
			continue
		}
		ls = append(ls, d)
	}
	rp.dns = ls
	p.trimNamespace(namespace)
	p.lock.Unlock()
	p.trigger()
}

func (p *_pool) UpdateUpstream(namespace string, f func([]*Upstream) ([]*Upstream, int)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	ups, ver := f(rp.ups)
	ls := make([]*Upstream, 0)
	for _, u := range ups {
		if u.M.Pending && u.M.Version < ver {
			continue
		}
		ls = append(ls, u)
	}
	rp.ups = ls
	p.trimNamespace(namespace)
	p.lock.Unlock()
	p.trigger()
}

func (p *_pool) UpdateServerGroup(namespace string, f func([]*ServerGroup) ([]*ServerGroup, int)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	sg, ver := f(rp.sg)
	ls := make([]*ServerGroup, 0)
	for _, s := range sg {
		if s.M.Pending && s.M.Version < ver {
			continue
		}
		ls = append(ls, s)
	}
	rp.sg = ls
	p.trimNamespace(namespace)
	p.lock.Unlock()
	p.trigger()
}

func (p *_pool) UpdateSecurityGroup(namespace string, f func([]*SecurityGroup) ([]*SecurityGroup, int)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	secg, ver := f(rp.secg)
	ls := make([]*SecurityGroup, 0)
	for _, s := range secg {
		if s.M.Pending && s.M.Version < ver {
			continue
		}
		ls = append(ls, s)
	}
	rp.secg = ls
	p.trimNamespace(namespace)
	p.lock.Unlock()
	p.trigger()
}

func (p *_pool) UpdateEndpoints(namespace string, f func([]*Endpoints) ([]*Endpoints, int)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	ep, ver := f(rp.ep)
	ls := make([]*Endpoints, 0)
	for _, e := range ep {
		if e.M.Pending && e.M.Version < ver {
			continue
		}
		ls = append(ls, e)
	}
	rp.ep = ls
	p.trimNamespace(namespace)
	p.lock.Unlock()
	p.trigger()
}

func (p *_pool) UpdateSecret(namespace string, f func([]*Secret) ([]*Secret, int)) {
	p.lock.Lock()
	rp := p.ensureAndGetNamespace(namespace)
	ck, ver := f(rp.ck)
	ls := make([]*Secret, 0)
	for _, c := range ck {
		if c.M.Pending && c.M.Version < ver {
			continue
		}
		ls = append(ls, c)
	}
	rp.ck = ls
	p.trimNamespace(namespace)
	p.lock.Unlock()
	p.trigger()
}
