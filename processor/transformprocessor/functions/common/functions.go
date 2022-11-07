// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/functions/common"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"
)

func Functions[K any]() map[string]interface{} {
	return map[string]interface{}{
		"TraceID":              ottlfuncs.TraceID[K],
		"SpanID":               ottlfuncs.SpanID[K],
		"IsMatch":              ottlfuncs.IsMatch[K],
		"Concat":               ottlfuncs.Concat[K],
		"Split":                ottlfuncs.Split[K],
		"Int":                  ottlfuncs.Int[K],
		"keep_keys":            ottlfuncs.KeepKeys[K],
		"set":                  ottlfuncs.Set[K],
		"truncate_all":         ottlfuncs.TruncateAll[K],
		"limit":                ottlfuncs.Limit[K],
		"replace_match":        ottlfuncs.ReplaceMatch[K],
		"replace_all_matches":  ottlfuncs.ReplaceAllMatches[K],
		"replace_pattern":      ottlfuncs.ReplacePattern[K],
		"replace_all_patterns": ottlfuncs.ReplaceAllPatterns[K],
		"delete_key":           ottlfuncs.DeleteKey[K],
		"delete_matching_keys": ottlfuncs.DeleteMatchingKeys[K],
	}
}

// Registry is used to track and retrieve known operator types
type Registry[K any] struct {
	Functions map[string]interface{}
}

// DefaultRegistry creates a new registry
func DefaultRegistry[K any]() *Registry[K] {
	return &Registry[K]{
		Functions: Functions[K](),
	}
}

// Register will register a function to an operator type.
// This function will return a builder for the supplied type.
func (r *Registry[K]) Register(key string, function interface{}) {
	r.Functions[key] = function
}
