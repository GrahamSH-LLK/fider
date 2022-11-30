package oauth

import "net/url"

var providerParams = map[string]map[string]string{
	"id.twitch.tv": {
		"claims": `{"userinfo":{"preferred_username":null,"email":null,"email_verified":null}}`,
	},
	"oauth2.scratch-wiki.info": {
		"title": `Special:ScratchOAuth2/authorize`,
	},

}

func getProviderInitialParams(u *url.URL) url.Values {
	v := url.Values{}
	if params, ok := providerParams[u.Hostname()]; ok {
		for key, value := range params {
			v.Add(key, value)
		}
	}
	return v
}
