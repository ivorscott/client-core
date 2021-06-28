// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	memberships "github.com/devpies/devpie-client-core/users/domain/memberships"
	database "github.com/devpies/devpie-client-core/users/platform/database"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MembershipQuerier is an autogenerated mock type for the MembershipQuerier type
type MembershipQuerier struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, repo, nm, now
func (_m *MembershipQuerier) Create(ctx context.Context, repo database.Storer, nm memberships.NewMembership, now time.Time) (memberships.Membership, error) {
	ret := _m.Called(ctx, repo, nm, now)

	var r0 memberships.Membership
	if rf, ok := ret.Get(0).(func(context.Context, database.Storer, memberships.NewMembership, time.Time) memberships.Membership); ok {
		r0 = rf(ctx, repo, nm, now)
	} else {
		r0 = ret.Get(0).(memberships.Membership)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Storer, memberships.NewMembership, time.Time) error); ok {
		r1 = rf(ctx, repo, nm, now)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, repo, tid, uid
func (_m *MembershipQuerier) Delete(ctx context.Context, repo database.Storer, tid string, uid string) (string, error) {
	ret := _m.Called(ctx, repo, tid, uid)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, database.Storer, string, string) string); ok {
		r0 = rf(ctx, repo, tid, uid)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Storer, string, string) error); ok {
		r1 = rf(ctx, repo, tid, uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveMembership provides a mock function with given fields: ctx, repo, uid, tid
func (_m *MembershipQuerier) RetrieveMembership(ctx context.Context, repo database.Storer, uid string, tid string) (memberships.Membership, error) {
	ret := _m.Called(ctx, repo, uid, tid)

	var r0 memberships.Membership
	if rf, ok := ret.Get(0).(func(context.Context, database.Storer, string, string) memberships.Membership); ok {
		r0 = rf(ctx, repo, uid, tid)
	} else {
		r0 = ret.Get(0).(memberships.Membership)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Storer, string, string) error); ok {
		r1 = rf(ctx, repo, uid, tid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveMemberships provides a mock function with given fields: ctx, repo, uid, tid
func (_m *MembershipQuerier) RetrieveMemberships(ctx context.Context, repo database.Storer, uid string, tid string) ([]memberships.MembershipEnhanced, error) {
	ret := _m.Called(ctx, repo, uid, tid)

	var r0 []memberships.MembershipEnhanced
	if rf, ok := ret.Get(0).(func(context.Context, database.Storer, string, string) []memberships.MembershipEnhanced); ok {
		r0 = rf(ctx, repo, uid, tid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]memberships.MembershipEnhanced)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, database.Storer, string, string) error); ok {
		r1 = rf(ctx, repo, uid, tid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, repo, tid, update, uid, now
func (_m *MembershipQuerier) Update(ctx context.Context, repo database.Storer, tid string, update memberships.UpdateMembership, uid string, now time.Time) error {
	ret := _m.Called(ctx, repo, tid, update, uid, now)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Storer, string, memberships.UpdateMembership, string, time.Time) error); ok {
		r0 = rf(ctx, repo, tid, update, uid, now)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}