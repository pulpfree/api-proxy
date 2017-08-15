package proxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

var portmapfile = "./portmap.json"

// Proxy struct
type Proxy struct {
	domainMap []dMap
	portMap   map[string]string
	proxy     *httputil.ReverseProxy
}

type dMap struct {
	Port   string `json:"port"`
	Domain string `json:"domain"`
}

// New function
func (p *Proxy) New() {
	p.getPortMap()
}

// Handle function
func (p *Proxy) Handle(w http.ResponseWriter, r *http.Request) {
	port := strings.Split(r.Host, ":")
	d := p.getDomain(port[1])
	var target = "http://" + d + ":3020"

	url, err := url.Parse(target)
	if err != nil {
		log.Fatal("URL failed to parse")
	}
	p.proxy = httputil.NewSingleHostReverseProxy(url)
	p.proxy.ServeHTTP(w, r)
}

func (p *Proxy) getPortMap() {
	raw, err := ioutil.ReadFile(portmapfile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &p.domainMap)
	p.portMap = map[string]string{}
	for _, pm := range p.domainMap {
		p.portMap[pm.Port] = pm.Domain
	}
}

func (p *Proxy) getDomain(port string) string {
	var d string
	if _, ok := p.portMap[port]; ok {
		d = p.portMap[port]
	}
	return d
}

// ============ Helper Methods ============ //

func (p *Proxy) toString(s interface{}) string {
	return toJSON(s)
}

func toJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}
