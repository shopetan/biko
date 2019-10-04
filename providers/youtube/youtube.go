package youtube

import (
	"net/url"
	"path"

	"github.com/urfave/cli"
)

type Provider struct {
	baseURL *url.URL
	URL     *url.URL
	Ctx     *cli.Context
}

// Init ...
func (p *Provider) Init(c *cli.Context) error {
	p.Ctx = c
	return nil
}

// GetTargetURL ...
func (p *Provider) GetTargetURL() (string, error) {
	const baseURL = "https://youtube.com"
	var err error
	if p.baseURL, err = url.Parse(baseURL); err != nil {
		return "", err
	}

	p.addProductPath(p.Ctx.Command.Name)
	return p.URL.String(), nil
}

func (p *Provider) addProductPath(product string) {
	switch product {
	case "search":
		p.join("results")
		param := url.Values{}
		var query string
		if query = p.Ctx.String("query"); query != "" {
			param.Add("search_query", query)
			p.URL.RawQuery = param.Encode()
		}
	}
}

func (p *Provider) join(additionPath string) {
	if p.URL == nil {
		p.URL = p.baseURL
	}
	p.URL.Path = path.Join(p.URL.Path, additionPath)
}
