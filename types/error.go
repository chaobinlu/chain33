// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"errors"
)

<<<<<<< HEAD
var (
=======
// chain33定义的错误类型
var (
	ErrTooManySeqCB            = errors.New("ErrTooManySeqCB")
	ErrPushSeqPostData         = errors.New("ErrPushSeqPostData")
>>>>>>> upstream/master
	ErrMethodReturnType        = errors.New("ErrMethodReturnType")
	ErrMethodNotFound          = errors.New("ErrMethodNotFound")
	ErrExecBlockNil            = errors.New("ErrExecBlockNil")
	ErrNotAllow                = errors.New("ErrNotAllow")
	ErrCanOnlyDelTopVersion    = errors.New("ErrCanOnlyDelTopVersion")
	ErrPrevVersion             = errors.New("ErrPrevVersion")
	ErrNoExecerInMavlKey       = errors.New("ErrNoExecerInMavlKey")
	ErrMavlKeyNotStartWithMavl = errors.New("ErrMavlKeyNotStartWithMavl")
	ErrNotFound                = errors.New("ErrNotFound")
	ErrBlockExec               = errors.New("ErrBlockExec")
	ErrCheckStateHash          = errors.New("ErrCheckStateHash")
	ErrCheckTxHash             = errors.New("ErrCheckTxHash")
	ErrReRunGenesis            = errors.New("ErrReRunGenesis")
	ErrActionNotSupport        = errors.New("ErrActionNotSupport")
	ErrQueryNotSupport         = errors.New("ErrQueryNotSupport")
	ErrChannelFull             = errors.New("ErrChannelFull")
	ErrAmount                  = errors.New("ErrAmount")
	ErrMinerIsStared           = errors.New("ErrMinerIsStared")
	ErrMinerNotStared          = errors.New("ErrMinerNotStared")
	ErrMinerNotClosed          = errors.New("ErrMinerNotClosed")
	ErrNoPeer                  = errors.New("ErrNoPeer")
	ErrExecNameNotMatch        = errors.New("ErrExecNameNotMatch")
	ErrChannelClosed           = errors.New("ErrChannelClosed")
	ErrNotMinered              = errors.New("ErrNotMinered")
	ErrFromAddr                = errors.New("ErrFromAddr")
	ErrBlockHeight             = errors.New("ErrBlockHeight")
	ErrBlockTime               = errors.New("ErrBlockTime")
	ErrCoinBaseExecer          = errors.New("ErrCoinBaseExecer")
	ErrCoinBaseTxType          = errors.New("ErrCoinBaseTxType")
	ErrCoinBaseExecErr         = errors.New("ErrCoinBaseExecErr")
	ErrCoinBaseTarget          = errors.New("ErrCoinBaseTarget")
	ErrCoinbaseReward          = errors.New("ErrCoinbaseReward")
	ErrNotAllowDeposit         = errors.New("ErrNotAllowDeposit")
	ErrCoinBaseIndex           = errors.New("ErrCoinBaseIndex")
	ErrCoinBaseTicketStatus    = errors.New("ErrCoinBaseTicketStatus")
	ErrBlockNotFound           = errors.New("ErrBlockNotFound")
	ErrLogType                 = errors.New("ErrLogType")
	ErrInvalidParam            = errors.New("ErrInvalidParam")
	ErrInvalidAddress          = errors.New("ErrInvalidAddress")
<<<<<<< HEAD
=======
	ErrNotInited               = errors.New("ErrNotInited")
>>>>>>> upstream/master

	ErrStartBigThanEnd            = errors.New("ErrStartBigThanEnd")
	ErrToAddrNotSameToExecAddr    = errors.New("ErrToAddrNotSameToExecAddr")
	ErrTypeAsset                  = errors.New("ErrTypeAsset")
	ErrEmpty                      = errors.New("ErrEmpty")
	ErrSendSameToRecv             = errors.New("ErrSendSameToRecv")
	ErrExecNameNotAllow           = errors.New("ErrExecNameNotAllow")
<<<<<<< HEAD
=======
	ErrExecNotFound               = errors.New("ErrExecNotFound")
>>>>>>> upstream/master
	ErrLocalDBPerfix              = errors.New("ErrLocalDBPerfix")
	ErrTimeout                    = errors.New("ErrTimeout")
	ErrBlockHeaderDifficulty      = errors.New("ErrBlockHeaderDifficulty")
	ErrNoTx                       = errors.New("ErrNoTx")
	ErrTxExist                    = errors.New("ErrTxExist")
	ErrManyTx                     = errors.New("ErrManyTx")
	ErrDupTx                      = errors.New("ErrDupTx")
	ErrMemFull                    = errors.New("ErrMemFull")
	ErrNoBalance                  = errors.New("ErrNoBalance")
	ErrBalanceLessThanTenTimesFee = errors.New("ErrBalanceLessThanTenTimesFee")
	ErrTxExpire                   = errors.New("ErrTxExpire")
	ErrHeaderNotSet               = errors.New("ErrHeaderNotSet")
	ErrSign                       = errors.New("ErrSign")
	ErrFeeTooLow                  = errors.New("ErrFeeTooLow")
	ErrEmptyTx                    = errors.New("ErrEmptyTx")
	ErrTxFeeTooLow                = errors.New("ErrTxFeeTooLow")
	ErrTxMsgSizeTooBig            = errors.New("ErrTxMsgSizeTooBig")
	ErrFutureBlock                = errors.New("ErrFutureBlock")
	ErrHashNotFound               = errors.New("ErrHashNotFound")
	ErrTxDup                      = errors.New("ErrTxDup")
	ErrNotSync                    = errors.New("ErrNotSync")
	ErrSize                       = errors.New("ErrSize")

<<<<<<< HEAD
	// BlockChain Error Types
=======
	// ErrHashNotExist BlockChain Error Types
>>>>>>> upstream/master
	ErrHashNotExist           = errors.New("ErrHashNotExist")
	ErrHeightNotExist         = errors.New("ErrHeightNotExist")
	ErrTxNotExist             = errors.New("ErrTxNotExist")
	ErrAddrNotExist           = errors.New("ErrAddrNotExist")
	ErrStartHeight            = errors.New("ErrStartHeight")
	ErrEndLessThanStartHeight = errors.New("ErrEndLessThanStartHeight")
	ErrClientNotBindQueue     = errors.New("ErrClientNotBindQueue")
	ErrContinueBack           = errors.New("ErrContinueBack")
	ErrUnmarshal              = errors.New("ErrUnmarshal")
	ErrMarshal                = errors.New("ErrMarshal")
	ErrBlockExist             = errors.New("ErrBlockExist")
	ErrParentBlockNoExist     = errors.New("ErrParentBlockNoExist")
	ErrBlockHeightNoMatch     = errors.New("ErrBlockHeightNoEqual")
	ErrParentTdNoExist        = errors.New("ErrParentTdNoExist")
	ErrBlockHashNoMatch       = errors.New("ErrBlockHashNoMatch")
	ErrIsClosed               = errors.New("ErrIsClosed")
	ErrDecode                 = errors.New("ErrDecode")
	ErrNotRollBack            = errors.New("ErrNotRollBack")
	ErrPeerInfoIsNil          = errors.New("ErrPeerInfoIsNil")
<<<<<<< HEAD
	//wallet
=======
	//ErrWalletIsLocked wallet
>>>>>>> upstream/master
	ErrWalletIsLocked       = errors.New("ErrWalletIsLocked")
	ErrSaveSeedFirst        = errors.New("ErrSaveSeedFirst")
	ErrUnLockFirst          = errors.New("ErrUnLockFirst")
	ErrLabelHasUsed         = errors.New("ErrLabelHasUsed")
	ErrPrivkeyExist         = errors.New("ErrPrivkeyExist")
	ErrPrivkey              = errors.New("ErrPrivkey")
	ErrInsufficientBalance  = errors.New("ErrInsufficientBalance")
	ErrInsufficientTokenBal = errors.New("ErrInsufficientTokenBalance")
	ErrInsuffSellOrder      = errors.New("ErrInsufficientSellOrder2buy")
	ErrVerifyOldpasswdFail  = errors.New("ErrVerifyOldpasswdFail")
	ErrInputPassword        = errors.New("ErrInputPassword")
	ErrSeedlang             = errors.New("ErrSeedlang")
	ErrSeedNotExist         = errors.New("ErrSeedNotExist")
	ErrSubPubKeyVerifyFail  = errors.New("ErrSubPubKeyVerifyFail")
	ErrLabelNotExist        = errors.New("ErrLabelNotExist")
	ErrAccountNotExist      = errors.New("ErrAccountNotExist")
	ErrSeedExist            = errors.New("ErrSeedExist")
	ErrNotSupport           = errors.New("ErrNotSupport")
	ErrSeedWordNum          = errors.New("ErrSeedWordNum")
	ErrPubKeyLen            = errors.New("ErrPublicKeyLen")
	ErrPrivateKeyLen        = errors.New("ErrPrivateKeyLen")
	ErrSeedWord             = errors.New("ErrSeedWord")
	ErrNoPrivKeyOrAddr      = errors.New("ErrNoPrivKeyOrAddr")
	ErrNewWalletFromSeed    = errors.New("ErrNewWalletFromSeed")
	ErrNewKeyPair           = errors.New("ErrNewKeyPair")
	ErrPrivkeyToPub         = errors.New("ErrPrivkeyToPub")

	ErrOnlyTicketUnLocked = errors.New("ErrOnlyTicketUnLocked")
	ErrNewCrypto          = errors.New("ErrNewCrypto")
	ErrFromHex            = errors.New("ErrFromHex")
	ErrPrivKeyFromBytes   = errors.New("ErrFromHex")
	ErrParentHash         = errors.New("ErrParentHash")

<<<<<<< HEAD
	//p2p
=======
	//ErrPing p2p模块错误类型
>>>>>>> upstream/master
	ErrPing       = errors.New("ErrPingSignature")
	ErrVersion    = errors.New("ErrVersionNoSupport")
	ErrStreamPing = errors.New("ErrStreamPing")
	ErrPeerStop   = errors.New("ErrPeerStop")

	ErrBlockSize                  = errors.New("ErrBlockSize")
	ErrTxGroupIndex               = errors.New("ErrTxGroupIndex")
	ErrTxGroupFormat              = errors.New("ErrTxGroupFormat")
	ErrTxGroupCountLessThanTwo    = errors.New("ErrTxGroupCountLessThanTwo")
	ErrTxGroupHeader              = errors.New("ErrTxGroupHeader")
	ErrTxGroupNext                = errors.New("ErrTxGroupNext")
	ErrTxGroupCountBigThanMaxSize = errors.New("ErrTxGroupCountBigThanMaxSize")
	ErrTxGroupEmpty               = errors.New("ErrTxGroupEmpty")
	ErrTxGroupCount               = errors.New("ErrTxGroupCount")
	ErrTxGroupFeeNotZero          = errors.New("ErrTxGroupFeeNotZero")
	ErrNomalTx                    = errors.New("ErrNomalTx")
	ErrUnknowDriver               = errors.New("ErrUnknowDriver")
	ErrUnRegistedDriver           = errors.New("ErrUnRegistedDriver")
	ErrSymbolNameNotAllow         = errors.New("ErrSymbolNameNotAllow")
	ErrTxGroupNotSupport          = errors.New("ErrTxGroupNotSupport")
	ErrNotAllowKey                = errors.New("ErrNotAllowKey")
	ErrNotAllowMemSetKey          = errors.New("ErrNotAllowMemSetKey")
	ErrDataBaseDamage             = errors.New("ErrDataBaseDamage")
	ErrIndex                      = errors.New("ErrIndex")
	ErrTxGroupParaCount           = errors.New("ErrTxGroupParaCount")

<<<<<<< HEAD
	//rpc
	ErrInvalidMainnetRpcAddr = errors.New("ErrInvalidMainnetRpcAddr")
=======
	//ErrInvalidMainnetRPCAddr rpc模块的错误类型
	ErrInvalidMainnetRPCAddr = errors.New("ErrInvalidMainnetRPCAddr")
>>>>>>> upstream/master

	ErrDBFlag      = errors.New("ErrDBFlag")
	ErrLocalPrefix = errors.New("ErrLocalPrefix")
	ErrLocalKeyLen = errors.New("ErrLocalKeyLen")

	ErrCloneForkFrom      = errors.New("ErrCloneForkFrom")
	ErrCloneForkToExist   = errors.New("ErrCloneForkToExist")
	ErrQueryThistIsNotSet = errors.New("ErrQueryThistIsNotSet")
)
