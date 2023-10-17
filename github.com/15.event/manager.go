package event

import "sync"

var once sync.Once
var eventManager *EventManager

type Event