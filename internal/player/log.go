/*
 *   Copyright 2020 Dmitry Kann
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package player

import (
	"errors"
	"fmt"
	"github.com/op/go-logging"
	"github.com/yktoo/ymuse/internal/util"
)

// Package-wide Logger instance
var log *logging.Logger

func init() {
	// Init logging
	log = logging.MustGetLogger("player")
	logging.SetLevel(util.GetConfig().LogLevel, "player")
}

// errCheck() logs a warning if the error is not nil.
func errCheck(err error, message string) bool {
	if err != nil {
		log.Warning(errors.New(fmt.Sprintf("%v: %v", message, err)))
		return true
	}
	return false
}
