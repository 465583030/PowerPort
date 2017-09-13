package powerportd

import "net"

func ServeData(fd net.Listener) {
	defer fd.Close()
}