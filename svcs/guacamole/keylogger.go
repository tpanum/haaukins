package guacamole

import (
	"bytes"
	"github.com/aau-network-security/go-ntp/store"
	"github.com/aau-network-security/go-ntp/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

var (
	keyOpcode   = []byte("3.key")
	mouseOpcode = []byte("5.mouse")

	KeyPressed = func(kf *KeyFrame) bool {
		return kf.Pressed == "1"
	}
	MouseClicked = func(mf *MouseFrame) bool {
		return mf.Button == "1" || mf.Button == "4" || mf.Button == "2"
	}
)

type KeyFrameFilterCondition func(*KeyFrame) bool

type KeyFrameFilter interface {
	Filter(RawFrame) (*KeyFrame, bool, error)
}

type keyFrameFilter struct {
	conditions []KeyFrameFilterCondition
}

func (kff *keyFrameFilter) Filter(rawFrame RawFrame) (*KeyFrame, bool, error) {
	if len(rawFrame) < len(keyOpcode) {
		return nil, false, nil
	}
	h := []byte(rawFrame)[:len(keyOpcode)]
	if bytes.Compare(keyOpcode, h) != 0 {
		return nil, false, nil
	}
	f, err := NewFrame(rawFrame)
	if err != nil {
		return nil, false, err
	}
	kf, err := NewKeyFrame(f)
	if err != nil {
		return nil, false, err
	}

	for _, c := range kff.conditions {
		if !c(kf) {
			return nil, false, nil
		}
	}
	return kf, true, nil
}

func NewKeyFrameFilter(conditions ...KeyFrameFilterCondition) KeyFrameFilter {
	return &keyFrameFilter{
		conditions: conditions,
	}
}

type MouseFrameFilterCondition func(*MouseFrame) bool

type MouseFrameFilter interface {
	Filter(RawFrame) (*MouseFrame, bool, error)
}

type mouseFrameFilter struct {
	conditions []MouseFrameFilterCondition
}

func (mff *mouseFrameFilter) Filter(rawFrame RawFrame) (*MouseFrame, bool, error) {
	if len(rawFrame) < len(mouseOpcode) {
		return nil, false, nil
	}
	h := []byte(rawFrame)[:len(mouseOpcode)]
	if bytes.Compare(mouseOpcode, h) != 0 {
		return nil, false, nil
	}
	f, err := NewFrame(rawFrame)
	if err != nil {
		return nil, false, err
	}
	mf, err := NewMouseFrame(f)
	if err != nil {
		return nil, false, err
	}
	for _, c := range mff.conditions {
		if !c(mf) {
			return nil, false, nil
		}
	}

	return mf, true, nil
}

func NewMouseFrameFilter(conditions ...MouseFrameFilterCondition) MouseFrameFilter {
	return &mouseFrameFilter{
		conditions: conditions,
	}
}

type logEvent struct {
	timestamp time.Time
	rawFrame  RawFrame
}

type KeyLogger interface {
	Log(rm RawFrame)
}

type keyLogger struct {
	ch     chan logEvent
	logger *zerolog.Logger
	kff    KeyFrameFilter
	mff    MouseFrameFilter
}

func (k keyLogger) run() {
	for {
		event := <-k.ch

		kf, ok, err := k.kff.Filter(event.rawFrame)
		if err != nil {
			log.Warn().Msgf("Failed to filter raw message: %s", err)
		} else if ok {
			k.logger.Log().
				Time("t", event.timestamp).
				Str("k", string(kf.Key)).
				Str("p", string(kf.Pressed)).
				Msg("key")
			continue
		}

		mf, ok, err := k.mff.Filter(event.rawFrame)
		if err != nil {
			log.Warn().Msgf("Failed to filter raw message: %s", err)
		} else if ok {
			k.logger.Log().
				Time("t", event.timestamp).
				Str("x", string(mf.X)).
				Str("y", string(mf.Y)).
				Str("b", string(mf.Button)).
				Msg("mouse")
		}
	}
}

func (k keyLogger) Log(rawFrame RawFrame) {
	timestamp := time.Now()

	k.ch <- logEvent{
		timestamp: timestamp,
		rawFrame:  rawFrame,
	}
}

func NewKeyLogger(logger *zerolog.Logger) (KeyLogger, error) {
	c := make(chan logEvent)
	kff := NewKeyFrameFilter(KeyPressed)
	mff := NewMouseFrameFilter(MouseClicked)

	kl := keyLogger{
		ch:     c,
		logger: logger,
		kff:    kff,
		mff:    mff,
	}
	go kl.run()
	return kl, nil
}

type KeyLoggerPool interface {
	GetLogger(t store.Team) (KeyLogger, error)
}

type keyLoggerPool struct {
	logpool util.LogPool
}

func (klp keyLoggerPool) GetLogger(t store.Team) (KeyLogger, error) {
	logger, err := klp.logpool.GetLogger(t.Id)
	if err != nil {
		return nil, err
	}
	return NewKeyLogger(logger)
}

func NewKeyLoggerPool(dir string) (KeyLoggerPool, error) {
	logpool, err := util.NewLogPool(dir)
	if err != nil {
		return nil, err
	}

	return keyLoggerPool{
		logpool: logpool,
	}, nil
}
