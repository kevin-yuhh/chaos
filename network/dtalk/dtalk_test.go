package dtalk

import "testing"

func TestSendTextMsg(t *testing.T) {
	Warning("测试信息", PRIORITY_0, "tronins.web.1")
	SendText("今天处理了 n 个 tx", "tronins.web.1")
}
