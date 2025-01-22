package route

import (
	socketio "github.com/googollee/go-socket.io"
)

// SocketRoute đại diện cho một sự kiện trong Socket.IO
type SocketRoute struct {
	Event       string                                                                                   // Tên sự kiện (event)
	Handler     func(s socketio.Conn, data any)                                                          // Hàm xử lý sự kiện
	Middlewares []func(next func(s socketio.Conn, data any) error) func(s socketio.Conn, data any) error // Middleware
}

// GroupSocketRoute đại diện cho một nhóm sự kiện
type GroupSocketRoute struct {
	Namespace   string                                                                                   // Namespace trong Socket.IO (ví dụ: "/chat")
	Middlewares []func(next func(s socketio.Conn, data any) error) func(s socketio.Conn, data any) error // Middleware
	Routes      []SocketRoute                                                                            // Danh sách các sự kiện
}

func (r SocketRoute) AddSocketRouteToServer(server *socketio.Server, namespace string) {
	server.OnEvent(namespace, r.Event, r.Handler)
}

func (r GroupSocketRoute) AddGroupSocketRouteToServer(server *socketio.Server) {
	server.OnConnect(r.Namespace, func(s socketio.Conn) error {
		for _, mw := range r.Middlewares {
			s.SetContext(mw)
		}
		return nil
	})

	for _, route := range r.Routes {
		route.AddSocketRouteToServer(server, r.Namespace)
	}
}

func RegisterSocketRoutes(server *socketio.Server, routes []GroupSocketRoute) {
	for _, group := range routes {
		group.AddGroupSocketRouteToServer(server)
	}
}
