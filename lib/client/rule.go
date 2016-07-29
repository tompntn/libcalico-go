// Copyright (c) 2016 Tigera, Inc. All rights reserved.

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

package client

import (
	"github.com/tigera/libcalico-go/lib/api"
	"github.com/tigera/libcalico-go/lib/backend/model"
)

func ruleActionAPIToBackend(action string) string {
	return action
}

func ruleActionBackendToAPI(action string) string {
	return action
}

// Convert an API Rule structure to a Backend Rule structure
func ruleAPIToBackend(ar api.Rule) model.Rule {
	return model.Rule{
		Action:      ruleActionAPIToBackend(ar.Action),
		Protocol:    ar.Protocol,
		ICMPCode:    ar.ICMPCode,
		ICMPType:    ar.ICMPType,
		NotProtocol: ar.NotProtocol,
		NotICMPCode: ar.NotICMPCode,
		NotICMPType: ar.NotICMPType,

		SrcTag:      ar.Source.Tag,
		SrcNet:      ar.Source.Net,
		SrcSelector: ar.Source.Selector,
		SrcPorts:    ar.Source.Ports,
		DstTag:      ar.Destination.Tag,
		DstNet:      ar.Destination.Net,
		DstSelector: ar.Destination.Selector,
		DstPorts:    ar.Destination.Ports,

		NotSrcTag:      ar.Source.NotTag,
		NotSrcNet:      ar.Source.NotNet,
		NotSrcSelector: ar.Source.NotSelector,
		NotSrcPorts:    ar.Source.NotPorts,
		NotDstTag:      ar.Destination.NotTag,
		NotDstNet:      ar.Destination.NotNet,
		NotDstSelector: ar.Destination.NotSelector,
		NotDstPorts:    ar.Destination.NotPorts,
	}
}

// Convert a Backend Rule structure to an API Rule structure
func ruleBackendToAPI(br model.Rule) api.Rule {
	return api.Rule{
		Action:      ruleActionBackendToAPI(br.Action),
		Protocol:    br.Protocol,
		ICMPCode:    br.ICMPCode,
		ICMPType:    br.ICMPType,
		NotProtocol: br.NotProtocol,
		NotICMPCode: br.NotICMPCode,
		NotICMPType: br.NotICMPType,

		Source: api.EntityRule{
			Tag:         br.SrcTag,
			Net:         br.SrcNet,
			Selector:    br.SrcSelector,
			Ports:       br.SrcPorts,
			NotTag:      br.NotSrcTag,
			NotNet:      br.NotSrcNet,
			NotSelector: br.NotSrcSelector,
			NotPorts:    br.NotSrcPorts,
		},

		Destination: api.EntityRule{
			Tag:         br.DstTag,
			Net:         br.DstNet,
			Selector:    br.DstSelector,
			Ports:       br.DstPorts,
			NotTag:      br.NotDstTag,
			NotNet:      br.NotDstNet,
			NotSelector: br.NotDstSelector,
			NotPorts:    br.NotDstPorts,
		},
	}
}

// Convert an API Rule structure slice to a Backend Rule structure slice
func rulesAPIToBackend(ars []api.Rule) []model.Rule {
	if ars == nil {
		return []model.Rule{}
	}

	brs := make([]model.Rule, len(ars))
	for idx, ar := range ars {
		brs[idx] = ruleAPIToBackend(ar)
	}
	return brs
}

// Convert a Backend Rule structure slice to an API Rule structure slice
func rulesBackendToAPI(brs []model.Rule) []api.Rule {
	if brs == nil {
		return nil
	}

	ars := make([]api.Rule, len(brs))
	for idx, br := range brs {
		ars[idx] = ruleBackendToAPI(br)
	}
	return ars
}
