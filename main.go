package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"

	discord "github.com/bwmarrin/discordgo"
	"github.com/monzo/typhon"
	"github.com/ameliaikeda/cynthia/handler"
)

var (
	token, guild, appID string
)

func init() {
	flag.StringVar(&token, "secret", "", "OAuth2 client secret.")
	flag.StringVar(&guild, "guild", "", "Discord guild")
	flag.StringVar(&appID, "app", "", "Bot app ID")
	flag.Parse()

	if token == "" {
		fmt.Println("client secret is required.")
		os.Exit(1)
	}
	if appID == "" {
		fmt.Println("client ID is required.")
		os.Exit(1)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🔍 looking up access token")
	t := accessToken(appID, token)
	fmt.Printf("💖 found an access token: %s\n", t.Token)

	app, err := discord.New(fmt.Sprintf("%s %s", t.Type, t.Token))
	if err != nil {
		fmt.Printf("error starting bot: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("🕸  creating a web handler on :3030")
	srv, err := typhon.Listen(handler.Service(os.Getenv("DISCORD_PUBLIC_KEY")), ":3030")
	if err != nil {
		fmt.Printf("🕸 web server failed to start: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("🐈 creating the /roll command.")
	_, err = app.ApplicationCommandCreate(appID, guild, &discord.ApplicationCommand{
		Name:        "roll",
		Description: "Roll a d20 in the current channel",
		Version:     "1.0.0",
	})
	if err != nil {
		fmt.Printf("error registering slash command: %v\n", err)
		os.Exit(1)
	}

	defer app.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	select {
	case <-stop:
	case <-srv.Done():
	}
	fmt.Println("👋 shutting down, bye")
}

type AccessToken struct {
	Token     string `json:"access_token"`
	Type      string `json:"token_type"`
	ExpiresIn int64  `json:"expires_in"`
	Scope     string `json:"scope"`
}

// take a client ID and secret and exchange it for a token.
func accessToken(id, secret string) *AccessToken {
	const apiURL = "https://discord.com/api/v8"
	const payload = `grant_type=client_credentials&scope=applications.commands+applications.commands.update`
	app, _ := discord.New()

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/oauth2/token", apiURL), bytes.NewReader([]byte(payload)))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(id, secret)

	defer req.Body.Close()

	rsp, err := app.Client.Do(req)
	if err != nil {
		fmt.Printf("error sending token request: %v\n", err)
		os.Exit(1)
	}
	defer rsp.Body.Close()

	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Printf("error reading body: %v\n", err)
		os.Exit(1)
	}

	if rsp.StatusCode != 200 {
		fmt.Printf("error fetching access token: %s\n", string(b))
		os.Exit(1)
	}

	token := &AccessToken{}
	if err := json.Unmarshal(b, token); err != nil {
		fmt.Printf("error decoding body: %v\n", err)
		os.Exit(1)
	}

	return token
}
