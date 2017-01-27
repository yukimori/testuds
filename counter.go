package testuds

import (
	"sync/atomic"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
)

type Counter struct {
	c int64
}

type CounterCreator struct {
}

func (c *CounterCreator) CreateState(ctx *core.Context, params data.Map) (core.SharedState, error) {
	cnt := &Counter{}
	if v, ok := params["start"]; ok {
		i, err := data.ToInt(v)
		if err != nil {
			return nil, err
		}
		cnt.c = i - 1
	}
	return cnt, nil
}

func (c *Counter) Terminate(ctx *core.Context) error {
	return nil
}

func (c *Counter) Next() int64 {
	return atomic.AddInt64(&c.c, 1)
}
