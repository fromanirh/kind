/*
Copyright 2018 The Kubernetes Authors.

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

package docker

import (
	"sigs.k8s.io/kind/pkg/exec"
)

// Run creates a container with "docker run", with some error handling
func Run(image string, runArgs []string, containerArgs []string) error {
	// construct the actual docker run argv
	args := []string{"run"}
	args = append(args, runArgs...)
	args = append(args, image)
	args = append(args, containerArgs...)
	cmd := exec.Command("docker", args...)
	_, err := exec.CombinedOutputLines(cmd)
	return err
}
