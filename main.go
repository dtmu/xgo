package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"strings"
)

type buildPattern struct {
	goos   string
	goarch string
}

type buildPatterns []buildPattern

var bps buildPatterns = buildPatterns{
	buildPattern{"aix", "ppc64"},
	buildPattern{"android", "386"},
	buildPattern{"android", "amd64"},
	buildPattern{"android", "arm"},
	buildPattern{"android", "arm64"},
	buildPattern{"darwin", "386"},
	buildPattern{"darwin", "amd64"},
	buildPattern{"darwin", "arm"},
	buildPattern{"darwin", "arm64"},
	buildPattern{"dragonfly", "amd64"},
	buildPattern{"freebsd", "386"},
	buildPattern{"freebsd", "amd64"},
	buildPattern{"freebsd", "arm"},
	buildPattern{"illumos", "amd64"},
	buildPattern{"js", "wasm"},
	buildPattern{"linux", "386"},
	buildPattern{"linux", "amd64"},
	buildPattern{"linux", "arm"},
	buildPattern{"linux", "arm64"},
	buildPattern{"linux", "ppc64"},
	buildPattern{"linux", "ppc64le"},
	buildPattern{"linux", "mips"},
	buildPattern{"linux", "mipsle"},
	buildPattern{"linux", "mips64"},
	buildPattern{"linux", "mips64le"},
	buildPattern{"linux", "s390x"},
	buildPattern{"netbsd", "386"},
	buildPattern{"netbsd", "amd64"},
	buildPattern{"netbsd", "arm"},
	buildPattern{"openbsd", "386"},
	buildPattern{"openbsd", "amd64"},
	buildPattern{"openbsd", "arm"},
	buildPattern{"openbsd", "arm64"},
	buildPattern{"plan9", "386"},
	buildPattern{"plan9", "amd64"},
	buildPattern{"plan9", "arm"},
	buildPattern{"solaris", "amd64"},
	buildPattern{"windows", "386"},
	buildPattern{"windows", "amd64"},
}

func main() {
	app := cli.NewApp()
	app.Name = "xgo"
	app.Usage = "This is cli tool to build all binary for multi platform."
	app.Version = "0.0.1"
	app.Action = action
	app.Run(os.Args)
}

func action(c *cli.Context) error {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return err
	}
	binName := strings.Trim(c.Args().Get(0), " ")
	if binName == "" {
		binName = pop(strings.Split(pwd, "/"))
	}
	fmt.Println("binary name: " + binName)
	exe := ""
	for i := 0; i < len(bps); i++ {
		if bps[i].goos == "windows" {
			exe = ".exe"
		} else {
			exe = ""
		}
		binPath := pwd + "/bin/" + bps[i].goos + "/" + bps[i].goarch + "/" + binName + exe
		cmd := exec.Command("go", "build", "-o", binPath)
		cmd.Env = append(os.Environ(), "GOOS="+bps[i].goos, "GOARCH="+bps[i].goarch, binPath)
		err := cmd.Run()
		if err != nil {
			fmt.Println("made binary fail:", binPath)
			continue
		}
		fmt.Println("made binary:", binPath)
	}
	return nil
}

func pop(slice []string) string {
	ans := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return ans
}
