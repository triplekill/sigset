package loader

import (
	domain_suffix_trie "github.com/golang-infrastructure/go-domain-suffix-trie"
	"google.golang.org/protobuf/proto"
	pb "matcher/pb"
	"os"
	"path"
)

func LoadGeoSite(datDir string, name string) (*domain_suffix_trie.DomainSuffixTrieNode[string], error) {
	datName := name + ".dat"
	datPath := path.Join(datDir, datName)
	geositeBytes, err := os.ReadFile(datPath)
	if err != nil {
		return nil, err
	}

	var rulelist pb.RuleList

	if err := proto.Unmarshal(geositeBytes, &rulelist); err != nil {
		return nil, err
	}
	ruleTree := domain_suffix_trie.NewDomainSuffixTrie[string]()
	for _, site := range rulelist.Entry {
		//log.Println(site.Domain, "site.Type:", site.Type)
		ruleTree.AddDomainSuffix(site.Domain, site.Type)
	}

	return ruleTree, nil
}
