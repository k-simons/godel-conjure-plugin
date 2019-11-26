// Copyright (c) 2018 Palantir Technologies. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/palantir/godel-conjure-plugin/v4/conjureplugin/config/internal/legacy"
	v1 "github.com/palantir/godel-conjure-plugin/v4/conjureplugin/config/internal/v1"
	"github.com/palantir/godel/v2/pkg/versionedconfig"
	"github.com/pkg/errors"
)

func UpgradeConfig(cfgBytes []byte) ([]byte, error) {
	if versionedconfig.IsLegacyConfig(cfgBytes) {
		v0Bytes, err := legacy.UpgradeConfig(cfgBytes)
		if err != nil {
			return nil, err
		}
		cfgBytes = v0Bytes
	}
	version, err := versionedconfig.ConfigVersion(cfgBytes)
	if err != nil {
		return nil, err
	}
	switch version {
	case "", "0":
		if len(cfgBytes) == 0 {
			return cfgBytes, nil
		}
		return nil, errors.Errorf("v0 configuration is not supported")
	case "1":
		return v1.UpgradeConfig(cfgBytes)
	default:
		return nil, errors.Errorf("unsupported version: %s", version)
	}
}
