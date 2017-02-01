package GoJob

import (
	"log"
)

type Logger struct {
	DEBUG_ENABLED bool
	INFORMATION_ENABLED bool;
}

var logger = Logger{DEBUG_ENABLED:true,INFORMATION_ENABLED:true};

func (logger * Logger)  Infof(format string, p ...interface{}) {
	if logger.INFORMATION_ENABLED {
		log.Printf(format,p...);
	}
}

func (logger * Logger) Debugf(format string, p ...interface{})  {
	if logger.INFORMATION_ENABLED {
		log.Printf(format,p...);
	}
}

func (logger * Logger) Info(description string) {
	if logger.INFORMATION_ENABLED {
		log.Println(description);
	}
}

func (logger * Logger) Debug(description string) {
	if logger.DEBUG_ENABLED {
		log.Println(description);
	}
}
