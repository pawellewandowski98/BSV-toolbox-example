package _const

import (
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/brc29"
)

const TxIDToInternalize = "299ac36833a5ffe6ae30e4d9fcebd6328a5fd4e6cae5dc4d18bda95adc1bbad1"

var AliceTxIDsToInternalize = []string{
	//"e2ced8fa7cacb3deac90ba19aca8dd5ef1280f9f8f76c0ff314700121a3fe4e6", // alice 100
	//"56142a70cc753e6ecb252d8aab6fac3d85a87c0a9c8b3dfbc7d2aa749c927788", // alice 500
	//"220e5262647615549f6310260d6b37da23c1ed972e94c12edf8191c14f1fb9e4",
	//"e624f731861a004f58439af3caa050d99604fddea7fd1bc975dad5bbc464d9bc",
	//"be4b937debb7237a94e6e47f2d797e8298a2d512064d7acf700d4fa36c589b3a",
	//"5790fef96c99e3e7fcc52c484be73ceeb8e2eb2e615ff0c1d9cd85ebd101b5be",
	//"074d1aaafbe68716557b2942de003743164419256048293792c4e0c835e96a84",
	//"ebc6b813e677fa72e988c28aed1c4206c055d6ed4ca897905126aa8f3884c5f5",
	//"d939ae40f52939783672e52362ec9dafc55bdbb50c7e2bca883337469b4083a2",
	"b49384ba63e214ac729bb2d44f612d7f5569df3d16c8592015e99cef6ec73acc",
	"ea64e777edd49b6505bca22b91f5725091882e8ca1ba12e29b47145b341ce1a3",
}

var BobTxIDsToInternalize = []string{
	//"220e5262647615549f6310260d6b37da23c1ed972e94c12edf8191c14f1fb9e4",
	//"fb6671546ac071f720bbc3beb5bbee406e0a705f2d8a3b28f2cf22ac67d4acca",
	//"b0f2756dfd42f901d6313fb60970bfb116ab913fa1ed41e2003a58937cba46f3",
	//"fe9f9c6c10acd43068a3f61052c2db77f35b844807861e6881e689a892e198c5",
}

const (
	DefaultBase64Prefix = "SfKxPIJNgdI="
	DefaultBase64Suffix = "NaGLC6fMH50="
)

var KeyID = brc29.KeyID{
	DerivationPrefix: DefaultBase64Prefix,
	DerivationSuffix: DefaultBase64Suffix,
}
