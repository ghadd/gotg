package gotg

type Client struct {
	token string
}

func (c *Client) SetToken(token string) {
	c.token = token
}
