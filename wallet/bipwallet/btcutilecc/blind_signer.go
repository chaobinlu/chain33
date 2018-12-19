// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package btcutil

import "crypto/ecdsa"
import "crypto/rand"
import "fmt"
import "math/big"

<<<<<<< HEAD
=======
// BlindSignerState blind signer state
>>>>>>> upstream/master
type BlindSignerState struct {
	// secret stuff
	d, k *big.Int

	// shareable stuff
	Q *ecdsa.PublicKey
}

<<<<<<< HEAD
// Request that the signer start a blind signature protocol.  Returns
=======
// BlindSession Request that the signer start a blind signature protocol.  Returns
>>>>>>> upstream/master
// the signer's public key and an EC point named R.
func BlindSession(sState *BlindSignerState) (*ecdsa.PublicKey, *ecdsa.PublicKey) {

	// generate signer's private & public key pair
	if sState.Q == nil {
		keys, err := GenerateKey(rand.Reader)
		maybePanic(err)
		sState.d = keys.D
		sState.Q = &keys.PublicKey
		fmt.Printf("Signer:\t%x\n\t%x\n", sState.d, sState.Q.X)
	}

	// generate k and R for each user request (§4.2)
	request, err := GenerateKey(rand.Reader)
	maybePanic(err)
	sState.k = request.D
	R := &request.PublicKey

	return sState.Q, R
}

<<<<<<< HEAD
// Signs a blinded message
=======
// BlindSign Signs a blinded message
>>>>>>> upstream/master
func BlindSign(sState *BlindSignerState, R *ecdsa.PublicKey, mHat *big.Int) *big.Int {
	crv := Secp256k1().Params()

	// verify that R matches our secret k
<<<<<<< HEAD
	R_ := ScalarBaseMult(sState.k)
	if !KeysEqual(R, R_) {
=======
	RSM := ScalarBaseMult(sState.k)
	if !KeysEqual(R, RSM) {
>>>>>>> upstream/master
		panic("unknown R")
	}

	// signer generates signature (§4.3)
	sHat := new(big.Int).Mul(sState.d, mHat)
	sHat.Add(sHat, sState.k)
	sHat.Mod(sHat, crv.N)

	return sHat
}
