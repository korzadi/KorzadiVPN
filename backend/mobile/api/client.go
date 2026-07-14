package api

type Client struct {
	BaseURL string
}

func NewClient(
	url string,
) Client {

	return Client{

		BaseURL: url,
	}
}
