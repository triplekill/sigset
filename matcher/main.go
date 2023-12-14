package main

import (
	domain_suffix_trie "github.com/golang-infrastructure/go-domain-suffix-trie"
	"log"
	"matcher/loader"
)

//func DomainMatcher(input *domain_suffix_trie.DomainSuffixTrieNode[string]) bool {
//
//}

type DomainMatcher struct {
	WhiteRules *domain_suffix_trie.DomainSuffixTrieNode[string]
	BlackRules *domain_suffix_trie.DomainSuffixTrieNode[string]
}

func (d DomainMatcher) MatchWhite(domain string) bool {

	if len(d.WhiteRules.FindMatchDomainSuffixPayload(domain)) > 0 {
		return true
	}

	return false
}

func (d DomainMatcher) MatchBlack(domain string) bool {

	if len(d.BlackRules.FindMatchDomainSuffixPayload(domain)) > 0 {
		return true
	}

	return false
}

func main() {

	var domainMatcher DomainMatcher
	var err error
	domainMatcher.WhiteRules, err = loader.LoadGeoSite("./", "white")
	if err != nil {
		log.Panic("err:", err)
	}

	domainMatcher.BlackRules, err = loader.LoadGeoSite("./", "black")
	if err != nil {
		log.Panic("err:", err)
	}

	log.Println(domainMatcher.MatchWhite("www.xiaohongshu.com"))
	log.Println(domainMatcher.MatchWhite("xiaohongshu.com"))
	log.Println(domainMatcher.MatchWhite("1.xiaohongshu.com"))
	log.Println(domainMatcher.MatchWhite("www.xiaohongshu.com.cn"))
	log.Println(domainMatcher.MatchWhite("www.cdnctrmi.cn"))

}
