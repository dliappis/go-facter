package main

import (
	"flag"

	"github.com/dliappis/go-facter/lib/cpu"
	"github.com/dliappis/go-facter/lib/disk"
	"github.com/dliappis/go-facter/lib/facter"
	"github.com/dliappis/go-facter/lib/formatter"
	"github.com/dliappis/go-facter/lib/host"
	"github.com/dliappis/go-facter/lib/keyfilter"
	"github.com/dliappis/go-facter/lib/mem"
	"github.com/dliappis/go-facter/lib/net"
)

func main() {
	conf := facter.Config{}
	ptFormat := flag.Bool("plaintext", false,
		"Emit facts as key => value pairs")
	kvFormat := flag.Bool("keyvalue", false,
		"Emit facts as key:value pairs")
	jsonFormat := flag.Bool("json", false,
		"Emit facts as a JSON")
	conf.KeyFilter = keyfilter.NewFilter()

	flag.Parse()

	queryArgs := flag.Args()

	if *ptFormat == true {
		conf.Formatter = formatter.NewFormatter()
	} else if *kvFormat == true {
		conf.Formatter = formatter.NewKeyValueFormatter()
	} else if *jsonFormat == true {
		conf.Formatter = formatter.NewJSONFormatter()
	} else {
		conf.Formatter = formatter.NewFormatter()
	}

	conf.KeyFilter.AddMany(queryArgs)

	facter := facter.New(&conf)
	_ = cpu.GetCPUFacts(facter)
	_ = disk.GetDiskFacts(facter)
	_ = host.GetHostFacts(facter)
	_ = mem.GetMemoryFacts(facter)
	_ = net.GetNetFacts(facter)
	facter.Print()
}
