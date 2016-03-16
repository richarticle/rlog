package rlog

import (
	"testing"
)

func TestSyslog(t *testing.T) {
	log, err := NewLogger("syslog", "app", WARN, 0)
	if err != nil {
		t.Error(err)
	}

	log.Emergf("Emergency %d", 0)
	log.Alertf("Alert %d", 1)
	log.Critf("Critical %d", 2)
	log.Errorf("Error %d", 3)
	log.Warnf("Warning %d", 4)
	log.Noticef("Notice %d", 5)
	log.Infof("Information %d", 6)
	log.Debugf("Debug %d", 7)

	log.Emergln("Emergency", 0)
	log.Alertln("Alert", 1)
	log.Critln("Critical", 2)
	log.Errorln("Error", 3)
	log.Warnln("Warning", 4)
	log.Noticeln("Notice", 5)
	log.Infoln("Information", 6)
	log.Debugln("Debug", 7)
}

func TestStdout(t *testing.T) {
	log, err := NewLogger("stdout", "app", NOTICE, PREFIX)
	if err != nil {
		t.Error(err)
	}

	log.Emergf("Emergency %d", 0)
	log.Alertf("Alert %d", 1)
	log.Critf("Critical %d", 2)
	log.Errorf("Error %d", 3)
	log.Warnf("Warning %d", 4)
	log.Noticef("Notice %d", 5)
	log.Infof("Information %d", 6)
	log.Debugf("Debug %d", 7)

	log.Emergln("Emergency", 0)
	log.Alertln("Alert", 1)
	log.Critln("Critical", 2)
	log.Errorln("Error", 3)
	log.Warnln("Warning", 4)
	log.Noticeln("Notice", 5)
	log.Infoln("Information", 6)
	log.Debugln("Debug", 7)
}

func TestStderr(t *testing.T) {
	log, err := NewLogger("stderr", "app", DEBUG, PREFIX)
	if err != nil {
		t.Error(err)
	}

	log.Emergf("Emergency %d", 0)
	log.Alertf("Alert %d", 1)
	log.Critf("Critical %d", 2)
	log.Errorf("Error %d", 3)
	log.Warnf("Warning %d", 4)
	log.Noticef("Notice %d", 5)
	log.Infof("Information %d", 6)
	log.Debugf("Debug %d", 7)

	log.Emergln("Emergency", 0)
	log.Alertln("Alert", 1)
	log.Critln("Critical", 2)
	log.Errorln("Error", 3)
	log.Warnln("Warning", 4)
	log.Noticeln("Notice", 5)
	log.Infoln("Information", 6)
	log.Debugln("Debug", 7)
}
