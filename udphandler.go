package log

import (
	"bytes"
	"fmt"
	oLog "github.com/go-playground/log"
	"libs/net"
)

func NewUdpHandler() *UdpHandler {
	h := new(UdpHandler)
	return h
}

// CustomHandler is your custom handler
type UdpHandler struct {
	timestampFormat string
	writer          *UdpWriter
}

func (this *UdpHandler) SetSvr(name, ip string, port int) {
	if this.writer == nil {
		this.writer = new(UdpWriter)
	}
	this.writer.svrIp = ip
	this.writer.svrPort = port
	this.writer.name = name
}

type UdpWriter struct {
	svrIp   string
	svrPort int
	name    string
}

func (this *UdpWriter) SetSvr(ip string, port int) {
	this.svrIp = ip
	this.svrPort = port
}

func (this *UdpWriter) Write(b []byte) (n int, err error) {
	if this.svrPort >= 0 {
		if len(b) > 1 {
			if e := net.UdpSend(fmt.Sprintf("%s:%d", this.svrIp, this.svrPort), []byte(this.name+"|"+string(b))); e != nil {
				fmt.Println(e)
			}
		}
	}
	return 0, nil
}

// SetTimestampFormat sets Console's timestamp output format
// Default is : "2006-01-02T15:04:05.000000000Z07:00"
func (this *UdpHandler) SetTimestampFormat(format string) {
	this.timestampFormat = format
}

// Log accepts log entries to be processed
func (this *UdpHandler) Log(e oLog.Entry) {
	b := new(bytes.Buffer)
	b.Reset()
	//b.WriteString(this.name + "|" + e.Timestamp.Format(this.timestampFormat) + ":" + "[" + e.Level.String() + "]" + e.Message)
	b.WriteString(e.Timestamp.Format(this.timestampFormat) + "[" + e.Level.String() + "]" + e.Message)
	for _, f := range e.Fields {
		fmt.Fprintf(b, " %s=%v", f.Key, f.Value)
	}
	this.writer.Write(b.Bytes())
}
