package swagger

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WatchInterface interface {
	Stop()

	ResultChan() <-chan WatchEvent
}

type EventType string

var WatchEventTypes = struct {
	Added, Modified, Deleted, Bookmark, Error EventType
}{"ADDED", "MODIFIED", "DELETED", "BOOKMARK", "ERROR"}

type WatchEvent struct {
	Type   EventType
	Object interface{}
}

// WatchDecoder is responsible for unmarshal []byte into the object
type WatchDecoder func([]byte) (interface{}, error)

type StreamWatcher struct {
	sync.Mutex
	reader  io.ReadCloser
	decoder WatchDecoder
	result  chan WatchEvent
	done    chan struct{}
}

func NewStreamWatcher(r io.ReadCloser, d WatchDecoder) *StreamWatcher {
	sw := &StreamWatcher{
		reader:  r,
		decoder: d,
		result:  make(chan WatchEvent),
		done:    make(chan struct{}),
	}
	go sw.receive()
	return sw
}

// ResultChan implements WatchInterface.
func (sw *StreamWatcher) ResultChan() <-chan WatchEvent {
	return sw.result
}

// Stop implements WatchInterface.
func (sw *StreamWatcher) Stop() {
	// Call Close() exactly once by locking and setting a flag.
	sw.Lock()
	defer sw.Unlock()
	// closing a closed channel always panics, therefore check before closing
	select {
	case <-sw.done:
	default:
		close(sw.done)
		sw.reader.Close()
	}
}

// receive reads result from the decoder in a loop and sends down the result channel.
func (sw *StreamWatcher) receive() {
	var err error
	var got metav1.WatchEvent

	defer close(sw.result)
	defer sw.Stop()

	jsonDecoder := json.NewDecoder(sw.reader)
	for {
		if err = jsonDecoder.Decode(&got); err != nil {
			sw.result <- WatchEvent{
				Type:   WatchEventTypes.Error,
				Object: fmt.Errorf("unable to decode an event from the watch stream: %v", err),
			}
			return
		}
		switch got.Type {
		case string(WatchEventTypes.Added), string(WatchEventTypes.Modified), string(WatchEventTypes.Deleted), string(WatchEventTypes.Error), string(WatchEventTypes.Bookmark):
		default:
			sw.result <- WatchEvent{
				Type:   WatchEventTypes.Error,
				Object: fmt.Errorf("invalid event type %s from the watch stream", got.Type),
			}
			return
		}

		var obj interface{}
		if obj, err = sw.decoder(got.Object.Raw); err != nil {
			sw.result <- WatchEvent{
				Type:   WatchEventTypes.Error,
				Object: fmt.Errorf("unable to decode an object from the watch stream: %v", err),
			}
			return
		}
		select {
		case <-sw.done:
			return
		case sw.result <- WatchEvent{
			Type:   EventType(got.Type),
			Object: obj,
		}:
		}
	}
}
