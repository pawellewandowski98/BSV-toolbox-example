package _const

import (
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/brc29"
)

const TxIDToInternalize = "299ac36833a5ffe6ae30e4d9fcebd6328a5fd4e6cae5dc4d18bda95adc1bbad1"

var TxIDsToInternalize = []string{
	"e2ced8fa7cacb3deac90ba19aca8dd5ef1280f9f8f76c0ff314700121a3fe4e6",
	//"5cdd1c148994bd4ea9018b043e98d47eb868abfdb36f0ded96002bb1f2e1c7d8",
	//"220e5262647615549f6310260d6b37da23c1ed972e94c12edf8191c14f1fb9e4",
	//"e624f731861a004f58439af3caa050d99604fddea7fd1bc975dad5bbc464d9bc",
	//"be4b937debb7237a94e6e47f2d797e8298a2d512064d7acf700d4fa36c589b3a",
	//"5790fef96c99e3e7fcc52c484be73ceeb8e2eb2e615ff0c1d9cd85ebd101b5be",
	//"a2fa4e67c9bd730a78c0c65f17dc61154d1227a8b20398b63cabc9ad327499db",
}

const (
	DefaultBase64Prefix = "SfKxPIJNgdI="
	DefaultBase64Suffix = "NaGLC6fMH50="
)

var KeyID = brc29.KeyID{
	DerivationPrefix: DefaultBase64Prefix,
	DerivationSuffix: DefaultBase64Suffix,
}
