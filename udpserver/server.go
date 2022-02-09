package udpserver

import (
	"context"
	"log"
	"net"
	"sync"
	"time"
)

const (
	Localhost   = "127.0.0.1"
	Multicast   = "224.0.0.101"
	DefaultPort = 2237
)

type UDPServer struct {
	conn   *net.UDPConn
	start  time.Time
	remote *net.UDPAddr
	ctx    context.Context
	r      chan []byte
	w      chan []byte
	log    logger
	wg     sync.WaitGroup
	cancel context.CancelFunc
	stats  stats
}

type stats struct {
	rxMessages int64
	txMessages int64
	started    time.Time
	sync.RWMutex
}

type Status struct {
	RxMessages int64
	TxMessages int64
	Started    time.Time
	Uptime     time.Duration
}

func (s *stats) IncrRxMessages() {
	s.Lock()
	defer s.Unlock()
	s.rxMessages++
}

func (s *stats) IncrTxMessages() {
	s.Lock()
	defer s.Unlock()
	s.txMessages++
}

func (s *stats) GetStats() Status {
	s.RLock()
	defer s.RUnlock()
	return Status{
		RxMessages: s.rxMessages,
		TxMessages: s.txMessages,
		Uptime:     time.Since(s.started),
		Started:    s.started,
	}
}

type logger interface {
	Println(v ...interface{})
}

func NewServer(ctx context.Context, ip string, port int, logger logger) (*UDPServer, error) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	})
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	u := UDPServer{
		conn:   conn,
		start:  time.Now(),
		w:      make(chan []byte),
		r:      make(chan []byte),
		ctx:    ctx,
		log:    logger,
		cancel: cancel,
		stats: stats{
			started: time.Now(),
		},
	}
	u.wg.Add(2)
	go u.reader()
	go u.writer()
	return &u, nil
}

func (u *UDPServer) Close() {
	u.cancel()
	u.conn.Close()

	u.wg.Wait()
	close(u.r)
	close(u.w)
}

func (u *UDPServer) reader() error {
	defer u.wg.Done()

	for {
		select {
		case <-u.ctx.Done():
			u.log.Println("reader: closing")

			return nil
		default:
		}
		buf := make([]byte, 1024)
		rlen, addr, err := u.conn.ReadFromUDP(buf)
		if err != nil {
			u.cancel()
			u.log.Println("close reader:", err)
			return err
		}
		u.remote = addr
		u.stats.IncrRxMessages()
		u.r <- buf[:rlen]
	}

	return nil
}

func (u *UDPServer) writer() {
	defer u.wg.Done()
	for {
		select {
		case <-u.ctx.Done():
			u.log.Println("writer: closing")

			return
		case w, ok := <-u.w:
			if !ok {
				u.cancel()

				return
			}
			_, err := u.conn.WriteToUDP(w, u.remote)
			if err != nil {
				log.Println("Cannot write to Remote UDP Server:" + err.Error())
			}
			u.stats.IncrTxMessages()

		}
	}
}

func (u *UDPServer) Read() chan []byte {
	return u.r
}

func (u *UDPServer) Write(w []byte) {
	u.w <- w
}

func (u *UDPServer) GetStatus() Status {
	return u.stats.GetStats()
}
