package avahi

import (
	"fmt"
	"sync"

	dbus "github.com/godbus/dbus/v5"
)

// A ServiceBrowser browses for mDNS services
type ServiceBrowser struct {
	object        dbus.BusObject
	mu            sync.Mutex
	freeOnce      sync.Once
	addChannel    chan Service
	removeChannel chan Service
	wake          chan struct{}
	stop          chan struct{}
	stopped       chan struct{}
	pending       []serviceBrowserEvent
}

type serviceBrowserEvent struct {
	service Service
	channel chan Service
}

// ServiceBrowserNew creates a new browser for mDNS records
func ServiceBrowserNew(addChan, removeChan chan Service, conn *dbus.Conn, path dbus.ObjectPath) (ServiceBrowserInterface, error) {
	c := new(ServiceBrowser)

	c.object = conn.Object("org.freedesktop.Avahi", path)
	c.addChannel = addChan
	c.removeChannel = removeChan
	c.wake = make(chan struct{}, 1)
	c.stop = make(chan struct{})
	c.stopped = make(chan struct{})
	go c.dispatchPending()

	return c, nil
}

var _ ServiceBrowserInterface = (*ServiceBrowser)(nil)

func (c *ServiceBrowser) interfaceForMember(method string) string {
	return fmt.Sprintf("%s.%s", "org.freedesktop.Avahi.ServiceBrowser", method)
}

func (c *ServiceBrowser) Free() {
	c.freeOnce.Do(func() {
		c.mu.Lock()
		close(c.stop)
		c.addChannel = nil
		c.removeChannel = nil
		c.pending = nil
		c.mu.Unlock()
		<-c.stopped
		c.object.Call(c.interfaceForMember("Free"), 0)
	})
}

func (c *ServiceBrowser) GetObjectPath() dbus.ObjectPath {
	return c.object.Path()
}

func (c *ServiceBrowser) DispatchSignal(signal *dbus.Signal) error {
	if signal.Name == c.interfaceForMember("ItemNew") || signal.Name == c.interfaceForMember("ItemRemove") {
		var service Service
		err := dbus.Store(signal.Body, &service.Interface, &service.Protocol, &service.Name, &service.Type, &service.Domain, &service.Flags)
		if err != nil {
			return err
		}

		c.mu.Lock()
		defer c.mu.Unlock()

		if signal.Name == c.interfaceForMember("ItemNew") {
			if c.addChannel != nil {
				c.pending = append(c.pending, serviceBrowserEvent{service: service, channel: c.addChannel})
			}
		} else {
			if c.removeChannel != nil {
				c.pending = append(c.pending, serviceBrowserEvent{service: service, channel: c.removeChannel})
			}
		}
		select {
		case c.wake <- struct{}{}:
		default:
		}
	}

	return nil
}

func (c *ServiceBrowser) dispatchPending() {
	defer close(c.stopped)

	for {
		c.mu.Lock()
		if len(c.pending) == 0 {
			c.mu.Unlock()
			select {
			case <-c.stop:
				return
			case <-c.wake:
			}
			continue
		}
		event := c.pending[0]
		c.pending[0] = serviceBrowserEvent{}
		c.pending = c.pending[1:]
		if len(c.pending) == 0 {
			c.pending = nil
		}
		c.mu.Unlock()

		if event.channel != nil {
			select {
			case <-c.stop:
				return
			case event.channel <- event.service:
			}
		}
	}
}
