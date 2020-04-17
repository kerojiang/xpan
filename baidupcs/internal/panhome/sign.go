package panhome

import (
	"github.com/iikira/Baidu-Login/bdcrypto"
	"xpan/baidupcs/netdisksign"
	"xpan/pcsutil/converter"
)

// SignRes
type (
	// SignRes 签名结果
	SignRes interface {
		Sign() string
		Timestamp() string
	}

	signRes struct {
		sign      string
		timestamp string
	}
)

// Sign
func (sr *signRes) Sign() string {
	return sr.sign
}

// Timestamp
func (sr *signRes) Timestamp() string {
	return sr.timestamp
}

// Signature
func (ph *PanHome) Signature() (sign SignRes, err error) {
	err = ph.getSignInfo()
	if err != nil {
		return nil, err
	}

	o := netdisksign.Sign2(ph.sign3, ph.sign1)
	signed := bdcrypto.Base64Encode(o)
	return &signRes{
		sign:      converter.ToString(signed),
		timestamp: ph.timestamp,
	}, nil
}
