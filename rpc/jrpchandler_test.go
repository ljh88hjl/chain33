// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"encoding/hex"

	"github.com/33cn/chain33/account"
	"github.com/33cn/chain33/client"
	"github.com/33cn/chain33/client/mocks"
	"github.com/33cn/chain33/common"
	rpctypes "github.com/33cn/chain33/rpc/types"
	_ "github.com/33cn/chain33/system"
	cty "github.com/33cn/chain33/system/dapp/coins/types"
	mty "github.com/33cn/chain33/system/dapp/manage/types"
	"github.com/33cn/chain33/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDecodeLogErr(t *testing.T) {
	enc := "0001020304050607"
	dec := []byte{0, 1, 2, 3, 4, 5, 6, 7}

	hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogErr,
		Log: "0x" + enc,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   1,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogErr", result.Logs[0].TyName)
	assert.Equal(t, int32(types.TyLogErr), result.Logs[0].Ty)
}

func TestDecodeLogFee(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogFee,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogFee", result.Logs[0].TyName)
}

func TestDecodeLogTransfer(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogTransfer,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogTransfer", result.Logs[0].TyName)
}

func TestDecodeLogGenesis(t *testing.T) {
	enc := "0001020304050607"

	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogGenesis,
		Log: "0x" + enc,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	//这个已经废弃
	assert.Equal(t, "unkownType", result.Logs[0].TyName)
}

func TestDecodeLogDeposit(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogDeposit,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogDeposit", result.Logs[0].TyName)
}

func TestDecodeLogExecTransfer(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptExecAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogExecTransfer,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogExecTransfer", result.Logs[0].TyName)
}

func TestDecodeLogExecWithdraw(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptExecAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogExecWithdraw,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogExecWithdraw", result.Logs[0].TyName)
}

func TestDecodeLogExecDeposit(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptExecAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogExecDeposit,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogExecDeposit", result.Logs[0].TyName)
}

func TestDecodeLogExecFrozen(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptExecAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogExecFrozen,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogExecFrozen", result.Logs[0].TyName)
}

func TestDecodeLogExecActive(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptExecAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogExecActive,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogExecActive", result.Logs[0].TyName)
}

func TestDecodeLogGenesisTransfer(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogGenesisTransfer,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogGenesisTransfer", result.Logs[0].TyName)
}

func TestDecodeLogGenesisDeposit(t *testing.T) {
	var account = &types.Account{}
	var logTmp = &types.ReceiptAccountTransfer{
		Prev:    account,
		Current: account,
	}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  types.TyLogGenesisDeposit,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("coins"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogGenesisDeposit", result.Logs[0].TyName)
}

func TestDecodeLogModifyConfig(t *testing.T) {
	var logTmp = &types.ReceiptConfig{}
	dec := types.Encode(logTmp)
	strdec := hex.EncodeToString(dec)
	rlog := &rpctypes.ReceiptLog{
		Ty:  mty.TyLogModifyConfig,
		Log: "0x" + strdec,
	}

	logs := []*rpctypes.ReceiptLog{}
	logs = append(logs, rlog)

	var data = &rpctypes.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	result, err := rpctypes.DecodeLog([]byte("manage"), data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "LogModifyConfig", result.Logs[0].TyName)
}

func newTestChain33(api client.QueueProtocolAPI) *Chain33 {
	types.AssertConfig(api)
	return &Chain33{
		cli: channelClient{
			QueueProtocolAPI: api,
			accountdb:        account.NewCoinsAccount(api.GetConfig()),
		},
	}
}

func TestChain33_CreateRawTransaction(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	// var result interface{}
	// api.On("CreateRawTransaction", nil, &result).Return()
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)
	var testResult interface{}
	err := testChain33.CreateRawTransaction(nil, &testResult)
	assert.Nil(t, testResult)
	assert.NotNil(t, err)

	tx := &rpctypes.CreateTx{
		To:          "184wj4nsgVxKyz2NhM3Yb5RK5Ap6AFRFq2",
		Amount:      10,
		Fee:         1,
		Note:        "12312",
		IsWithdraw:  false,
		IsToken:     false,
		TokenSymbol: "",
		ExecName:    cfg.ExecName("coins"),
	}

	err = testChain33.CreateRawTransaction(tx, &testResult)
	assert.NotNil(t, testResult)
	assert.Nil(t, err)
}

func TestChain33_ReWriteRawTx(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)
	txHex1 := "0a05636f696e73122c18010a281080c2d72f222131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b7120a08d0630a696c0b3f78dd9ec083a2131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b71"
	//txHex2 := "0a05636f696e73122d18010a29108084af5f222231484c53426e7437486e486a7857797a636a6f573863663259745550663337594d6320a08d0630dbc4cbf6fbc4e1d0533a2231484c53426e7437486e486a7857797a636a6f573863663259745550663337594d63"

	reTx := &rpctypes.ReWriteRawTx{
		Tx:     txHex1,
		Fee:    29977777777,
		Expire: "130s",
		To:     "aabbccdd",
		Index:  0,
	}
	var testResult interface{}
	err := testChain33.ReWriteRawTx(reTx, &testResult)
	assert.Nil(t, err)
	assert.NotNil(t, testResult)
	assert.NotEqual(t, txHex1, testResult)
	txData, err := common.FromHex(testResult.(string))
	assert.Nil(t, err)
	tx := &types.Transaction{}
	err = types.Decode(txData, tx)
	assert.Nil(t, err)
	assert.Equal(t, tx.Fee, reTx.Fee)
	assert.Equal(t, reTx.To, tx.To)

}

func TestChain33_CreateTxGroup(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)
	var testResult interface{}
	api.On("GetProperFee", mock.Anything).Return(nil, nil)
	err := testChain33.CreateRawTxGroup(nil, &testResult)
	assert.Nil(t, testResult)
	assert.NotNil(t, err)

	txHex1 := "0a05636f696e73122c18010a281080c2d72f222131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b7120a08d0630a696c0b3f78dd9ec083a2131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b71"
	txHex2 := "0a05636f696e73122d18010a29108084af5f222231484c53426e7437486e486a7857797a636a6f573863663259745550663337594d6320a08d0630dbc4cbf6fbc4e1d0533a2231484c53426e7437486e486a7857797a636a6f573863663259745550663337594d63"
	txs := &types.CreateTransactionGroup{
		Txs: []string{txHex1, txHex2},
	}
	err = testChain33.CreateRawTxGroup(txs, &testResult)
	assert.Nil(t, err)
	tx, err := decodeTx(testResult.(string))
	assert.Nil(t, err)
	tg, err := tx.GetTxGroup()
	assert.Nil(t, err)
	if len(tg.GetTxs()) != 2 {
		t.Error("Test createtxgroup failed")
		return
	}
	err = tx.Check(cfg, 0, cfg.GetMinTxFeeRate(), cfg.GetMaxTxFee())
	assert.Nil(t, err)
}

func TestChain33_SendTransaction(t *testing.T) {
	//if types.IsPara() {
	//	t.Skip()
	//	return
	//}
	api := new(mocks.QueueProtocolAPI)
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api.On("GetConfig", mock.Anything).Return(cfg)
	tx := &types.Transaction{}
	api.On("SendTx", tx).Return(nil, errors.New("error value"))
	testChain33 := newTestChain33(api)
	var testResult interface{}
	data := rpctypes.RawParm{
		Data: "",
	}
	err := testChain33.SendTransaction(data, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_SendTransactions(t *testing.T) {

	api := new(mocks.QueueProtocolAPI)
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api.On("GetConfig", mock.Anything).Return(cfg)

	api.On("SendTx", mock.Anything).Return(&types.Reply{
		IsOk: true,
		Msg:  []byte("test"),
	}, nil)
	testChain33 := newTestChain33(api)
	var testResult interface{}
	txCount := 10
	data := rpctypes.ReqStrings{
		Datas: make([]string, txCount),
	}
	err := testChain33.SendTransactions(data, &testResult)
	require.Nil(t, err)
	reply, ok := testResult.(rpctypes.ReplyHashes)
	require.True(t, ok)
	require.Equal(t, txCount, len(reply.Hashes))
	require.Equal(t, common.ToHex([]byte("test")), reply.Hashes[0])
}

func TestChain33_SendTransactionSync(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api.On("GetConfig", mock.Anything).Return(cfg)
	tx := &types.Transaction{}
	hash := tx.Hash()
	api.On("SendTx", tx).Return(&types.Reply{IsOk: true, Msg: hash}, nil)
	api.On("QueryTx", mock.Anything).Return(&types.TransactionDetail{}, nil)
	testChain33 := newTestChain33(api)
	var testResult interface{}
	data := rpctypes.RawParm{
		Data: common.ToHex(types.Encode(tx)),
	}
	err := testChain33.SendTransactionSync(data, &testResult)
	t.Log(err)
	assert.Equal(t, common.ToHex(hash), testResult.(string))
	assert.Nil(t, err)
}

func TestChain33_GetHexTxByHash(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	api.On("QueryTx", &types.ReqHash{Hash: []byte("")}).Return(nil, errors.New("error value"))
	testChain33 := newTestChain33(api)
	var testResult interface{}
	data := rpctypes.QueryParm{
		Hash: "",
	}
	err := testChain33.GetHexTxByHash(data, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_QueryTransaction(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	api.On("QueryTx", &types.ReqHash{Hash: []byte("")}).Return(nil, errors.New("error value"))
	testChain33 := newTestChain33(api)
	var testResult interface{}
	data := rpctypes.QueryParm{
		Hash: "",
	}
	err := testChain33.QueryTransaction(data, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_QueryTransactionOk(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	data := rpctypes.QueryParm{
		Hash: "",
	}
	var act = &cty.CoinsAction{
		Ty: 1,
	}
	payload := types.Encode(act)
	var tx = &types.Transaction{
		Execer:  []byte(cfg.ExecName("ticket")),
		Payload: payload,
	}

	var logTmp = &types.ReceiptAccountTransfer{}

	dec := types.Encode(logTmp)

	strdec := hex.EncodeToString(dec)
	strdec = "0x" + strdec

	rlog := &types.ReceiptLog{
		Ty:  types.TyLogTransfer,
		Log: []byte(strdec),
	}

	logs := []*types.ReceiptLog{}
	logs = append(logs, rlog)

	var rdata = &types.ReceiptData{
		Ty:   5,
		Logs: logs,
	}
	reply := types.TransactionDetail{
		Tx:      tx,
		Receipt: rdata,
		Height:  10,
	}

	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	api.On("QueryTx", &types.ReqHash{Hash: []byte("")}).Return(&reply, nil)
	testChain33 := newTestChain33(api)
	var testResult interface{}

	err := testChain33.QueryTransaction(data, &testResult)
	t.Log(err)
	assert.Nil(t, err)
	assert.Equal(t, testResult.(*rpctypes.TransactionDetail).Height, reply.Height)
	assert.Equal(t, testResult.(*rpctypes.TransactionDetail).Tx.Execer, string(tx.Execer))

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetBlocks(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	api.On("GetBlocks", &types.ReqBlocks{Pid: []string{""}}).Return(&types.BlockDetails{Items: []*types.BlockDetail{{}}}, nil)
	testChain33 := newTestChain33(api)
	var testResult interface{}
	data := rpctypes.BlockParam{}
	err := testChain33.GetBlocks(data, &testResult)
	t.Log(err)
	assert.NoError(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetLastHeader(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	api.On("GetLastHeader", mock.Anything).Return(&types.Header{}, nil)
	testChain33 := newTestChain33(api)
	var testResult interface{}
	data := &types.ReqNil{}
	err := testChain33.GetLastHeader(data, &testResult)
	t.Log(err)
	assert.NotNil(t, &testResult)
	assert.NoError(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetTxByAddr(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	api.On("GetTransactionByAddr", mock.Anything).Return(&types.ReplyTxInfos{TxInfos: []*types.ReplyTxInfo{{}}}, nil)
	var testResult interface{}
	data := types.ReqAddr{}
	err := testChain33.GetTxByAddr(&data, &testResult)
	t.Log(err)
	assert.NotNil(t, testResult)
	assert.NoError(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetTxByHashes(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	api.On("GetTransactionByHash", mock.Anything).Return(&types.TransactionDetails{}, nil)
	var testResult interface{}
	data := rpctypes.ReqHashes{}
	data.Hashes = append(data.Hashes, "0xdcf13a93e3bf58534c773e13d339894c18dafbd3ff273a9d1caa0c2bec8e8cd6")
	err := testChain33.GetTxByHashes(data, &testResult)
	t.Log(err)
	assert.NotNil(t, testResult)
	assert.NoError(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetMempool(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	api.On("GetMempool", &types.ReqGetMempool{}).Return(&types.ReplyTxList{Txs: []*types.Transaction{{}}}, nil)
	var testResult interface{}
	data := &types.ReqGetMempool{IsAll: false}
	err := testChain33.GetMempool(data, &testResult)
	t.Log(err)
	assert.NotNil(t, testResult)
	assert.NoError(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetAccountsV2(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	api.On("ExecWalletFunc", "wallet", "WalletGetAccountList", mock.Anything).Return(&types.WalletAccounts{Wallets: []*types.WalletAccount{{}}}, nil)
	var testResult interface{}
	err := testChain33.GetAccountsV2(nil, &testResult)
	t.Log(err)
	assert.NotNil(t, testResult)
	assert.NoError(t, err)
}

func TestChain33_GetAccounts(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	api.On("ExecWalletFunc", "wallet", "WalletGetAccountList", mock.Anything).Return(nil, errors.New("error value"))
	var testResult interface{}
	data := &types.ReqAccountList{}
	err := testChain33.GetAccounts(data, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_NewAccount(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	api.On("ExecWalletFunc", "wallet", "NewAccount", &types.ReqNewAccount{}).Return(nil, errors.New("error value"))

	var testResult interface{}
	err := testChain33.NewAccount(&types.ReqNewAccount{}, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)
	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetAccount(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	api.On("ExecWalletFunc", "wallet", "WalletGetAccount", &types.ReqGetAccount{}).Return(nil, errors.New("error value"))

	var testResult interface{}
	err := testChain33.GetAccount(&types.ReqGetAccount{}, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)
	mock.AssertExpectationsForObjects(t, api)
}
func TestChain33_WalletTxList(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.ReqWalletTransactionList{FromTx: []byte("")}
	api.On("ExecWalletFunc", "wallet", "WalletTransactionList", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := rpctypes.ReqWalletTransactionList{}
	err := testChain33.WalletTxList(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_ImportPrivkey(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.ReqWalletImportPrivkey{}
	api.On("ExecWalletFunc", "wallet", "WalletImportPrivkey", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqWalletImportPrivkey{}
	err := testChain33.ImportPrivkey(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_SendToAddress(t *testing.T) {
	//if types.IsPara() {
	//	t.Skip()
	//	return
	//}
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.ReqWalletSendToAddress{}
	api.On("ExecWalletFunc", "wallet", "WalletSendToAddress", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqWalletSendToAddress{}
	err := testChain33.SendToAddress(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_SetTxFee(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.ReqWalletSetFee{}
	api.On("ExecWalletFunc", "wallet", "WalletSetFee", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqWalletSetFee{}
	err := testChain33.SetTxFee(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_SetLabl(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.ReqWalletSetLabel{}
	api.On("ExecWalletFunc", "wallet", "WalletSetLabel", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqWalletSetLabel{}
	err := testChain33.SetLabl(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_MergeBalance(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.ReqWalletMergeBalance{}
	api.On("ExecWalletFunc", "wallet", "WalletMergeBalance", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqWalletMergeBalance{}
	err := testChain33.MergeBalance(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_SetPasswd(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.ReqWalletSetPasswd{}
	api.On("ExecWalletFunc", "wallet", "WalletSetPasswd", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqWalletSetPasswd{}
	err := testChain33.SetPasswd(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_Lock(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := types.ReqNil{}
	api.On("ExecWalletFunc", "wallet", "WalletLock", &expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqNil{}
	err := testChain33.Lock(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_UnLock(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.WalletUnLock{}
	api.On("ExecWalletFunc", "wallet", "WalletUnLock", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.WalletUnLock{}
	err := testChain33.UnLock(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetPeerInfo(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	api.On("PeerInfo", mock.Anything).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.P2PGetPeerReq{}
	err := testChain33.GetPeerInfo(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetPeerInfoOk(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	var peerlist types.PeerList
	var pr = &types.Peer{
		Addr: "abcdsd",
	}
	peerlist.Peers = append(peerlist.Peers, pr)

	api.On("PeerInfo", mock.Anything).Return(&peerlist, nil)
	var testResult interface{}
	var in types.P2PGetPeerReq
	_ = testChain33.GetPeerInfo(&in, &testResult)
	assert.Equal(t, testResult.(*rpctypes.PeerList).Peers[0].Addr, peerlist.Peers[0].Addr)
}

func TestChain33_GetHeaders(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.ReqBlocks{}
	api.On("GetHeaders", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqBlocks{}
	err := testChain33.GetHeaders(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetHeadersOk(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	var headers types.Headers
	var header = &types.Header{
		TxCount: 10,
	}
	headers.Items = append(headers.Items, header)

	expected := &types.ReqBlocks{}
	api.On("GetHeaders", expected).Return(&headers, nil)

	var testResult interface{}
	actual := types.ReqBlocks{}
	err := testChain33.GetHeaders(&actual, &testResult)
	assert.Nil(t, err)
	assert.Equal(t, testResult.(*rpctypes.Headers).Items[0].TxCount, header.TxCount)

}

func TestChain33_GetLastMemPool(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	// expected := &types.ReqBlocks{}
	api.On("GetLastMempool").Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqNil{}
	err := testChain33.GetLastMemPool(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetProperFee(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := types.ReqProperFee{}
	api.On("GetProperFee", &expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	err := testChain33.GetProperFee(&expected, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetBlockOverview(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.ReqHash{Hash: []byte{}}
	api.On("GetBlockOverview", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := rpctypes.QueryParm{}
	err := testChain33.GetBlockOverview(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetBlockOverviewOk(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)
	var head = &types.Header{
		Hash: []byte("123456"),
	}
	var replyblock = &types.BlockOverview{
		Head:    head,
		TxCount: 1,
	}

	expected := &types.ReqHash{Hash: []byte{0x12, 0x34, 0x56}}
	api.On("GetBlockOverview", expected).Return(replyblock, nil)

	var testResult interface{}
	actual := rpctypes.QueryParm{Hash: "123456"}

	err := testChain33.GetBlockOverview(actual, &testResult)
	t.Log(err)
	assert.Nil(t, err)
	assert.Equal(t, testResult.(*rpctypes.BlockOverview).TxCount, replyblock.TxCount)
	assert.Equal(t, testResult.(*rpctypes.BlockOverview).Head.Hash, common.ToHex(replyblock.Head.Hash))
	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetAddrOverview(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.ReqAddr{}
	api.On("GetAddrOverview", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqAddr{}
	err := testChain33.GetAddrOverview(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	// mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetBlockHash(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.ReqInt{}
	api.On("GetBlockHash", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.ReqInt{}
	err := testChain33.GetBlockHash(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GenSeed(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.GenSeedLang{}
	api.On("ExecWalletFunc", "wallet", "GenSeed", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.GenSeedLang{}
	err := testChain33.GenSeed(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_SaveSeed(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.SaveSeedByPw{}
	api.On("ExecWalletFunc", "wallet", "SaveSeed", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.SaveSeedByPw{}
	err := testChain33.SaveSeed(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetSeed(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := &types.GetSeedByPw{}
	api.On("ExecWalletFunc", "wallet", "GetSeed", expected).Return(nil, errors.New("error value"))

	var testResult interface{}
	actual := types.GetSeedByPw{}
	err := testChain33.GetSeed(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}

func TestChain33_GetWalletStatus(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)

	expected := types.ReqNil{}
	api.On("ExecWalletFunc", "wallet", "GetWalletStatus", &expected).Return(nil, errors.New("error value")).Once()

	var testResult interface{}
	actual := types.ReqNil{}
	err := testChain33.GetWalletStatus(&actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	expect := types.WalletStatus{
		IsWalletLock: true,
		IsAutoMining: true,
		IsHasSeed:    false,
		IsTicketLock: false,
	}

	api.On("ExecWalletFunc", "wallet", "GetWalletStatus", &expected).Return(&expect, nil).Once()
	err = testChain33.GetWalletStatus(&actual, &testResult)
	t.Log(err)
	assert.Nil(t, err)
	status, ok := testResult.(*rpctypes.WalletStatus)
	if !ok {
		t.Error("GetWalletStatus type error")
	}
	assert.Equal(t, expect.IsWalletLock, status.IsWalletLock)
	assert.Equal(t, expect.IsAutoMining, status.IsAutoMining)
	assert.Equal(t, expect.IsHasSeed, status.IsHasSeed)
	assert.Equal(t, expect.IsTicketLock, status.IsTicketLock)

	mock.AssertExpectationsForObjects(t, api)
}

// ----------------------------

func TestChain33_Version(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)
	var testResult interface{}
	in := &types.ReqNil{}
	ver := &types.VersionInfo{Chain33: "6.0.2"}
	api.On("Version", mock.Anything).Return(ver, nil)
	err := testChain33.Version(in, &testResult)
	t.Log(err)
	t.Log(testResult)
	assert.Equal(t, nil, err)
	assert.NotNil(t, testResult)
}

func TestChain33_GetTimeStatus(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var result interface{}
	err := client.GetTimeStatus(&types.ReqNil{}, &result)
	assert.Nil(t, err)
}

func TestChain33_GetServerTime(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var result interface{}
	err := client.GetServerTime(&types.ReqNil{}, &result)
	assert.Nil(t, err)
}

func TestChain33_GetLastBlockSequence(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var result interface{}
	api.On("GetLastBlockSequence", mock.Anything).Return(nil, types.ErrInvalidParam)
	err := client.GetLastBlockSequence(&types.ReqNil{}, &result)
	assert.NotNil(t, err)

	api = new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client = newTestChain33(api)
	var result2 interface{}
	lastSeq := types.Int64{Data: 1}
	api.On("GetLastBlockSequence", mock.Anything).Return(&lastSeq, nil)
	err = client.GetLastBlockSequence(&types.ReqNil{}, &result2)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), result2)
}

func TestChain33_GetBlockSequences(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var result interface{}
	api.On("GetBlockSequences", mock.Anything).Return(nil, types.ErrInvalidParam)
	err := client.GetBlockSequences(rpctypes.BlockParam{}, &result)
	assert.NotNil(t, err)

	api = new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client = newTestChain33(api)
	var result2 interface{}
	blocks := types.BlockSequences{}
	blocks.Items = make([]*types.BlockSequence, 0)
	blocks.Items = append(blocks.Items, &types.BlockSequence{Hash: []byte("h1"), Type: 1})
	api.On("GetBlockSequences", mock.Anything).Return(&blocks, nil)
	err = client.GetBlockSequences(rpctypes.BlockParam{}, &result2)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result2.(*rpctypes.ReplyBlkSeqs).BlkSeqInfos))
}

func TestChain33_GetBlockByHashes(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	in := rpctypes.ReqHashes{Hashes: []string{}}
	in.Hashes = append(in.Hashes, common.ToHex([]byte("h1")))
	api.On("GetBlockByHashes", mock.Anything).Return(&types.BlockDetails{}, nil)
	err := client.GetBlockByHashes(in, &testResult)
	assert.Nil(t, err)

	api = new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client = newTestChain33(api)
	var testResult2 interface{}
	api.On("GetBlockByHashes", mock.Anything).Return(nil, types.ErrInvalidParam)
	err = client.GetBlockByHashes(in, &testResult2)
	assert.NotNil(t, err)
}

func TestChain33_CreateTransaction(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)

	var result interface{}
	err := client.CreateTransaction(nil, &result)
	assert.NotNil(t, err)

	in := &rpctypes.CreateTxIn{Execer: "notExist", ActionName: "x", Payload: []byte("x")}
	err = client.CreateTransaction(in, &result)
	assert.Equal(t, types.ErrExecNotFound, err)

	in = &rpctypes.CreateTxIn{Execer: cfg.ExecName("coins"), ActionName: "notExist", Payload: []byte("x")}
	err = client.CreateTransaction(in, &result)
	assert.Equal(t, types.ErrActionNotSupport, err)

	in = &rpctypes.CreateTxIn{
		Execer:     cfg.ExecName("coins"),
		ActionName: "Transfer",
		Payload:    []byte("{\"to\": \"addr\", \"amount\":\"10\"}"),
	}
	err = client.CreateTransaction(in, &result)
	assert.Nil(t, err)
}

func TestChain33_GetExecBalance(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	in := &types.ReqGetExecBalance{}
	api.On("StoreList", mock.Anything).Return(&types.StoreListReply{}, nil)
	err := client.GetExecBalance(in, &testResult)
	assert.Nil(t, err)

	api = new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client = newTestChain33(api)
	var testResult2 interface{}
	api.On("StoreList", mock.Anything).Return(nil, types.ErrInvalidParam)
	err = client.GetExecBalance(in, &testResult2)
	assert.NotNil(t, err)
}

func TestChain33_GetBalance(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)

	var addrs = []string{"1Jn2qu84Z1SUUosWjySggBS9pKWdAP3tZt"}
	cases := []struct {
		In *types.ReqBalance
	}{
		{In: &types.ReqBalance{
			Execer:    cfg.ExecName("coins"),
			Addresses: addrs,
		}},
		{In: &types.ReqBalance{
			Execer:    cfg.ExecName("ticket"),
			Addresses: addrs,
		}},

		{In: &types.ReqBalance{
			AssetSymbol: "bty",
			AssetExec:   "coins",
			Execer:      cfg.ExecName("ticket"),
			Addresses:   addrs,
		}},
		{In: &types.ReqBalance{
			AssetSymbol: "bty",
			AssetExec:   "coins",
			Execer:      cfg.ExecName("coins"),
			Addresses:   addrs,
		}},
	}

	for _, c := range cases {
		c := c
		t.Run("test GetBalance", func(t *testing.T) {
			head := &types.Header{StateHash: []byte("sdfadasds")}
			api.On("GetLastHeader").Return(head, nil)

			var acc = &types.Account{Addr: "1Jn2qu84Z1SUUosWjySggBS9pKWdAP3tZt", Balance: 100}
			accv := types.Encode(acc)
			storevalue := &types.StoreReplyValue{}
			storevalue.Values = append(storevalue.Values, accv)
			api.On("StoreGet", mock.Anything).Return(storevalue, nil)

			var data interface{}
			err := client.GetBalance(c.In, &data)
			assert.Nil(t, err)
			result := data.([]*rpctypes.Account)
			assert.Equal(t, 1, len(result))
			//t.Error("result", "x", result)
			assert.Equal(t, acc.Addr, result[0].Addr)
			assert.Equal(t, int64(100), result[0].Balance)
		})
	}

	//测试地址不合法返回ErrInvalidAddress
	var data1 interface{}
	var addrs1 = []string{"17n2qu84Z1SUUosWjySggBS9pKWdAP3tZt"}
	input := types.ReqBalance{
		Execer:    cfg.ExecName("coins"),
		Addresses: addrs1,
	}

	err := client.GetBalance(&input, &data1)
	assert.Equal(t, err, types.ErrInvalidAddress)

	//测试多重签名地址不合法返回ErrInvalidAddress
	var addrs2 = []string{"3BJqXn4v741wDJY6Fzb4YbLSftXwgDzFE8"}
	input = types.ReqBalance{
		Execer:    cfg.ExecName("coins"),
		Addresses: addrs2,
	}

	err = client.GetBalance(&input, &data1)
	assert.Equal(t, err, types.ErrInvalidAddress)

	//测试多重签名地址合法
	var addrs3 = []string{"3BJqXn4v741wDJY6Fzb4YbLSftXwgDzFE7"}
	input = types.ReqBalance{
		Execer:    cfg.ExecName("coins"),
		Addresses: addrs3,
	}

	err = client.GetBalance(&input, &data1)
	assert.Nil(t, err)

}

func TestChain33_CreateNoBalanceTransaction(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	chain33 := newTestChain33(api)
	api.On("GetProperFee", mock.Anything).Return(&types.ReplyProperFee{ProperFee: 1000000}, nil)
	var result string
	err := chain33.CreateNoBalanceTransaction(&types.NoBalanceTx{TxHex: "0a05636f696e73122c18010a281080c2d72f222131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b7120a08d0630a696c0b3f78dd9ec083a2131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b71"}, &result)
	assert.NoError(t, err)
}

func TestChain33_CreateNoBalanceTxs(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	chain33 := newTestChain33(api)
	api.On("GetProperFee", mock.Anything).Return(&types.ReplyProperFee{ProperFee: 1000000}, nil)
	var result string
	err := chain33.CreateNoBlanaceTxs(&types.NoBalanceTxs{TxHexs: []string{"0a05746f6b656e12413804223d0a0443434e5910a09c011a0d74657374207472616e73666572222231333559774e715367694551787577586650626d526d48325935334564673864343820a08d0630969a9fe6c4b9c7ba5d3a2231333559774e715367694551787577586650626d526d483259353345646738643438", "0a05746f6b656e12413804223d0a0443434e5910b0ea011a0d74657374207472616e73666572222231333559774e715367694551787577586650626d526d48325935334564673864343820a08d0630bca0a2dbc0f182e06f3a2231333559774e715367694551787577586650626d526d483259353345646738643438"}}, &result)
	assert.NoError(t, err)
}

func TestChain33_ExecWallet(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	in := &rpctypes.ChainExecutor{}
	api.On("ExecWallet", mock.Anything).Return(nil, nil)
	err := client.ExecWallet(in, &testResult)
	assert.NotNil(t, err)
}

func TestChain33_Query(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	in := rpctypes.Query4Jrpc{Execer: "coins"}
	api.On("Query", mock.Anything).Return(nil, nil)
	err := client.Query(in, &testResult)
	assert.NotNil(t, err)
}

func TestChain33_DumpPrivkey(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	api.On("ExecWalletFunc", "wallet", "DumpPrivkey", mock.Anything).Return(nil, nil)
	err := client.DumpPrivkey(&types.ReqString{}, &testResult)
	assert.NoError(t, err)
}

func TestChain33_DumpPrivkeysFile(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	api.On("ExecWalletFunc", "wallet", "DumpPrivkeysFile", mock.Anything).Return(&types.Reply{}, nil)
	err := client.DumpPrivkeysFile(&types.ReqPrivkeysFile{}, &testResult)
	assert.NoError(t, err)
}

func TestChain33_ImportPrivkeysFile(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	api.On("ExecWalletFunc", "wallet", "ImportPrivkeysFile", mock.Anything).Return(&types.Reply{}, nil)
	err := client.ImportPrivkeysFile(&types.ReqPrivkeysFile{}, &testResult)
	assert.NoError(t, err)
}

func TestChain33_GetTotalCoins(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	api.On("StoreGetTotalCoins", mock.Anything).Return(nil, nil)
	err := client.GetTotalCoins(&types.ReqGetTotalCoins{}, &testResult)
	assert.NoError(t, err)
}

func TestChain33_GetFatalFailure(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}

	expected := types.ReqNil{}
	api.On("ExecWalletFunc", "wallet", "FatalFailure", &expected).Return(&types.Int32{}, nil)
	err := client.GetFatalFailure(&expected, &testResult)
	assert.NoError(t, err)
}

func TestChain33_DecodeRawTransaction(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	err := client.DecodeRawTransaction(&types.ReqDecodeRawTransaction{TxHex: "0a05636f696e73122c18010a281080c2d72f222131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b7120a08d0630a696c0b3f78dd9ec083a2131477444795771577233553637656a7663776d333867396e7a6e7a434b58434b71"}, &testResult)
	assert.NoError(t, err)
}

func TestChain33_CloseQueue(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	api.On("CloseQueue", mock.Anything).Return(nil, nil)
	err := client.CloseQueue(nil, &testResult)
	assert.True(t, testResult.(*types.Reply).IsOk)
	assert.NoError(t, err)
}

func TestChain33_AddSeqCallBack(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	api.On("AddPushSubscribe", mock.Anything).Return(&types.ReplySubscribePush{}, nil)
	err := client.AddPushSubscribe(nil, &testResult)
	assert.NoError(t, err)
}

func TestChain33_ListSeqCallBack(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	api.On("ListPushes", mock.Anything).Return(&types.PushSubscribes{}, nil)
	err := client.ListPushes(nil, &testResult)
	assert.NoError(t, err)
}

func TestChain33_GetSeqCallBackLastNum(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult interface{}
	api.On("GetPushSeqLastNum", mock.Anything).Return(&types.Int64{}, nil)
	err := client.GetPushSeqLastNum(nil, &testResult)
	assert.NoError(t, err)
}

func TestChain33_ConvertExectoAddr(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var testResult string
	err := client.ConvertExectoAddr(rpctypes.ExecNameParm{ExecName: "coins"}, &testResult)
	assert.NoError(t, err)
}

func Test_fmtTxDetail(t *testing.T) {

	tx := &types.Transaction{Execer: []byte("coins")}
	log := &types.ReceiptLog{Ty: 0, Log: []byte("test")}
	receipt := &types.ReceiptData{Ty: 0, Logs: []*types.ReceiptLog{log}}
	detail := &types.TransactionDetail{Tx: tx, Receipt: receipt}
	var err error
	//test withdraw swap from to
	detail.Fromaddr = "from"
	detail.Tx.Payload, err = common.FromHex("0x180322301080c2d72f2205636f696e732a22314761485970576d71414a7371527772706f4e6342385676674b7453776a63487174")
	assert.NoError(t, err)
	tx.To = "to"
	tran, err := fmtTxDetail(detail, false, "coins", types.DefaultCoinPrecision)
	assert.NoError(t, err)
	assert.Equal(t, "to", tran.Fromaddr)
	assert.Equal(t, "from", tx.To)
}

func queryTotalFee(client *Chain33, req *types.LocalDBGet, t *testing.T) int64 {
	var testResult interface{}
	err := client.QueryTotalFee(req, &testResult)
	assert.NoError(t, err)
	fee, _ := testResult.(*types.TotalFee)
	return fee.Fee
}

func TestChain33_QueryTotalFee(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)

	total := &types.TotalFee{TxCount: 1, Fee: 10000}
	api.On("LocalGet", mock.Anything).Return(&types.LocalReplyValue{Values: [][]byte{types.Encode(total)}}, nil)
	req := &types.LocalDBGet{Keys: [][]byte{types.TotalFeeKey([]byte("testHash"))}}
	req1 := &types.LocalDBGet{Keys: [][]byte{[]byte("testHash")}}

	assert.Equal(t, total.Fee, queryTotalFee(client, req, t))
	assert.Equal(t, total.Fee, queryTotalFee(client, req1, t))
	assert.True(t, bytes.Equal(req.Keys[0], req1.Keys[0]))
}

func TestChain33_GetSequenceByHash(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	api.On("GetSequenceByHash", mock.Anything).Return(&types.Int64{}, nil)
	var testResult interface{}
	err := client.GetSequenceByHash(rpctypes.ReqHashes{Hashes: []string{"testhash"}}, &testResult)
	assert.Error(t, err)
	hash := "0x06a9f4ae07dd8a9b5f7f01ed23084880209d5ddff7195a8515ce43da218e8aa7"
	err = client.GetSequenceByHash(rpctypes.ReqHashes{Hashes: []string{hash}}, &testResult)
	assert.Nil(t, err)

}

func TestChain33_GetBlockBySeq(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	api.On("GetBlockBySeq", mock.Anything).Return(&types.BlockSeq{Num: 12345, Seq: &types.BlockSequence{Type: 1, Hash: []byte("1111")}}, nil)
	var testResult interface{}
	err := client.GetBlockBySeq(&types.Int64{Data: 11}, &testResult)
	assert.Nil(t, err)

}

func TestChain33_GetParaTxByTitle(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	api.On("GetParaTxByTitle", mock.Anything).Return(&types.ParaTxDetails{Items: []*types.ParaTxDetail{}}, nil)
	var testResult interface{}
	err := client.GetParaTxByTitle(&types.ReqParaTxByTitle{Start: 11, End: 11, Title: "2323"}, &testResult)
	assert.Nil(t, err)
}

func TestChain33_LoadParaTxByTitle(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	api.On("LoadParaTxByTitle", mock.Anything).Return(&types.ReplyHeightByTitle{}, nil)
	var testResult interface{}
	err := client.LoadParaTxByTitle(&types.ReqHeightByTitle{}, &testResult)
	assert.Nil(t, err)
}

func TestChain33_GetParaTxByHeight(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	api.On("GetParaTxByHeight", mock.Anything).Return(&types.ParaTxDetails{}, nil)
	var testResult interface{}
	err := client.GetParaTxByHeight(&types.ReqParaTxByHeight{}, &testResult)
	assert.Nil(t, err)
}

func TestChain33_QueryChain(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	api.On("QueryChain", mock.Anything).Return(nil, types.ErrInvalidParam)
	var testResult interface{}
	err := client.QueryChain(rpctypes.ChainExecutor{}, &testResult)
	assert.NotNil(t, err)
}

func TestChain33_convertParaTxDetails(t *testing.T) {

	var details types.ParaTxDetails
	var detail types.ParaTxDetail
	details.Items = append(details.Items, &detail)
	detail.Type = 123
	txhash := "0x7feb86911f2143b992c5d543cc7314f24c3f94535f1beb38f781f2a0d72ae918"
	hashBs, err := common.FromHex(txhash)
	assert.Nil(t, err)
	detail.Header = &types.Header{Height: 555, BlockTime: 39169, TxHash: hashBs}
	var rmsg rpctypes.ParaTxDetails
	convertParaTxDetails(&details, &rmsg, types.DefaultCoinPrecision)
	assert.Equal(t, 555, int(rmsg.Items[0].Header.Height))
	assert.Equal(t, 39169, int(rmsg.Items[0].Header.BlockTime))
	assert.Equal(t, txhash, rmsg.Items[0].Header.TxHash)
}

func TestChain33_convertHeader(t *testing.T) {
	var header types.Header
	var reheader rpctypes.Header
	header.TxHash, _ = hex.DecodeString("7feb86911f2143b992c5d543cc7314f24c3f94535f1beb38f781f2a0d72ae918")
	header.Height = 666
	header.BlockTime = 1234567
	header.Signature = nil
	header.TxCount = 9
	convertHeader(&header, &reheader)
	assert.Equal(t, "0x7feb86911f2143b992c5d543cc7314f24c3f94535f1beb38f781f2a0d72ae918", reheader.TxHash)
	assert.Equal(t, header.GetTxCount(), reheader.TxCount)
	assert.Equal(t, header.GetBlockTime(), reheader.BlockTime)
	assert.Equal(t, header.GetHeight(), reheader.Height)
}

func TestChain33_ChainID(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	testChain33 := newTestChain33(api)
	var testResult interface{}

	in := &types.ReqNil{}
	chainID := int32(33)
	api.On("GetChainID", mock.Anything).Return(chainID, nil)
	err := testChain33.GetChainID(in, &testResult)
	assert.Equal(t, nil, err)
	assert.Equal(t, int32(33), testResult.(*rpctypes.ChainIDInfo).ChainID)
}

func TestChain33_GetCryptoList(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var result interface{}
	api.On("GetCryptoList").Return(nil)
	err := client.GetCryptoList(&types.ReqNil{}, &result)
	assert.Nil(t, err)
}

func TestChain33_SendDelayTransaction(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var result interface{}
	testData := []byte("testTxHash")
	api.On("SendDelayTx", mock.Anything, true).Return(&types.Reply{Msg: testData}, nil)
	err := client.SendDelayTransaction(&types.ReqString{Data: "1234"}, &result)
	require.NotNil(t, err)
	err = client.SendDelayTransaction(
		&types.ReqString{Data: common.ToHex(testData)}, &result)
	require.NotNil(t, err)
	err = client.SendDelayTransaction(
		&types.ReqString{Data: common.ToHex(types.Encode(&types.DelayTx{}))}, &result)
	require.Nil(t, err)
	require.Equal(t, common.ToHex(testData), result.(string))
}

func TestChain33_WalletRecoverScript(t *testing.T) {

	chain33 := &Chain33{cli: channelClient{}}
	var result interface{}
	err := chain33.GetWalletRecoverAddress(nil, &result)
	require.Equal(t, types.ErrInvalidParam, err)
	err = chain33.SignWalletRecoverTx(nil, &result)
	require.Equal(t, types.ErrInvalidParam, err)
}

func TestChain33_GetChainConfig(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	client := newTestChain33(api)
	var result interface{}
	api.On("GetChainConfig").Return(nil)
	err := client.GetChainConfig(&types.ReqNil{}, &result)
	assert.Nil(t, err)
}

func TestChain33_AddBlacklist(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)

	testChain33 := newTestChain33(api)
	var peer types.BlackPeer
	peer.PeerAddr = "192.168.0.31:13802"
	peer.Lifetime = "1min"
	peer.PeerName = "16Uiu2HAm2HS2VPChZsfAk3yKwfKncpCFAW9WT2L4UQQGtTtUEk8k"

	expected := &types.Reply{IsOk: true, Msg: []byte("success")}
	api.On("AddBlacklist", mock.Anything).Return(expected, nil)

	var testResult interface{}
	err := testChain33.AddBlacklist(&peer, &testResult)
	assert.Nil(t, err)
	reply, ok := testResult.(*rpctypes.Reply)
	assert.True(t, ok)
	assert.Equal(t, reply.Msg, "success")

	peer.PeerAddr = "192.168.0.312222"
	peer.Lifetime = "1hour"
	expected.IsOk = false
	expected.Msg = []byte("ip format err")
	err = testChain33.AddBlacklist(&peer, &testResult)
	assert.Nil(t, err)
	reply, ok = testResult.(*rpctypes.Reply)
	assert.True(t, ok)
	assert.Equal(t, reply.Msg, "ip format err")
}

func TestChain33_DelBlacklist(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)

	testChain33 := newTestChain33(api)
	var peer types.BlackPeer
	peer.PeerName = "1222234serwer23423e2334123421"
	peer.Lifetime = "1min"

	expected := &types.Reply{IsOk: false, Msg: []byte("no this peerName")}
	api.On("DelBlacklist", mock.Anything).Return(expected, nil)

	var testResult interface{}
	err := testChain33.DelBlacklist(&peer, &testResult)
	assert.Nil(t, err)
	reply, ok := testResult.(*rpctypes.Reply)
	assert.True(t, ok)
	assert.Equal(t, reply.Msg, "no this peerName")

}

func TestChain33_ShowBlacklist(t *testing.T) {
	cfg := types.NewChain33Config(types.GetDefaultCfgstring())
	api := new(mocks.QueueProtocolAPI)
	api.On("GetConfig", mock.Anything).Return(cfg)
	expected := &types.Blacklist{}
	api.On("ShowBlacklist", mock.Anything).Return(expected, nil)
	testChain33 := newTestChain33(api)
	var testResult interface{}
	err := testChain33.ShowBlacklist(&types.ReqNil{}, &testResult)
	assert.Nil(t, err)
	_, ok := testResult.([]*types.BlackInfo)
	assert.True(t, ok)

}
