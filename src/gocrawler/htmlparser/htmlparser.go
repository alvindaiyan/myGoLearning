package htmlparser

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	BaseUrl = "****"
)

type Catelog struct {
	Name     string
	Url      string
	SubCates []SubCatelog
}

type SubCatelog struct {
	Name string
	Url  string
}

func GetCategory(url string) []*Catelog {
	page := GetByUrl(url)

	reader := bytes.NewReader(page)
	doc, err := html.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}

	cates := GetAllHtmlHrefs(doc)

	return cates
}

func WriteToFile(filename string, content string) {
	// f, err := os.OpenFile(filename, os.O_CREATE, 0600)
	// f, err := os.Create(filename)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	if _, err = f.WriteString(content); err != nil {
		log.Println("...")
		panic(err)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetAllHtmlHrefs(doc *html.Node) (cates []*Catelog) {
	var parentCate *Catelog
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" && n.Parent != nil && n.Parent.Parent != nil && n.Parent.Parent.Parent != nil && n.Parent.Parent.Parent.Type == html.ElementNode && n.Parent.Parent.Parent.Data == "div" {
			toProcess := false
			for _, a := range n.Parent.Parent.Parent.Attr {
				if a.Key == "class" && a.Val == "sitemap_container" {
					toProcess = true
					break
				}
			}
			if toProcess {
				for _, a := range n.Attr {
					if a.Key == "href" {
						if a.Val != "javascript:void(0);" && a.Val != "#" && !strings.HasPrefix(a.Val, "http://") && !strings.HasPrefix(a.Val, "//") && !strings.Contains(a.Val, ";") && !strings.Contains(a.Val, "?") && !strings.Contains(a.Val, "#") && !strings.Contains(a.Val, ".jsp") && strings.HasPrefix(a.Val, "/nz/en") && !strings.Contains(a.Val, "customer-service") && strings.Index(a.Val, "/nz/en/")+7 != len(a.Val) {
							if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
								isHeading := false
								if n.Parent != nil && n.Parent.Type == html.ElementNode && n.Parent.Data == "li" {
									for _, attr := range n.Parent.Attr {
										if attr.Key == "class" && attr.Val == "sitemap_heading" {
											isHeading = true
											break
										}
									}
								}
								if isHeading {
									var subcates []SubCatelog
									parentCate = &Catelog{n.FirstChild.Data, a.Val, subcates}
									cates = append(cates, parentCate)
								} else {
									if parentCate.Name != "" {
										cate := SubCatelog{n.FirstChild.Data, a.Val}
										parentCate.SubCates = append(parentCate.SubCates, cate)
									}
								}
							}
						}
						break
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return
}

type ItemUrl struct {
	Url  string
	Name string
}

type Item struct {
	Name        string
	CatelogName string
	Price       string
	Description string
}

func GetItem(url string) (item Item) {
	if strings.HasPrefix(url, "/") {
		url = BaseUrl + url
	}
	page := GetByUrl(url)

	reader := bytes.NewReader(page)
	doc, err := html.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}

	f = func(n *html.Node) {
		// to get the price
		if n != nil && n.Type == html.ElementNode && n.Data == "span" {
			for _, id := range n.Attr {
				if id.Key == "id" && id.Val == "crc_product_rp" {
					if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
						item.Price = strings.TrimSpace(n.FirstChild.Data)
						break
					}
				}
			}
		}

		// set item name
		if n != nil && n.Type == html.ElementNode && n.Data == "li" {
			for _, class := range n.Attr {
				if class.Key == "class" && class.Val == "product_title" {
					if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
						item.Name = strings.TrimSpace(n.FirstChild.Data)
						break
					}
				}
			}
		}

		// set item description
		if n != nil && n.Type == html.ElementNode && n.Data == "div" {
			for _, class := range n.Attr {
				if class.Key == "class" && class.Val == "short_desc" {
					var getAllText func(*html.Node) string

					for span := n.FirstChild; span != nil; span = span.NextSibling {
						if span.Type == html.ElementNode && span.Data == "span" {

							// get all the text data
							getAllText = func(node *html.Node) (text string) {
								if node != nil && node.Type == html.TextNode {
									text = text + strings.TrimSpace(node.Data) + "\n"
								}

								for spanc := node.FirstChild; spanc != nil; spanc = spanc.NextSibling {
									text = text + getAllText(spanc)
								}
								return
							} // endfunc

							item.Description = strings.TrimSpace(getAllText(span))

							break
						}
					}
				}
			}
		} // endif
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return
}

func GetAllItemHtmlHrefs(doc *html.Node) (itemUrls []ItemUrl) {
	f = func(n *html.Node) {
		if n != nil && n.Parent != nil && n.Parent.Type == html.ElementNode && n.Parent.Data == "li" && n.Type == html.ElementNode && n.Data == "a" && n.Parent.Parent.Parent.Parent.Parent.Data == "div" {
			for _, li := range n.Parent.Attr {
				if li.Key == "class" && li.Val == "description" {
					// the current one is what we what maybe?
					for _, a := range n.Attr {
						if a.Key == "href" {
							if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
								itemUrl := ItemUrl{strings.TrimSpace(a.Val), strings.TrimSpace(n.FirstChild.Data)}
								itemUrls = append(itemUrls, itemUrl)
							}
							break
						}
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return
}

func GetByUrl(url string) []byte {
	log.Println("download start download url: ", url)
	response, err := http.Get(url)
	log.Println("download finished url: ", url)

	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
		}
		// fmt.Printf("%s\n", string(contents))
		return contents
	}
	return nil
}

var f func(*html.Node)
