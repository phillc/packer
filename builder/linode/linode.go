package linode

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/packer/builder/linode/version"
	"github.com/linode/linodego"
	"golang.org/x/oauth2"
)

func newLinodeClient(pat string) linodego.Client {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: pat})

	oauthTransport := &oauth2.Transport{
		Source: tokenSource,
	}
	oauth2Client := &http.Client{
		Transport: oauthTransport,
	}

	client := linodego.NewClient(oauth2Client)

	projectURL := "https://www.packer.io"
	userAgent := fmt.Sprintf("Packer/%s (+%s)",
		version.LinodePluginVersion.FormattedVersion(), projectURL)

	client.SetUserAgent(userAgent)
	return client
}
