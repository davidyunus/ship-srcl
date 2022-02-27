package cruiseship

import "github.com/ship-srcl/src/domain"

type cruiseShip struct {
}

func NewCruiseShip() domain.Shipping {
	return &cruiseShip{}
}

func (c *cruiseShip) StartEngine() (*domain.Ship, error) {
	// do something
	return nil, nil
}

func (c *cruiseShip) SpeedUp() error {
	// do something
	return nil
}

func (c *cruiseShip) Break() error {
	// do something
	return nil
}
