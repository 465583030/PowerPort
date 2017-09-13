package powerportd

import "net"
import "loger"

func ServeAdmin(fd net.Listener) {
	defer fd.Close()
	for {
		conn, err := fd.Accept()
		if err != nil {
			loger.Log(err)
			continue
		}

		conn.Close()
	}
}