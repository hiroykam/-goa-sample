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

type SampleLogger struct {
	logTime     string
	requestID   string
	requestBody string
	requestPath string
	uid         string
	message     string
	Logger      log15.Logger
}

type Option func(*SampleLogger)

func LogTime(l string) Option {
	return func(args *SampleLogger) {
		args.logTime = l
	}
}

func RequestId(id string) Option {
	return func(args *SampleLogger) {
		args.requestID = id
	}
}

func RequestPath(p string) Option {
	return func(args *SampleLogger) {
		args.requestPath = p
	}
}

func UID(uid string) Option {
	return func(args *SampleLogger) {
		args.uid = uid
	}
}

func NewSampleLooger(ctx context.Context, opts ...Option) (*SampleLogger, *sample_error.SampleError) {
	p := &SampleLogger{
		logTime:     time.Now().Format(time.RFC3339),
		requestID:   middleware.ContextRequestID(ctx),
		requestPath: "",
		requestBody: "",
		uid:         "",
		message:     "",

		Logger: log15.New(log15.Ctx{"module": "goa-sample"}),
	}
	p.Logger.SetHandler(log15.StreamHandler(os.Stdout, log15.JsonFormat()))

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

func (p *SampleLogger) Info(msg string) {
	p.Logger.Info(msg,
		"logTime", p.logTime,
		"RequestID", p.requestID,
		"RequestBody", p.requestBody,
		"RequestPath", p.requestPath,
		"UID", p.uid)
}

func (p *SampleLogger) Debug(msg string) {
	p.Logger.Debug(msg,
		"logTime", p.logTime,
		"RequestID", p.requestID,
		"RequestBody", p.requestBody,
		"RequestPath", p.requestPath,
		"UID", p.uid)
}

func (p *SampleLogger) Error(msg string) {
	p.Logger.Error(msg,
		"logTime", p.logTime,
		"RequestID", p.requestID,
		"RequestBody", p.requestBody,
		"RequestPath", p.requestPath,
		"UID", p.uid)
}

func (p *SampleLogger) Warn(msg string) {
	p.Logger.Warn(msg,
		"logTime", p.logTime,
		"RequestID", p.requestID,
		"RequestBody", p.requestBody,
		"RequestPath", p.requestPath,
		"UID", p.uid)
}

func (p *SampleLogger) SampleError(err *sample_error.SampleError) {
	if err.Code == sample_error.InternalError {
		p.Error(err.Msg)
	} else {
		p.Warn(err.Msg)
	}
}
