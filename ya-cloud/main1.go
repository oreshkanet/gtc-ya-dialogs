package main

import (
	"context"
	"fmt"
)

type requestType string

const (
	Discovery requestType = "discovery"
)

func Handler(ctx context.Context, req RequestDTO) (*ResponseDTO, error) {
	print(fmt.Sprintf("%v", req))
	res := &ResponseDTO{
		RequestId: req.Headers.RequestId,
		Payload: &DiscoveryResponseDTO{
			UserId:  "1",
			Devices: make([]DeviceDTO, 0),
		},
	}
	res.Payload.Devices = append(res.Payload.Devices, DeviceDTO{
		Id:          "1",
		Name:        "GTC_вентиляция",
		Description: "GTC_вентиляция (+-)",
		Room:        "Дом",
		Type:        "devices.types.thermostat",
		Capabilities: []CapabilityDTO{
			{
				Type:        "devices.capabilities.on_off",
				Retrievable: false,
				Reportable:  false,
				Parameters:  nil,
			},
		},
	})

	print(fmt.Sprintf("%v", res))

	return res, nil
}

func Handler1(ctx context.Context, req []byte) (*ResponseDTO, error) {
	print(fmt.Sprintf("%v", string(req)))
	return &ResponseDTO{}, nil
}

type RequestDTO struct {
	Headers struct {
		RequestId     string `json:"request_id"`
		Authorization string `json:"authorization"`
	} `json:"headers"`
	RequestType string  `json:"request_type"`
	ApiVersion  float64 `json:"api_version"`
}

type ResponseDTO struct {
	RequestId string                `json:"request_id"`
	Payload   *DiscoveryResponseDTO `json:"payload"`
}

type DiscoveryResponseDTO struct {
	UserId  string      `json:"user_id"`
	Devices []DeviceDTO `json:"devices"`
}

type DeviceDTO struct {
	Id           string          `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Room         string          `json:"room"`
	Type         string          `json:"type"`
	CustomData   string          `json:"custom_data"`
	Capabilities []CapabilityDTO `json:"capabilities"`
	Properties   []interface{}   `json:"properties"`
	DeviceInfo   struct {
		Manufacturer string `json:"manufacturer"`
		Model        string `json:"model"`
		HwVersion    string `json:"hw_version"`
		SwVersion    string `json:"sw_version"`
	} `json:"device_info"`
}

type CapabilityDTO struct {
	Type        string      `json:"type"`
	Retrievable bool        `json:"retrievable"`
	Reportable  bool        `json:"reportable"`
	Parameters  interface{} `json:"parameters"`
}
