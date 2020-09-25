/*
Copyright paskal.maksim@gmail.com
Licensed under the Apache License, Version 2.0 (the "License")
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package sluglify

import (
	"testing"
)

type TestStrings struct {
	input string
	want  string
}

func TestGetSlugString(t *testing.T) {
	tests := make([]TestStrings, 0)

	tests = append(tests, TestStrings{
		input: "-123abcdeЦУЦ&#&&$^$^7-azcx23",
		want:  "abcdetsutsand-andzzz",
	})

	tests = append(tests, TestStrings{
		input: "--123test-$$$^$^",
		want:  "test",
	})

	tests = append(tests, TestStrings{
		input: "test-12345",
		want:  "test-12345",
	})

	tests = append(tests, TestStrings{
		input: "$^$^-test-12345",
		want:  "test-12345",
	})

	for _, test := range tests {
		stringLen := 20
		got := GetSlugString(test.input, stringLen, "zzz")

		if got != test.want {
			t.Errorf("GetSlugString() = %q, want %q", got, test.want)
		}

		if len(got) > stringLen {
			t.Errorf("GetSlugString() = %q, length must be %d", got, stringLen)
		}
	}
}
