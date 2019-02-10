package sample_middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/goadesign/goa"
	"github.com/hiroykam/goa-sample/sample_logger"
)

func LogRequest(verbose bool, sensitiveHeaders ...string) goa.Middleware {
	var suppressed map[string]struct{}
	if len(sensitiveHeaders) > 0 {
		suppressed = make(map[string]struct{}, len(sensitiveHeaders))
		for _, sh := range sensitiveHeaders {
			suppressed[strings.ToLower(sh)] = struct{}{}
		}
	}

	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			r := goa.ContextRequest(ctx)
			l, err := sample_logger.NewSampleLooger(ctx, sample_logger.LogTime(time.Now().Format(time.RFC3339)))
			if err != nil {
				return err
			}

			l.Logger.Info("started", "method", r.Method, "url", r.URL.String(), "from", from(req),
				"ctrl", goa.ContextController(ctx), "action", goa.ContextAction(ctx))
			if verbose {
				if len(r.Header) > 0 {
					logCtx := make([]interface{}, 2*len(r.Header))
					i := 0
					keys := make([]string, len(r.Header))
					for k := range r.Header {
						keys[i] = k
						i++
					}
					sort.Strings(keys)
					i = 0
					for _, k := range keys {
						v := r.Header[k]
						logCtx[i] = k
						if _, ok := suppressed[strings.ToLower(k)]; ok {
							logCtx[i+1] = "<hidden>"
						} else {
							logCtx[i+1] = interface{}(strings.Join(v, ", "))
						}
						i = i + 2
					}
					fmt.Println(logCtx)
					l.Logger.Info("headers", "ctx", logCtx)
				}
				if len(r.Params) > 0 {
					logCtx := make([]interface{}, 2*len(r.Params))
					i := 0
					for k, v := range r.Params {
						logCtx[i] = k
						logCtx[i+1] = interface{}(strings.Join(v, ", "))
						i = i + 2
					}
					l.Logger.Info("params", "ctx", logCtx)
				}
				if r.ContentLength > 0 {
					if mp, ok := r.Payload.(map[string]interface{}); ok {
						logCtx := make([]interface{}, 2*len(mp))
						i := 0
						for k, v := range mp {
							logCtx[i] = k
							logCtx[i+1] = interface{}(v)
							i = i + 2
						}
						j, err := json.Marshal(logCtx)
						if err != nil {
							return err
						}
						l.Info(string(j))
					} else {
						// Not the most efficient but this is used for debugging
						js, err := json.Marshal(r.Payload)
						if err != nil {
							js = []byte("<invalid JSON>")
						}
						goa.LogInfo(ctx, "payload", "raw", string(js))
					}
				}
			}
			hErr := h(ctx, rw, req)
			resp := goa.ContextResponse(ctx)
			msg := fmt.Sprintf("status: %d", resp.Status)
			if code := resp.ErrorCode; code != "" {
				l.Error(msg)
			} else {
				l.Info(msg)
			}
			return hErr
		}
	}
}

func from(req *http.Request) string {
	if f := req.Header.Get("X-Forwarded-For"); f != "" {
		return f
	}
	f := req.RemoteAddr
	ip, _, err := net.SplitHostPort(f)
	if err != nil {
		return f
	}
	return ip
}
