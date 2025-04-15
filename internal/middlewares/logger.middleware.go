package middlewares

import (
	"bytes"
	"io"
	"time"

	"github.com/DoHongKien/go-structure/global"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		bodyWriter := &responseBodyWriter{
			ResponseWriter: ctx.Writer,
			body:           bytes.NewBuffer([]byte{}),
		}

		ctx.Writer = bodyWriter

		var requestBodyBytes []byte
		if ctx.Request.Body != nil {
			requestBodyBytes, _ = io.ReadAll(ctx.Request.Body)
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBodyBytes))
		}

		ctx.Next()

		end := time.Now()
		latency := end.Sub(start)

		requestID := ctx.Request.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
			ctx.Request.Header.Set("X-Request-ID", requestID)
		}

		global.Logger.Info("Request",
			zap.String("request_id", requestID),
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.Int("status_code", ctx.Writer.Status()),
			zap.Duration("latency", latency),
			zap.String("request_body", string(requestBodyBytes)),
			zap.String("response_body", bodyWriter.body.String()),
			// zap.String("client_ip", ctx.ClientIP()),
			// zap.String("user_agent", ctx.Request.UserAgent()),
		)
	}
}