package main

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/consul/api"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func checkServiceHealth(client *api.Client) {
	for {
		services, _, err := client.Health().Service("payhub", "", true, nil)
		if err != nil {
			log.Fatalf("Failed to check service health: %v", err)
		}

		for _, entry := range services {
			fmt.Printf("Service ID: %s, Status: %s\n", entry.Service.ID, entry.Checks.AggregatedStatus())
		}

		time.Sleep(10 * time.Second) // 每10秒检查一次
	}
}

func monitorConfigChanges(client *api.Client) {
	var lastIndex uint64

	for {
		kvPair, meta, err := client.KV().Get("config/payhub", &api.QueryOptions{
			WaitIndex: lastIndex,
		})
		if err != nil {
			log.Fatalf("Failed to retrieve KV from Consul: %v", err)
		}

		if kvPair != nil && meta.LastIndex > lastIndex {
			fmt.Printf("Config changed: %s\n", string(kvPair.Value))
			lastIndex = meta.LastIndex
		}
		time.Sleep(5 * time.Second) // 每5秒检查一次配置变化
	}
}
func main1() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	checkServiceHealth(client)
	//monitorConfigChanges(client)
}
func main() {
	// 初始化 Consul 客户端
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}
	// 获取服务地址

	services, _, err := client.Catalog().Service("payhub", "", nil)
	if err != nil {
		log.Fatalf("Failed to retrieve service from Consul: %v", err)
	}
	httpServices := make([]*api.CatalogService, 0)
	for _, service := range services {
		fmt.Println(service.Checks.AggregatedStatus())
		if service.ServicePort == 9000 {
			continue
		} else {
			httpServices = append(httpServices, service)
		}
	}
	if len(httpServices) == 0 {
		log.Fatalf("No service found")
	}
	rand.Seed(time.Now().Unix())
	fmt.Println(len(httpServices))
	selectedService := httpServices[rand.Intn(len(httpServices))]
	//service := services[1]
	serviceAddress := fmt.Sprintf("http://%s:%d", selectedService.ServiceAddress, selectedService.ServicePort)
	fmt.Println("serviceAddress", serviceAddress)
	jsonData := []byte(`{"merchantid":"1000011","amount":"123456"}`)
	// 发送 HTTP 请求到 Kratos 服务
	req, err := http.NewRequest("POST", serviceAddress+"/payment/create", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	// 设置请求头为 JSON
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	http_client := &http.Client{}
	resp, err := http_client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make request to service: %v", err)
	}
	defer resp.Body.Close()
	// 处理响应
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request successful")
	} else {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	}
	response, err := io.ReadAll(resp.Body)
	fmt.Printf("Response: %s\n", string(response))
}
