// Copyright 2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package modularis

import "os"

func Collect(dir string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		path := dir + "/" + entry.Name()

		if entry.IsDir() {
			deeperFiles, err := Collect(path)
			if err != nil {
				return nil, err
			}
			files = append(files, deeperFiles...)
		} else {
			files = append(files, path)
		}
	}

	return files, nil
}
