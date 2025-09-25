package sniff

import (
	std_bufio "bufio"
	"context"
	"errors"
	"io"
	"strings"

	"github.com/sagernet/sing-box/adapter"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/log"
	E "github.com/sagernet/sing/common/exceptions"
	M "github.com/sagernet/sing/common/metadata"
	"github.com/sagernet/sing/protocol/http"
)

func HTTPHost(ctx context.Context, metadata *adapter.InboundContext, reader io.Reader) error {
	request, err := http.ReadRequest(std_bufio.NewReader(reader))
	if err != nil {
		if errors.Is(err, io.ErrUnexpectedEOF) {
			return E.Cause1(ErrNeedMoreData, err)
		} else {
			return err
		}
	}
	metadata.Protocol = C.ProtocolHTTP
	metadata.Domain = M.ParseSocksaddr(request.Host).AddrString()

	// 记录HTTP流量的详细信息
	log.InfoContext(ctx, "[HTTP Traffic] Method: ", request.Method, " Host: ", request.Host, " URL: ", request.URL.String())

	if metadata.Destination.Port == 80 || metadata.Destination.Port == 443 {
		// 构建完整URL用于记录
		scheme := "http"
		if metadata.Destination.Port == 443 {
			scheme = "https"
		}
		fullURL := strings.TrimSpace(scheme + "://" + request.Host + request.URL.String())
		log.InfoContext(ctx, "[HTTP Traffic] Full URL: ", fullURL)

		if request.URL != nil && request.URL.Path != "" {
			log.InfoContext(ctx, "[HTTP Traffic] Path: ", request.URL.Path)
		}

		if request.URL != nil && request.URL.RawQuery != "" {
			log.InfoContext(ctx, "[HTTP Traffic] Query: ", request.URL.RawQuery)
		}
	}

	return nil
}
