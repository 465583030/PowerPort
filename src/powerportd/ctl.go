package powerportd

import "net"
import "loger"

func ServeCtl(fd net.Listener) {
	defer fd.Close()

	for {
		conn, err := fd.Accept()
		if err != nil {
			loger.Log(err)
			continue
		}

		go handleCtlConn(conn)
	}
}

func handleCtlConn(c net.Conn){
	//类型：
	//1.s端regist请求
	//2.c端info请求
	//3.c端conn请求
	
	loop := true
	//若类型为1，来自s端，则无须循环读取，该c只发不收
	//若类型为2、3，来自c端，循环读取数据，看是否有新连接需求，并在超时后关闭连接
	for loop {
		data := make([]byte, 256)

		l, err := c.Read(data)
		if err != nil {
			loger.Log(err)
			continue
		}

		
	}

}