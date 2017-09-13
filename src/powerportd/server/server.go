package server

type Server struct {
	Id string
	Pwd string
	CtlPwd string
	Ports []int
	isRegisted bool
}

func NewServer() *Server {
	s := &Server{}
	s.Ports = make([]int,0)
	s.isRegisted = false
	return s
}

func (s *Server)Regist {

}