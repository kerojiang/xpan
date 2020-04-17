package pcsdownload

import (
	"net/url"
	"strconv"
	"xpan/baidupcs"
	"xpan/internal/pcsconfig"
)

// GetLocateDownloadLinks
func GetLocateDownloadLinks(pcs *baidupcs.BaiduPCS, pcspath string) (dlinks []*url.URL, err error) {
	dInfo, pcsError := pcs.LocateDownload(pcspath)
	if pcsError != nil {
		return nil, pcsError
	}

	us := dInfo.URLStrings(pcsconfig.Config.EnableHTTPS)
	if len(us) == 0 {
		return nil, ErrDlinkNotFound
	}

	return us, nil
}

// GetLocatePanLink
func GetLocatePanLink(pcs *baidupcs.BaiduPCS, fsID int64) (dlink string, err error) {
	list, err := pcs.LocatePanAPIDownload(fsID)
	if err != nil {
		return
	}

	var link string
	for k := range list {
		if strconv.FormatInt(fsID, 10) == list[k].FsID {
			link = list[k].Dlink
		}
	}

	if link == "" {
		return "", ErrDlinkNotFound
	}

	dc := pcsconfig.Config.DlinkClient()
	dlink, err = dc.CacheLinkRedirectPr(link)
	return
}
