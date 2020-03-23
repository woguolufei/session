package session

var providers = make(map[string]Provider)

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session 服务提供者不能为nil")
	}
	if _, dup := providers[name]; dup {
		panic("session 服务提供者不能注册两次")
	}
	providers[name] = provider
}
