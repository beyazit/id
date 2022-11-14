package id

import (
	"crypto/rand"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sony/sonyflake"
)

type PrefixRecord struct {
	Prefix      string
	Description string
	Secure      bool
}

type ID struct {
	prefixes []*PrefixRecord
	sf       *sonyflake.Sonyflake
}

func New(prefixes []*PrefixRecord, st sonyflake.Settings) *ID {
	sf := sonyflake.NewSonyflake(st)

	return &ID{
		prefixes,
		sf,
	}
}

func (i *ID) Generate(prefix string) (string, error) {
	var id string

	if match, _ := regexp.MatchString("^[a-zA-Z0-9]{1,32}$", prefix); !match {
		return "", errors.New("invalid prefix")
	}

	founded, prefixRecord := find(i.prefixes, prefix)
	if !founded {
		return "", errors.New("prefix not found")
	}

	sf, err := i.sf.NextID()
	if err != nil {
		return "", err
	}

	if prefixRecord.Secure {
		token := make([]byte, 16)
		rand.Read(token)
		id = fmt.Sprintf("%s_%s", prefix, strings.Replace(b64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("s_%s", fmt.Sprintf("%x_%s", token, strconv.Itoa(int(sf)))))), "=", "", -1))
	} else {
		id = fmt.Sprintf("%s_%s", prefix, strings.Replace(b64.URLEncoding.EncodeToString([]byte(strconv.Itoa(int(sf)))), "=", "", -1))
	}

	return id, nil
}

func find(prefixes []*PrefixRecord, prefix string) (bool, *PrefixRecord) {
	for _, v := range prefixes {
		if prefix == v.Prefix {
			return true, v
		}
	}
	return false, &PrefixRecord{}
}
