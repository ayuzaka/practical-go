package chapter02

type SensorData struct {
	SensorType string
	ModelID    string
	Value      float32
}

// BAD
func ReadValue(r SensorData) float32 {
	if r.SensorType == "Fahrenheit" {
		return (r.Value * 9 / 5) + 32
	}

	return r.Value
}

// GOOD
func (r SensorData) ReadValue() float32 {
	if r.SensorType == "Fahrenheit" {
		return (r.Value * 9 / 5) + 32
	}

	return r.Value
}
