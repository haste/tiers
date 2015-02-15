// Slightly modified version of Firebase's Push IDs:
// https://gist.github.com/mikelehen/3596a30bd69384624c11
package id

import (
	"math/rand"
	"sync"
	"time"
)

// Modeled after base62 web-safe chars, but ordered by ASCII.
const pushChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var (
	// Timestamp of last push, used to prevent local collisions if you push twice in one ms.
	lastPushTimeMs int64

	// We generate 72-bits of randomness which get turned into 12 characters and appended to the
	// timestamp to prevent collisions with other clients.  We store the last characters we
	// generated because in the event of a collision, we'll use those same characters except
	// "incremented" by one.
	lastRandChars [12]int
	mu            sync.Mutex
)

func init() {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 12; i++ {
		lastRandChars[i] = rand.Intn(62)
	}
}

func New() string {
	var id [20]byte
	mu.Lock()

	timeMs := time.Now().UTC().UnixNano() / 1e6
	if timeMs == lastPushTimeMs {
		for i := 0; i < 12; i++ {
			lastRandChars[i]++
			if lastRandChars[i] < 62 {
				break
			}
			lastRandChars[i] = 0
		}
	}

	lastPushTimeMs = timeMs
	for i := 0; i < 12; i++ {
		id[19-i] = pushChars[lastRandChars[i]]
	}
	mu.Unlock()

	for i := 7; i >= 0; i-- {
		n := int(timeMs % 62)
		id[i] = pushChars[n]
		timeMs /= 62
	}

	return string(id[:])
}
