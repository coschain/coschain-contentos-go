package utils

import (
	"fmt"
	"github.com/coschain/contentos-go/cmd/wallet-cli/wallet"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/prototype"
	"hash/crc32"
	"math/rand"
	"syscall"
	"time"
)

func GenerateSignedTxAndValidate(ops []interface{}, signers ...*wallet.PrivAccount) (*prototype.SignedTransaction, error) {
	privKeys := []*prototype.PrivateKeyType{}
	for _, acc := range signers {
		privKey, err := prototype.PrivateKeyFromWIF(acc.PrivKey)
		if err != nil {
			return nil, err
		}
		privKeys = append(privKeys, privKey)
	}
	// occupant implement
	tx := &prototype.Transaction{RefBlockNum: 0, RefBlockPrefix: 0, Expiration: &prototype.TimePointSec{UtcSeconds: uint32(time.Now().Unix()) + constants.TRX_MAX_EXPIRATION_TIME}}
	for _, op := range ops {
		tx.AddOperation(op)
	}

	signTx := prototype.SignedTransaction{Trx: tx}
	for _, privkey := range privKeys {
		res := signTx.Sign(privkey, prototype.ChainId{Value: 0})
		signTx.Signatures = append(signTx.Signatures, &prototype.SignatureType{Sig: res})
	}

	if err := signTx.Validate(); err != nil {
		return nil, err
	}

	return &signTx, nil
}

func GenerateUUID(content string) uint64 {
	crc32q := crc32.MakeTable(0xD5828281)
	randContent := content + string(rand.Intn(1e5))
	return uint64(time.Now().Unix()*1e9) + uint64(crc32.Checksum([]byte(randContent), crc32q))
}

func GetPassphrase(reader PasswordReader) (string, error) {
	fmt.Print("Enter passphrase > ")
	bytePassphrase, err := reader.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", err
	}
	passphrase := string(bytePassphrase)
	return passphrase, nil
}
