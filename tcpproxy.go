package main

import (
	"fmt"
	"net"
	"os"
	"github.com/golang/glog"
)

type tcpproxy struct {
	config   *Config
	backends *backends
}

func NewTCPProxy(c *Config, backends *backends) *tcpproxy {
	return &tcpproxy{c, backends}
}

func (p *tcpproxy) start() {
	local, err := net.Listen("tcp", p.config.acceptAddr)
	glog.Infof("Listening on %s : ", p.config.acceptAddr)
	if local == nil {
		die("cannot listen: %v", err)
	}
	for {
		conn, err := local.Accept()
		if conn == nil {
			die("accept failed: %v", err)
		}

		remoteAddr := p.backends.Next()
		if "" != remoteAddr {
			go forward(conn, remoteAddr)
		} else {
			glog.Errorf("No host found for service")
			conn.Close()
		}
	}
}

func forward(local net.Conn, remoteAddr string) {
	remote, err := net.Dial("tcp", remoteAddr)
	if remote == nil {
		glog.Fatalf("remote dial failed: %v\n", err)
		return
	}

	localTCP := local.(*net.TCPConn)
	remoteTCP := remote.(*net.TCPConn)

	go copy(localTCP, remoteTCP)
	go copy(remoteTCP, localTCP)
}


func copy(dst, src *net.TCPConn) {
	dst.ReadFrom(src)
	src.Close()
	dst.Close()
}

func die(s string, a ...interface{}) {
	glog.Fatalf("netfwd: %s\n", fmt.Sprintf(s, a))
	os.Exit(2)
}
