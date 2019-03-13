
# Cast
[![Circle CI](https://circleci.com/gh/datasweet/cast.svg?style=svg)](https://circleci.com/gh/datasweet/cast) [![Go Report Card](https://goreportcard.com/badge/github.com/datasweet/cast)](https://goreportcard.com/report/github.com/datasweet/cast) [![GoDoc](https://godoc.org/github.com/datasweet/cast?status.png)](https://godoc.org/github.com/datasweet/cast) [![GitHub stars](https://img.shields.io/github/stars/datasweet/cast.svg)](https://github.com/datasweet/cast/stargazers)
[![GitHub license](https://img.shields.io/github/license/datasweet/cast.svg)](https://github.com/datasweet/cast/blob/master/LICENSE)

[![datasweet-logo](https://www.datasweet.fr/wp-content/uploads/2019/02/datasweet-black.png)](http://www.datasweet.fr)

Cast is a Go package to simply cast an interface to a specific type in go, without panic ! 

## Installation
```go
go get github.com/datasweet/cast
```

## Usage

### Parsing a new json
```
import (
  "fmt"
  "github.com/datasweet/cast"
)

func main() {
  j := cast.AsDatetime("13/03/2019")
  fmt.Println(j)
}
```

## Who are we ?
We are Datasweet, a french startup providing full service (big) data solutions.

## Questions ? problems ? suggestions ?
If you find a bug or want to request a feature, please create a [GitHub Issue](https://github.com/datasweet/cast/issues/new).

## License
```
This software is licensed under the Apache License, version 2 ("ALv2"), quoted below.

Copyright 2019 Datasweet <http://www.datasweet.fr>

Licensed under the Apache License, Version 2.0 (the "License"); you may not
use this file except in compliance with the License. You may obtain a copy of
the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
License for the specific language governing permissions and limitations under
the License.
```
