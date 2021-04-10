package design

import (
	"testing"
)

func TestAliyunClientAdapter_CreateServer(t *testing.T) {
	var a ICreateServer = &AliYunClientAdapter{
		Client: AliYunClient{},
	}
	a.CreateServer(1.0, 2.0)
}

func TestQiNiuClientAdapter_CreateServer(t *testing.T) {
	var a ICreateServer = &QiNiuClientAdapter{
		Client: QiNiuClient{},
	}

	a.CreateServer(1.0, 2.0)
}
