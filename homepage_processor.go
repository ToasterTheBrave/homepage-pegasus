package main

import (
	"github.com/docker/go-dockercloud/dockercloud"
	"log"
	"fmt"
)

type HomepageProcessor struct{
	topic string
	offset int64
	dockercloud_user string
	dockercloud_apikey string
}

func (processor HomepageProcessor) GetTopic() string {
	return processor.topic
}

func (processor HomepageProcessor) GetOffset() int64 {
	return processor.offset
}

// TODO : use a mock for dockercloud so that I can test this
func NewHomepageProcessor(topic string, offset int64, dockercloud_user string, dockercloud_apikey string) (*HomepageProcessor, error) {
	c := new(HomepageProcessor)
	if topic == "" {
		return nil, fmt.Errorf("topic can not be empty")
	}
	c.topic = topic

	c.offset = offset

	if dockercloud_user == "" {
		return nil, fmt.Errorf("dockercloud_user can not be empty")
	}
	c.dockercloud_user = dockercloud_user

	if dockercloud_apikey == "" {
		return nil, fmt.Errorf("dockercloud_apikey can not be empty")
	}
	c.dockercloud_apikey = dockercloud_apikey

	return c, nil
}

func (processor HomepageProcessor) ProcessMessage(key string, value string) (int, error) {
	log.Printf("Message received: '%s'\n", key)
	switch key {
		case "build-complete":
			return processor.redeployService(value)
		default:
			return 0, fmt.Errorf("Key not implemented: '%s'\n", key)
	}

}

func (processor HomepageProcessor) redeployService(serviceId string) (int, error) {

	log.Printf("Redeploying '%s'\n", serviceId);

	dockercloud.User = processor.dockercloud_user
	dockercloud.ApiKey = processor.dockercloud_apikey

	service, err := dockercloud.GetService(serviceId)
	if err != nil {
		return 0, err
	}

	reuseVolumes := dockercloud.ReuseVolumesOption{Reuse: false}

	err = service.Redeploy(reuseVolumes)
	if err != nil {
		return 0, err
	}

	return 1, nil
}
