package plugin

import (
	"github.com/yukimori/testuds"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
)

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


func init() {
	udf.MustRegisterGlobalUDSCreator("my_counter", &CounterCreator{})
}
