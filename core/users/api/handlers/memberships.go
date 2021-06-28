package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devpies/devpie-client-core/users/domain/memberships"
	"github.com/devpies/devpie-client-core/users/platform/auth0"
	"github.com/devpies/devpie-client-core/users/platform/database"
	"github.com/devpies/devpie-client-core/users/platform/web"
	"github.com/devpies/devpie-client-events/go/events"
	"github.com/go-chi/chi"
)

type Membership struct {
	repo  database.Storer
	log   *log.Logger
	auth0 *auth0.Auth0
	nats  *events.Client
	query MembershipQueries
}

type MembershipQueries struct {
	membership memberships.MembershipQuerier
}

func (m *Membership) RetrieveMembers(w http.ResponseWriter, r *http.Request) error {
	uid := m.auth0.UserByID(r.Context())
	tid := chi.URLParam(r, "tid")

	ms, err := m.query.membership.RetrieveMemberships(r.Context(), m.repo, uid, tid)
	if err != nil {
		switch err {
		case memberships.ErrInvalidID:
			return web.NewRequestError(err, http.StatusBadRequest)
		case memberships.ErrNotFound:
			return web.NewRequestError(err, http.StatusNotFound)
		default:
		}
		return fmt.Errorf("failed to retrieve memberships: %w", err)
	}

	return web.Respond(r.Context(), w, ms, http.StatusOK)
}
