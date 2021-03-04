package handler

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	discord "github.com/bwmarrin/discordgo"
	"github.com/monzo/typhon"
)

func Service(publicKey string) typhon.Service {
	return func(req typhon.Request) typhon.Response {
		b, _ := req.BodyBytes(false)

		// signature validation.
		if !verify(req, b, publicKey) {
			rsp := typhon.NewResponse(req)
			rsp.StatusCode = 401

			return rsp
		}

		i := &discord.Interaction{}
		if err := req.Decode(i); err != nil {
			// can't respond here.
			fmt.Printf("error decoding request: %v\n", err)
			return typhon.Response{Error: err}
		}

		switch i.Type {
		case discord.InteractionPing:
			return req.Response(&discord.InteractionResponse{
				Type: discord.InteractionResponsePong,
			})
		case discord.InteractionApplicationCommand:
			return executeCommand(req, i)
		}

		return req.Response("")
	}
}

func verify(req typhon.Request, body []byte, publicKey string) bool {
	// parts:
	s := req.Header.Get("X-Signature-Ed25519")
	ts := req.Header.Get("X-Signature-Timestamp")

	if s == "" || ts == "" {
		return false
	}

	var pub = make(ed25519.PublicKey, ed25519.PublicKeySize)
	var sig = make([]byte, ed25519.SignatureSize)

	hex.Decode(pub, []byte(publicKey))
	hex.Decode(sig, []byte(s))

	return ed25519.Verify(pub, []byte(ts+string(body)), sig)
}

func executeCommand(req typhon.Request, i *discord.Interaction) typhon.Response {
	switch i.Data.Name {
	case "roll":
		return req.Response(roll())
	}

	return req.Response("")
}

func roll() *discord.InteractionResponse {
	fmt.Println("🎲 rolling a die")

	n := random()
	if n == 0 {
		return oops("Oops, something went wrong.")
	}

	return respond(fmt.Sprintf(":game_die: **%d**!", n))
}

func random() int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(20))
	if err != nil {
		return 0
	}

	return n.Int64() + 1
}

func respond(msg string) *discord.InteractionResponse {
	return &discord.InteractionResponse{
		Type: discord.InteractionResponseChannelMessage,
		Data: &discord.InteractionApplicationCommandResponseData{
			Content: msg,
		},
	}
}

func oops(msg string) *discord.InteractionResponse {
	return &discord.InteractionResponse{
		Type: discord.InteractionResponseChannelMessage,
		Data: &discord.InteractionApplicationCommandResponseData{
			Flags:   1 << 6,
			Content: msg,
		},
	}
}
