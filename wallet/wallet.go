package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/michael_cho77/go-michael-coin/utils"
)

const (
	signature     string = "77ab4f88fb7d1650bca967516e036cd99036f486320c7969ed9a05e0b9319382674e52a96b8c8254d5aa3d949afe5aa7ae849637e4f4078c23c51e01619023dc"
	privateKey    string = "30770201010420c3b823e4be15904db8dbc06478ea6073996a41b2a4f2368e64e5278a3d52fb2ea00a06082a8648ce3d030107a14403420004d2a59141caf62966c21fda3cbd7f7fe734843d6dcaec9aa2609f036da8203b2274cfea14f564846e8a2751a1f4e5801a84e4f516a6c18c6d05cc50fd6d09f342"
	hashedMessage string = "0c48bdfc98b783b5d8e0f550c8296b3dcc685096109f4cbaf42bf9f3d815de21"
)

func Start() {
	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	private, err := x509.ParseECPrivateKey(privBytes)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)
	fmt.Println("Private key:", private)

}
