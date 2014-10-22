package page

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"tiers/model"
	"tiers/session"

	"code.google.com/p/goauth2/oauth"
)

const (
	clientID        = "447339154437-o09jreqv5rujsrtujg4u3rl80c5qcii9.apps.googleusercontent.com"
	clientSecret    = "9JvJpKfPlNJnx77fO2gh16pt"
	applicationName = "tie.rs"
)

// config is the configuration specification supplied to the OAuth package.
var config = &oauth.Config{
	ClientId:     clientID,
	ClientSecret: clientSecret,
	// Scope determines which API calls you are authorized to make
	Scope:    "profile",
	AuthURL:  "https://accounts.google.com/o/oauth2/auth",
	TokenURL: "https://accounts.google.com/o/oauth2/token",
	// Use "postmessage" for the code-flow for server side apps
	RedirectURL: "postmessage",
}

// ClaimSet represents an IdToken response.
type ClaimSet struct {
	Sub string
}

// Token represents an OAuth token response.
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	IdToken     string `json:"id_token"`
}

func base64Decode(s string) ([]byte, error) {
	// add back missing padding
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return base64.URLEncoding.DecodeString(s)
}

// decodeIdToken takes an ID Token and decodes it to fetch the Google+ ID within
func decodeIdToken(idToken string) (gplusID string, err error) {
	// An ID token is a cryptographically-signed JSON object encoded in base 64.
	// Normally, it is critical that you validate an ID token before you use it,
	// but since you are communicating directly with Google over an
	// intermediary-free HTTPS channel and using your Client Secret to
	// authenticate yourself to Google, you can be confident that the token you
	// receive really comes from Google and is valid. If your server passes the ID
	// token to other components of your app, it is extremely important that the
	// other components validate the token before using it.
	var set ClaimSet
	if idToken != "" {
		// Check that the padding is correct for a base64decode
		parts := strings.Split(idToken, ".")
		if len(parts) < 2 {
			return "", fmt.Errorf("Malformed ID token")
		}
		// Decode the ID token
		b, err := base64Decode(parts[1])
		if err != nil {
			return "", fmt.Errorf("Malformed ID token: %v", err)
		}
		err = json.Unmarshal(b, &set)
		if err != nil {
			return "", fmt.Errorf("Malformed ID token: %v", err)
		}
	}
	return set.Sub, nil
}

// exchange takes an authentication code and exchanges it with the OAuth
// endpoint for a Google API bearer token and a Google+ ID
func exchange(code string) (accessToken string, idToken string, err error) {
	// Exchange the authorization code for a credentials object via a POST request
	addr := "https://accounts.google.com/o/oauth2/token"
	values := url.Values{
		"Content-Type":  {"application/x-www-form-urlencoded"},
		"code":          {code},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"redirect_uri":  {config.RedirectURL},
		"grant_type":    {"authorization_code"},
	}
	resp, err := http.PostForm(addr, values)
	if err != nil {
		return "", "", fmt.Errorf("Exchanging code: %v", err)
	}
	defer resp.Body.Close()

	// Decode the response body into a token object
	var token Token
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return "", "", fmt.Errorf("Decoding access token: %v", err)
	}

	return token.AccessToken, token.IdToken, nil
}

func GPlusHandler(w http.ResponseWriter, r *http.Request) {
	code, _ := ioutil.ReadAll(r.Body)
	accessToken, idToken, err := exchange(string(code))

	if err != nil {
		fmt.Println(err)
	}

	gplusid, _ := decodeIdToken(idToken)
	u, err := model.GetUserByGPlusId(gplusid)
	if err != nil {
		u = model.CreateGPlusUser(gplusid, accessToken)
	} else {
		model.UpdateGPlusToken(u.Id, accessToken)
	}

	session.Set(w, r, u.Id)
	fmt.Fprintln(w, `{ "success": true }`)
}
