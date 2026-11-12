package mailer

type MailReport struct {
	Accepted []string
	Failed   []Report
}

type Report struct {
	RcptMail string
	Err      error
}
