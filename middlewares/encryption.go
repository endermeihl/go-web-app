package middlewares

import (
	"bytes"
	"go-web-app/logger"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

// 示例加密和解密函数
func encrypt(data []byte) ([]byte, error) {
	// 简单的加密逻辑（这里只是反转字节，实际使用时请使用更安全的加密方法）
	encrypted := make([]byte, len(data))
	for i := range data {
		encrypted[i] = data[len(data)-1-i]
	}
	return encrypted, nil
}

func decrypt(data []byte) ([]byte, error) {
	// 简单的解密逻辑（这里只是反转字节，实际使用时请使用更安全的解密方法）
	decrypted := make([]byte, len(data))
	for i := range data {
		decrypted[i] = data[len(data)-1-i]
	}
	return decrypted, nil
}

func EncryptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解密请求体
		if c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err != nil {
				logger.Logger.Error("Failed to read request body", zap.Error(err))
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			decryptedBytes, err := decrypt(bodyBytes)
			if err != nil {
				logger.Logger.Error("Failed to decrypt request body", zap.Error(err))
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(decryptedBytes))
		}

		// 创建一个新的 responseWriter
		rw := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = rw

		c.Next()

		// 加密响应体
		if c.Writer.Status() == http.StatusOK {
			encryptedBytes, err := encrypt(rw.body.Bytes())
			if err != nil {
				logger.Logger.Error("Failed to encrypt response body", zap.Error(err))
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			c.Writer.Header().Set("Content-Length", string(len(encryptedBytes)))
			c.Writer.WriteHeader(c.Writer.Status())
			c.Writer.Write(encryptedBytes)
		}
	}
}
