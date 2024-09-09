package avahi

import (
	"fmt"
	"sync"

	dbus "github.com/godbus/dbus/v5"
)

type Event uint

const (
	// DBUS or Avahi got disconnected
	Disconnected = 0
)

// A Server is the cental object of an Avahi connection
type Server struct {
	conn          *dbus.Conn
	object        dbus.BusObject
	signalChannel chan *dbus.Signal
	quitChannel   chan struct{}

	eventCB EventCB

	mutex          sync.Mutex
	signalEmitters map[dbus.ObjectPath]SignalEmitter
}

// ServerNew returns a new Server object
//   - retryUntilConnect: if true, the server will retry to connect to DBus and Avahi if they are not available at the start
//   - closeCB: is invoked when either DBus or Avahi disconnects
func ServerNew() ServerInterface {
	return &Server{
		signalChannel:  make(chan *dbus.Signal, 10),
		quitChannel:    make(chan struct{}),
		signalEmitters: make(map[dbus.ObjectPath]SignalEmitter),
	}
}

var _ ServerInterface = (*Server)(nil)

func (c *Server) Setup(eventCB EventCB) error {
	conn, err := dbus.SystemBus()
	if err != nil {
		return err
	}

	c.conn = conn
	c.eventCB = eventCB

	return nil
}

// Start the server
//
// returns an error if the dbus connection failed
func (c *Server) Start() {
	c.object = c.conn.Object("org.freedesktop.Avahi", dbus.ObjectPath("/"))
	// Get signals for DBus Disconnects
	c.conn.Signal(c.signalChannel)
	// Get signals for DBus Disconnects
	c.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.freedesktop.DBus.Local'")
	// Get signals for Avahi NameOwnerChanged, for avahi (dis)connects
	c.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.freedesktop.DBus'")
	// Get signals for Avahi updates
	c.conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, "type='signal',interface='org.freedesktop.Avahi'")

	go c.handleSignals()
}

func (c *Server) handleSignals() {
	shutdownFunc := func() {
		c.shutdown()
	}

	for {
		select {
		case signal, ok := <-c.signalChannel:
			if !ok {
				continue
			}

			switch signal.Name {
			case "org.freedesktop.DBus.Local.Disconnected":
				// DBus disconneted
				defer shutdownFunc()
				return

			case "org.freedesktop.DBus.NameOwnerChanged":
				if signal.Path != "/org/freedesktop/DBus" {
					break
				}

				var name, old, new *string

				if err := dbus.Store(signal.Body, &name, &old, &new); err != nil {
					break
				}

				if name == nil || *name != "org.freedesktop.Avahi" {
					break
				}

				if old != nil && *old != "" {
					// Avahi Daemon disconneted
					defer shutdownFunc()
					return
				}

			default:
				c.mutex.Lock()
				for path, obj := range c.signalEmitters {
					if path == signal.Path {
						_ = obj.DispatchSignal(signal)
					}
				}
				c.mutex.Unlock()
			}

		case <-c.quitChannel:
			return
		}
	}
}

func (c *Server) shutdown() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for path, obj := range c.signalEmitters {
		obj.Free()
		delete(c.signalEmitters, path)
	}

	if c.conn != nil {
		c.conn.Close()
		c.conn = nil

		if c.eventCB != nil {
			go c.eventCB(Disconnected)
		}
	}
}

// Close the connection to a dbus server
func (c *Server) Shutdown() {
	if c.quitChannel != nil {
		c.quitChannel <- struct{}{}
		close(c.quitChannel)
		c.quitChannel = nil
	}

	c.shutdown()
}

func (c *Server) interfaceForMember(method string) string {
	return fmt.Sprintf("%s.%s", "org.freedesktop.Avahi.Server", method)
}

// EntryGroupNew returns a new and empty EntryGroup
func (c *Server) EntryGroupNew() (EntryGroupInterface, error) {
	var o dbus.ObjectPath

	c.mutex.Lock()
	defer c.mutex.Unlock()

	err := c.object.Call(c.interfaceForMember("EntryGroupNew"), 0).Store(&o)
	if err != nil {
		return nil, err
	}

	r, err := EntryGroupNew(c.conn, o)
	if err != nil {
		return nil, err
	}

	c.signalEmitters[o] = r

	return r, nil
}

// EntryGroupFree frees an entry group and releases its resources on the service
func (c *Server) EntryGroupFree(r EntryGroupInterface) {
	c.signalEmitterFree(r)
}

// ResolveHostName ...
func (c *Server) ResolveHostName(iface, protocol int32, name string, aprotocol int32, flags uint32) (reply HostName, err error) {
	err = c.object.Call(c.interfaceForMember("ResolveHostName"), 0, iface, protocol, name, aprotocol, flags).
		Store(&reply.Interface, &reply.Protocol, &reply.Name, &reply.Aprotocol, &reply.Address, &reply.Flags)
	return reply, err
}

// ResolveAddress ...
func (c *Server) ResolveAddress(iface, protocol int32, address string, flags uint32) (reply Address, err error) {
	err = c.object.Call(c.interfaceForMember("ResolveAddress"), 0, iface, protocol, address, flags).
		Store(&reply.Interface, &reply.Protocol, &reply.Aprotocol, &reply.Address, &reply.Name, &reply.Flags)
	return reply, err
}

// ResolveService ...
func (c *Server) ResolveService(iface, protocol int32, name, serviceType, domain string, aprotocol int32, flags uint32) (reply Service, err error) {
	err = c.object.Call(c.interfaceForMember("ResolveService"), 0, iface, protocol, name, serviceType, domain, aprotocol, flags).
		Store(&reply.Interface, &reply.Protocol, &reply.Name, &reply.Type, &reply.Domain,
			&reply.Host, &reply.Aprotocol, &reply.Address, &reply.Port, &reply.Txt, &reply.Flags)
	return reply, err
}

// DomainBrowserNew ...
func (c *Server) DomainBrowserNew(iface, protocol int32, domain string, btype int32, flags uint32) (DomainBrowserInterface, error) {
	var o dbus.ObjectPath

	err := c.object.Call(c.interfaceForMember("DomainBrowserNew"), 0, iface, protocol, domain, btype, flags).Store(&o)
	if err != nil {
		return nil, err
	}

	r, err := DomainBrowserNew(c.conn, o)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// DomainBrowserFree ...
func (c *Server) DomainBrowserFree(r DomainBrowserInterface) {
	c.signalEmitterFree(r)
}

// ServiceTypeBrowserNew ...
func (c *Server) ServiceTypeBrowserNew(iface, protocol int32, domain string, flags uint32) (ServiceTypeBrowserInterface, error) {
	var o dbus.ObjectPath

	c.mutex.Lock()
	defer c.mutex.Unlock()

	err := c.object.Call(c.interfaceForMember("ServiceTypeBrowserNew"), 0, iface, protocol, domain, flags).Store(&o)
	if err != nil {
		return nil, err
	}

	r, err := ServiceTypeBrowserNew(c.conn, o)
	if err != nil {
		return nil, err
	}

	c.signalEmitters[o] = r

	return r, nil
}

// ServiceTypeBrowserFree ...
func (c *Server) ServiceTypeBrowserFree(r ServiceTypeBrowserInterface) {
	c.signalEmitterFree(r)
}

// ServiceBrowserNew ...
func (c *Server) ServiceBrowserNew(addChan, removeChan chan Service, iface, protocol int32, serviceType string, domain string, flags uint32) (ServiceBrowserInterface, error) {
	var o dbus.ObjectPath

	c.mutex.Lock()
	defer c.mutex.Unlock()

	err := c.object.Call(c.interfaceForMember("ServiceBrowserNew"), 0, iface, protocol, serviceType, domain, flags).Store(&o)
	if err != nil {
		return nil, err
	}

	r, err := ServiceBrowserNew(addChan, removeChan, c.conn, o)
	if err != nil {
		return nil, err
	}

	c.signalEmitters[o] = r

	return r, nil
}

// ServiceBrowserFree ...
func (c *Server) ServiceBrowserFree(r ServiceBrowserInterface) {
	c.signalEmitterFree(r)
}

// ServiceResolverNew ...
func (c *Server) ServiceResolverNew(iface, protocol int32, name, serviceType, domain string, aprotocol int32, flags uint32) (ServiceResolverInterface, error) {
	var o dbus.ObjectPath

	c.mutex.Lock()
	defer c.mutex.Unlock()

	err := c.object.Call(c.interfaceForMember("ServiceResolverNew"), 0, iface, protocol, name, serviceType, domain, aprotocol, flags).Store(&o)
	if err != nil {
		return nil, err
	}

	r, err := ServiceResolverNew(c.conn, o)
	if err != nil {
		return nil, err
	}

	c.signalEmitters[o] = r

	return r, nil
}

// ServiceResolverFree ...
func (c *Server) ServiceResolverFree(r ServiceResolverInterface) {
	c.signalEmitterFree(r)
}

// HostNameResolverNew ...
func (c *Server) HostNameResolverNew(iface, protocol int32, name string, aprotocol int32, flags uint32) (HostNameResolverInterface, error) {
	var o dbus.ObjectPath

	c.mutex.Lock()
	defer c.mutex.Unlock()

	err := c.object.Call(c.interfaceForMember("HostNameResolverNew"), 0, iface, protocol, name, aprotocol, flags).Store(&o)
	if err != nil {
		return nil, err
	}

	r, err := HostNameResolverNew(c.conn, o)
	if err != nil {
		return nil, err
	}

	c.signalEmitters[o] = r

	return r, nil
}

// AddressResolverNew ...
func (c *Server) AddressResolverNew(iface, protocol int32, address string, flags uint32) (AddressResolverInterface, error) {
	var o dbus.ObjectPath

	c.mutex.Lock()
	defer c.mutex.Unlock()

	err := c.object.Call(c.interfaceForMember("AddressResolverNew"), 0, iface, protocol, address, flags).Store(&o)
	if err != nil {
		return nil, err
	}

	r, err := AddressResolverNew(c.conn, o)
	if err != nil {
		return nil, err
	}

	c.signalEmitters[o] = r

	return r, nil
}

// AddressResolverFree ...
func (c *Server) AddressResolverFree(r AddressResolverInterface) {
	c.signalEmitterFree(r)
}

// RecordBrowserNew ...
func (c *Server) RecordBrowserNew(iface, protocol int32, name string, class uint16, recordType uint16, flags uint32) (RecordBrowserInterface, error) {
	var o dbus.ObjectPath

	c.mutex.Lock()
	defer c.mutex.Unlock()

	err := c.object.Call(c.interfaceForMember("RecordBrowserNew"), 0, iface, protocol, name, class, recordType, flags).Store(&o)
	if err != nil {
		return nil, err
	}

	r, err := RecordBrowserNew(c.conn, o)
	if err != nil {
		return nil, err
	}

	c.signalEmitters[o] = r

	return r, nil
}

// RecordBrowserFree ...
func (c *Server) RecordBrowserFree(r RecordBrowserInterface) {
	c.signalEmitterFree(r)
}

// GetAPIVersion ...
func (c *Server) GetAPIVersion() (int32, error) {
	var i int32

	err := c.object.Call(c.interfaceForMember("GetAPIVersion"), 0).Store(&i)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// GetAlternativeHostName ...
func (c *Server) GetAlternativeHostName(name string) (string, error) {
	var s string

	err := c.object.Call(c.interfaceForMember("GetAlternativeHostName"), 0, name).Store(&s)
	if err != nil {
		return "", err
	}

	return s, nil
}

// GetAlternativeServiceName ...
func (c *Server) GetAlternativeServiceName(name string) (string, error) {
	var s string

	err := c.object.Call(c.interfaceForMember("GetAlternativeServiceName"), 0, name).Store(&s)
	if err != nil {
		return "", err
	}

	return s, nil
}

// GetDomainName ...
func (c *Server) GetDomainName() (string, error) {
	var s string

	err := c.object.Call(c.interfaceForMember("GetDomainName"), 0).Store(&s)
	if err != nil {
		return "", err
	}

	return s, nil
}

// GetHostName ...
func (c *Server) GetHostName() (string, error) {
	var s string

	err := c.object.Call(c.interfaceForMember("GetHostName"), 0).Store(&s)
	if err != nil {
		return "", err
	}

	return s, nil
}

// GetHostNameFqdn ...
func (c *Server) GetHostNameFqdn() (string, error) {
	var s string

	err := c.object.Call(c.interfaceForMember("GetHostNameFqdn"), 0).Store(&s)
	if err != nil {
		return "", err
	}

	return s, nil
}

// GetLocalServiceCookie ...
func (c *Server) GetLocalServiceCookie() (int32, error) {
	var i int32

	err := c.object.Call(c.interfaceForMember("GetLocalServiceCookie"), 0).Store(&i)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// GetNetworkInterfaceIndexByName -...
func (c *Server) GetNetworkInterfaceIndexByName(name string) (int32, error) {
	var i int32

	err := c.object.Call(c.interfaceForMember("GetNetworkInterfaceIndexByName"), 0, name).Store(&i)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// GetNetworkInterfaceNameByIndex ...
func (c *Server) GetNetworkInterfaceNameByIndex(index int32) (string, error) {
	var s string

	err := c.object.Call(c.interfaceForMember("GetNetworkInterfaceNameByIndex"), 0, index).Store(&s)
	if err != nil {
		return "", err
	}

	return s, nil
}

// GetState ...
func (c *Server) GetState() (int32, error) {
	var i int32

	err := c.object.Call(c.interfaceForMember("GetState"), 0).Store(&i)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// GetVersionString ...
func (c *Server) GetVersionString() (string, error) {
	var s string

	err := c.object.Call(c.interfaceForMember("GetVersionString"), 0).Store(&s)
	if err != nil {
		return "", err
	}

	return s, nil
}

// IsNSSSupportAvailable ...
func (c *Server) IsNSSSupportAvailable() (bool, error) {
	var b bool

	err := c.object.Call(c.interfaceForMember("IsNSSSupportAvailable"), 0).Store(&b)
	if err != nil {
		return false, err
	}

	return b, nil
}

// SetServerName ...
func (c *Server) SetServerName(name string) error {
	return c.object.Call(c.interfaceForMember("SetServerName"), 0, name).Err
}
