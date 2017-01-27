package plugin

import (
	"github.com/yukimori/testuds"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
)


func init() {
	udf.MustRegisterGlobalUDSCreator("my_counter", &testuds.CounterCreator{})
}
