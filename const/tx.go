package _const

import (
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/brc29"
)

const TxIDToInternalize = "299ac36833a5ffe6ae30e4d9fcebd6328a5fd4e6cae5dc4d18bda95adc1bbad1"

var TxIDsToInternalize = []string{
	"e2ced8fa7cacb3deac90ba19aca8dd5ef1280f9f8f76c0ff314700121a3fe4e6",
	"5cdd1c148994bd4ea9018b043e98d47eb868abfdb36f0ded96002bb1f2e1c7d8",
}

const (
	DefaultBase64Prefix = "SfKxPIJNgdI="
	DefaultBase64Suffix = "NaGLC6fMH50="
)

var KeyID = brc29.KeyID{
	DerivationPrefix: DefaultBase64Prefix,
	DerivationSuffix: DefaultBase64Suffix,
}
