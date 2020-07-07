package gvm

type GVM interface {
	Start()
	Stop()
}

type Context struct {

}

func (c *Context) Start() {
	panic("implement me")
}

func (c *Context) Stop() {
	panic("implement me")
}

func NewGVM() GVM {
	return &Context{}
}