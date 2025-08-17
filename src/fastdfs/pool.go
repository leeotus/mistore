package fastdfs

import (
	"net"
	"sync"
)

// tcpConnBaseInfo 连接信息
type tcpConnBaseInfo struct {
	net.Conn
}

// tcpConnPool 连接池
type tcpConnPool struct {
	addr     string
	maxConns int
	conns    chan *tcpConnBaseInfo
	mu       sync.Mutex
}

// initTcpConnPool 初始化连接池
func initTcpConnPool(addr string, maxConns int) (*tcpConnPool, error) {
	pool := &tcpConnPool{
		addr:     addr,
		maxConns: maxConns,
		conns:    make(chan *tcpConnBaseInfo, maxConns),
	}

	// 预先创建一些连接
	for i := 0; i < 3; i++ {
		conn, err := net.DialTimeout("tcp", addr, TCP_CONN_TIMEOUT)
		if err != nil {
			return nil, err
		}
		pool.conns <- &tcpConnBaseInfo{conn}
	}
	return pool, nil
}

// get 获取连接
func (p *tcpConnPool) get() (*tcpConnBaseInfo, error) {
	select {
	case conn := <-p.conns:
		return conn, nil
	default:
		// 如果没有空闲连接，尝试创建新连接
		p.mu.Lock()
		defer p.mu.Unlock()
		if len(p.conns) < p.maxConns {
			conn, err := net.DialTimeout("tcp", p.addr, TCP_CONN_TIMEOUT)
			if err != nil {
				return nil, err
			}
			return &tcpConnBaseInfo{conn}, nil
		}
		return nil, nil
	}
}

// put 归还连接
func (p *tcpConnPool) put(conn *tcpConnBaseInfo) {
	select {
	case p.conns <- conn:
	default:
		// 连接池已满，关闭连接
		conn.Close()
	}
}

// Destroy 销毁连接池
func (p *tcpConnPool) Destroy() {
	close(p.conns)
	for conn := range p.conns {
		conn.Close()
	}
}
