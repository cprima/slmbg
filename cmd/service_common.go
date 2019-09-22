package cmd

import (
	"log"
	"time"

	"github.com/kardianos/service"
	"github.com/spf13/viper"
)

// var logger service.Logger

// @see https://github.com/kardianos/service/blob/master/example/logging/main.go
// Program structures.
//  Define Start and Stop methods.
type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	if service.Interactive() {
		// logger.Info("Running in terminal.")
	} else {
		// logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() error {
	// logger.Infof("I'm running %v.", service.Platform())
	tick()
	// systray.SetTooltip("Last run: " + viper.GetString("last_run"))
	ticker := time.NewTicker(time.Duration(viper.GetInt64("interval") * int64(time.Second)))
	for {
		select {
		case <-ticker.C:
			// case tm := <-ticker.C:
			// logger.Infof("Still running at %v...", tm)
			tick()
			// systray.SetTooltip("Last run: " + viper.GetString("last_run"))
			// systray.SetTooltip("Last run: " + viper.GetString("last_run"))
			// logger.Infof("Still running at %v...", tm)
		case <-p.exit:
			ticker.Stop()
			// systray.Quit()
			// logger.Infof("Exiting in run at case p.exit")
			return nil
		}
	}
}

func (p *program) Stop(s service.Service) error {
	// Any work in Stop should be quick, usually a few seconds at most.
	// logger.Info("I'm Stopping!")
	close(p.exit)
	return nil
}

func serviceHelper(control string) {
	// fmt.Println("myHelperTest called: ", control)

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	svcConfig := &service.Config{
		Name:        "Sunlightmap",
		DisplayName: "Sunlightmap Service",
		Description: "Generates a day- and night map of the planet.",
		//Dependencies: []string{
		//	"Requires=network.target",
		//	"After=network-online.target syslog.target"},
		Option: options,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	// logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if control != "service" {
		err = service.Control(s, control)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		// logger.Error(err)
	}

}
