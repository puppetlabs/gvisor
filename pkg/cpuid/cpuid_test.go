// Copyright 2020 The gVisor Authors.
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

package cpuid

import "testing"

func TestFeatureFromString(t *testing.T) {
	// Check that known features do match.
	for feature, _ := range allFeatures {
		f, ok := FeatureFromString(feature.String())
		if f != feature || !ok {
			t.Errorf("got %v, %v want %v, true", f, ok, feature)
		}
	}

	// Check that "bad" doesn't match.
	f, ok := FeatureFromString("bad")
	if ok {
		t.Errorf("got %v, %v want false", f, ok)
	}
}
