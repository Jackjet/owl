package main

import (
	"owl/common/types"

	"github.com/wuyingsong/tcp"
)

type callback struct {
}

func (cb *callback) OnConnected(conn *tcp.TCPConn) {
	lg.Info("callback:%s connected", conn.GetRemoteAddr().String())
}

//链接断开回调
func (cb *callback) OnDisconnected(conn *tcp.TCPConn) {
	lg.Info("callback:%s disconnect ", conn.GetRemoteAddr().String())
}

//错误回调
func (cb *callback) OnError(err error) {
	lg.Error("callback: %s", err)
}

//消息处理回调
func (cb *callback) OnMessage(conn *tcp.TCPConn, p tcp.Packet) {
	defer func() {
		if r := recover(); r != nil {
			lg.Error("Recovered in OnMessage", r)
		}
	}()
	pkt := p.(*tcp.DefaultPacket)
	lg.Info("receive %v %v", types.MsgTextMap[pkt.Type], string(pkt.Body))
	switch pkt.Type {
	case types.MsgAgentSendTimeSeriesData, types.MsgRepeaterPostTimeSeriesData:
		// TODO：是否需要考虑超时？
		repeater.buffer <- pkt.Body
	default:
		lg.Error("Unknown Option: %v", pkt.Type)
	}
}
