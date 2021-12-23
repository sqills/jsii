package friendlyRandom

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/aws/jsii/jsii-calc/go/scopejsiicalclib"
)

type SubclassFriendlyRandom struct {
	scopejsiicalclib.Number
	next float64
}

func NewSubclass() *SubclassFriendlyRandom {
	s := &SubclassFriendlyRandom{next: 100}
	scopejsiicalclib.NewNumber_Override(s, jsii.Number(0))
	return s
}

func (s *SubclassFriendlyRandom) Hello() jsii.String {
	return jsii.String("SubclassNativeFriendlyRandom")
}

func (s *SubclassFriendlyRandom) Next() jsii.Number {
	defer func() { s.next += 100 }()
	return jsii.Number(s.next)
}

type PureFriendlyRandom struct {
	next float64
}

func NewPure() *PureFriendlyRandom {
	return &PureFriendlyRandom{next: 1000}
}

func (p *PureFriendlyRandom) Hello() jsii.String {
	return jsii.String("I am a native!")
}

func (p *PureFriendlyRandom) Next() jsii.Number {
	defer func() { p.next += 1000 }()
	return jsii.Number(p.next)
}
