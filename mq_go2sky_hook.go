package mq_go2sky_hook

import (
	"context"
	"fmt"
	"time"

	"github.com/SkyAPM/go2sky"
	v3 "github.com/SkyAPM/go2sky/reporter/grpc/language-agent"
	"github.com/qit-team/work"
)

type Hook struct {
	tracer *go2sky.Tracer
}

func NewHook(tracer *go2sky.Tracer) *Hook {
	return &Hook{tracer: tracer}
}

func (h *Hook) BeforeProcess(c *work.ContextHook) error {
	peer := "queue"
	if p, ok := c.Ctx.Value("peer").(string); ok {
		peer = p
	}

	args := fmt.Sprintf("%v", c.Args)
	operateName := fmt.Sprintf("enqueue_%v", c.Topic)
	span, nCtx, err := h.tracer.CreateLocalSpan(c.Ctx, go2sky.WithSpanType(go2sky.SpanTypeLocal), go2sky.WithOperationName(operateName))
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	span.SetPeer(peer)
	span.Tag("args", args)
	span.Tag("topic", c.Topic)
	span.SetSpanLayer(v3.SpanLayer_MQ)
	c.Ctx = nCtx
	c.Ctx = context.WithValue(c.Ctx, operateName, span)
	return nil
}

func (h *Hook) AfterProcess(c *work.ContextHook) error {
	span := c.Ctx.Value(fmt.Sprintf("enqueue_%v", c.Topic)).(go2sky.Span)
	if c.ExecuteTime > 0 {
		span.Tag("elapsed", fmt.Sprintf("%d ms", c.ExecuteTime.Milliseconds()))
	}
	if c.Err != nil {
		timeNow := time.Now()
		span.Error(timeNow, c.Err.Error())
	}
	span.End()
	return nil
}
