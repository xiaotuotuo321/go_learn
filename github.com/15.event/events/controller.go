package events

import (
	"bytes"
	"container/list"
	"errors"
	"fmt"
	zap "go.uber.org/zap"
	statsdv2 "gopkg.in/alexcesaro/statsd.v2"
	"reflect"
	"regexp"
	"sync"
	"time"
)

type Client struct {
	serverName string

	db  *statsdv2.Client
	log Logger
}

type TimeMonitorNode struct {
	timing   *statsdv2.Timing
	start    time.Time
	srvName  string
	mod      string
	funcName string
}

func (t *TimeMonitorNode) Send(extra ...string) {
	if t.timing == nil {
		return
	}

	srvName := t.srvName
	if srvName == "" {
		srvName = "game_server"
	}

	bucket := fmt.Sprintf("game_server.%s.%s.%s.timecost", srvName, t.mod, t.funcName)
	bucket = formatBucket(bucket)
	t.timing.Send(bucket)

	// 上报额外的统计
	for _, b := range extra {
		bucket = fmt.Sprintf("%s.%s.%s", srvName, t.mod, b)
		bucket = formatBucket(bucket)
		t.timing.Send(bucket)
	}
}

func NewTimeMonitor(cli *Client, srvName string, ty string, funcName string) *TimeMonitorNode {
	if cli != nil && cli.db != nil {
		t := cli.db.NewTiming()
		return &TimeMonitorNode{
			timing:   &t,
			srvName:  srvName,
			mod:      ty,
			funcName: funcName,
			start:    time.Now(),
		}
	}
	return &TimeMonitorNode{
		srvName:  srvName,
		mod:      ty,
		funcName: funcName,
		start:    time.Now(),
	}
}

func (c *Client) StartTiming(bucket string) *TimeMonitorNode {
	if c == nil || c.db == nil {
		return nil
	}

	ret := NewTimeMonitor(c, c.serverName, "time", bucket)
	return ret
}

func (c *Client) Incr(bucket string) {
	if c == nil || c.db == nil {
		return
	}

	srvName := c.serverName
	if srvName == "" {
		srvName = "game_server"
	}

	// !有些版本的statsd服务会把第一段丢弃掉，没有深究是为什么，如果发现多了一个点，可以把前面的点去掉
	bucket = fmt.Sprintf("%s.count.%s", srvName, bucket)
	bucket = formatBucket(bucket)
	c.db.Increment(bucket)
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

func (c *EventController) Publish(topic Topic, waitAsync bool, arg ...interface{}) {
	c.l.RLock()
	defer c.l.Unlock()

	tmpList, ok := c.subscribers[topic]
	if !ok {
		return
	}

	bucket := fmt.Sprintf("event.publish.%s", topic)
	tm := c.statsd.StartTiming(bucket)
	defer tm.Send()

	c.statsd.Incr(bucket)
	var wg sync.WaitGroup

	for i := tmpList.Front(); i != nil; i = i.Next() {
		tmpV, _ := i.Value.(*Subs)
	}
}

func (c *EventController) handlerIsFunc(handler EventHandler) bool {
	if handler == nil {
		return false
	}

	return reflect.TypeOf(handler).Kind() == reflect.Func
}

func formatBucket(arg string) string {
	tmp := []byte(arg)
	var keyMatchRegex = regexp.MustCompile(`(\w+)`)
	var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)
	converted := keyMatchRegex.ReplaceAllFunc(
		tmp,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)
	return string(converted)
}
