package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

func main1() {

	cmd := exec.Command("ls", "-alh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to call cmd.Run(): %v", err)
	}

}

func main2() {

	cmd := exec.Command("ls", "-lah")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "/"
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to call cmd.Run(): %v", err)
	}
}

func main4() {
	cmd := exec.Command("/opt/homebrew/go/bin/go", "env")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("path: %s", cmd.Path)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to call cmd.Run(): %v", err)
	}
}

func main5() {
	cmd := exec.Command("bash", "-c", "echo $myvar")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = []string{"myvar=abc"}
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to call cmd.Run(): %v", err)
	}

}

func main6() {
	cmd := exec.Command("bash", "-c", "sleep 1;echo $myvar")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatalf("failed to call cmd.Start(): %v", err)
	}
	log.Printf("pid: %d", cmd.Process.Pid)
	log.Printf("pid: %d", cmd.ProcessState)
	cmd.Process.Wait()
	log.Printf("exitcode: %d", cmd.ProcessState.ExitCode())
}

func main7() {
	path, err := exec.LookPath("ls")
	if err != nil {
		log.Printf("'ls' not found")
	} else {
		log.Printf("'ls' is in '%s'\n", path)
	}
}

func main8() {
	cmd := exec.Command("ls", "-lah")
	data, err := cmd.Output()
	if err != nil {
		log.Fatalf("failed to call Output(): %v", err)
	}
	log.Printf("output: %s", data)
}

func main9() {

	cmd := exec.Command("ls", "-lahxyz")
	cmd.Stderr = os.Stderr
	data, err := cmd.Output()
	if err != nil {
		log.Fatalf("failed to call Output(): %v", err)
	}
	log.Printf("output: %s", data)

}

func main10() {
	cmd := exec.Command("curl", "-o", "go1.15.6.linux-amd64.tar.gz", "https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz")
	err := cmd.Start()
	if err != nil {
		log.Fatalf("failed to call Run(): %v", err)
	}

}

func main11() {
	cmdCat := exec.Command("cat", "main.go")
	catout, err := cmdCat.StdoutPipe()
	if err != nil {
		log.Fatalf("failed to get StdoutPipe of cat: %v", err)
	}
	cmdWC := exec.Command("wc", "-l")
	cmdWC.Stdin = catout
	cmdWC.Stdout = os.Stdout
	err = cmdCat.Start()
	if err != nil {
		log.Fatalf("failed to call cmdCat.Run(): %v", err)
	}
	err = cmdWC.Run()
	if err != nil {
		log.Fatalf("failed to call cmdWC.Start(): %v", err)
	}
	cmdCat.Wait()
	cmdWC.Wait()
}

func main12() {
	cmdCat := exec.Command("cat", "main.go")
	cmdWC := exec.Command("wc", "-l")
	data, err := pipeCommands(cmdCat, cmdWC)
	if err != nil {
		log.Fatalf("failed to call pipeCommands(): %v", err)
	}
	log.Printf("output: %s", data)
}
func pipeCommands(commands ...*exec.Cmd) ([]byte, error) {
	for i, command := range commands[:len(commands)-1] {
		out, err := command.StdoutPipe()
		if err != nil {
			return nil, err
		}
		command.Start()
		commands[i+1].Stdin = out
	}
	final, err := commands[len(commands)-1].Output()
	fmt.Println(string(final[:len(final)-1]))
	if err != nil {
		return nil, err
	}
	return final, nil
}

func main13() {
	cmd := exec.Command("bash", "-c", "cat main.go| wc -l")
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("failed to call pipeCommands(): %v", err)
	}
	log.Printf("output: %s", data)

}

func main14() {
	stdin, err := os.Open("main.go")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	cmd := exec.Command("wc", "-l")
	cmd.Stdin = stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("failed to call cmd.Run(): %v", err)
	}
}

func main15() {

	cmd := exec.Command("curl", "https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz")
	var stdoutProcessStatus bytes.Buffer
	cmd.Stdout = io.MultiWriter(ioutil.Discard, &stdoutProcessStatus)
	done := make(chan struct{})
	go func() {
		tick := time.NewTicker(time.Second)
		defer tick.Stop()
		for {
			select {
			case <-done:
				return
			case <-tick.C:
				log.Printf("downloaded: %d", stdoutProcessStatus.Len())
			}
		}
	}()
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to call Run(): %v", err)
	}
	close(done)

}

func main() {
	cmd := exec.Command("ls", "-lah")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to call Run(): %v", err)
	}
	log.Printf("out:\n%s\nerr:\n%s", stdout.String(), stderr.String())
}
