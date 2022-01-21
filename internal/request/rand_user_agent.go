package request

import (
	"math/rand"
	"time"
)

func randUserAgent() string {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	n := r.Intn(len(userAgents))

	return userAgents[n]
}
