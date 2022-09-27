package utils

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"net/url"
	"os"
	"strings"

	"github.com/itchyny/base58-go"
)

type LinkHasher interface {
	GenerateShortLink(url *url.URL) string
	IsValidLink(link string) (*url.URL, error)
}

type linkHasher struct {
	domain string
}

func NewLinkHasher(domain string) *linkHasher {
	return &linkHasher{
		domain: domain,
	}
}

func (l *linkHasher) IsValidLink(link string) (*url.URL, error) {
	return url.ParseRequestURI(link)
}

func (l *linkHasher) GenerateShortLink(url *url.URL) string {
	urlHashBytes := sha256Of(url.RequestURI())
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))

	var b strings.Builder
	b.WriteString(l.domain)
	b.WriteString("/")
	b.WriteString(finalString[:10])

	return b.String()
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}
