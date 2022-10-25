package main

import (
	"fmt"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/vke"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	ak := "your_ak"
	sk := "your_sk"
	var region = "your_region"

	//use env Credentials
	//export VOLCSTACK_ACCESS_KEY=AK
	//export VOLCSTACK_SECRET_KEY=SK
	// and WithCredentials(credentials.NewEnvCredentials())

	// Init client
	config := volcengine.NewConfig().
		WithRegion(region).
		WithAkSk(ak, sk).
		//WithCredentials(credentials.NewEnvCredentials()).
		//WithLogLevel(volcengine.LogDebugWithHTTPBody).
		WithDisableSSL(true)
	sess, _ := session.NewSession(config)
	svc := vke.New(sess)

	// 1. 本示例将创建一个Fannel网络的VKE集群。并创建一个节点池，配置节点池的自动扩容策略。
	// 2. 本示例将涉及云服务器等付费产品资源的创建，请注意资源的及时释放。
	// CreateFlannelVKE(svc)

	// 1. 本示例将创建一个VPC-CNI网络的VKE集群。并创建一个节点池，配置节点池的自动扩容策略。
	// 2. 本示例将涉及云服务器等付费产品资源的创建，请注意资源的及时释放。
	CreateVpcCniVKE(svc)
}

func CreateFlannelVKE(svc *vke.VKE) {
	CreateVKE(svc, vke.EnumOfPodNetworkModeForCreateClusterInputFlannel)
}

func CreateVpcCniVKE(svc *vke.VKE) {
	CreateVKE(svc, vke.EnumOfPodNetworkModeForCreateClusterInputVpcCniShared)
}

func CreateVKE(svc *vke.VKE, networkMode string) {
	clusterName := "your_cluster_name"
	subnetID := "your_subnet_id"
	podCidr := "your_pod_cidr" // flannel cluster only
	serviceCidr := "your_service_cidr"
	nodePoolName := "your_node_pool_name"
	instanceType := "your_instance_type"
	password := "your_password"
	autoScaling := true

	// 1. 创建集群
	fmt.Println("create cluster")
	clusterID, err := CreateCluster(svc, &CreateClusterOption{
		ClusterName: clusterName,
		SubnetID:    subnetID,
		PodCidr:     podCidr,
		ServiceCidr: serviceCidr,
		NetworkMode: networkMode,
	})
	if err != nil {
		panic(err)
	}

	// 2. 查询集群状态
	fmt.Println("check cluster status")
	var cluster *vke.ItemForListClustersOutput
	for cluster == nil || *cluster.Status.Phase == vke.EnumOfPhaseForListClustersInputCreating {
		cluster, err = GetCluster(svc, &GetClusterOption{
			ClusterID: clusterID,
		})
		if err != nil {
			panic(err)
		}
		if cluster.Status == nil || cluster.Status.Phase == nil {
			panic("cluster status or phase nil")
		}
		time.Sleep(time.Second * 10)
	}

	// 3. 创建节点池
	fmt.Println("create node pool")
	nodePoolID, err := CreateNodePool(svc, &CreateNodePoolOption{
		ClusterID:    clusterID,
		NodePoolName: nodePoolName,
		InstanceType: instanceType,
		Password:     password,
		SubnetID:     subnetID,
	})
	if err != nil {
		panic(err)
	}

	// 4. 查看节点池
	fmt.Println("check node pool status")
	var nodePool *vke.ItemForListNodePoolsOutput
	for nodePool == nil || *nodePool.Status.Phase == vke.EnumOfPhaseForListNodePoolsInputCreating {
		nodePool, err = GetNodePool(svc, &GetNodePoolOption{
			NodePoolID: nodePoolID,
		})
		if err != nil {
			panic(err)
		}
		if nodePool.Status == nil || nodePool.Status.Phase == nil {
			panic("node pool status of phase nil")
		}
		time.Sleep(time.Second)
	}
	fmt.Println("node pool:")
	fmt.Println(*nodePool)

	// 5. 查看节点列表
	fmt.Println("list nodes")
	nodes, err := ListNodes(svc, &ListNodesOption{
		NodePoolID: nodePoolID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("node list:")
	fmt.Println(*nodes)

	// 6. 安装组件 and 设置节点池弹性策略
	fmt.Println("install addon")
	err = InstallAddon(svc, &InstallAddonOption{
		ClusterID: clusterID,
	})
	if err != nil {
		panic(err)
	}

	// 7. 修改节点数量 or 设置节点池弹性伸缩
	fmt.Println("update node pool")
	err = UpdateNodePool(svc, &UpdateNodePoolOption{
		ClusterID:   clusterID,
		NodePoolID:  nodePoolID,
		AutoScaling: autoScaling,
	})
	if err != nil {
		panic(err)
	}

	// 8. 查看节点列表
	fmt.Println("check nodes count")
	for len(nodes.Items) == 0 {
		nodes, err = ListNodes(svc, &ListNodesOption{
			NodePoolID: nodePoolID,
		})
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
	fmt.Println("node list:")
	fmt.Println(*nodes)
}
