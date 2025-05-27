package uuid

import (
	"github.com/3JoB/unsafeConvert"
	"github.com/google/uuid"
	"golang.org/x/crypto/sha3"
)

func GenerateUUIDv9(name string) uuid.UUID {
	return uuid.NewHash(sha3.NewShake128(), uuid.NameSpaceURL, unsafeConvert.BytePointer(name), 9)
}

func GenerateUUIDv3(name string) uuid.UUID {
	return uuid.NewMD5(uuid.NameSpaceDNS, unsafeConvert.BytePointer(name))
}

func GenerateUUIDv5(name string) uuid.UUID {
	return uuid.NewSHA1(uuid.NameSpaceDNS, unsafeConvert.BytePointer(name))
}
