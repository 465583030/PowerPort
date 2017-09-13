package powerportd

import "loger"
import "net"
import "powerportd/config"
import "crypto/tls"

func Main() {

	//1. 读取配置文件
	conf := config.GetConfig()

	//2.监听ctlport端口，接收注册和连接请求，TLS加密
	ctlfd, err := tls.Listen("tcp", conf.CtlListen, conf.TlsConfig)
	loger.CheckError(err)

	//3.监听dataport端口，接收数据传输请求，TLS加密
	datafd, err := tls.Listen("tcp", conf.DataListen, conf.TlsConfig)
	loger.CheckError(err)

	//4.监听adminport端口，接收控制数据，不加密
	adminfd, err := net.Listen("tcp", conf.AdminListen)

	//5.启动伺服
	go ServeCtl(ctlfd)
	go ServeData(datafd)

	//5.自身作为控制线程
	ServeAdmin(adminfd)
}