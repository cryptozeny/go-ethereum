// Copyright 2017 The ethereumprogpow Authors
// This file is part of the ethereumprogpow library.
//
// The ethereumprogpow library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The ethereumprogpow library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the ethereumprogpow library. If not, see <http://www.gnu.org/licenses/>.

package accounts

import (
	"testing"
)

func TestURLParsing(t *testing.T) {
	url, err := parseURL("https://ethereumprogpow.com")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "ethereumprogpow.com" {
		t.Errorf("expected: %v, got: %v", "ethereumprogpow.com", url.Path)
	}

	_, err = parseURL("ethereumprogpow.com")
	if err == nil {
		t.Error("expected err, got: nil")
	}
}

func TestURLString(t *testing.T) {
	url := URL{Scheme: "https", Path: "ethereumprogpow.com"}
	if url.String() != "https://ethereumprogpow.com" {
		t.Errorf("expected: %v, got: %v", "https://ethereumprogpow.com", url.String())
	}

	url = URL{Scheme: "", Path: "ethereumprogpow.com"}
	if url.String() != "ethereumprogpow.com" {
		t.Errorf("expected: %v, got: %v", "ethereumprogpow.com", url.String())
	}
}

func TestURLMarshalJSON(t *testing.T) {
	url := URL{Scheme: "https", Path: "ethereumprogpow.com"}
	json, err := url.MarshalJSON()
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if string(json) != "\"https://ethereumprogpow.com\"" {
		t.Errorf("expected: %v, got: %v", "\"https://ethereumprogpow.com\"", string(json))
	}
}

func TestURLUnmarshalJSON(t *testing.T) {
	url := &URL{}
	err := url.UnmarshalJSON([]byte("\"https://ethereumprogpow.com\""))
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "ethereumprogpow.com" {
		t.Errorf("expected: %v, got: %v", "https", url.Path)
	}
}

func TestURLComparison(t *testing.T) {
	tests := []struct {
		urlA   URL
		urlB   URL
		expect int
	}{
		{URL{"https", "ethereumprogpow.com"}, URL{"https", "ethereumprogpow.com"}, 0},
		{URL{"http", "ethereumprogpow.com"}, URL{"https", "ethereumprogpow.com"}, -1},
		{URL{"https", "ethereumprogpow.com/a"}, URL{"https", "ethereumprogpow.com"}, 1},
		{URL{"https", "abc.org"}, URL{"https", "ethereumprogpow.com"}, -1},
	}

	for i, tt := range tests {
		result := tt.urlA.Cmp(tt.urlB)
		if result != tt.expect {
			t.Errorf("test %d: cmp mismatch: expected: %d, got: %d", i, tt.expect, result)
		}
	}
}
