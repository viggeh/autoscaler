/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"github.com/stretchr/testify/mock"
)

type MockHistogram struct {
	mock.Mock
}

func (m *MockHistogram) Percentile(percentile float64) float64 {
	args := m.Called(percentile)
	return args.Get(0).(float64)
}

func (m *MockHistogram) AddSample(value float64, weight float64) {
	m.Called(value, weight)
}

func (m *MockHistogram) SubtractSample(value float64, weight float64) {
	m.Called(value, weight)
}

func (m *MockHistogram) IsEmpty() bool {
	args := m.Called()
	return args.Bool(0)

}

