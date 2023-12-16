package main

import (
	"fmt"
	"net"
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	var err error
	http.Handle("/", &helloHandler{})
	//err = http.ListenAndServe(":8083", nil) // IPv4 或 IPv6
	err = http.ListenAndServe("[::1]:8083", nil) // 具体指定，仅 IPv6
	//err = ListenAndServe(":8083", nil) // 重构 ListenAndServe 函数
	if err != nil {
		fmt.Println(err)
	}

}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func ListenAndServe(addr string, handler http.Handler) error {
	srv := &http.Server{Addr: addr, Handler: handler}
	addr = srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp4", addr) // 仅指定 IPv4
	//ln, err := net.Listen("tcp6", addr) // 仅指定 IPv6
	if err != nil {
		return err
	}
	return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}
