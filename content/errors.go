/*
Copyright © 2020 Red Hat, Inc.

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

// Package content contains logic for parsing rule content.
package content

import "fmt"

// MissingMandatoryFile is an error raised while parsing, when a mandatory file is missing
type MissingMandatoryFile struct {
	FileName string
}

// InvalidItem is an error raised when unexpected type is found when parsing
type InvalidItem struct {
	FileName string
	KeyName  string
}

func (err MissingMandatoryFile) Error() string {
	return fmt.Sprintf("Missing required file: %s", err.FileName)
}

func (err InvalidItem) Error() string {
	return fmt.Sprintf("Invalid item `%s` of file %s", err.KeyName, err.FileName)
}
