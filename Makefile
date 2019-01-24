# Copyright © 2018 inwinSTACK Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

all: out/elector

TAG = v0.2.0
PREFIX = inwinstack/ipvs-elector

$(shell mkdir -p ./out)

.PHONY: out/elector
out/elector:
	CGO_ENABLED=0 GOOS=linux \
	  go build -a -installsuffix cgo -ldflags '-s -w' -o $@ cmd/main.go

.PHONY: test
test:
	./hack/test-go.sh

.PHONY: build_image
build_image: out/elector
	docker build -t $(PREFIX):$(TAG) .

.PHONY: push_image
push_image:
	docker push $(PREFIX):$(TAG)

clean:
	@rm -f server