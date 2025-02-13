//go:build cgo
// +build cgo

package ffi

// #cgo linux LDFLAGS: ${SRCDIR}/libfilcrypto.a -Wl,-unresolved-symbols=ignore-all
// #cgo darwin LDFLAGS: ${SRCDIR}/libfilcrypto.a -Wl,-undefined,dynamic_lookup
// #cgo pkg-config: ${SRCDIR}/filcrypto.pc
// #include "./filcrypto.h"
import "C"
import (
	"github.com/filecoin-project/filecoin-ffi/generated"
)

// Hash computes the digest of a message
func Hash(message Message) Digest {
	resp := generated.FilHash(message, uint(len(message)))
	resp.Deref()
	resp.Digest.Deref()

	defer generated.FilDestroyHashResponse(resp)

	var out Digest
	copy(out[:], resp.Digest.Inner[:])
	return out
}

// Verify verifies that a signature is the aggregated signature of digests - pubkeys
func Verify(signature *Signature, digests []Digest, publicKeys []PublicKey) bool {
	// prep data
	flattenedDigests := make([]byte, DigestBytes*len(digests))
	for idx, digest := range digests {
		copy(flattenedDigests[(DigestBytes*idx):(DigestBytes*(1+idx))], digest[:])
	}

	flattenedPublicKeys := make([]byte, PublicKeyBytes*len(publicKeys))
	for idx, publicKey := range publicKeys {
		copy(flattenedPublicKeys[(PublicKeyBytes*idx):(PublicKeyBytes*(1+idx))], publicKey[:])
	}

	isValid := generated.FilVerify(signature[:], flattenedDigests, uint(len(flattenedDigests)), flattenedPublicKeys, uint(len(flattenedPublicKeys)))

	return isValid > 0
}

// HashVerify verifies that a signature is the aggregated signature of hashed messages.
func HashVerify(signature *Signature, messages []Message, publicKeys []PublicKey) bool {
	var flattenedMessages []byte
	messagesSizes := make([]uint, len(messages))
	for idx := range messages {
		flattenedMessages = append(flattenedMessages, messages[idx]...)
		messagesSizes[idx] = uint(len(messages[idx]))
	}

	flattenedPublicKeys := make([]byte, PublicKeyBytes*len(publicKeys))
	for idx, publicKey := range publicKeys {
		copy(flattenedPublicKeys[(PublicKeyBytes*idx):(PublicKeyBytes*(1+idx))], publicKey[:])
	}

	isValid := generated.FilHashVerify(signature[:], flattenedMessages, uint(len(flattenedMessages)), messagesSizes, uint(len(messagesSizes)), flattenedPublicKeys, uint(len(flattenedPublicKeys)))

	return isValid > 0
}

// Aggregate aggregates signatures together into a new signature. If the
// provided signatures cannot be aggregated (due to invalid input or an
// an operational error), Aggregate will return nil.
func Aggregate(signatures []Signature) *Signature {
	// prep data
	flattenedSignatures := make([]byte, SignatureBytes*len(signatures))
	for idx, sig := range signatures {
		copy(flattenedSignatures[(SignatureBytes*idx):(SignatureBytes*(1+idx))], sig[:])
	}

	resp := generated.FilAggregate(flattenedSignatures, uint(len(flattenedSignatures)))
	if resp == nil {
		return nil
	}

	defer generated.FilDestroyAggregateResponse(resp)

	resp.Deref()
	resp.Signature.Deref()

	var out Signature
	copy(out[:], resp.Signature.Inner[:])
	return &out
}

// PrivateKeyGenerate generates a private key
func PrivateKeyGenerate() PrivateKey {
	resp := generated.FilPrivateKeyGenerate()
	resp.Deref()
	resp.PrivateKey.Deref()
	defer generated.FilDestroyPrivateKeyGenerateResponse(resp)

	var out PrivateKey
	copy(out[:], resp.PrivateKey.Inner[:])
	return out
}

// PrivateKeyGenerate generates a private key in a predictable manner
func PrivateKeyGenerateWithSeed(seed PrivateKeyGenSeed) PrivateKey {
	var ary generated.Fil32ByteArray
	copy(ary.Inner[:], seed[:])

	resp := generated.FilPrivateKeyGenerateWithSeed(ary)
	resp.Deref()
	resp.PrivateKey.Deref()
	defer generated.FilDestroyPrivateKeyGenerateResponse(resp)

	var out PrivateKey
	copy(out[:], resp.PrivateKey.Inner[:])
	return out
}

// PrivateKeySign signs a message
func PrivateKeySign(privateKey PrivateKey, message Message) *Signature {
	resp := generated.FilPrivateKeySign(privateKey[:], message, uint(len(message)))
	resp.Deref()
	resp.Signature.Deref()

	defer generated.FilDestroyPrivateKeySignResponse(resp)

	var signature Signature
	copy(signature[:], resp.Signature.Inner[:])
	return &signature
}

// PrivateKeyPublicKey gets the public key for a private key
func PrivateKeyPublicKey(privateKey PrivateKey) PublicKey {
	resp := generated.FilPrivateKeyPublicKey(privateKey[:])
	resp.Deref()
	resp.PublicKey.Deref()

	defer generated.FilDestroyPrivateKeyPublicKeyResponse(resp)

	var publicKey PublicKey
	copy(publicKey[:], resp.PublicKey.Inner[:])
	return publicKey
}

// CreateZeroSignature creates a zero signature, used as placeholder in filecoin.
func CreateZeroSignature() Signature {
	resp := generated.FilCreateZeroSignature()
	resp.Deref()
	resp.Signature.Deref()

	defer generated.FilDestroyZeroSignatureResponse(resp)

	var sig Signature
	copy(sig[:], resp.Signature.Inner[:])

	return sig
}
