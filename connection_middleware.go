import (
	"net"
)

func CompressionMiddleware(c net.Connection) net.Connection {
	return &compressionMiddleware{c}
}
