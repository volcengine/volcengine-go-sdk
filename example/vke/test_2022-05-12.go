package main

import (
	"errors"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/vke"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

type CreateClusterOption struct {
	ClusterName string
	SubnetID    string
	PodCidr     string
	ServiceCidr string
	NetworkMode string
}

func CreateCluster(svc *vke.VKE, opt *CreateClusterOption) (string, error) {
	input := &vke.CreateClusterInput{
		ClusterConfig: &vke.ClusterConfigForCreateClusterInput{
			ApiServerPublicAccessEnabled: volcengine.Bool(false),
			SubnetIds:                    []*string{&opt.SubnetID},
		},
		Name: &opt.ClusterName,
		PodsConfig: &vke.PodsConfigForCreateClusterInput{
			PodNetworkMode: &opt.NetworkMode,
		},
		ServicesConfig: &vke.ServicesConfigForCreateClusterInput{
			ServiceCidrsv4: []*string{&opt.ServiceCidr},
		},
	}
	if opt.NetworkMode == vke.EnumOfPodNetworkModeForCreateClusterInputFlannel {
		input.PodsConfig.FlannelConfig = &vke.FlannelConfigForCreateClusterInput{
			PodCidrs:       []*string{&opt.PodCidr},
			MaxPodsPerNode: volcengine.Int32(64),
		}
	} else if opt.NetworkMode == vke.EnumOfPodNetworkModeForCreateClusterInputVpcCniShared {
		input.PodsConfig.VpcCniConfig = &vke.VpcCniConfigForCreateClusterInput{
			SubnetIds: []*string{&opt.SubnetID},
		}
	}
	resp, err := svc.CreateCluster(input)
	if err != nil {
		return "", err
	}
	return *resp.Id, nil
}

type GetClusterOption struct {
	ClusterID string
}

func GetCluster(svc *vke.VKE, opt *GetClusterOption) (*vke.ItemForListClustersOutput, error) {
	input := &vke.ListClustersInput{
		Filter: &vke.FilterForListClustersInput{
			CreateClientToken:        nil,
			DeleteProtectionEnabled:  nil,
			Ids:                      []*string{&opt.ClusterID},
			Name:                     nil,
			PodsConfigPodNetworkMode: nil,
			Statuses:                 nil,
			UpdateClientToken:        nil,
		},
		PageNumber: nil,
		PageSize:   nil,
	}

	resp, err := svc.ListClusters(input)
	if err != nil {
		return nil, err
	}
	if len(resp.Items) != 1 {
		return nil, errors.New("get cluster resp items must be 1")
	}
	return resp.Items[0], nil
}

type CreateNodePoolOption struct {
	ClusterID    string
	NodePoolName string
	InstanceType string
	Password     string
	SubnetID     string
}

func CreateNodePool(svc *vke.VKE, opt *CreateNodePoolOption) (string, error) {
	input := &vke.CreateNodePoolInput{
		ClusterId: &opt.ClusterID,
		Name:      &opt.NodePoolName,
		NodeConfig: &vke.NodeConfigForCreateNodePoolInput{
			SystemVolume: &vke.SystemVolumeForCreateNodePoolInput{
				Size: volcengine.Int32(40),
				Type: volcengine.String(vke.EnumOfTypeForCreateNodePoolInputEssdPl0),
			},
			InstanceTypeIds: []*string{&opt.InstanceType},
			Security: &vke.SecurityForCreateNodePoolInput{
				Login: &vke.LoginForCreateNodePoolInput{
					Password: &opt.Password,
				},
			},
			SubnetIds: []*string{&opt.SubnetID},
		},
	}

	resp, err := svc.CreateNodePool(input)
	if err != nil {
		return "", err
	}

	return *resp.Id, nil
}

type GetNodePoolOption struct {
	NodePoolID string
}

func GetNodePool(svc *vke.VKE, opt *GetNodePoolOption) (*vke.ItemForListNodePoolsOutput, error) {
	input := &vke.ListNodePoolsInput{
		Filter: &vke.FilterForListNodePoolsInput{
			Ids: []*string{&opt.NodePoolID},
		},
		PageNumber: nil,
		PageSize:   nil,
	}

	resp, err := svc.ListNodePools(input)
	if err != nil {
		return nil, err
	}
	if len(resp.Items) != 1 {
		return nil, errors.New("get node pool resp items must be 1")
	}
	return resp.Items[0], nil
}

type ListNodesOption struct {
	NodePoolID string
}

func ListNodes(svc *vke.VKE, opt *ListNodesOption) (*vke.ListNodesOutput, error) {
	input := &vke.ListNodesInput{
		Filter: &vke.FilterForListNodesInput{
			NodePoolIds: []*string{&opt.NodePoolID},
		},
		PageNumber: nil,
		PageSize:   nil,
	}

	return svc.ListNodes(input)
}

type InstallAddonOption struct {
	ClusterID string
}

func InstallAddon(svc *vke.VKE, opt *InstallAddonOption) error {
	const installAddonInterval = 5 * time.Second
	_, err := svc.CreateAddon(&vke.CreateAddonInput{
		ClusterId: &opt.ClusterID,
		Name:      volcengine.String("metrics-server"),
	})
	if err != nil {
		return err
	}
	time.Sleep(installAddonInterval)
	_, err = svc.CreateAddon(&vke.CreateAddonInput{
		ClusterId: &opt.ClusterID,
		Name:      volcengine.String("core-dns"),
	})
	if err != nil {
		return err
	}
	time.Sleep(installAddonInterval)
	_, err = svc.CreateAddon(&vke.CreateAddonInput{
		ClusterId: &opt.ClusterID,
		Name:      volcengine.String("cluster-autoscaler"),
		Config:    volcengine.String(`{"Expander":"random"}`),
	})
	if err != nil {
		return err
	}
	return nil
}

type UpdateNodePoolOption struct {
	ClusterID   string
	NodePoolID  string
	AutoScaling bool
}

func UpdateNodePool(svc *vke.VKE, opt *UpdateNodePoolOption) error {
	var num int32
	if opt.AutoScaling {
		num = 0
	} else {
		num = 1
	}
	input := &vke.UpdateNodePoolConfigInput{
		ClusterId: &opt.ClusterID,
		Id:        &opt.NodePoolID,
		AutoScaling: &vke.AutoScalingForUpdateNodePoolConfigInput{
			Enabled:         &opt.AutoScaling,
			DesiredReplicas: &num,
		},
	}

	_, err := svc.UpdateNodePoolConfig(input)
	return err
}
