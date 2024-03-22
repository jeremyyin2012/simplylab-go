package providers

type Providers struct {
}

func NewProviders() *Providers {
	return &Providers{}
}

func (p Providers) Chat() ChatProvider {
	return ChatProvider{}
}

func (p Providers) User() UserProvider {
	return UserProvider{}
}

func (p Providers) OpenRouter() OpenRouterProvider {
	return OpenRouterProvider{}
}

func (p Providers) Ping() PingProvider {
	return PingProvider{}
}
