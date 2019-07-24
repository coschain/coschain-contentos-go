package dandelion

import (
	"github.com/coschain/contentos-go/prototype"
)

func AccountCreate(creator, account string, owner *prototype.PublicKeyType, fee uint64, jsonMeta string) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.AccountCreateOperation{
		Creator: prototype.NewAccountName(creator),
		NewAccountName: prototype.NewAccountName(account),
		Owner: owner,
		Fee: prototype.NewCoin(fee),
		JsonMetadata: jsonMeta,
	})
}

func Transfer(from, to string, amount uint64, memo string) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.TransferOperation{
		From: prototype.NewAccountName(from),
		To: prototype.NewAccountName(to),
		Amount: prototype.NewCoin(amount),
		Memo: memo,
	})
}

func AccountUpdate(name string, pubkey *prototype.PublicKeyType) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.AccountUpdateOperation{
		Owner: prototype.NewAccountName(name),
		Pubkey: pubkey,
	})
}

func TransferToVesting(from, to string, amount uint64) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.TransferToVestingOperation{
		From: prototype.NewAccountName(from),
		To: prototype.NewAccountName(to),
		Amount: prototype.NewCoin(amount),
	})
}

func Vote(voter string, postId uint64) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.VoteOperation{
		Voter: prototype.NewAccountName(voter),
		Idx: postId,
	})
}

func BpRegister(name, url, desc string, signingKey *prototype.PublicKeyType, props *prototype.ChainProperties) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.BpRegisterOperation{
		Owner: prototype.NewAccountName(name),
		Url: url,
		Desc: desc,
		BlockSigningKey: signingKey,
		Props: props,
	})
}

func BpUpdate(name string, props *prototype.ChainProperties) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.BpUpdateOperation{
		Owner: prototype.NewAccountName(name),
		ProposedStaminaFree: props.StaminaFree,
		TpsExpected: props.TpsExpected,
		AccountCreationFee: props.AccountCreationFee,
		TopNAcquireFreeToken: props.TopNAcquireFreeToken,
		EpochDuration: props.EpochDuration,
		PerTicketPrice: props.PerTicketPrice,
		PerTicketWeight: props.PerTicketWeight,
	})
}

func BpUnregister(name string) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.BpUnregisterOperation{
		Owner: prototype.NewAccountName(name),
	})
}

func BpVote(voter, bp string) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.BpVoteOperation{
		Voter: prototype.NewAccountName(voter),
		Witness: prototype.NewAccountName(bp),
	})
}

func Follow(follower, followee string, cancel bool) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.FollowOperation{
		Account: prototype.NewAccountName(follower),
		FAccount: prototype.NewAccountName(followee),
		Cancel: cancel,
	})
}

func ContractDeploy(owner, contract string, abi, code []byte, upgradable bool, url, desc string) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.ContractDeployOperation{
		Owner: prototype.NewAccountName(owner),
		Contract: contract,
		Abi: abi,
		Code: code,
		Upgradeable: upgradable,
		Url: url,
		Describe: desc,
	})
}

func ContractApply(caller, owner, contract, method, jsonParams string, coins uint64) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.ContractApplyOperation{
		Caller: prototype.NewAccountName(caller),
		Owner: prototype.NewAccountName(owner),
		Contract: contract,
		Method: method,
		Params: jsonParams,
		Amount: prototype.NewCoin(coins),
	})
}

func Post(postId uint64, author, title, content string, tags []string, beneficiaries map[string]int) *prototype.Operation {
	var benefits []*prototype.BeneficiaryRouteType
	if beneficiaries != nil {
		for name, weight := range beneficiaries {
			benefits = append(benefits, &prototype.BeneficiaryRouteType{
				Name: prototype.NewAccountName(name),
				Weight: uint32(weight),
			})
		}
	}
	return prototype.GetPbOperation(&prototype.PostOperation{
		Uuid: postId,
		Owner: prototype.NewAccountName(author),
		Title: title,
		Content: content,
		Tags: tags,
		Beneficiaries: benefits,
	})
}

func Reply(postId, parentId uint64, author, content string, beneficiaries map[string]int) *prototype.Operation {
	var benefits []*prototype.BeneficiaryRouteType
	if beneficiaries != nil {
		for name, weight := range beneficiaries {
			benefits = append(benefits, &prototype.BeneficiaryRouteType{
				Name: prototype.NewAccountName(name),
				Weight: uint32(weight),
			})
		}
	}
	return prototype.GetPbOperation(&prototype.ReplyOperation{
		Uuid: postId,
		ParentUuid: parentId,
		Owner: prototype.NewAccountName(author),
		Content: content,
		Beneficiaries: benefits,
	})
}

func Report(reporter string, postId uint64, reason []prototype.ReportOperationTag, arbitration, approved bool) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.ReportOperation{
		Reporter: prototype.NewAccountName(reporter),
		Reported: postId,
		ReportTag: reason,
		IsArbitration: arbitration,
		IsApproved: approved,
	})
}

func ConvertVesting(name string, vests uint64) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.ConvertVestingOperation{
		From: prototype.NewAccountName(name),
		Amount: prototype.NewVest(vests),
	})
}

func Stake(from, to string, coins uint64) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.StakeOperation{
		From: prototype.NewAccountName(from),
		To: prototype.NewAccountName(to),
		Amount: prototype.NewCoin(coins),
	})
}

func UnStake(creditor, debtor string, coins uint64) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.UnStakeOperation{
		Creditor: prototype.NewAccountName(creditor),
		Debtor: prototype.NewAccountName(debtor),
		Amount: prototype.NewCoin(coins),
	})
}

func AcquireTicket(name string, count uint64) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.AcquireTicketOperation{
		Account: prototype.NewAccountName(name),
		Count: count,
	})
}

func VoteByTicket(name string, idx, count uint64) *prototype.Operation {
	return prototype.GetPbOperation(&prototype.VoteByTicketOperation{
		Account: prototype.NewAccountName(name),
		Idx: idx,
		Count: count,
	})
}