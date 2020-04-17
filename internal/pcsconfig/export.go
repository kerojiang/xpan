package pcsconfig

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"time"
	"xpan/baidupcs"
	"xpan/baidupcs/dlinkclient"
	"xpan/pcstable"
	"xpan/pcsutil/converter"
	"xpan/requester"
)

// ActiveUser 获取当前登录的用户
func (c *PCSConfig) ActiveUser() *Baidu {
	if c.activeUser == nil {
		return &Baidu{}
	}
	return c.activeUser
}

// ActiveUserBaiduPCS 获取当前登录的用户的baidupcs.BaiduPCS
func (c *PCSConfig) ActiveUserBaiduPCS() *baidupcs.BaiduPCS {
	if c.pcs == nil {
		c.pcs = c.ActiveUser().BaiduPCS()
	}
	return c.pcs
}

func (c *PCSConfig) httpClientWithUA(ua string) *requester.HTTPClient {
	client := requester.NewHTTPClient()
	client.SetHTTPSecure(c.EnableHTTPS)
	client.SetUserAgent(ua)
	return client
}

// HTTPClient 返回设置好的 HTTPClient
func (c *PCSConfig) HTTPClient() *requester.HTTPClient {
	return c.httpClientWithUA(c.UserAgent)
}

// PCSHTTPClient 返回设置好的 PCS HTTPClient
func (c *PCSConfig) PCSHTTPClient() *requester.HTTPClient {
	return c.httpClientWithUA(c.PCSUA)
}

// PanHTTPClient 返回设置好的 Pan HTTPClient
func (c *PCSConfig) PanHTTPClient() *requester.HTTPClient {
	return c.httpClientWithUA(c.PanUA)
}

// DlinkClient 返回设置好的DlinkClient
func (c *PCSConfig) DlinkClient() *dlinkclient.DlinkClient {
	if c.dc == nil {
		dc := dlinkclient.NewDlinkClient()
		client := c.PanHTTPClient()
		client.SetResponseHeaderTimeout(30 * time.Second)
		dc.SetClient(client)
		c.dc = dc
	}
	return c.dc
}

// NumLogins 获取登录的用户数量
func (c *PCSConfig) NumLogins() int {
	return len(c.BaiduUserList)
}

// AverageParallel 返回平均的下载最大并发量
func (c *PCSConfig) AverageParallel() int {
	return AverageParallel(c.MaxParallel, c.MaxDownloadLoad)
}

// PrintTable 输出表格
func (c *PCSConfig) PrintTable() {
	tb := pcstable.NewTable(os.Stdout)
	tb.SetHeader([]string{"名称", "值", "建议值", "描述"})
	tb.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	tb.SetColumnAlignment([]int{tablewriter.ALIGN_DEFAULT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
	tb.AppendBulk([][]string{
		{"appid", fmt.Sprint(c.AppID), "", "百度 PCS 应用ID"},
		{"cache_size", converter.ConvertFileSize(int64(c.CacheSize), 2), "1KB ~ 256KB", "下载缓存, 如果硬盘占用高或下载速度慢, 请尝试调大此值"},
		{"max_parallel", strconv.Itoa(c.MaxParallel), "50 ~ 500", "下载最大并发量"},
		{"max_upload_parallel", strconv.Itoa(c.MaxUploadParallel), "1 ~ 100", "上传最大并发量"},
		{"max_download_load", strconv.Itoa(c.MaxDownloadLoad), "1 ~ 5", "同时进行下载文件的最大数量"},
		{"max_download_rate", showMaxRate(c.MaxDownloadRate), "", "限制最大下载速度, 0代表不限制"},
		{"max_upload_rate", showMaxRate(c.MaxUploadRate), "", "限制最大上传速度, 0代表不限制"},
		{"savedir", c.SaveDir, "", "下载文件的储存目录"},
		{"enable_https", fmt.Sprint(c.EnableHTTPS), "true", "启用 https"},
		{"user_agent", c.UserAgent, requester.DefaultUserAgent, "浏览器标识"},
		{"pcs_ua", c.PCSUA, "", "PCS 浏览器标识"},
		{"pan_ua", c.PanUA, baidupcs.NetdiskUA, "Pan 浏览器标识"},
		{"proxy", c.Proxy, "", "设置代理, 支持 http/socks5 代理"},
		{"local_addrs", c.LocalAddrs, "", "设置本地网卡地址, 多个地址用逗号隔开"},
	})
	tb.Render()
}
