package streamer

import (
	"context"
	"net/http"
	"time"

	tezos "github.com/ecadlabs/go-tezos"
	"github.com/ecadlabs/tezos-streamer/config"
	log "github.com/sirupsen/logrus"
)

type Subscription struct {
	Stream chan interface{}
	filter string
	close  chan *Subscription
}

func (s *Subscription) Close() error {
	errChan := make(chan error)
	// Since Close can be called when nextHash is occuring run the close logic in a go routine to avoid deadlock
	go func() {
		s.close <- s
		close(s.Stream)
		log.WithField("sub", s).Debug("Closed subscription")
		errChan <- nil
	}()
	return <-errChan
}

type Streamer struct {
	newHeadHash chan string
	client      *tezos.Service
	closeSub    chan *Subscription
	newSub      chan *Subscription
	subs        map[string]map[*Subscription]bool
}

func NewStreamer(cfg *config.Config) (*Streamer, error) {
	c, err := tezos.NewRPCClient(http.DefaultClient, cfg.RPCUrl)

	if err != nil {
		return nil, err
	}

	return &Streamer{
		client:      &tezos.Service{Client: c},
		newHeadHash: make(chan string),
		closeSub:    make(chan *Subscription),
		newSub:      make(chan *Subscription),
		subs:        make(map[string]map[*Subscription]bool),
	}, nil
}

func (s *Streamer) Start() {
	cMonitorBlock := make(chan *tezos.MonitorBlock)
	errCount := 0
	defer close(cMonitorBlock)
	go func() {
		for block := range cMonitorBlock {
			// Reset the error count on new block
			errCount = 0
			if block != nil {
				log.WithField("hash", block.Hash).Info("Received new hash from RPC")
				s.newHeadHash <- block.Hash
			}
		}
	}()

	go func() {
		for {
			select {
			case hash := <-s.newHeadHash:
				for _, subs := range s.subs {
					for v := range subs {
						v.Stream <- hash
					}
				}
			case sub := <-s.closeSub:
				filter := sub.filter
				if len(s.subs[filter]) == 0 {
					continue
				}

				if _, ok := s.subs[filter][sub]; ok {
					log.WithField("sub", sub).Debug("Removed subscription from subscriptions map")
					delete(s.subs[filter], sub)
				}
			case sub := <-s.newSub:
				filter := sub.filter

				if s.subs[filter] == nil {
					s.subs[filter] = make(map[*Subscription]bool)
				}

				s.subs[filter][sub] = true
			}
		}
	}()

	for {
		err := s.client.GetMonitorHeads(context.Background(), "main", cMonitorBlock)
		if err != nil {
			errCount++
			log.Errorf("Error encountered while trying to connect to rpc node (err count: %d): %s", errCount, err.Error())
			time.Sleep(time.Duration(errCount) * time.Second)
		}
	}
}

func (s *Streamer) NewHeadSubscription() *Subscription {
	sub := &Subscription{
		Stream: make(chan interface{}),
		close:  s.closeSub,
		filter: "default",
	}
	s.newSub <- sub
	return sub
}

func (s *Streamer) HeadChan() chan string {
	return s.newHeadHash
}
