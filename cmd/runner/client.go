package runner

import (
	"context"
	"fmt"
	"io"
	"os"

	pb "github.com/parthpower/loonabot/cmd/runner/api"
	"google.golang.org/grpc"
)

func UploadExec(dial *grpc.ClientConn, path string, args ...string) error {
	if dial == nil {
		return fmt.Errorf("empty dial")
	}
	client := pb.NewRunnerClient(dial)

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	stream, err := client.Update(context.Background(), grpc.WaitForReady(true))
	if err != nil {
		return err
	}

	buf := make([]byte, 1024*1000)
	md := &pb.FileMd{
		Name: f.Name(),
		Args: args,
	}
	fmt.Print("tx:")
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			if n > 0 {
				stream.Send(&pb.UpdateRequest{
					Buf: buf[:n],
					Md:  md,
				})
			}
			break
		}
		err = stream.Send(&pb.UpdateRequest{
			Buf: buf[:n],
			Md:  md,
		})
		if err != nil {
			return err
		}
		fmt.Print(".")
		_, err = stream.Recv()
		if err != nil {
			return err
		}
	}
	stream.CloseSend()
	fmt.Println()
	return nil
}

func Restart(dial *grpc.ClientConn) error {
	if dial == nil {
		return fmt.Errorf("empty dial")
	}
	client := pb.NewRunnerClient(dial)
	p := &pb.RunRequest{Restart: true}
	fmt.Println(p)
	_, err := client.Run(context.Background(), p, grpc.WaitForReady(true))
	return err
}
