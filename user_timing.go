/* Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gameasure

import (
	"fmt"
	"net/url"
	"time"
)

// UserTiming sends a timing hit type.
type UserTiming struct {
	// Category is the timing category. e.g. jsonLoader
	Category string `ga:"utc"`
	// Variable is the timing variable. e.g. load
	Variable string `ga:"utv"`
	// Time is the time it took in milliseconds.
	Time time.Duration `ga:"utt"`
	// Label is the timing label. e.g jQuery
	Label string `ga:"utl"`

	startTime time.Time
}

// Begin starts the timer.
func (u *UserTiming) Begin() {
	u.startTime = time.Now()
}

// Calculate the timing
func (u *UserTiming) End() {
	u.Time = time.Since(u.startTime)
}

// UserTiming sends user timings to Google Analytics
func (g *GA) UserTiming(e UserTiming) error {
	data := url.Values{}

	data.Add("t", "timing")
	data.Add("utc", e.Category)
	data.Add("utv", e.Variable)

	data.Add("utt", fmt.Sprintf("%d", int64(e.Time.Seconds()*1000)))
	data.Add("utl", e.Label)

	return g.send(data)
}
