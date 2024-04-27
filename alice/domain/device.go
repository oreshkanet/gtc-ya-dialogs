// https://yandex.ru/dev/dialogs/smart-home/doc/reference/get-devices-jrpc.html?lang=ru

package domain

type DeviceType string

const (
	DeviceType_Socket     DeviceType = "devices.types.socket"     // Умная розетка
	DeviceType_Switch     DeviceType = "devices.types.switch"     // Выключатель
	DeviceType_Thermostat DeviceType = "devices.types.thermostat" // Устройство с возможностью регулирования температуры
)

type DeviceCapability string

const (
	Capability_OnOff DeviceCapability = "devices.capabilities.on_off"
	Capability_Range DeviceCapability = "devices.capabilities.on_off"
)

type Device struct {
	Id           string         `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Room         string         `json:"room"`
	Type         DeviceType     `json:"type"`
	CustomData   any            `json:"custom_data"`
	Capabilities map[string]any `json:"capabilities"`
	Properties   map[string]any `json:"properties"`
	DeviceInfo   DeviceInfo     `json:"device_info"`
}

type DeviceInfo struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	HwVersion    string `json:"hw_version"`
	SwVersion    string `json:"sw_version"`
}
