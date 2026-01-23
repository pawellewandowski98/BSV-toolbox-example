package _const

import (
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/brc29"
)

const TxIDToInternalize = "299ac36833a5ffe6ae30e4d9fcebd6328a5fd4e6cae5dc4d18bda95adc1bbad1"

var AliceTxIDsToInternalize = []string{
	//"e2ced8fa7cacb3deac90ba19aca8dd5ef1280f9f8f76c0ff314700121a3fe4e6",
	//"5790fef96c99e3e7fcc52c484be73ceeb8e2eb2e615ff0c1d9cd85ebd101b5be",
	//"074d1aaafbe68716557b2942de003743164419256048293792c4e0c835e96a84",
	//"ebc6b813e677fa72e988c28aed1c4206c055d6ed4ca897905126aa8f3884c5f5",
	//"e2bf6acf6666ee06e652a256128ba86308053fa2e03bd574790590dbe41daf00",
	"18e55d95961245041875664d7598cfeb0dfea15668b901a24dfdbae68d199a8f",
}

var BobTxIDsToInternalize = []string{
	//"220e5262647615549f6310260d6b37da23c1ed972e94c12edf8191c14f1fb9e4",
	//"fb6671546ac071f720bbc3beb5bbee406e0a705f2d8a3b28f2cf22ac67d4acca",
	//"b0f2756dfd42f901d6313fb60970bfb116ab913fa1ed41e2003a58937cba46f3",
	//"fe9f9c6c10acd43068a3f61052c2db77f35b844807861e6881e689a892e198c5",
	//"3892aa7eb226129829628bc4292cf24a1efe01b61b20446eb1c6006e825574a2",
}

const (
	DefaultBase64Prefix = "SfKxPIJNgdI="
	DefaultBase64Suffix = "NaGLC6fMH50="
)

var KeyID = brc29.KeyID{
	DerivationPrefix: DefaultBase64Prefix,
	DerivationSuffix: DefaultBase64Suffix,
}
