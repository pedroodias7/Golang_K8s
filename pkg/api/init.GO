package api


type Options struct {
	Password string
	LoginURL string
}

type APIIface interface {
	return 
}


type API struct {
	Options Options
	Client http.client
}

func New(options Options) APIIface {
	return API{
		Options: options,
		Client: http.client{
			Transport: myJWTTransport{
				transport: http.DefaultTransport,
				password: options.Password,
				loginURL: options.loginURL,
			}
		}
	}
}