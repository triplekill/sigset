package main

import (
	pb "builder/pb"
	"flag"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"strings"
)

var (
	dattype  = flag.String("type", "white", "datfile's name")
	sitefile = flag.String("dir", "sites", "sites file folder")

	whitedat = "./white.dat"
	blackdat = "./black.dat"
)

func GetSitesList(fileName, domaintype string) []*pb.Rule {
	//func GetSitesList(fileName string) []*pb.Rule {

	var ruleSitesDomains []*pb.Rule
	d, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	domains := strings.Split(string(d), "\n")

	for _, pattern := range domains {
		if len(pattern) != 0 {
			log.Println(strings.Split(domaintype, ".")[0], pattern)
			item := &pb.Rule{
				Type:   strings.Split(domaintype, ".")[0],
				Domain: pattern,
			}
			ruleSitesDomains = append(ruleSitesDomains, item)
		}

	}

	return ruleSitesDomains
}

func main() {

	flag.CommandLine.Usage = func() {
		fmt.Println(`./builder -type white -dir sites`)
	}
	flag.Parse()

	if *dattype == "" {
		return
	}
	var outpath string

	if *dattype == "white" {
		outpath = whitedat
	} else {
		outpath = blackdat
	}

	rulefiles, err := os.ReadDir(*sitefile)

	if err != nil {
		panic(err)
	}
	rulelist := new(pb.RuleList)

	for _, rf := range rulefiles {
		tmplist := GetSitesList(*sitefile+"/"+rf.Name(), rf.Name())
		log.Println("tmplist:", len(tmplist))
		for _, rule := range tmplist {
			rulelist.Entry = append(rulelist.Entry, rule)
		}
	}
	log.Println(len(rulelist.Entry))
	SiteListBytes, err := proto.Marshal(rulelist)
	if err != nil {
		fmt.Println("failed to marshal v2sites:", err)
		return
	}
	if err := os.WriteFile(outpath, SiteListBytes, 0777); err != nil {
		fmt.Println("failed to write %s.", outpath, err)
	}
	fmt.Println(*sitefile, "-->", outpath)
}
