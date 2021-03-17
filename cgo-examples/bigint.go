package cgo_examples

// #cgo LDFLAGS: -lcrypto
// #include "bn.h"
import "C"

import (
	"math/big"
	"reflect"
	"unsafe"
)

/*
 * rsa algorithm
 *  c = m^e (mod N)
 * where N = p*q.
 * note that p, q are two large prime number of same bit-length.
 */

func bigint_to_bn(x *big.Int) *C.BIGNUM {
	bytes := x.Bytes()
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	len := C.int(hdr.Len * int(unsafe.Sizeof(byte(0))))
	ptr := (*C.uchar)(unsafe.Pointer(hdr.Data))

	_p := C.BN_new()
	C.BN_bin2bn(ptr, len, _p)
	return _p
}

func bn_to_bigint(x *C.BIGNUM) *big.Int {
	len := (int(C.BN_num_bits(x)) + 7) / 8
	bytes := make([]byte, len)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	ptr := (*C.uchar)(unsafe.Pointer(hdr.Data))

	C.BN_bn2bin(x, ptr)
	C.BN_free(x)

	return new(big.Int).SetBytes(bytes)
}

func BnMult(p *big.Int, q *big.Int) *big.Int {
	_p := bigint_to_bn(p)
	_q := bigint_to_bn(q)

	ctx := C.BN_CTX_new()
	r := C.BN_new()
	C.BN_mul(r, _p, _q, ctx)

	C.BN_free(_p)
	C.BN_free(_q)
	C.BN_CTX_free(ctx)

	return bn_to_bigint(r)
}
