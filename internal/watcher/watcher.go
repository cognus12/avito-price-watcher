package watcher

import (
	"apricescrapper/internal/advt"
	"apricescrapper/pkg/logger"
	"apricescrapper/pkg/slices"
	"apricescrapper/pkg/timer"
)

type watcher struct {
	Url          string
	Subscribers  []string
	Milliseconds int
	logger       logger.Logger
	service      advt.Service
	stop         chan bool
	val          uint64
}

type Watcher interface {
	Run()
	Stop()
	AddSubscriber(email string)
	RemoveSubscriber(email string)
}

func New(url string, subscribers []string, ms int, l logger.Logger, s advt.Service) *watcher {

	return &watcher{Url: url, Subscribers: subscribers, Milliseconds: ms, stop: make(chan bool), logger: l, service: s}
}

func (w *watcher) Run() {
	w.logger.Info("Start observing URL: %v", w.Url)

	w.observ()
}

func (w *watcher) Stop() {
	// to be implemented
}

func (w *watcher) AddSubscriber(email string) {
	w.Subscribers = append(w.Subscribers, email)
}

func (w *watcher) RemoveSubscriber(email string) {
	index := slices.IndexOf(w.Subscribers, email)

	w.Subscribers = append(w.Subscribers[:index], w.Subscribers[index+1:]...)
}

func (w *watcher) observ() {
	timer.SetInterval(func() {
		ad, err := w.service.GetAdInfo(w.Url)

		if err == nil {
			w.updateVal(ad.Price)

			w.logger.Info("Ad checked, current price: %v", w.val)
		}

		if err != nil {
			w.logger.Errorf(err)
		}

	}, w.Milliseconds, true)
}

func (w *watcher) updateVal(v uint64) {
	if w.val == 0 {
		w.val = v
		return
	}

	if w.val != v {
		w.notify(w.val, v)
		w.val = v
	}
}

func (w *watcher) notify(prev, new uint64) {
	for _, s := range w.Subscribers {
		w.logger.Info("Notyfy email %v, prev price: %v, new price: %v", s, prev, new)
	}
}
