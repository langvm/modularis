// Copyright 2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package modularis

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"sort"
)

func HashAllSha256(files []string) (string, error) {
	sort.Strings(files)

	h := sha256.New()

	for _, path := range files {
		f, err := os.Open(path)
		if err != nil {
			return "", err
		}

		hf := sha256.New()
		_, err = io.Copy(hf, f)
		if err != nil {
			return "", err
		}

		h.Write(hf.Sum(nil))
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func HashDirSha256(dir string) (string, error) {
	files, err := Collect(dir)
	if err != nil {
		return "", err
	}

	return HashAllSha256(files)
}
