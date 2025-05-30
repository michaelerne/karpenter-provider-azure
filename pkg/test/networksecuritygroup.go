/*
Portions Copyright (c) Microsoft Corporation.

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

package test

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/Azure/karpenter-provider-azure/pkg/fake"
	"github.com/samber/lo"
)

func MakeNetworkSecurityGroup(resourceGroup string, name string) armnetwork.SecurityGroup {
	nsgID := fake.MakeNetworkSecurityGroupID(resourceGroup, name)

	result := armnetwork.SecurityGroup{
		ID:   &nsgID,
		Name: &name,
		Properties: &armnetwork.SecurityGroupPropertiesFormat{
			SecurityRules: []*armnetwork.SecurityRule{
				{
					Name: lo.ToPtr("k8s-azure-lb_allow_IPv4_0000"),
					// TODO: Not filling this in now, can later if we need it
				},
			},
		},
	}

	return result
}
