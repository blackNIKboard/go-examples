package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type (
	inventory struct {
		All all `yaml:"all"`
	}
	all struct {
		Children children `yaml:"children"`
	}
	children struct {
		K3SCluster k3sCluster `yaml:"k3s_cluster"`
	}
	k3sCluster struct {
		Children nodes `yaml:"children"`
	}
	nodes struct {
		Master host `yaml:"master"`
		Node   host `yaml:"node"`
	}
	host struct {
		Hosts map[string]map[string]interface{} `yaml:"hosts"` // for 1.1.1.1: {} format
	}
)

func main() {
	inv := initInventory()
	inv.addNode("1.1.1.1", true)
	inv.addNode("2.2.2.2", false)
	inv.addNode("3.3.3.3", false)

	marshal, err := yaml.Marshal(&inv)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))
}

func initInventory() inventory {
	return inventory{
		all{
			children{
				k3sCluster{
					nodes{
						host{map[string]map[string]interface{}{}},
						host{map[string]map[string]interface{}{}},
					},
				},
			},
		},
	}
}

func (i *inventory) addNode(addr string, isMaster bool) {
	if isMaster {
		i.All.Children.K3SCluster.Children.Master.Hosts[addr] = map[string]interface{}{}
	} else {
		i.All.Children.K3SCluster.Children.Node.Hosts[addr] = map[string]interface{}{}
	}
}
