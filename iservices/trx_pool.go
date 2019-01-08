package iservices

import (
	"github.com/coschain/contentos-go/common"
	"github.com/coschain/contentos-go/prototype"
)

//
// This file defines interfaces of Database service.
//

var TxPoolServerName = "ctrl"

type ITrxPool interface {
	PushTrx(trx *prototype.SignedTransaction) *prototype.TransactionReceiptWithInfo
	PushBlock(blk *prototype.SignedBlock, skip prototype.SkipFlag) error
	HeadBlockTime() *prototype.TimePointSec
	GenerateBlock(witness string, pre *prototype.Sha256, timestamp uint32, priKey *prototype.PrivateKeyType, skip prototype.SkipFlag) *prototype.SignedBlock
	GetWitnessTopN(n uint32) []string
	SetShuffledWitness(names []string)
	GetShuffledWitness() []string
	SetShuffle(s common.ShuffleFunc)
	// will set DB status to num
	PopBlockTo(num uint64)
	// will cut off DB status that before num
	Commit(num uint64)

	GetProps() *prototype.DynamicProperties

	TransferToVest(value *prototype.Coin)
	TransferFromVest(value *prototype.Vest)

	AddWeightedVP(value uint64)
	// put trx into pending directly, no return value, so should be used by witness node to collect p2p trx
	PushTrxToPending(trx *prototype.SignedTransaction)
	GenerateAndApplyBlock(witness string, pre *prototype.Sha256, timestamp uint32, priKey *prototype.PrivateKeyType, skip prototype.SkipFlag) (*prototype.SignedBlock, error)
	VerifySig(name *prototype.AccountName, digest []byte, sig []byte) bool
	Sign(priv *prototype.PrivateKeyType, digest []byte) []byte
}
