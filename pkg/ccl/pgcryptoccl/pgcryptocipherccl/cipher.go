// Copyright 2023 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

package pgcryptocipherccl

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgcode"
	"github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgerror"
	"github.com/cockroachdb/errors"
)

// Encrypt returns the ciphertext obtained by running the encryption
// algorithm for the specified cipher type with the provided key and
// initialization vector over the provided data.
func Encrypt(data []byte, key []byte, iv []byte, cipherType string) ([]byte, error) {
	method, err := parseCipherMethod(cipherType)
	if err != nil {
		return nil, err
	}
	block, err := newCipher(method, key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	data, err = padData(method, data, blockSize)
	if err != nil {
		return nil, err
	}
	err = validateDataLength(data, blockSize)
	if err != nil {
		return nil, err
	}
	return encrypt(method, block, iv, data)
}

// Decrypt returns the plaintext obtained by running the decryption
// algorithm for the specified cipher type with the provided key and
// initialization vector over the provided data.
func Decrypt(data []byte, key []byte, iv []byte, cipherType string) ([]byte, error) {
	method, err := parseCipherMethod(cipherType)
	if err != nil {
		return nil, err
	}
	block, err := newCipher(method, key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	err = validateDataLength(data, blockSize)
	if err != nil {
		return nil, err
	}
	data, err = decrypt(method, block, iv, data)
	if err != nil {
		return nil, err
	}
	return unpadData(method, data)
}

func newCipher(method cipherMethod, key []byte) (cipher.Block, error) {
	switch a := method.algorithm; a {
	case aesCipher:
		switch l := len(key); {
		case l >= 32:
			key = zeroPadOrTruncate(key, 32)
		case l >= 24:
			key = zeroPadOrTruncate(key, 24)
		default:
			key = zeroPadOrTruncate(key, 16)
		}
		return aes.NewCipher(key)
	default:
		return nil, errors.Newf("cannot create new cipher for unknown algorithm: %d", a)
	}
}

func padData(method cipherMethod, data []byte, blockSize int) ([]byte, error) {
	switch p := method.padding; p {
	case pkcsPadding:
		return pkcsPad(data, blockSize)
	case noPadding:
		return data, nil
	default:
		return nil, errors.Newf("cannot pad for unknown padding: %d", p)
	}
}

func unpadData(method cipherMethod, data []byte) ([]byte, error) {
	switch p := method.padding; p {
	case pkcsPadding:
		return pkcsUnpad(data)
	case noPadding:
		return data, nil
	default:
		return nil, errors.Newf("cannot unpad for unknown padding: %d", p)
	}
}

func validateDataLength(data []byte, blockSize int) error {
	if len(data)%blockSize != 0 {
		// TODO(yang): Not sure if this is the right pgcode since Postgres
		// returns pgcode.ExternalRoutineException.
		return pgerror.Newf(
			pgcode.InvalidParameterValue,
			`data has length %d, which is not a multiple of block size %d`,
			len(data), blockSize,
		)
	}
	return nil
}

func encrypt(method cipherMethod, block cipher.Block, iv []byte, data []byte) ([]byte, error) {
	switch m := method.mode; m {
	case cbcMode:
		ret := make([]byte, len(data))
		iv = zeroPadOrTruncate(iv, block.BlockSize())
		mode := cipher.NewCBCEncrypter(block, iv)
		mode.CryptBlocks(ret, data)
		return ret, nil
	default:
		return nil, errors.Newf("cannot encrypt for unknown mode: %d", m)
	}
}

func decrypt(method cipherMethod, block cipher.Block, iv []byte, data []byte) ([]byte, error) {
	switch m := method.mode; m {
	case cbcMode:
		ret := make([]byte, len(data))
		iv = zeroPadOrTruncate(iv, block.BlockSize())
		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(ret, data)
		return ret, nil
	default:
		return nil, errors.Newf("cannot encrypt for unknown mode: %d", m)
	}
}

func zeroPadOrTruncate(data []byte, size int) []byte {
	if len(data) >= size {
		return data[:size]
	}
	paddedData := make([]byte, size)
	copy(paddedData, data)
	return paddedData
}
