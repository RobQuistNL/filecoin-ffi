// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated

/*
#cgo LDFLAGS: -L${SRCDIR}/..
#cgo pkg-config: ${SRCDIR}/../filcrypto.pc
#include "../filcrypto.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// FCPResponseStatus as declared in filecoin-ffi/filcrypto.h:31
type FCPResponseStatus int32

// FCPResponseStatus enumeration from filecoin-ffi/filcrypto.h:31
const (
	FCPResponseStatusFCPNoError           FCPResponseStatus = C.FCPResponseStatus_FCPNoError
	FCPResponseStatusFCPUnclassifiedError FCPResponseStatus = C.FCPResponseStatus_FCPUnclassifiedError
	FCPResponseStatusFCPCallerError       FCPResponseStatus = C.FCPResponseStatus_FCPCallerError
	FCPResponseStatusFCPReceiverError     FCPResponseStatus = C.FCPResponseStatus_FCPReceiverError
)

// FilFvmRegisteredVersion as declared in filecoin-ffi/filcrypto.h:35
type FilFvmRegisteredVersion int32

// FilFvmRegisteredVersion enumeration from filecoin-ffi/filcrypto.h:35
const (
	FilFvmRegisteredVersionV1 FilFvmRegisteredVersion = C.fil_FvmRegisteredVersion_V1
)

// FilRegisteredAggregationProof as declared in filecoin-ffi/filcrypto.h:39
type FilRegisteredAggregationProof int32

// FilRegisteredAggregationProof enumeration from filecoin-ffi/filcrypto.h:39
const (
	FilRegisteredAggregationProofSnarkPackV1 FilRegisteredAggregationProof = C.fil_RegisteredAggregationProof_SnarkPackV1
)

// FilRegisteredPoStProof as declared in filecoin-ffi/filcrypto.h:52
type FilRegisteredPoStProof int32

// FilRegisteredPoStProof enumeration from filecoin-ffi/filcrypto.h:52
const (
	FilRegisteredPoStProofStackedDrgWinning2KiBV1   FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWinning2KiBV1
	FilRegisteredPoStProofStackedDrgWinning8MiBV1   FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWinning8MiBV1
	FilRegisteredPoStProofStackedDrgWinning512MiBV1 FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWinning512MiBV1
	FilRegisteredPoStProofStackedDrgWinning32GiBV1  FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWinning32GiBV1
	FilRegisteredPoStProofStackedDrgWinning64GiBV1  FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWinning64GiBV1
	FilRegisteredPoStProofStackedDrgWindow2KiBV1    FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWindow2KiBV1
	FilRegisteredPoStProofStackedDrgWindow8MiBV1    FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWindow8MiBV1
	FilRegisteredPoStProofStackedDrgWindow512MiBV1  FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWindow512MiBV1
	FilRegisteredPoStProofStackedDrgWindow32GiBV1   FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWindow32GiBV1
	FilRegisteredPoStProofStackedDrgWindow64GiBV1   FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWindow64GiBV1
)

// FilRegisteredSealProof as declared in filecoin-ffi/filcrypto.h:65
type FilRegisteredSealProof int32

// FilRegisteredSealProof enumeration from filecoin-ffi/filcrypto.h:65
const (
	FilRegisteredSealProofStackedDrg2KiBV1    FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg2KiBV1
	FilRegisteredSealProofStackedDrg8MiBV1    FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg8MiBV1
	FilRegisteredSealProofStackedDrg512MiBV1  FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg512MiBV1
	FilRegisteredSealProofStackedDrg32GiBV1   FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg32GiBV1
	FilRegisteredSealProofStackedDrg64GiBV1   FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg64GiBV1
	FilRegisteredSealProofStackedDrg2KiBV11   FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg2KiBV1_1
	FilRegisteredSealProofStackedDrg8MiBV11   FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg8MiBV1_1
	FilRegisteredSealProofStackedDrg512MiBV11 FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg512MiBV1_1
	FilRegisteredSealProofStackedDrg32GiBV11  FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg32GiBV1_1
	FilRegisteredSealProofStackedDrg64GiBV11  FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg64GiBV1_1
)

// FilRegisteredUpdateProof as declared in filecoin-ffi/filcrypto.h:73
type FilRegisteredUpdateProof int32

// FilRegisteredUpdateProof enumeration from filecoin-ffi/filcrypto.h:73
const (
	FilRegisteredUpdateProofStackedDrg2KiBV1   FilRegisteredUpdateProof = C.fil_RegisteredUpdateProof_StackedDrg2KiBV1
	FilRegisteredUpdateProofStackedDrg8MiBV1   FilRegisteredUpdateProof = C.fil_RegisteredUpdateProof_StackedDrg8MiBV1
	FilRegisteredUpdateProofStackedDrg512MiBV1 FilRegisteredUpdateProof = C.fil_RegisteredUpdateProof_StackedDrg512MiBV1
	FilRegisteredUpdateProofStackedDrg32GiBV1  FilRegisteredUpdateProof = C.fil_RegisteredUpdateProof_StackedDrg32GiBV1
	FilRegisteredUpdateProofStackedDrg64GiBV1  FilRegisteredUpdateProof = C.fil_RegisteredUpdateProof_StackedDrg64GiBV1
)
