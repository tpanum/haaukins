// Copyright (c) 2018-2019 Aalborg University
// Use of this source code is governed by a GPLv3
// license that can be found in the LICENSE file.

package lab

import (
	"context"
	"math/rand"
	"time"

	"io"
	"sync"

	"github.com/tpanum/haaukins/exercise"
	"github.com/tpanum/haaukins/store"
	"github.com/tpanum/haaukins/virtual"
	"github.com/tpanum/haaukins/virtual/docker"
	"github.com/tpanum/haaukins/virtual/vbox"
	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/rs/zerolog/log"
)

var (
	newEnvironment = exercise.NewEnvironment
)

type Config struct {
	Frontends []store.InstanceConfig
	Exercises []store.Exercise
}

func (conf Config) Flags() []store.FlagConfig {
	var res []store.FlagConfig
	for _, exercise := range conf.Exercises {
		res = append(res, exercise.Flags()...)
	}
	return res
}

type Creator interface {
	NewLab(context.Context) (Lab, error)
}

type LabHost struct {
	Vlib vbox.Library
	Conf Config
}

func (lh *LabHost) NewLab(ctx context.Context) (Lab, error) {
	env := newEnvironment(lh.Vlib)
	if err := env.Create(ctx); err != nil {
		return nil, err
	}

	if err := env.Add(ctx, lh.Conf.Exercises...); err != nil {
		return nil, err
	}

	dockerHost := docker.NewHost()
	l := &lab{
		tag:         generateTag(),
		lib:         lh.Vlib,
		environment: env,
		dockerHost:  dockerHost,
		frontends:   map[uint]frontendConf{},
	}

	for _, f := range lh.Conf.Frontends {
		port := virtual.GetAvailablePort()
		if _, err := l.addFrontend(ctx, f, port); err != nil {
			return nil, err
		}
	}

	return l, nil
}

type Lab interface {
	Start(context.Context) error
	Stop() error
	Restart(context.Context) error
	Environment() exercise.Environment
	ResetFrontends(ctx context.Context) error
	RdpConnPorts() []uint
	Tag() string
	InstanceInfo() []virtual.InstanceInfo
	io.Closer
}

type lab struct {
	tag         string
	lib         vbox.Library
	environment exercise.Environment
	frontends   map[uint]frontendConf
	dockerHost  docker.Host
}

type frontendConf struct {
	vm   vbox.VM
	conf store.InstanceConfig
}

func (l *lab) addFrontend(ctx context.Context, conf store.InstanceConfig, rdpPort uint) (vbox.VM, error) {
	hostIp, err := l.dockerHost.GetDockerHostIP()
	if err != nil {
		return nil, err
	}

	vm, err := l.lib.GetCopy(ctx,
		conf,
		vbox.SetBridge(l.environment.NetworkInterface()),
		vbox.SetLocalRDP(hostIp, rdpPort),
		vbox.SetRAM(conf.MemoryMB),
	)
	if err != nil {
		return nil, err
	}

	l.frontends[rdpPort] = frontendConf{
		vm:   vm,
		conf: conf,
	}

	log.Debug().Msgf("Created lab frontend on port %d", rdpPort)

	return vm, nil
}

func (l *lab) Environment() exercise.Environment {
	return l.environment
}

func (l *lab) ResetFrontends(ctx context.Context) error {
	var errs []error
	for p, vmConf := range l.frontends {
		err := vmConf.vm.Close()
		if err != nil {
			errs = append(errs, err)
			continue
		}

		vm, err := l.addFrontend(ctx, vmConf.conf, p)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = vm.Start(ctx)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

func (l *lab) Start(ctx context.Context) error {
	if err := l.environment.Start(ctx); err != nil {
		return err
	}

	for _, fconf := range l.frontends {
		if err := fconf.vm.Start(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (l *lab) Stop() error {
	if err := l.environment.Stop(); err != nil {
		return err
	}

	for _, fconf := range l.frontends {
		if err := fconf.vm.Stop(); err != nil {
			return err
		}
	}

	return nil
}

func (l *lab) Restart(ctx context.Context) error {
	if err := l.environment.Stop(); err != nil {
		return err
	}

	if err := l.environment.Start(ctx); err != nil {
		return err
	}

	for _, fconf := range l.frontends {
		if err := fconf.vm.Stop(); err != nil {
			return err
		}

		if err := fconf.vm.Start(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (l *lab) Close() error {
	var wg sync.WaitGroup

	closer := func(c io.Closer) {
		if err := c.Close(); err != nil {
			log.Warn().Msgf("error while closing lab: %s", err)
		}
		wg.Done()
	}

	if l.environment != nil {
		wg.Add(1)
		go closer(l.environment)
	}

	for _, lab := range l.frontends {
		wg.Add(1)
		go closer(lab.vm)
	}

	wg.Wait()

	return nil
}

func (l *lab) RdpConnPorts() []uint {
	var ports []uint
	for p, _ := range l.frontends {
		ports = append(ports, p)
	}

	return ports
}

func (l *lab) Tag() string {
	return l.tag
}

func (l *lab) InstanceInfo() []virtual.InstanceInfo {
	var instances []virtual.InstanceInfo
	for _, fconf := range l.frontends {
		instances = append(instances, fconf.vm.Info())
	}
	instances = append(instances, l.environment.InstanceInfo()...)
	return instances
}

func generateTag() string {
	// seed for our GetRandomName
	rand.Seed(time.Now().UnixNano())
	return namesgenerator.GetRandomName(0)
}
