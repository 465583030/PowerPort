package config

import "flag"
import "loger"
import "io/ioutil"
import "powerportd/server"
import "encoding/json"
import "crypto/tls"
import "github.com/bitly/go-simplejson"

type Config struct {
	CtlListen string
	DataListen string
	AdminListen string
	Servers map[string] *server.Server
	TlsConfig *tls.Config
}

var conf *Config

func init() {

	//读配置文件路径
	confPath := flag.String("c",
							"../config/PowerPortd.json",
							"PowerPortd config file path.")
	flag.Parse()

	//解析配置文件为json
	confData, err := ioutil.ReadFile(*confPath)
	loger.CheckError(err)

	jsonData, err := simplejson.NewJson(confData)
	loger.CheckError(err)

	conf = NewConfig()

	//读取控制端口号和数据传输端口号
	conf.CtlListen = jsonData.Get("ctllisten").MustString()
	conf.DataListen = jsonData.Get("datalisten").MustString()
	conf.AdminListen = jsonData.Get("adminlisten").MustString()

	//解析加密传输需要的证书
	crtPath := jsonData.Get("crtpath").MustString()
	keyPath := jsonData.Get("keypath").MustString()

	cert, err := tls.LoadX509KeyPair(crtPath, keyPath)
	loger.CheckError(err)
	conf.TlsConfig = &tls.Config{
		Certificates:[]tls.Certificate{cert},
		InsecureSkipVerify:true}

	//解析预配置的s端信息
	servers := jsonData.Get("servers").MustArray()

	for _, v := range servers {
		s := v.(map[string] interface{})

		id := s["id"].(string)
		pwd := s["pwd"].(string)
		ctlpwd := s["ctlpwd"].(string)
		ports := s["ports"].([]interface{})

		serv := server.NewServer()
		serv.Id = id
		serv.Pwd = pwd
		serv.CtlPwd = ctlpwd

		for _, p := range ports {
			port, err := p.(json.Number).Int64()
			loger.CheckError(err)
			serv.Ports = append(serv.Ports, int(port))
		}

		conf.Servers[id] = serv
	}
}

func NewConfig() *Config {
	return &Config{Servers: make(map[string] *server.Server)}
}

func GetConfig() *Config {
	return conf
}