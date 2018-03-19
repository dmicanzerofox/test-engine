package transcons

import (
	"context"
	"fmt"
	"github.com/dm03514/test-engine/ids"
	"github.com/dm03514/test-engine/results"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

type IntEqual struct {
	UsingProperty string `mapstructure:"using_property"`
	ToEqual       int    `mapstructure:"to_equal"`
	Type          string
}

func (ie IntEqual) Evaluate(ctx context.Context, r results.Result) results.Result {
	v, err := r.ValueOfProperty(ie.UsingProperty)
	log.WithFields(log.Fields{
		"component":      ie.Type,
		"using_property": ie.UsingProperty,
		"to_equal":       ie.ToEqual,
		"against":        v.Int(),
		"execution_id":   ctx.Value(ids.Execution("execution_id")),
	}).Info("Evaluate()")
	if err != nil {
		return results.ErrorResult{
			From: r,
			Err:  err,
		}
	}

	if v.Int() != ie.ToEqual {
		return results.ErrorResult{
			From: r,
			Err:  fmt.Errorf("%+v != %+v", ie.ToEqual, v.Int()),
		}
	}
	return r
}

func NewIntEqualFromMap(m map[string]interface{}) (TransCon, error) {
	var ie IntEqual
	err := mapstructure.Decode(m, &ie)
	return ie, err
}

type StringEqual struct {
	UsingProperty string `mapstructure:"using_property"`
	ToEqual       string `mapstructure:"to_equal"`

	Type string
}

func (se StringEqual) Evaluate(ctx context.Context, r results.Result) results.Result {
	v, err := r.ValueOfProperty(se.UsingProperty)
	log.WithFields(log.Fields{
		"component":      se.Type,
		"using_property": se.UsingProperty,
		"to_equal":       se.ToEqual,
		"against":        v.String(),
		"execution_id":   ctx.Value(ids.Execution("execution_id")),
	}).Info("Evaluate()")
	if err != nil {
		return results.ErrorResult{
			From: r,
			Err:  err,
		}
	}
	if v.String() != se.ToEqual {
		return results.ErrorResult{
			From: r,
			Err: fmt.Errorf("%q != %q, expected %q, received %q",
				se.ToEqual, v.String(), se.ToEqual, v.String()),
		}
	}
	return r
}

func NewStringEqualFromMap(m map[string]interface{}) (TransCon, error) {
	var se StringEqual
	err := mapstructure.Decode(m, &se)
	return se, err
}
