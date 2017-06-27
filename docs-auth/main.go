package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/coreos/go-oidc/jose"
	"github.com/coreos/go-oidc/oidc"
	"github.com/urfave/cli"
)

type OIDCAuthenticator struct {
	usernameClaim string
	client        *oidc.Client
}

var oidcAuthenticator *OIDCAuthenticator

const (
	appName    = "docs-auth"
	appVersion = "dev"

	bindAddressFlag      = "bind-address"
	authClientDomainFlag = "auth-client-domain"
	authClientIDFlag     = "auth-client-id"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = appVersion
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   bindAddressFlag,
			Usage:  "bind address",
			EnvVar: "BIND_ADDRESS",
			Value:  "0.0.0.0:9000",
		},
		cli.StringFlag{
			Name:   authClientDomainFlag,
			Usage:  "OIDC issuer address",
			EnvVar: "AUTH_CLIENT_DOMAIN",
		},
		cli.StringFlag{
			Name:   authClientIDFlag,
			Usage:  "OIDC trusted client ID",
			EnvVar: "AUTH_CLIENT_ID",
		},
	}
	app.Action = start
	app.Run(os.Args)
}

func start(ctx *cli.Context) error {
	log.Info("Starting...")
	oidcAuthenticator = &OIDCAuthenticator{}
	oidcAuthenticator.configure(ctx)

	http.HandleFunc("/", tokenVerificationHandler)

	address := ctx.String(bindAddressFlag)
	log.Infof("Listening on %s", address)
	return http.ListenAndServe(address, nil)
}

func tryTokenCookie(w http.ResponseWriter, r *http.Request) error {
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		return err
	}
	token := tokenCookie.Value
	_, err = oidcAuthenticator.authenticate(token)
	if err != nil {
		return err
	}
	return nil
}

func getRedirect(r *http.Request) string {
	u := r.URL
	q := r.URL.Query()
	q.Del("id_token")
	u.RawQuery = q.Encode()
	return u.String()
}

func tokenVerificationHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Auth-Token")
	if token == "" {
		err := tryTokenCookie(w, r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		redirectURI := getRedirect(r)
		http.Redirect(w, r, redirectURI, http.StatusMovedPermanently)
		return
	}
	claims, err := oidcAuthenticator.authenticate(token)
	if err != nil {
		cookie := http.Cookie{Name: "token", Value: "", Expires: time.Now(), HttpOnly: true, Path: "/"}
		http.SetCookie(w, &cookie)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	expireCookie, ok, _ := claims.TimeClaim("exp")
	if !ok {
		expireCookie = time.Now().Add(time.Hour * 1)
	}
	cookie := http.Cookie{Name: "token", Value: token, Expires: expireCookie, HttpOnly: true, Path: "/"}
	http.SetCookie(w, &cookie)

	redirectURI := getRedirect(r)
	http.Redirect(w, r, redirectURI, http.StatusMovedPermanently)
	return
}

func (a *OIDCAuthenticator) configure(ctx *cli.Context) {
	httpClient := &http.Client{}
	issuerURL := ctx.String(authClientDomainFlag)
	clientID := ctx.String(authClientIDFlag)

	providerConfig, err := oidc.FetchProviderConfig(httpClient, issuerURL)
	if err != nil {
		log.Fatalf("Unable to fetch provider: Will not be able to verify RSA signed tokens. %v", err)
	}
	clientConfig := oidc.ClientConfig{
		HTTPClient:     httpClient,
		Credentials:    oidc.ClientCredentials{ID: clientID},
		ProviderConfig: providerConfig,
	}

	client, err := oidc.NewClient(clientConfig)
	if err != nil {
		log.Fatalf("Failed to create OIDC client: %v", err)
	}
	a.client = client
	a.usernameClaim = "email"
}

func (a *OIDCAuthenticator) authenticate(token string) (jose.Claims, error) {
	jwt, err := jose.ParseJWT(token)
	if err != nil {
		return nil, err
	}

	if err := a.client.VerifyJWT(jwt); err != nil {
		return nil, err
	}

	claims, err := jwt.Claims()
	if err != nil {
		return nil, err
	}

	_, ok, err := claims.StringClaim(a.usernameClaim)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("cannot find %v in JWT claims", a.usernameClaim)
	}

	return claims, nil
}
