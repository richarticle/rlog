# rlog
This is a simple logging package which can output logs to syslog, stdout, or stderr.

The definition of log level follows the spec of syslog.
- Emergency: 0
- Alert: 1
- Critical: 2
- Error: 3
- Warning: 4
- Notice: 5
- Information: 6
- Debug: 7

# Usage
First, you need to import this package

	import "github.com/richarticle/rlog"


Then, you need to create a Logger instance, and specify the settings.

	log, err := rlog.NewLogger("syslog", "app", 5, 0)
	log, err := rlog.NewLogger("stdout", "app", 7, rlog.PREFIX)

After successfully creating a Logger instance, you can use the following function to log messages.

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
