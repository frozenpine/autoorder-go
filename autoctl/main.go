// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"gitlab.quantdo.cn/yuanyang/autoorder/autoctl/cmd"
	"gitlab.quantdo.cn/yuanyang/autoorder/protocol"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *protocol.HelloRequest) (*protocol.HelloReply, error) {
	return &protocol.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	sock, _ := net.Listen("tcp", ":9000")

	s := grpc.NewServer()

	protocol.RegisterGreetServer(s, &server{})

	s.Serve(sock)

	cmd.Execute()
}
