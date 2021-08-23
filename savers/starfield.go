package savers

import (
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/vilmibm/gh-screensaver/savers/shared"
)

type StarfieldSaver struct {
	screen tcell.Screen
	style  tcell.Style
}

func NewStarfieldSaver(opts shared.ScreensaverOpts) (shared.Screensaver, error) {
	s := &StarfieldSaver{}

	if err := s.Initialize(opts); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *StarfieldSaver) Initialize(opts shared.ScreensaverOpts) error {
	s.screen = opts.Screen
	s.style = opts.Style

	rand.Seed(time.Now().UTC().UnixNano())

	// TODO calcuate spawn area
	//   - square like 10x10 roughly centered on screen

	return nil
}

func (s *StarfieldSaver) Inputs() map[string]shared.SaverInput {
	// TODO support rainbow mode
	return map[string]shared.SaverInput{}
}

func (s *StarfieldSaver) SetInputs(inputs map[string]string) error {
	return nil
}

func (s *StarfieldSaver) Update() error {
	// TODO star spawn chance
	// TODO compute movement vector
	// TODO spawn in random spot in spawn square

	// TODO how to compute vector? don't overthink it

	// TODO for each star
	//        move along vector one unit
	//        possibly change character based on proximity to wall
	//        if moved beyond walls, delete
	return nil
}
