package command

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
	"github.com/mitchellh/cli"
	v1auth "github.com/nerdalize/nerd/nerd/client/auth/v1"
	"github.com/nerdalize/nerd/nerd/conf"
	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"
)

const (
	authorizeEndpoint = "o/authorize"
)

//LoginOpts describes command options
type LoginOpts struct {
	NerdOpts
}

//Login command
type Login struct {
	*command

	opts   *LoginOpts
	parser *flags.Parser
}

//LoginFactory returns a factory method for the join command
func LoginFactory() (cli.Command, error) {
	base, err := baseCommand()
	if err != nil {
		return nil, err
	}
	base.help = ""
	base.synopsis = "start a new authorized session"
	base.parser = flags.NewNamedParser("nerd login", flags.Default)
	base.ui = &cli.BasicUi{
		Reader: os.Stdin,
		Writer: os.Stderr,
	}

	cmd := &Login{
		command: base,
		opts:    &LoginOpts{},
	}

	cmd.runFunc = cmd.DoRun
	_, err = cmd.command.parser.AddGroup("options", "options", cmd.opts)
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

//DoRun is called by run and allows an error to be returned
func (cmd *Login) DoRun(args []string) error {
	c, err := cmd.conf.Read()
	if err != nil {
		HandleError(errors.Wrap(err, "failed to read config"))
	}
	authbase, err := url.Parse(c.Auth.APIEndpoint)
	if err != nil {
		HandleError(errors.Wrapf(err, "auth endpoint '%v' is not a valid URL", c.Auth.APIEndpoint))
	}
	authOpsClient := v1auth.NewOpsClient(v1auth.OpsClientConfig{
		Base:   authbase,
		Logger: logrus.StandardLogger(),
	})
	randomState := randomString(32)
	doneCh := make(chan response)
	svr := &http.Server{
		Addr: c.Auth.OAuthLocalServer,
	}
	defer svr.Close()
	go spawnServer(svr, c.Auth, randomState, doneCh)

	err = open.Run(fmt.Sprintf("http://%s/oauth?state=%s", c.Auth.OAuthLocalServer, randomState))
	if err != nil {
		HandleError(errors.Wrap(err, "failed to open browser window"))
	}

	oauthResponse := <-doneCh
	if oauthResponse.err != nil {
		HandleError(errors.Wrap(oauthResponse.err, "failed to do oauth login"))
	}

	out, err := authOpsClient.GetOAuthCredentials(oauthResponse.code, c.Auth.ClientID, fmt.Sprintf("http://%s/oauth/callback", c.Auth.OAuthLocalServer))
	if err != nil {
		HandleError(errors.Wrap(err, "failed to get oauth credentials"))
	}

	expiration := time.Unix(time.Now().Unix()+int64(out.ExpiresIn), 0)
	err = conf.WriteOAuth(out.AccessToken, out.RefreshToken, expiration, out.Scope, out.TokenType)
	if err != nil {
		HandleError(errors.Wrap(err, "failed to write oauth tokens to config"))
	}
	return nil
}

type response struct {
	err  error
	code string
}

//spawnServer spawns a local server http server to guide the login flow
func spawnServer(svr *http.Server, cfg conf.AuthConfig, randomState string, doneCh chan response) {
	mux := http.NewServeMux()
	mux.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
		if randomState != r.URL.Query().Get("state") {
			doneCh <- response{
				err: errors.Errorf("oauth state does not match provided state (expected: %v, actual: %v)", randomState, r.URL.Query().Get("state")),
			}
			return
		}

		w.Header().Set("Location", cfg.OAuthSuccessURL)
		w.WriteHeader(http.StatusFound)

		doneCh <- response{
			code: r.URL.Query().Get("code"),
		}
	})

	mux.HandleFunc("/oauth", func(w http.ResponseWriter, r *http.Request) {
		base, err := url.Parse(cfg.APIEndpoint)
		if err != nil {
			doneCh <- response{
				err: errors.Wrapf(err, "failed to parse auth api endpoint (%v)", cfg.APIEndpoint),
			}
			return
		}
		path, err := url.Parse(authorizeEndpoint)
		if err != nil {
			doneCh <- response{
				err: errors.Wrapf(err, "failed to parse authorize endpoint (%v)", authorizeEndpoint),
			}
			return
		}

		resolved := base.ResolveReference(path)
		q := resolved.Query()
		q.Set("state", r.URL.Query().Get("state"))
		q.Set("client_id", cfg.ClientID)
		q.Set("response_type", "code")
		q.Set("redirect_uri", fmt.Sprintf("http://%s/oauth/callback", cfg.OAuthLocalServer))
		resolved.RawQuery = q.Encode()

		w.Header().Set("Location", resolved.String())
		w.WriteHeader(http.StatusFound)
	})
	svr.Handler = mux
	err := svr.ListenAndServe()
	if err != nil {
		doneCh <- response{
			err: errors.Wrap(err, "failed to spawn local server"),
		}
	}
}

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

//randomString returns a random string of `n` characters
func randomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
