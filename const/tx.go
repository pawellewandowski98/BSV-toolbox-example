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
	//"e554d280bcde0685dc25a3829b1a3a7e1dbe681353fb3f4558221e3e6aad4ecd",
	//"8e1990441a2fe04fec267b55cbea53064d9786dc1363c972c0cace55ee3e08b6",
	//"4c221282369b29e35a59b5da8596d9cabeff1e680e15f1e6d7ec9db4b9720f7f",
	"de61ed21f90d934a35b7528a20e9baac021628ee3de6c46b46a833b3470cb4d7",
}

var BobTxIDsToInternalize = []string{
	//"220e5262647615549f6310260d6b37da23c1ed972e94c12edf8191c14f1fb9e4",
	//"fb6671546ac071f720bbc3beb5bbee406e0a705f2d8a3b28f2cf22ac67d4acca",
	//"b0f2756dfd42f901d6313fb60970bfb116ab913fa1ed41e2003a58937cba46f3",
	//"fe9f9c6c10acd43068a3f61052c2db77f35b844807861e6881e689a892e198c5",
	//"3892aa7eb226129829628bc4292cf24a1efe01b61b20446eb1c6006e825574a2",
	//"48f0fecf968925c6d2af83cf1e08f24966d1910bd9db9b3db08122fd340a74c3",
}

const (
	DefaultBase64Prefix = "SfKxPIJNgdI="
	DefaultBase64Suffix = "NaGLC6fMH50="
)

var KeyID = brc29.KeyID{
	DerivationPrefix: DefaultBase64Prefix,
	DerivationSuffix: DefaultBase64Suffix,
}
