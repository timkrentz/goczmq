package goczmq

/*
#cgo !windows pkg-config: libczmq
#cgo windows CFLAGS: -I/usr/local/include
#cgo windows LDFLAGS: -L/usr/local/lib -lczmq
#include "czmq.h"

zactor_t *Gossip_new (char *name) { zactor_t *gossip = zactor_new(zgossip, name); return gossip; }
*/
import "C"

import (
	"unsafe"
)

// Gossip actors use a gossip protocol for decentralized configuration management.
// Gossip nodes form a loosely connected network that publishes and redistributed
// name/value tuples.  A network of Gossip actors will eventually achieve
// a consistent state
type Gossip struct {
	zactor_t *C.struct__zactor_t
}

// NewGossip creates a new Gossip actor
func NewGossip(name string) *Gossip {
	g := &Gossip{}
	g.zactor_t = C.Gossip_new(C.CString(name))
	return g
}

// Bind binds the gossip service to a specified endpoint
func (g *Gossip) Bind(endpoint string) error {
	rc := C.zstr_sendm(unsafe.Pointer(g.zactor_t), C.CString("BIND"))
	if rc == -1 {
		return ErrActorCmd
	}

	rc = C.zstr_send(unsafe.Pointer(g.zactor_t), C.CString(endpoint))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Verbose sets the gossip actor to log information to stdout.
func (g *Gossip) Verbose() error {
	rc := C.zstr_send(unsafe.Pointer(g.zactor_t), C.CString("VERBOSE"))
	if rc == -1 {
		return ErrActorCmd
	}

	return nil
}

// Destroy destroys the gossip actor.
func (g *Gossip) Destroy() {
	C.zactor_destroy(&g.zactor_t)
}
