// Copyright © 2017 Ahmed T. Ali <ah.tajelsir@gmail.com>
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

package todos

// ImportList returns title -> Todo map for provider issues list
func ImportList(term string, client Provider) map[string]Todo {
	// Title => Todo Map
	all := make(map[string]Todo)

	// Errors should be handled in Provider.Search!
	todos := client.Search(term)

	for _, t := range todos {
		all[t.Title] = t
	}
	return all
}