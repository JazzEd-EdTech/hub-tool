/*
   Copyright 2020 Docker Hub Tool authors

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

package hub

import (
	"os"

	"github.com/docker/docker/api/types/registry"
)

//Instance stores all the specific pieces needed to dialog with Hub
type Instance struct {
	APIHubBaseURL string
	RegistryInfo  *registry.IndexInfo
}

//GetInstance returns the current hub instance, which can be overridden by
// DOCKER_HUB_INSTANCE env var
func getInstance() *Instance {
	override := os.Getenv("DOCKER_SCAN_HUB_INSTANCE")
	switch override {
	case "staging":
		return &staging
	case "prod":
		return &prod
	default:
		return &prod
	}
}

var (
	staging = Instance{
		APIHubBaseURL: "https://hub-stage.docker.com",
		RegistryInfo: &registry.IndexInfo{
			Name:     "index-stage.docker.io",
			Mirrors:  nil,
			Secure:   true,
			Official: false,
		},
	}

	prod = Instance{
		APIHubBaseURL: "https://hub.docker.com",
		RegistryInfo: &registry.IndexInfo{
			Name:     "index.docker.io",
			Mirrors:  nil,
			Secure:   true,
			Official: true,
		},
	}
)
