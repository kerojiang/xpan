package dlinkclient

import (
	"net/url"
	"xpan/baidupcs/expires/cachemap"
	"xpan/requester"
)

const (
	// DlinkHost 服务器
	DlinkHost = "dlink.iikira.com"

	// OperationReg
	OperationReg = "注册分享信息"
	// OperationList
	OperationList = "获取目录下的文件列表"
	// OperationRedirect
	OperationRedirect = "重定向"
	// OperationRedirectPr
	OperationRedirectPr = "重定向(pr)"
)

// DlinkClient 客户端
type DlinkClient struct {
	client     *requester.HTTPClient
	cacheOpMap cachemap.CacheOpMap
}

// NewDlinkClient
func NewDlinkClient() *DlinkClient {
	return &DlinkClient{}
}

func (dc *DlinkClient) lazyInit() {
	if dc.client == nil {
		dc.client = requester.NewHTTPClient()
	}
}

// SetClient
func (dc *DlinkClient) SetClient(client *requester.HTTPClient) {
	dc.client = client
}

func (dc *DlinkClient) genShareURL(method string, param map[string]string) *url.URL {
	dlinkURL := url.URL{
		Scheme: "https",
		Host:   DlinkHost,
		Path:   "/api/v1.1/pan/share/" + method,
	}

	if param != nil {
		uv := url.Values{}
		for k := range param {
			uv.Set(k, param[k])
		}
		dlinkURL.RawQuery = uv.Encode()
	}

	return &dlinkURL
}

func (dc *DlinkClient) genCgiBinURL(method string, param map[string]string) *url.URL {
	cgiBinURL := url.URL{
		Scheme: "https",
		Host:   DlinkHost,
		Path:   "/cgi-bin/" + method,
	}

	if param != nil {
		uv := url.Values{}
		for k := range param {
			uv.Set(k, param[k])
		}
		cgiBinURL.RawQuery = uv.Encode()
	}

	return &cgiBinURL
}
