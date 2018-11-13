package led

import (
	"github.com/tarm/serial"
)

type LED struct {
	p *serial.Port
}

func NewLED(port string) (*LED, error) {

	c := &serial.Config{Name: port, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	return &LED{p: s}, nil
}

func (l *LED) Send(R, G, B uint8) error {
	var res []byte
	res = append(res, ToByte(R))
	res = append(res, ToByte(G))
	res = append(res, ToByte(B))
	_, err := l.p.Write(res)
	return err
}

func ToByte(i uint8) byte {
	return byte(i)
}
