// @author AlphaSnow

package g

import (
	"github.com/samber/lo"
	"sync"
)

type Container interface {
	Has(string) bool
	Get(string) interface{}
	Set(string, interface{})
}

type container struct {
	instances map[string]interface{}
	rwmutex   *sync.RWMutex
}

func (c *container) Has(k string) bool {
	return lo.HasKey(c.instances, k)
}

func (c *container) Get(k string) interface{} {
	c.rwmutex.RLock()
	defer c.rwmutex.RUnlock()
	return lo.ValueOr(c.instances, k, nil)
}

func (c *container) Set(k string, v interface{}) {
	c.rwmutex.Lock()
	defer c.rwmutex.Unlock()
	c.instances[k] = v
}

var _ Container = (*container)(nil)

func NewContainer() Container {
	return &container{
		instances: map[string]interface{}{},
		rwmutex:   new(sync.RWMutex),
	}
}
