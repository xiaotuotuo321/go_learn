package events

import (
	"container/list"
	"errors"
	zap "go.uber.org/zap"
	statsdv2 "gopkg.in/alexcesaro/statsd.v2"
	"reflect"
	"sync"
)

type Client struct {
	serverName string

	db  *statsdv2.Client
	log Logger
}

type StdLogger struct {
	logger *zap.SugaredLogger
	Logger
}

type Logger interface {
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
	Panic(msg string, fields ...interface{})
}

type TopicSubStatistics struct {
	RunTimes     int64
	TotalRunTime int64
	AvgRunTime   int64
}

type Topic string

type EventController struct {
	subscribers  map[Topic]*list.List
	topicHandler map[Topic]interface{}

	statistics     map[string]*TopicSubStatistics
	statisticsOpen bool
	handlerCheck   bool
	l              sync.RWMutex
	sl             sync.RWMutex

	statsd *Client
	logger *StdLogger
}

func NewEventController(l *StdLogger, handlerCheck bool, statisticsOpen bool, StatsdClient *Client) *EventController {
	return &EventController{
		logger:         l,
		subscribers:    make(map[Topic]*list.List),
		handlerCheck:   handlerCheck,
		statisticsOpen: statisticsOpen,
		statistics:     make(map[string]*TopicSubStatistics),
		statsd:         StatsdClient,
	}
}

func (c *EventController) TopicRegister(topic Topic, handler interface{}) error {
	if !c.handlerIsFunc(handler) {
		return errors.New("handler not func")
	}

	c.l.Lock()
	defer c.l.Unlock()

	c.topicHandler[topic] = handler

	return nil
}

func (c *EventController) GetTopicRegisteredHandler(topic Topic) interface{} {
	c.l.RLock()
	defer c.l.Unlock()

	res, ok := c.topicHandler[topic]
	if !ok {
		return nil
	}
	return res
}

func (c *EventController) handlerIsFunc(handler EventHandler) bool {
	if handler == nil {
		return false
	}

	return reflect.TypeOf(handler).Kind() == reflect.Func
}
