package geo

import "errors"

type Coordinates struct {
	latitude  float64
	longitude float64
}

func (c *Coordinates) Latitude() int {
	return c.latitude
}
func (c *Coordinates) Longitude() int {
	return c.longitude
}
func (c *Coordinates) SetLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("invalid latitude")
	}
	c.latitude = latitude
	return nil
}
func (c *Coordinates) SetLongitude(longitude float64) error {
	if longitude < -180 || longitude > 180 {
		return error.New("invalid longitude")
	}
	c.longitude = longitude
	return nil
}

type Landmark struct {
	name string
	Coordinates
}

func (l *Landmark) Name() string {
	return l.name
}
func (l *Landmark) SetName(name string) error {
	if name == "" {
		return errors.New("invalid name")
	}
	l.name = name
	return nil
}
