/*
Copyright 2012 Google Inc.

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

/*
Package local implements functionality common to both the "localdisk" and
"diskpacked" storage mechanisms.
*/
package local

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// The file containing the salt for scrypt will be named this in the
// local (localdisk and diskpacked) blobservers. Normal enumeration
// and blob streaming must ignore this file. There is a special
// interface, Salter, to manipulate the data in this file.
const SaltFileName = "SALT"

// TODO (marete): Store the salt in hex-encoding.

// Implements PutSalt in the blobserver.Salter interface.
func PutSalt(salt []byte, dirRoot string) error {
	tmpFile, err := ioutil.TempFile(dirRoot, SaltFileName)
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	_, err = tmpFile.WriteString(hex.EncodeToString(salt) + "\n")
	if err != nil {
		return err
	}
	err = tmpFile.Sync()
	if err != nil {
		return err
	}

	saltPath := filepath.Join(dirRoot, SaltFileName)
	err = os.Rename(tmpFile.Name(), saltPath)
	if err != nil {
		return err
	}

	return nil
}

// Implements GetSalt in the blobserver.Salter interface.
func GetSalt(dirRoot string) ([]byte, error) {
	fname := filepath.Join(dirRoot, SaltFileName)

	data, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	bin, err := hex.DecodeString(strings.TrimSpace(string(data)))
	if err != nil {
		return nil, err
	}

	return bin, nil
}

//Implements HasSalt in the blobserver.Salter interface.
func HasSalt(dirRoot string) (bool, error) {
	fi, err := os.Stat(filepath.Join(dirRoot, SaltFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return fi.Size() > 0, nil
}
