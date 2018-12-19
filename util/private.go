// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/common/crypto"
	"github.com/33cn/chain33/types"
)

<<<<<<< HEAD
=======
//TestPrivkeyHex ：
>>>>>>> upstream/master
var TestPrivkeyHex = []string{
	"4257D8692EF7FE13C68B65D6A52F03933DB2FA5CE8FAF210B5B8B80C721CED01",
	"CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944",
	"B0BB75BC49A787A71F4834DA18614763B53A18291ECE6B5EDEC3AD19D150C3E7",
	"56942AD84CCF4788ED6DACBC005A1D0C4F91B63BCF0C99A02BE03C8DEAE71138",
	"2AFF1981291355322C7A6308D46A9C9BA311AA21D94F36B43FC6A6021A1334CF",
	"2116459C0EC8ED01AA0EEAE35CAC5C96F94473F7816F114873291217303F6989",
}

<<<<<<< HEAD
=======
//TestPrivkeyList ：
>>>>>>> upstream/master
var TestPrivkeyList = []crypto.PrivKey{
	HexToPrivkey("4257D8692EF7FE13C68B65D6A52F03933DB2FA5CE8FAF210B5B8B80C721CED01"),
	HexToPrivkey("CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944"),
	HexToPrivkey("B0BB75BC49A787A71F4834DA18614763B53A18291ECE6B5EDEC3AD19D150C3E7"),
	HexToPrivkey("56942AD84CCF4788ED6DACBC005A1D0C4F91B63BCF0C99A02BE03C8DEAE71138"),
	HexToPrivkey("2AFF1981291355322C7A6308D46A9C9BA311AA21D94F36B43FC6A6021A1334CF"),
	HexToPrivkey("2116459C0EC8ED01AA0EEAE35CAC5C96F94473F7816F114873291217303F6989"),
}

<<<<<<< HEAD
=======
//HexToPrivkey ： convert hex string to private key
>>>>>>> upstream/master
func HexToPrivkey(key string) crypto.PrivKey {
	cr, err := crypto.New(types.GetSignName("", types.SECP256K1))
	if err != nil {
		panic(err)
	}
	bkey, err := common.FromHex(key)
	if err != nil {
		panic(err)
	}
	priv, err := cr.PrivKeyFromBytes(bkey)
	if err != nil {
		panic(err)
	}
	return priv
}
