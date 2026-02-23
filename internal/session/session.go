package session

import (
	"encoding/gob"
	"net/http"
	"nexample/internal/config"
	"nexample/internal/database"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

type Data struct {
	AccountID uint
	Email     string
	Active    bool
}

func init() {
	gob.Register(Data{})
}

func Init() {
	cfg := config.Load()
	store = sessions.NewCookieStore(
		[]byte(cfg.SessionAuthKey),
		[]byte(cfg.SessionEncKey),
	)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60,
		HttpOnly: true,
		Secure:   !config.IsDev(),
		SameSite: http.SameSiteLaxMode,
	}
}

func Get(r *http.Request) database.Account {
	if store == nil {
		return database.Account{}
	}

	sess, err := store.Get(r, config.Load().CookieName)
	if err != nil {
		return database.Account{}
	}

	raw, ok := sess.Values["data"]
	if !ok {
		return database.Account{}
	}

	data, ok := raw.(Data)
	if !ok {
		return database.Account{}
	}

	acc := database.Account{
		Email:  data.Email,
		Active: data.Active,
	}
	acc.ID = data.AccountID
	return acc
}

func Set(w http.ResponseWriter, r *http.Request, acc database.Account) error {
	sess, _ := store.Get(r, config.Load().CookieName)
	sess.Values["data"] = Data{
		AccountID: acc.ID,
		Email:     acc.Email,
		Active:    acc.Active,
	}
	return sess.Save(r, w)
}

func Destroy(w http.ResponseWriter, r *http.Request) error {
	sess, _ := store.Get(r, config.Load().CookieName)
	sess.Options.MaxAge = -1
	return sess.Save(r, w)
}
