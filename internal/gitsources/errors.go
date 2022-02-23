// Copyright 2019 Sorint.lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
// See the License for the specific language governing permissions and
// limitations under the License.

package gitsource

import (
	"net/http"

	errors "golang.org/x/xerrors"
)

//var ErrUnauthorized = errors.New("unauthorized")
//var ErrForbidden = errors.New("forbidden")
//var ErrNotExist = errors.New("not exists")
//var ErrUnprocessableEntity = errors.New("unprocessable entity")
//var ErrServiceUnavailable = errors.New("service unavailable")

//
type ErrUnauthorized struct {
	Err error
}

func (e *ErrUnauthorized) Error() string {
	return e.Err.Error()
}

func NewErrUnauthorized(err error) *ErrUnauthorized {
	return &ErrUnauthorized{Err: err}
}

func (*ErrUnauthorized) Is(err error) bool {
	_, ok := err.(*ErrUnauthorized)
	return ok
}

func IsErrUnauthorized(err error) bool {
	return errors.Is(err, &ErrUnauthorized{})
}

//
type ErrForbidden struct {
	Err error
}

func (e *ErrForbidden) Error() string {
	return e.Err.Error()
}

func NewErrForbidden(err error) *ErrForbidden {
	return &ErrForbidden{Err: err}
}

func (*ErrForbidden) Is(err error) bool {
	_, ok := err.(*ErrForbidden)
	return ok
}

func IsErrForbidden(err error) bool {
	return errors.Is(err, &ErrForbidden{})
}

//
type ErrNotExist struct {
	Err error
}

func (e *ErrNotExist) Error() string {
	return e.Err.Error()
}

func NewErrNotExist(err error) *ErrNotExist {
	return &ErrNotExist{Err: err}
}

func (*ErrNotExist) Is(err error) bool {
	_, ok := err.(*ErrNotExist)
	return ok
}

func IsErrNotExist(err error) bool {
	return errors.Is(err, &ErrNotExist{})
}

//
type ErrUnprocessableEntity struct {
	Err error
}

func (e *ErrUnprocessableEntity) Error() string {
	return e.Err.Error()
}

func NewErrUnprocessableEntity(err error) *ErrUnprocessableEntity {
	return &ErrUnprocessableEntity{Err: err}
}

func (*ErrUnprocessableEntity) Is(err error) bool {
	_, ok := err.(*ErrUnprocessableEntity)
	return ok
}

func IsErrUnprocessableEntity(err error) bool {
	return errors.Is(err, &ErrUnprocessableEntity{})
}

//
type ErrServiceUnavailable struct {
	Err error
}

func (e *ErrServiceUnavailable) Error() string {
	return e.Err.Error()
}

func NewErrServiceUnavailable(err error) *ErrServiceUnavailable {
	return &ErrServiceUnavailable{Err: err}
}

func (*ErrServiceUnavailable) Is(err error) bool {
	_, ok := err.(*ErrServiceUnavailable)
	return ok
}

func IsErrServiceUnavailable(err error) bool {
	return errors.Is(err, &ErrServiceUnavailable{})
}

func GitSourceError(statusCode int, err error) error {
	switch statusCode {
	case http.StatusUnauthorized:
		return NewErrUnauthorized(err)
	case http.StatusForbidden:
		return NewErrForbidden(err)
	case http.StatusNotFound:
		return NewErrNotExist(err)
	case http.StatusUnprocessableEntity:
		return NewErrUnprocessableEntity(err)
	case http.StatusServiceUnavailable:
		return NewErrServiceUnavailable(err)
	}

	if statusCode/100 != 2 {
		return errors.Errorf("gitsource api status code %d", statusCode)
	}

	return err
}
