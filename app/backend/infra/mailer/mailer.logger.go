package mailer

import "log"

func (s *SMTPInfra) logWarn(msg string) {
	log.Printf("WARN: %s\n", msg)
}
func (s *SMTPInfra) logErr(msg string) {
	log.Printf("ERR: %s\n", msg)
}
func (s *SMTPInfra) logFatal(msg string) {
	log.Printf("ERR: %s\n", msg)
}
func (s *SMTPInfra) logInfo(msg string) {
	log.Printf("INFO: %s\n", msg)
}
