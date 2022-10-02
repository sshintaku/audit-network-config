package main

import (
	"github.com/sshintaku/prisma_session"
)

func main() {
	var params = prisma_session.ReadParameters()
	session := prisma_session.Session{ApiUrl: params.ApiUrl}
	session.CreateSession()
	session.GetAuditEvents()
}
