package observer

import (
	"apricescrapper/internal/advt"
	"apricescrapper/internal/subscription"
	"apricescrapper/internal/watcher"
	"apricescrapper/pkg/logger"
)

type UrlsMap map[string]watcher.Watcher

type Observer struct {
	Urls                UrlsMap
	subscribtionService subscription.Service
	advtService         advt.Service
	logger              logger.Logger
}

func NewObserver(subscribtionService subscription.Service, advtService advt.Service, logger logger.Logger) *Observer {
	return &Observer{
		subscribtionService: subscribtionService,
		Urls:                make(map[string]watcher.Watcher),
		advtService:         advtService,
		logger:              logger}
}

func (o *Observer) Prepare() *Observer {
	m := o.subscribtionService.GetAllSubscribtions()

	for k, v := range m {

		o.Urls[k] = watcher.New(k, v, 60000, o.logger, o.advtService)
	}

	return o
}

func (o *Observer) Run() {
	for _, w := range o.Urls {
		w.Run()
	}
}
