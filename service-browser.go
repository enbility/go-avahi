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
	addChannel    chan Service
	removeChannel chan Service
}

// ServiceBrowserNew creates a new browser for mDNS records
func ServiceBrowserNew(addChan, removeChan chan Service, conn *dbus.Conn, path dbus.ObjectPath) (ServiceBrowserInterface, error) {
	c := new(ServiceBrowser)

	c.object = conn.Object("org.freedesktop.Avahi", path)
	c.addChannel = addChan
	c.removeChannel = removeChan

	return c, nil
}

var _ ServiceBrowserInterface = (*ServiceBrowser)(nil)

func (c *ServiceBrowser) interfaceForMember(method string) string {
	return fmt.Sprintf("%s.%s", "org.freedesktop.Avahi.ServiceBrowser", method)
}

func (c *ServiceBrowser) Free() {
	// Nil channels before calling D-Bus Free so any concurrent DispatchSignal
	// (dispatched after the server snapshot) sees nil and skips the send.
	c.mu.Lock()
	c.addChannel = nil
	c.removeChannel = nil
	c.mu.Unlock()
	c.object.Call(c.interfaceForMember("Free"), 0)
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

		// Read channel pointers under mu so we see any nil written by Free().
		c.mu.Lock()
		addCh := c.addChannel
		removeCh := c.removeChannel
		c.mu.Unlock()

		if signal.Name == c.interfaceForMember("ItemNew") {
			if addCh != nil {
				// Non-blocking send: the channel is buffered (see AvahiProvider.Start).
				// Dropping a signal here is far safer than blocking handleSignals while
				// it no longer holds c.mutex, which would create a goroutine leak.
				select {
				case addCh <- service:
				default:
				}
			}
		} else {
			if removeCh != nil {
				select {
				case removeCh <- service:
				default:
				}
			}
		}
	}

	return nil
}
