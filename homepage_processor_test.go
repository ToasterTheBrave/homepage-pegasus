package main

import (
	"testing"
)

func TestNewHomepageProcessor(t *testing.T) {

	var err error = nil

	processor, _ := NewHomepageProcessor("topic", 123, "username", "password")

	if(processor.topic != "topic") {
		t.Error("Expected \"topic\", got: ", processor.topic)
	}

	if(processor.offset != 123) {
		t.Error("Expected 123, got: ", processor.offset)
	}

	_, err = NewHomepageProcessor("", 123, "username", "password")
	if(err.Error() != "topic can not be empty") {
		t.Error("Expected \"topic can not be empty\", got: ", err.Error())
	}

	_, err = NewHomepageProcessor("topic", 123, "", "password")
	if(err.Error() != "dockercloud_user can not be empty") {
		t.Error("Expected \"topic can not be empty\", got: ", err.Error())
	}

	_, err = NewHomepageProcessor("topic", 123, "username", "")
	if(err.Error() != "dockercloud_apikey can not be empty") {
		t.Error("Expected \"topic can not be empty\", got: ", err.Error())
	}

}

func TestGetTopic(t *testing.T) {
	processor, _ := NewHomepageProcessor("topic", 123, "username", "password")
	if(processor.GetTopic() != "topic") {
		t.Error("Expected \"topic\", got: ", processor.GetTopic())
	}
}

func TestGetOffset(t *testing.T) {
	processor, _ := NewHomepageProcessor("topic", 123, "username", "password")
	if(processor.GetOffset() != 123) {
		t.Error("Expected 123, got: ", processor.GetOffset())
	}
}

// func TestProcessMessage(t *testing.T) {
// 	processor, _ := NewHomepageProcessor("topic", 123, "username", "password")
// 	success, _ := processor.ProcessMessage("key", "value")
// 	if(success != 1) {
// 		t.Error("Expected 1, got: ", success)
// 	}
// }