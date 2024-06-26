// /*
// Copyright The Kubernetes Authors.
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
// */

// Code generated by client-gen. DO NOT EDIT.
package azclient

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"

	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/availabilitysetclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/deploymentclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/diskclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/interfaceclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/loadbalancerclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/managedclusterclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/policy/ratelimit"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/privateendpointclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatelinkserviceclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatezoneclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipaddressclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipprefixclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/routetableclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/securitygroupclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/snapshotclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/subnetclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachineclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetclient"
	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetvmclient"
)

type ClientFactoryImpl struct {
	*ClientFactoryConfig
	cred                                    azcore.TokenCredential
	availabilitysetclientInterface          availabilitysetclient.Interface
	deploymentclientInterface               deploymentclient.Interface
	diskclientInterface                     diskclient.Interface
	interfaceclientInterface                interfaceclient.Interface
	loadbalancerclientInterface             loadbalancerclient.Interface
	managedclusterclientInterface           managedclusterclient.Interface
	privateendpointclientInterface          privateendpointclient.Interface
	privatelinkserviceclientInterface       privatelinkserviceclient.Interface
	privatezoneclientInterface              privatezoneclient.Interface
	publicipaddressclientInterface          publicipaddressclient.Interface
	publicipprefixclientInterface           publicipprefixclient.Interface
	routetableclientInterface               routetableclient.Interface
	securitygroupclientInterface            securitygroupclient.Interface
	snapshotclientInterface                 snapshotclient.Interface
	subnetclientInterface                   subnetclient.Interface
	virtualmachineclientInterface           virtualmachineclient.Interface
	virtualmachinescalesetclientInterface   virtualmachinescalesetclient.Interface
	virtualmachinescalesetvmclientInterface virtualmachinescalesetvmclient.Interface
}

func NewClientFactory(config *ClientFactoryConfig, armConfig *ARMClientConfig, cred azcore.TokenCredential) (ClientFactory, error) {
	if config == nil {
		config = &ClientFactoryConfig{}
	}
	if cred == nil {
		cred = &azidentity.DefaultAzureCredential{}
	}

	var options *arm.ClientOptions
	var err error

	//initialize {availabilitysetclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/availabilitysetclient Interface availabilitySetRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}

	var ratelimitOption *ratelimit.Config
	var rateLimitPolicy policy.Policy
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("availabilitySetRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	availabilitysetclientInterface, err := availabilitysetclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {deploymentclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/deploymentclient Interface deploymentRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("deploymentRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	deploymentclientInterface, err := deploymentclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {diskclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/diskclient Interface diskRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("diskRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	diskclientInterface, err := diskclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {interfaceclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/interfaceclient Interface interfaceRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("interfaceRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	interfaceclientInterface, err := interfaceclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {loadbalancerclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/loadbalancerclient Interface loadBalancerRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("loadBalancerRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	loadbalancerclientInterface, err := loadbalancerclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {managedclusterclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/managedclusterclient Interface containerServiceRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("containerServiceRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	managedclusterclientInterface, err := managedclusterclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {privateendpointclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/privateendpointclient Interface privateEndpointRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("privateEndpointRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	privateendpointclientInterface, err := privateendpointclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {privatelinkserviceclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatelinkserviceclient Interface privateLinkServiceRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("privateLinkServiceRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	privatelinkserviceclientInterface, err := privatelinkserviceclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {privatezoneclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/privatezoneclient Interface privateDNSRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("privateDNSRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	privatezoneclientInterface, err := privatezoneclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {publicipaddressclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipaddressclient Interface publicIPAddressRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("publicIPAddressRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	publicipaddressclientInterface, err := publicipaddressclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {publicipprefixclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/publicipprefixclient Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}

	publicipprefixclientInterface, err := publicipprefixclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {routetableclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/routetableclient Interface routeTableRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("routeTableRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	routetableclientInterface, err := routetableclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {securitygroupclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/securitygroupclient Interface securityGroupRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("securityGroupRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	securitygroupclientInterface, err := securitygroupclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {snapshotclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/snapshotclient Interface snapshotRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("snapshotRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	snapshotclientInterface, err := snapshotclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {subnetclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/subnetclient Interface subnetsRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("subnetsRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	subnetclientInterface, err := subnetclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {virtualmachineclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachineclient Interface virtualMachineRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("virtualMachineRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	virtualmachineclientInterface, err := virtualmachineclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {virtualmachinescalesetclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetclient Interface virtualMachineSizesRateLimit}
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}
	//add ratelimit policy
	ratelimitOption = config.GetRateLimitConfig("virtualMachineSizesRateLimit")
	rateLimitPolicy = ratelimit.NewRateLimitPolicy(ratelimitOption)
	options.ClientOptions.PerCallPolicies = append(options.ClientOptions.PerCallPolicies, rateLimitPolicy)
	virtualmachinescalesetclientInterface, err := virtualmachinescalesetclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	//initialize {virtualmachinescalesetvmclient sigs.k8s.io/cloud-provider-azure/pkg/azclient/virtualmachinescalesetvmclient Interface }
	options, err = GetDefaultResourceClientOption(armConfig, config)
	if err != nil {
		return nil, err
	}

	virtualmachinescalesetvmclientInterface, err := virtualmachinescalesetvmclient.New(config.SubscriptionID, cred, options)
	if err != nil {
		return nil, err
	}

	return &ClientFactoryImpl{
		ClientFactoryConfig: config,
		cred:                cred, availabilitysetclientInterface: availabilitysetclientInterface,
		deploymentclientInterface:               deploymentclientInterface,
		diskclientInterface:                     diskclientInterface,
		interfaceclientInterface:                interfaceclientInterface,
		loadbalancerclientInterface:             loadbalancerclientInterface,
		managedclusterclientInterface:           managedclusterclientInterface,
		privateendpointclientInterface:          privateendpointclientInterface,
		privatelinkserviceclientInterface:       privatelinkserviceclientInterface,
		privatezoneclientInterface:              privatezoneclientInterface,
		publicipaddressclientInterface:          publicipaddressclientInterface,
		publicipprefixclientInterface:           publicipprefixclientInterface,
		routetableclientInterface:               routetableclientInterface,
		securitygroupclientInterface:            securitygroupclientInterface,
		snapshotclientInterface:                 snapshotclientInterface,
		subnetclientInterface:                   subnetclientInterface,
		virtualmachineclientInterface:           virtualmachineclientInterface,
		virtualmachinescalesetclientInterface:   virtualmachinescalesetclientInterface,
		virtualmachinescalesetvmclientInterface: virtualmachinescalesetvmclientInterface,
	}, nil
}

func (factory *ClientFactoryImpl) GetavailabilitysetclientInterface() availabilitysetclient.Interface {
	return factory.availabilitysetclientInterface
}

func (factory *ClientFactoryImpl) GetdeploymentclientInterface() deploymentclient.Interface {
	return factory.deploymentclientInterface
}

func (factory *ClientFactoryImpl) GetdiskclientInterface() diskclient.Interface {
	return factory.diskclientInterface
}

func (factory *ClientFactoryImpl) GetinterfaceclientInterface() interfaceclient.Interface {
	return factory.interfaceclientInterface
}

func (factory *ClientFactoryImpl) GetloadbalancerclientInterface() loadbalancerclient.Interface {
	return factory.loadbalancerclientInterface
}

func (factory *ClientFactoryImpl) GetmanagedclusterclientInterface() managedclusterclient.Interface {
	return factory.managedclusterclientInterface
}

func (factory *ClientFactoryImpl) GetprivateendpointclientInterface() privateendpointclient.Interface {
	return factory.privateendpointclientInterface
}

func (factory *ClientFactoryImpl) GetprivatelinkserviceclientInterface() privatelinkserviceclient.Interface {
	return factory.privatelinkserviceclientInterface
}

func (factory *ClientFactoryImpl) GetprivatezoneclientInterface() privatezoneclient.Interface {
	return factory.privatezoneclientInterface
}

func (factory *ClientFactoryImpl) GetpublicipaddressclientInterface() publicipaddressclient.Interface {
	return factory.publicipaddressclientInterface
}

func (factory *ClientFactoryImpl) GetpublicipprefixclientInterface() publicipprefixclient.Interface {
	return factory.publicipprefixclientInterface
}

func (factory *ClientFactoryImpl) GetroutetableclientInterface() routetableclient.Interface {
	return factory.routetableclientInterface
}

func (factory *ClientFactoryImpl) GetsecuritygroupclientInterface() securitygroupclient.Interface {
	return factory.securitygroupclientInterface
}

func (factory *ClientFactoryImpl) GetsnapshotclientInterface() snapshotclient.Interface {
	return factory.snapshotclientInterface
}

func (factory *ClientFactoryImpl) GetsubnetclientInterface() subnetclient.Interface {
	return factory.subnetclientInterface
}

func (factory *ClientFactoryImpl) GetvirtualmachineclientInterface() virtualmachineclient.Interface {
	return factory.virtualmachineclientInterface
}

func (factory *ClientFactoryImpl) GetvirtualmachinescalesetclientInterface() virtualmachinescalesetclient.Interface {
	return factory.virtualmachinescalesetclientInterface
}

func (factory *ClientFactoryImpl) GetvirtualmachinescalesetvmclientInterface() virtualmachinescalesetvmclient.Interface {
	return factory.virtualmachinescalesetvmclientInterface
}
