package sample_logger

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/hiroykam/goa-sample/sample_error"
	"github.com/inconshreveable/log15"
)

type Parameters struct {
	logTime     string
	requestID   string
	requestBody string
	requestPath string
	uid         string
	message     string
	logger      log15.Logger
}

type Option func(*Parameters)

func LogTime(l string) Option {
	return func(args *Parameters) {
		args.logTime = l
	}
}

func RequestId(id string) Option {
	return func(args *Parameters) {
		args.requestID = id
	}
}

func RequestPath(p string) Option {
	return func(args *Parameters) {
		args.requestPath = p
	}
}

func UID(uid string) Option {
	return func(args *Parameters) {
		args.uid = uid
	}
}

func New(ctx context.Context, opts ...Option) (*Parameters, *sample_error.SampleError) {
	p := &Parameters{
		logTime:     time.Now().Format(time.RFC3339),
		requestID:   middleware.ContextRequestID(ctx),
		requestPath: "",
		requestBody: "",
		uid:         "",
		message:     "",

		logger: log15.New(log15.Ctx{"module": "goa-sample"}),
	}
	p.logger.SetHandler(log15.StreamHandler(os.Stdout, log15.JsonFormat()))

	gc := goa.ContextRequest(ctx)
	j, err := json.Marshal(gc.Payload)
	if err != nil {
		return nil, sample_error.NewSampleError(sample_error.InternalError, err.Error())
	}
	p.requestPath = gc.RequestURI
	p.requestBody = string(j)
	for _, opt := range opts {
		opt(p)
	}
	return p, nil
}

func (p *Parameters) Info(msg string) {
	p.logger.Info(p.message,
		"logTime", p.logTime,
		"RequestID", p.requestID,
		"RequestBody", p.requestBody,
		"RequestPath", p.requestPath,
		"UID", p.uid)
}

func (p *Parameters) Debug(msg string) {
	p.logger.Debug(p.message,
		"logTime", p.logTime,
		"RequestID", p.requestID,
		"RequestBody", p.requestBody,
		"RequestPath", p.requestPath,
		"UID", p.uid)
}

func (p *Parameters) Error(msg string) {
	p.logger.Error(p.message,
		"logTime", p.logTime,
		"RequestID", p.requestID,
		"RequestBody", p.requestBody,
		"RequestPath", p.requestPath,
		"UID", p.uid)
}

func (p *Parameters) Warn(msg string) {
	p.logger.Warn(p.message,
		"logTime", p.logTime,
		"RequestID", p.requestID,
		"RequestBody", p.requestBody,
		"RequestPath", p.requestPath,
		"UID", p.uid)
}

func (p *Parameters) SampleError(err *sample_error.SampleError) {
	if err.Code == sample_error.InternalError {
		p.Error(err.Msg)
	} else {
		p.Warn(err.Msg)
	}
}

