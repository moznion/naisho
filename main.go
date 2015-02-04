package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Address  string
	Password string
}

type option struct {
	gmailConfPath string
	fromAddr      string
	pass          string
	subject       string
	body          string
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage:
  %s [OPTIONS] <Target GitHub ID> <Message>

Options:
`, os.Args[0])
		flag.PrintDefaults()
	}

	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	homeDir := usr.HomeDir

	opt := new(option)
	flag.StringVar(&opt.gmailConfPath, "conf", filepath.Join(homeDir, ".naisho"), "Path for the configuration file of gmail")
	flag.StringVar(&opt.fromAddr, "from", "", "Email address of Gmail to send (this overwrites the configuration file's one)")
	flag.StringVar(&opt.pass, "pass", "", "Password of your gmail acount (this overwrites the configuration file's one)")
	flag.StringVar(&opt.subject, "subject", "", "Subject of message. If this is empty, it will use default subject")
	flag.StringVar(&opt.body, "body", "", "Body of message. If this is empty, it will use default message")
	flag.Parse()

	args := flag.Args()

	if len(args) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	conf, err := readYamlConfig(opt.gmailConfPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if opt.fromAddr == "" {
		opt.fromAddr = conf.Address
	}

	if opt.pass == "" {
		opt.pass = conf.Password
	}

	targetGitHubId := args[0]
	pubkey, err := fetchPublicKeyByUserId(targetGitHubId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	toAddr, err := fetchEmailAddressByUserId(targetGitHubId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	msg := args[1]
	encrypted, err := encryptStringBySshRsaPublicKey(pubkey, msg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = sendByGmail(&mail{
		fromAddr: opt.fromAddr,
		pass:     opt.pass,
		toAddr:   toAddr,
		msg:      encrypted,
		subject:  opt.subject,
		body:     opt.body,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Sent Successfully!\n")
}

func readYamlConfig(configPath string) (*conf, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	conf := new(conf)
	err = yaml.Unmarshal(data, conf)
	return conf, nil
}
