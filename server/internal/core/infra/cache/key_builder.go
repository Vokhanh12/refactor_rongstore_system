package cache

import "strings"

type KeyBuilder struct {
	prefix string
}

func NewKeyBuilder(prefix string) *KeyBuilder {
	return &KeyBuilder{
		prefix: prefix,
	}
}

func (k *KeyBuilder) Build(parts ...string) string {

	if len(parts) == 0 {
		return k.prefix
	}

	return k.prefix + ":" + strings.Join(parts, ":")
}
