package rnd

import (
	"encoding/hex"

	zmc "github.com/0chain/gosdk/zmagmacore/magmasc"
	"golang.org/x/crypto/sha3"
)

func RandomConsumers(num int) []*zmc.Consumer {
	consumers := make([]*zmc.Consumer, num)
	for idx := range consumers {
		consumers[idx] = randomConsumer()
	}
	return consumers
}

func randomConsumer() *zmc.Consumer {
	id := RandString(32)
	hash := sha3.Sum256([]byte(id))
	return &zmc.Consumer{
		ID:    hex.EncodeToString(hash[:]),
		ExtID: "id:consumer:external:" + id,
		Host:  "host.consumer.local:" + id,
	}
}

func RandomProviders(num int) []*zmc.Provider {
	provider := make([]*zmc.Provider, num)
	for idx := range provider {
		provider[idx] = randomProvider()
	}
	return provider
}

func randomProvider() *zmc.Provider {
	id := RandString(32)
	hash := sha3.Sum256([]byte(id))
	return &zmc.Provider{
		ID:    hex.EncodeToString(hash[:]),
		ExtID: "id:provider:external:" + id,
		Host:  "host.provider.local:" + id,
	}
}
