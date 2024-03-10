// Copyright 2024 JetERA Creative
// This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0
// that can be found in the LICENSE file and https://mozilla.org/MPL/2.0/.

package modularis

import git "github.com/libgit2/git2go/v34"

type Environment struct {
	Root  string
	Proxy string
}

func (e *Environment) GetProxyOptions() git.ProxyOptions {
	if e.Proxy == "" {
		return git.ProxyOptions{
			Type: git.ProxyTypeNone,
		}
	}

	return git.ProxyOptions{
		Type: git.ProxyTypeSpecified,
		Url:  e.Proxy,
	}
}
