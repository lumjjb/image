/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package blockcipher

import (
	"bytes"
	"io"
	"testing"
)

func TestBlockCipherHandlerCreate(t *testing.T) {
	_, err := NewLayerBlockCipherHandler()
	if err != nil {
		t.Fatal(err)
	}
}

func TestBlockCipherEncryption(t *testing.T) {
	var (
		layerData = []byte("this is some data")
	)

	h, err := NewLayerBlockCipherHandler()
	if err != nil {
		t.Fatal(err)
	}

	layerDataReader := bytes.NewReader(layerData)

	ciphertextReader, finalizer, err := h.Encrypt(layerDataReader, AES256CTR)
	if err != nil {
		t.Fatal(err)
	}

	ciphertext := make([]byte, 1024)
	encsize, err := ciphertextReader.Read(ciphertext)
	if err != io.EOF {
		t.Fatal("Expected EOF")
	}

	ciphertextTestReader := bytes.NewReader(ciphertext[:encsize])
	lbco, err := finalizer()
	if err != nil {
		t.Fatal(err)
	}

	// Use a different instantiated object to indicate an invokation at a diff time
	plaintextReader, _, err := h.Decrypt(ciphertextTestReader, lbco)
	if err != nil {
		t.Fatal(err)
	}

	plaintext := make([]byte, 1024)
	decsize, err := plaintextReader.Read(plaintext)
	if err != nil && err != io.EOF {
		t.Fatal("Reading the plaintext should not have failed")
	}

	if string(plaintext[:decsize]) != string(layerData) {
		t.Fatal("Decrypted data is incorrect")
	}
}

func TestBlockCipherEncryptionInvalidKey(t *testing.T) {
	var (
		layerData = []byte("this is some data")
	)

	h, err := NewLayerBlockCipherHandler()
	if err != nil {
		t.Fatal(err)
	}

	layerDataReader := bytes.NewReader(layerData)

	ciphertextReader, finalizer, err := h.Encrypt(layerDataReader, AES256CTR)
	if err != nil {
		t.Fatal(err)
	}

	// Use a different instantiated object to indicate an invokation at a diff time
	bc2, err := NewAESCTRLayerBlockCipher(256)
	if err != nil {
		t.Fatal(err)
	}

	ciphertext := make([]byte, 1024)
	encsize, err := ciphertextReader.Read(ciphertext)
	if err != io.EOF {
		t.Fatal("Expected EOF")
	}

	ciphertextTestReader := bytes.NewReader(ciphertext[:encsize])
	lbco, err := finalizer()
	if err != nil {
		t.Fatal(err)
	}
	lbco.Private.SymmetricKey = []byte("aaa34567890123456789012345678901")

	plaintextReader, _, err := bc2.Decrypt(ciphertextTestReader, lbco)
	if err != nil {
		t.Fatal(err)
	}

	plaintext := make([]byte, 1024)
	_, err = plaintextReader.Read(plaintext)
	if err == nil || err == io.EOF {
		t.Fatal("Read() should have failed due to wrong key")
	}
}
