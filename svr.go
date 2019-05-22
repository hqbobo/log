package log

import lnet "libs/net"
import "net"

//收到日志后的处理
type LogSvrInterface interface {
	OnRead(data []byte, conn *net.UDPConn, remote *net.UDPAddr) error
}

type logSvr struct {
	h   LogSvrInterface
	svr *lnet.UdpSvr
}

const (
	logSvrIp   = "0.0.0.0"
	logSvrPort = 60000
)

var svr *logSvr

//初始化日志服务器
func LogSvrInit(h LogSvrInterface) error {
	var err error
	if svr == nil {
		svr = new(logSvr)
		svr.h = h
		svr.svr, err = lnet.NewUdpSvr(logSvrIp, logSvrPort, h.OnRead)
		if err != nil {
			return err
		}
	}
	return nil
}
