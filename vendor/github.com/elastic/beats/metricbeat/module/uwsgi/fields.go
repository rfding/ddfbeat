// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package uwsgi

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "uwsgi", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzEl8tu6zYQhvd+ikFWLWAb7daLAgGSRRdtUzvnslNocWQTpkhlOLQiP/0BdfFVVowDGoc7yfT/f8PLzGgCG6xm4Eu3UiMAVqxxBg/188MIQKJLSRWsrJnBXyMAaOZCbqXXOAIg1CgczmCJLEYAmUIt3ayeOgEjcjzIh8FVgTNYkfVF+6bH41TmWMqxYO/2r/v0wjin6kavVzNqyGmjDzkyqdQhtyBHM8/JjunYstBTwnePjt3JlI5UW7M6+2EAKozXIAqdKKyFkRrlgD9+pFirxSW4InvqXZJiTJDI0p3irx2gx+F8E4S8D4eQn7sXSkZyfSGbonOg+ne8tLRBmkaz+1brfeIWL7rWrrgpSJHWB/DC42fNFywYwWat/hj+BJWBY6U17L3AYHk4fJbXSKVyCH8MkUbOAP/6fIkUSPcgDmmLEpYV8Fq51ncISaJmkfwiMBAZI3UPyoX8bIVECeUaTbhT6Rol/PP4PZk////lefG6GIoleoJ73gsCCeWuZNfWfS1IbBSpJLXecCSCJ7JFgfKwjmEBVY7W8xCLUysjdLRlyBUzSmhV4TLAXvvk3aPH6EeqoyiFqi8iW1jiYPnrqM57hAPPBqvS0nn2ui1Jta3Bb+kaRTGGQniH40A5hqV31RiU1Pj7YFpwEanm6JREw7BAhoXa4RRyzC1VkHoiNKwr8K65iqLLsNOQ4YQuReVgh2SBqYK3yaT554SwsMRvYGvXsAnHrVtvTFu3i7TzXxWxF/oyHqd2CMKF84AyHIM7hkPeGGVWSbh7kSt4Kw0X0hfFwxWiNFHTy7zR/PxG80esbomEcft8shuMWGxXCcUK9XGLJFYIYRmtcXh9vVNLGK9zaovb30/XvZopScTmUISL3lm/DHl3hWVaN6jRkzWfNulNF3ADTUiqKo1fO2pZyJTGhqW/WT9lIesZY23NvBaD3oar395mWdMURSL4r9O7AeIOn22HzciE0qHVCx6DCxH7i+0CIVi40Y8AAAD//yl8fTE="
}