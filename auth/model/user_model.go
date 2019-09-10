package model

type AccountLoginSession struct {
	SessionKey string
	Guid       int64
}

func AccountLogin(session string, guid int64) AccountLoginSession {
	return AccountLoginSession{
		SessionKey: session,
		Guid:       guid,
	}
}
