//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package main

import (
	"flag"
	"io"
	"net/http"
	"strings"
)

var tiServer = flag.String("tuqqedin", "http://localhost:8093/default", "URL to tuqqedin")

func main() {
	flag.Parse()
	handleInteractiveMode()
}

func execute_internal(line string, w io.Writer) error {

	url := *tiServer + "_query"
	resp, err := http.Post(url, "text/plain", strings.NewReader(line))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)

	return nil
}
