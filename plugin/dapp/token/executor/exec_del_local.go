package executor

import (
	tokenty "gitlab.33.cn/chain33/chain33/plugin/dapp/token/types"
	"gitlab.33.cn/chain33/chain33/types"
)

func (t *token) execDelLocal(receiptData *types.ReceiptData) ([]*types.KeyValue, error) {
	var set []*types.KeyValue
	for i := 0; i < len(receiptData.Logs); i++ {
		item := receiptData.Logs[i]
		if item.Ty == tokenty.TyLogPreCreateToken || item.Ty == tokenty.TyLogFinishCreateToken || item.Ty == tokenty.TyLogRevokeCreateToken {
			var receipt tokenty.ReceiptToken
			err := types.Decode(item.Log, &receipt)
			if err != nil {
				tokenlog.Error("Failed to decode ReceiptToken in ExecDelLocal")
				continue
			}
			set = append(set, t.deleteLogs(&receipt)...)
		}
	}
	return set, nil
}

func (t *token) ExecDelLocal_Transfer(payload *types.AssetsTransfer, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := t.ExecDelLocalLocalTransWithdraw(tx, receiptData, index)
	if err != nil {
		return nil, err
	}
	if types.GetSaveTokenTxList() {
		tokenAction := tokenty.TokenAction{
			Ty: tokenty.ActionTransfer,
			Value: &tokenty.TokenAction_Transfer{
				payload,
			},
		}
		kvs, err := t.makeTokenTxKvs(tx, &tokenAction, receiptData, index, true)
		if err != nil {
			return nil, err
		}
		set.KV = append(set.KV, kvs...)
	}
	return set, nil
}

func (t *token) ExecDelLocal_Withdraw(payload *types.AssetsWithdraw, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := t.ExecDelLocalLocalTransWithdraw(tx, receiptData, index)
	if err != nil {
		return nil, err
	}
	if types.GetSaveTokenTxList() {
		tokenAction := tokenty.TokenAction{
			Ty: tokenty.ActionWithdraw,
			Value: &tokenty.TokenAction_Withdraw{
				payload,
			},
		}
		kvs, err := t.makeTokenTxKvs(tx, &tokenAction, receiptData, index, true)
		if err != nil {
			return nil, err
		}
		set.KV = append(set.KV, kvs...)
	}
	return set, nil
}

func (t *token) ExecDelLocal_TransferToExec(payload *types.AssetsTransferToExec, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := t.ExecDelLocalLocalTransWithdraw(tx, receiptData, index)
	if err != nil {
		return nil, err
	}
	return set, nil
}

func (t *token) ExecDelLocal_TokenPreCreate(payload *tokenty.TokenPreCreate, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	kv, err := t.execDelLocal(receiptData)
	if err != nil {
		return nil, err
	}
	return &types.LocalDBSet{KV: kv}, nil
}

func (t *token) ExecDelLocal_TokenFinishCreate(payload *tokenty.TokenFinishCreate, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	kv, err := t.execDelLocal(receiptData)
	if err != nil {
		return nil, err
	}
	return &types.LocalDBSet{KV: kv}, nil
}

func (t *token) ExecDelLocal_TokenRevokeCreate(payload *tokenty.TokenRevokeCreate, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	kv, err := t.execDelLocal(receiptData)
	if err != nil {
		return nil, err
	}
	return &types.LocalDBSet{KV: kv}, nil
}


