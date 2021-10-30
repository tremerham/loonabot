package runner

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"sync"

	pb "github.com/parthpower/loonabot/cmd/runner/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type runnerServer struct {
	pb.UnimplementedRunnerServer
	exec []*pb.UpdateRequest
	m    sync.Mutex
	proc *os.Process
}

func Start(host string, tlsCredentials credentials.TransportCredentials) error {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	server := grpc.NewServer(grpc.Creds(tlsCredentials))
	pb.RegisterRunnerServer(server, &runnerServer{})
	reflection.Register(server)
	if err = server.Serve(listener); err != nil {
		return err
	}
	return nil
}

func (r *runnerServer) Run(ctx context.Context, req *pb.RunRequest) (*pb.RunResponse, error) {
	if !req.Restart {
		return &pb.RunResponse{}, nil
	}
	if r.exec != nil {
		go r.runExec()
	}
	return &pb.RunResponse{Err: ""}, nil
}

func (r *runnerServer) Update(stream pb.Runner_UpdateServer) error {
	r.m.Lock()
	defer r.m.Unlock()
	r.exec = []*pb.UpdateRequest{}
	fmt.Printf("rx:")
	for {
		rd, err := stream.Recv()
		if err == io.EOF {
			if rd != nil {
				r.exec = append(r.exec, rd)
			}
			// stream.Send(&pb.UpdateResponse{Err: ""})
			break
		}
		if err != nil {
			r.exec = nil
			return err
		}
		r.exec = append(r.exec, rd)
		fmt.Printf(".")
		err = stream.Send(&pb.UpdateResponse{})
		if err != nil {
			r.exec = nil
			return err
		}
	}
	fmt.Println()
	return nil
}

func (r *runnerServer) runExec() error {
	if r.proc != nil {
		r.proc.Kill()
	}
	path, err := r.writeToFile()
	if err != nil {
		return err
	}
	fmt.Println(path)
	c := exec.Command(path, r.exec[0].Md.Args...)
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	fmt.Println("starting cmd:", c)
	err = c.Start()
	if err != nil {
		return err
	}

	r.proc = c.Process
	return nil
}

func (r *runnerServer) writeToFile() (string, error) {
	if r.exec == nil {
		return "", fmt.Errorf("exec protobuf empty")
	}

	execfile, err := ioutil.TempFile("", "exec")
	if err != nil {
		return "", err
	}
	defer execfile.Close()
	fullpath := execfile.Name()
	err = execfile.Chmod(0755)
	if err != nil {
		os.Remove(fullpath)
		return "", err
	}
	r.m.Lock()
	defer r.m.Unlock()
	for _, chunk := range r.exec {
		_, err := execfile.Write(chunk.GetBuf())
		if err != nil {
			os.Remove(fullpath)
			return "", err
		}
	}

	return fullpath, nil
}
