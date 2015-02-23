package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"gocrawler/htmlparser"
	"golang.org/x/net/html"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// type Asset struct {
// 	ParentCatelogue string
// 	CatelogName     string
// 	Items           []Item
// }

// type Item struct {
// 	Name     string
// 	Price    string
// 	Metadata map[string]string
// }

// func SaveAsset(asset Asset) {
// 	log.Println("save asset to db")
// }

type myFatcher struct{}

var (
	baseUrl string
)

func (m *myFatcher) FetchAllItemUrls(url string) []string {
	log.Println("---------------> fetch all item urls: ", url)

	fetched.Lock()
	if _, ok := fetched.m[url]; ok {
		fetched.Unlock()
		log.Printf("<- Done with %v, already fetched.\n", url)
		return nil
	}
	// We mark the url to be loading to avoid others reloading it at the same time.
	fetched.m[url] = loading
	fetched.Unlock()
	newUrl := url
	if strings.HasPrefix(url, "/") {
		newUrl = baseUrl + url
	}

	file := htmlparser.GetByUrl(newUrl)
	reader := bytes.NewReader(file)
	doc, err := html.Parse(reader)

	itemUrls := htmlparser.GetAllItemHtmlHrefs(doc)
	var urls []string
	for _, v := range itemUrls {
		urls = append(urls, v.Url)
	}

	// And update the status in a synced zone.
	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	return urls
}

// get all hrefs
func (m *myFatcher) Fetch(url string) (*html.Node, []string, error) {
	// log.Println("fetch url: ", url)
	if strings.HasPrefix(url, "/") {
		url = baseUrl + url
	}

	file := htmlparser.GetByUrl(url)
	reader := bytes.NewReader(file)
	doc, err := html.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}

	cates := htmlparser.GetAllHtmlHrefs(doc)
	var urls []string
	for _, v := range cates {
		// fmt.Println()
		// log.Println("--------------------------> url: ", v.Url)
		if url != baseUrl+"/sitemap" {
			log.Fatal("stop here")
		}
		for _, sub := range v.SubCates {
			// log.Printf("  --------------------------> name: %s, url: %s \n", sub.Name, sub.Url)
			// urls = append(urls, sub.Url)
			urls = append(urls, sub.Url)
		}
		// fmt.Println("................")
	}

	return doc, urls, nil
}

// func main() {
// 	// log.Printf("----------> process Item: %v, %v/%v of %v. \n", itemurl, idx+1, len(urls), url)
// 	item := htmlparser.GetItem("baseurl/nz/en/element-billings-3-shoes-ss14/rp-prod113407")
// 	item.CatelogName = "test url"
// 	log.Println("item catelog", item.CatelogName)
// 	log.Println("item Name: ", item.Name)
// 	log.Println("item Price: ", item.Price)
// 	log.Println("item Description: ", item.Description)

// }

func main() {
	crawlAll := false
	if len(os.Args) < 2 {
		log.Fatal("tell me the base url")
	}
	baseUrl = os.Args[1]
	htmlparser.BaseUrl = baseUrl
	if len(os.Args) >= 3 && os.Args[2] == "all" {
		crawlAll = true
	}

	if crawlAll {
		log.Println("crawl all product")
	} else {
		log.Println("only crawl Bike entry")
	}

	start := time.Now()
	siteMapUrl := baseUrl + "/sitemap"
	cates := htmlparser.GetCategory(siteMapUrl)

	// save as json str for catelog relationship
	js, _ := json.Marshal(cates)
	htmlparser.WriteToFile("data/catelog.json", string(js))
	myfetcher := myFatcher{}

	for _, cate := range cates {
		log.Println("cate.name = ", cate.Name)
		if !crawlAll && strings.TrimSpace(cate.Name) == "Bikes" {

			done := make(chan bool)
			for i, sub := range cate.SubCates {
				log.Printf("-> Crawling child %v/%v of %v : %v.\n", i+1, len(cate.SubCates), cate.Url, sub.Url)
				go func(url string, parentName string) {
					urls := myfetcher.FetchAllItemUrls(url) // the item summary page urls
					// could be pagination? fetch again!
					if len(urls) >= 20 {
						var urlsItemPage = myfetcher.FetchAllItemUrls(url + "?page=2")
						for _, u := range urlsItemPage {
							urls = append(urls, u)
						}
						for i := 3; len(urlsItemPage) > 24; i++ {
							urlsItemPage = urlsItemPage[:0] // clear the array
							urlsItemPage = myfetcher.FetchAllItemUrls(url + "?page=" + strconv.Itoa(i))
							// log.Println("...", len(urlsItemPage))
							for _, u := range urlsItemPage {
								urls = append(urls, u)
							}
						}
					}

					// get Items:
					for idx, itemurl := range urls {
						log.Printf("----------> process Item: %v, %v/%v of %v. \n", itemurl, idx+1, len(urls), url)
						item := htmlparser.GetItem(itemurl)
						item.CatelogName = url
						// log.Println("item catelog", item.CatelogName)
						// log.Println("item Name: ", item.Name)
						// log.Println("item Price: ", item.Price)
						// log.Println("item Description: ", item.Description)
						js, _ = json.Marshal(item)
						htmlparser.WriteToFile("data/_"+itemurl[strings.LastIndex(itemurl, "/")+1:]+".json", string(js))

					}

					// print the result
					// if len(urls) != 0 {
					// 	log.Println("\n/////////////////////////////////////")
					// 	log.Println("size(", url, ")", ": ", len(urls))
					//// print the result to log file
					////// 	htmlparser.WriteToFile("log", "size("+url+") parent: "+parentName+" : "+strconv.Itoa(len(urls))+"\n\n")
					// 	log.Println("\n/////////////////////////////////////")
					// }
					done <- true
				}(sub.Url, cate.Name)
			}

			for i, sub := range cate.SubCates {
				log.Printf("<- [%v] %v/%v Waiting for child %v.\n", cate.Url, i+1, len(cate.SubCates), sub.Url)
				<-done
			}

		}
	} // endfor

	log.Println("Fetching stats\n--------------")
	for url, err := range fetched.m {
		if err != nil {
			log.Printf("%v failed: %v\n", url, err)
		} else {
			log.Printf("%v was fetched\n", url)
		}
	}
	log.Println("total processed: ", len(fetched.m))
	elapsed := time.Since(start)
	log.Printf("time took: %s", elapsed)
}

// url is /xx/xx/xxxx
// body is in html.Node
func processPage(body *html.Node, url string) {
	log.Println("processing: ", url)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		log.Printf("<- Done with %v, depth 0.\n", url)
		return
	}

	fetched.Lock()
	if _, ok := fetched.m[url]; ok {
		fetched.Unlock()
		log.Printf("<- Done with %v, already fetched.\n", url)
		return
	}
	// We mark the url to be loading to avoid others reloading it at the same time.
	fetched.m[url] = loading
	fetched.Unlock()

	// We load it concurrently.
	body, urls, err := fetcher.Fetch(url)

	// And update the status in a synced zone.
	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	if err != nil {
		log.Printf("<- Error on %v: %v\n", url, err)
		return
	}

	log.Printf("Found: %s \n", url)
	processPage(body, url)

	done := make(chan bool)
	for i, u := range urls {
		log.Printf("-> Crawling child %v/%v of %v : %v.\n", i+1, len(urls), url, u)
		go func(url string) {
			Crawl(url, depth-1, fetcher)
			done <- true
		}(u)
	}
	for i, u := range urls {
		log.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i, len(urls), u)
		<-done
	}
	log.Printf("<- Done with %v\n", url)
}

const (
	MAX_DEPTH = 10
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body *html.Node, urls []string, err error)
}

// fetched tracks URLs that have been (or are being) fetched.
// The lock must be held while reading from or writing to the map.
// See http://golang.org/ref/spec#Struct_types section on embedded types.
var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

var loading = errors.New("url load in progress") // sentinel value
