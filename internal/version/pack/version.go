/*
Copyright 2024 The west2-online Authors.

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

package pack

import (
	"github.com/west2-online/fzuhelper-server/kitex_gen/model"
)

type Version struct {
	Version string
	Code    string
	Url     string
	Feature string
	Force   bool
}

func BuildVersion(version *Version) *model.Version {
	return &model.Version{
		VersionCode: &version.Code,
		VersionName: &version.Version,
		Force:       &version.Force,
		Changelog:   &version.Feature,
		Url:         &version.Url,
	}
}
