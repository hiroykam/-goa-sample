package sample_middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/goadesign/goa"
	"github.com/hiroykam/goa-sample/sample_logger"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	ctx context.Context
}

func (lrw *loggingResponseWriter) Write(buf []byte) (int, error) {
	l, err := sample_logger.NewSampleLooger(lrw.ctx, sample_logger.LogTime(time.Now().Format(time.RFC3339)))
	if err != nil {
		return 0, err
	}
	l.Logger.Debug("response", "body", string(buf))
	return lrw.ResponseWriter.Write(buf)
}

func LogResponse() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			// chain a new logging writer to the current response writer.
			resp := goa.ContextResponse(ctx)
			resp.SwitchWriter(
				&loggingResponseWriter{
					ResponseWriter: resp.SwitchWriter(nil),
					ctx:            ctx,
				})

			// next
			return h(ctx, rw, req)
		}
	}
}
