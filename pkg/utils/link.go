package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"net/url"
	"strings"
)

type LinkHasher interface {
	HashURI(url *url.URL) string
	IsValidLink(link string) (*url.URL, error)
}

type linkHasher struct {
	salt string
}

func NewLinkHasher(salt string) *linkHasher {
	return &linkHasher{salt: salt}
}

func (l *linkHasher) IsValidLink(link string) (*url.URL, error) {
	return url.ParseRequestURI(link)
}

func (l *linkHasher) HashURI(url *url.URL) string {
	h := sha1.New()
	h.Write([]byte(url.RequestURI()))
	shortLinkHash := hex.EncodeToString(h.Sum([]byte(l.salt)))

	var b strings.Builder
	b.WriteString(url.Host)
	b.WriteString("/")
	b.WriteString(shortLinkHash[:10])

	return b.String()
}
