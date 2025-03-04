// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client_test

import (
	"context"
	"encoding/json"

	"github.com/33cn/chain33/rpc/jsonclient"
	"github.com/33cn/chain33/types"

	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// TODO: SetPostRunCb()
type JSONRPCCtx struct {
	Addr   string
	Method string
	Params interface{}
	Res    interface{}

	cb Callback

	jsonClient *jsonclient.JSONClient
}

type Callback func(res interface{}) (interface{}, error)

func NewJSONRPCCtx(methed string, params, res interface{}) *JSONRPCCtx {
	return &JSONRPCCtx{
		Addr:   jrpcsite,
		Method: methed,
		Params: params,
		Res:    res,
	}
}

func (c *JSONRPCCtx) SetResultCb(cb Callback) {
	c.cb = cb
}

func (c *JSONRPCCtx) Run() (err error) {
	if c.jsonClient == nil {
		c.jsonClient, err = jsonclient.NewJSONClient(c.Addr)
		if err != nil {
			return
		}
	}
	err = c.jsonClient.Call(c.Method, c.Params, c.Res)
	if err != nil {
		return
	}
	// maybe format rpc result
	var result interface{}
	if c.cb != nil {
		result, err = c.cb(c.Res)
		if err != nil {
			return
		}
	} else {
		result = c.Res
	}

	_, err = json.MarshalIndent(result, "", "    ")
	if err != nil {
		return
	}
	return
}

type GrpcCtx struct {
	Method string
	Params interface{}
	Res    interface{}
}

func NewGRpcCtx(method string, params, res interface{}) *GrpcCtx {
	return &GrpcCtx{
		Method: method,
		Params: params,
		Res:    res,
	}
}

func (c *GrpcCtx) Run() (err error) {
	conn, errRet := grpc.Dial(grpcsite, grpc.WithInsecure())
	if errRet != nil {
		return errRet
	}
	defer conn.Close()

	rpc := types.NewChain33Client(conn)
	switch c.Method {
	case "GetBlocks":
		reply, err := rpc.GetBlocks(context.Background(), c.Params.(*types.ReqBlocks))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetLastHeader":
		reply, err := rpc.GetLastHeader(context.Background(), c.Params.(*types.ReqNil))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "CreateRawTransaction":
		reply, err := rpc.CreateRawTransaction(context.Background(), c.Params.(*types.CreateTx))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "QueryTransaction":
		reply, err := rpc.QueryTransaction(context.Background(), c.Params.(*types.ReqHash))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "SendTransaction":
		reply, err := rpc.SendTransaction(context.Background(), c.Params.(*types.Transaction))
		if err == nil {
			c.Res = reply
		}
	case "GetTransactionByAddr":
		reply, err := rpc.GetTransactionByAddr(context.Background(), c.Params.(*types.ReqAddr))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetTransactionByHashes":
		reply, err := rpc.GetTransactionByHashes(context.Background(), c.Params.(*types.ReqHashes))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetMemPool":
		reply, err := rpc.GetMemPool(context.Background(), c.Params.(*types.ReqGetMempool))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetAccounts":
		reply, err := rpc.GetAccounts(context.Background(), c.Params.(*types.ReqNil))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "NewAccount":
		reply, err := rpc.NewAccount(context.Background(), c.Params.(*types.ReqNewAccount))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "WalletTransactionList":
		reply, err := rpc.WalletTransactionList(context.Background(), c.Params.(*types.ReqWalletTransactionList))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "ImportPrivkey":
		reply, err := rpc.ImportPrivkey(context.Background(), c.Params.(*types.ReqWalletImportPrivkey))
		if err == nil {
			c.Res = reply
		}
	case "SendToAddress":
		reply, err := rpc.SendToAddress(context.Background(), c.Params.(*types.ReqWalletSendToAddress))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "SetTxFee":
		reply, err := rpc.SetTxFee(context.Background(), c.Params.(*types.ReqWalletSetFee))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "SetLabl":
		reply, err := rpc.SetLabl(context.Background(), c.Params.(*types.ReqWalletSetLabel))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "MergeBalance":
		reply, err := rpc.MergeBalance(context.Background(), c.Params.(*types.ReqWalletMergeBalance))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "SetPasswd":
		reply, err := rpc.SetPasswd(context.Background(), c.Params.(*types.ReqWalletSetPasswd))
		if err == nil {
			c.Res = reply
		}
	case "Lock":
		reply, err := rpc.Lock(context.Background(), c.Params.(*types.ReqNil))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "UnLock":
		reply, err := rpc.UnLock(context.Background(), c.Params.(*types.WalletUnLock))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetPeerInfo":
		reply, err := rpc.GetPeerInfo(context.Background(), c.Params.(*types.P2PGetPeerReq))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetLastMemPool":
		reply, err := rpc.GetLastMemPool(context.Background(), c.Params.(*types.ReqNil))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetProperFee":
		reply, err := rpc.GetProperFee(context.Background(), c.Params.(*types.ReqProperFee))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetWalletStatus":
		reply, err := rpc.GetWalletStatus(context.Background(), c.Params.(*types.ReqNil))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetBlockOverview":
		reply, err := rpc.GetBlockOverview(context.Background(), c.Params.(*types.ReqHash))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetAddrOverview":
		reply, err := rpc.GetAddrOverview(context.Background(), c.Params.(*types.ReqAddr))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetBlockHash":
		reply, err := rpc.GetBlockHash(context.Background(), c.Params.(*types.ReqInt))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GenSeed":
		reply, err := rpc.GenSeed(context.Background(), c.Params.(*types.GenSeedLang))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetSeed":
		reply, err := rpc.GetSeed(context.Background(), c.Params.(*types.GetSeedByPw))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "SaveSeed":
		reply, err := rpc.SaveSeed(context.Background(), c.Params.(*types.SaveSeedByPw))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetBalance":
		reply, err := rpc.GetBalance(context.Background(), c.Params.(*types.ReqBalance))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "QueryChain":
		reply, err := rpc.QueryChain(context.Background(), c.Params.(*types.ChainExecutor))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetHexTxByHash":
		reply, err := rpc.GetHexTxByHash(context.Background(), c.Params.(*types.ReqHash))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "DumpPrivkey":
		reply, err := rpc.DumpPrivkey(context.Background(), c.Params.(*types.ReqString))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "DumpPrivkeysFile":
		reply, err := rpc.DumpPrivkeysFile(context.Background(), c.Params.(*types.ReqPrivkeysFile))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "ImportPrivkeysFile":
		reply, err := rpc.ImportPrivkeysFile(context.Background(), c.Params.(*types.ReqPrivkeysFile))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "Version":
		reply, err := rpc.Version(context.Background(), c.Params.(*types.ReqNil))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "IsSync":
		reply, err := rpc.IsSync(context.Background(), c.Params.(*types.ReqNil))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "IsNtpClockSync":
		reply, err := rpc.IsNtpClockSync(context.Background(), c.Params.(*types.ReqNil))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "NetInfo":
		reply, err := rpc.NetInfo(context.Background(), c.Params.(*types.P2PGetNetInfoReq))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetSequenceByHash":
		reply, err := rpc.GetSequenceByHash(context.Background(), c.Params.(*types.ReqHash))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetBlockBySeq":
		reply, err := rpc.GetBlockBySeq(context.Background(), c.Params.(*types.Int64))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetParaTxByTitle":
		reply, err := rpc.GetParaTxByTitle(context.Background(), c.Params.(*types.ReqParaTxByTitle))
		if err == nil {
			c.Res = reply
		}
		errRet = err

	case "LoadParaTxByTitle":
		reply, err := rpc.LoadParaTxByTitle(context.Background(), c.Params.(*types.ReqHeightByTitle))
		if err == nil {
			c.Res = reply
		}
		errRet = err
	case "GetParaTxByHeight":
		reply, err := rpc.GetParaTxByHeight(context.Background(), c.Params.(*types.ReqParaTxByHeight))
		if err == nil {
			c.Res = reply
		}
		errRet = err

	default:
		errRet = errors.New(fmt.Sprintf("Unsupport method %v", c.Method))
	}
	return errRet
}
