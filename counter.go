package testuds

import (
	"sync/atomic"
	"gopkg.in/sensorbee/sensorbee.v0/core"
)

type Counter struct {
	c int64
}


func (c *Counter) Terminate(ctx *core.Context) error {
	return nil
}

func (c *Counter) Next() int64 {
	return atomic.AddInt64(&c.c, 1)
}
