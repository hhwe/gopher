// Package pool 管理用户定义的一组资源
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool 管理一组可以安全地在多个 goroutine 间
// 共享的资源。被管理的资源必须
// 实现 io.Closer 接口
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed 表示请求(Acquire)了一个
// 已经关闭的池
var ErrPoolClosed = errors.New("pool has been closed")

// New 创建一个用来管理资源的池。
// 这个池需要一个可以分配新资源的函数,
// 并规定池的大小
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}
	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 检查是否有空闲的资源
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	// 因为没有空闲资源可用,所以提供一个新资源
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release 将一个使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	// 保证本操作和 Close 操作的安全
	p.m.Lock()
	defer p.m.Unlock()
	// 如果池已经被关闭,销毁这个资源
	if p.closed {
		r.Close()
		return
	}
	select {
	// 试图将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:", "In Queue")
	// 如果队列已满,则关闭这个资源
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close 会让资源池停止工作,并关闭所有现有的资源
func (p *Pool) Close() {
	// 保证本操作与 Release 操作的安全
	p.m.Lock()
	defer p.m.Unlock()
	// 如果 pool 已经被关闭,什么也不做
	if p.closed {
		return
	}
	// 将池关闭
	p.closed = true
	// 在清空通道里的资源之前,将通道关闭
	// 如果不这样做,会发生死锁
	close(p.resources)
	// 关闭资源
	for r := range p.resources {
		r.Close()
	}
}
