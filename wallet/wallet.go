package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/michael_cho77/go-michael-coin/utils"
)

const (
	signature     string = "77ab4f88fb7d1650bca967516e036cd99036f486320c7969ed9a05e0b9319382674e52a96b8c8254d5aa3d949afe5aa7ae849637e4f4078c23c51e01619023dc"
	privateKey    string = "30770201010420c3b823e4be15904db8dbc06478ea6073996a41b2a4f2368e64e5278a3d52fb2ea00a06082a8648ce3d030107a14403420004d2a59141caf62966c21fda3cbd7f7fe734843d6dcaec9aa2609f036da8203b2274cfea14f564846e8a2751a1f4e5801a84e4f516a6c18c6d05cc50fd6d09f342"
	hashedMessage string = "0c48bdfc98b783b5d8e0f550c8296b3dcc685096109f4cbaf42bf9f3d815de21"
)

func Start() {

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)

	fmt.Printf("%x\n\n\n\n\n", keyAsBytes)

	utils.HandleErr(err)

	hashAsBytes, err := hex.DecodeString(hashedMessage)

	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)

	signature := append(r.Bytes(), s.Bytes()...)

	fmt.Printf("%x\n", signature)

	utils.HandleErr(err)
}
