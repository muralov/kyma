// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"
import mock "github.com/stretchr/testify/mock"

// apiPackageCredentialsRemover is an autogenerated mock type for the apiPackageCredentialsRemover type
type apiPackageCredentialsRemover struct {
	mock.Mock
}

// EnsureAPIPackageCredentialsDeleted provides a mock function with given fields: ctx, appID, pkgID, instanceID
func (_m *apiPackageCredentialsRemover) EnsureAPIPackageCredentialsDeleted(ctx context.Context, appID string, pkgID string, instanceID string) error {
	ret := _m.Called(ctx, appID, pkgID, instanceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, appID, pkgID, instanceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
