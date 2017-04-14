package types

const (
	// DeviceCLI is the device type for the cli tool
	DeviceCLI = iota
	// DeviceWeb is the device type for the web interface
	DeviceWeb
)

// Session is the type we use to store client sessions
type Session struct {
	ID         int
	UID        int
	User       *User
	Token      string `json:"token"`
	DeviceType int
}

// NewSession creates a new session object
func NewSession(uid int, token, deviceType string) *Session {
	return &Session{
		UID:        uid,
		Token:      token,
		User:       nil,
		DeviceType: GetDeviceType(deviceType),
	}
}

// GetDeviceType takes a string and returns the API's idiomatic type for a device
func GetDeviceType(deviceType string) int {
	switch deviceType {
	case "cli":
		return DeviceCLI
	default:
		return DeviceWeb
	}
}
