package events

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	rpchttp "github.com/cometbft/cometbft/rpc/client/http"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"

	nodeclient "pkg.akt.dev/go/node/client"
)

// The node URI schemes the service reaches a node through, and the networks it dials.
const (
	schemeTCP   = "tcp"
	schemeHTTP  = "http"
	schemeHTTPS = "https"
	schemeWS    = "ws"
	schemeWSS   = "wss"
	schemeUnix  = "unix"

	networkTCP  = "tcp"
	networkUnix = "unix"
)

// queryNode is what the service asks the node over HTTP.
type queryNode interface {
	Status(ctx context.Context) (*ctypes.ResultStatus, error)
	BlockResults(ctx context.Context, height *int64) (*ctypes.ResultBlockResults, error)
}

// wsTarget is where and how to open the node's websocket.
type wsTarget struct {
	// url is the ws:// or wss:// URL of the /websocket endpoint.
	url string
	// network is "tcp", or "unix" for a node on a unix socket.
	network string
	// address is host:port, or the socket path.
	address string
	// header carries Authorization when the node URI carries credentials, else nil.
	header http.Header
}

// parseNodeURI normalizes nodeURI, so it names an explicit port, and checks it is a
// node address the service can reach: tcp, http, https, ws or wss with a host, or unix
// with a socket path. A URI without a scheme, //host[:port], means tcp; host:port is not
// a URI and is rejected. Errors never repeat credentials the URI carries.
func parseNodeURI(nodeURI string) (*url.URL, error) {
	if nodeURI == "" {
		return nil, errors.New("events: node URI must not be empty")
	}

	raw := nodeURI
	if strings.HasPrefix(raw, "//") {
		raw = schemeTCP + ":" + raw
	}

	u, err := url.Parse(nodeclient.NormalizeEndpoint(raw))
	if err != nil {
		// A url.Error repeats the URI it failed on, credentials included.
		var uerr *url.Error
		if errors.As(err, &uerr) {
			err = uerr.Err
		}
		return nil, fmt.Errorf("events: invalid node URI %s: %w", redactURI(nodeURI), err)
	}

	switch u.Scheme {
	case schemeTCP, schemeHTTP, schemeHTTPS, schemeWS, schemeWSS:
		if u.Host == "" {
			return nil, fmt.Errorf("events: node URI %s has no host", redactURI(nodeURI))
		}
	case schemeUnix:
		if u.Path == "" {
			return nil, fmt.Errorf("events: node URI %s has no socket path", redactURI(nodeURI))
		}
	default:
		return nil, fmt.Errorf("events: node URI %s has unsupported scheme %q, want tcp://, http://, https://, ws://, wss:// or unix://",
			redactURI(nodeURI), u.Scheme)
	}

	return u, nil
}

// redactURI quotes nodeURI for an error message, with any credentials in it masked: the
// user part as well as the password, since a token often sits in the user part.
func redactURI(nodeURI string) string {
	at := strings.LastIndex(nodeURI, "@")
	if at < 0 {
		return strconv.Quote(nodeURI)
	}

	start := 0
	if i := strings.Index(nodeURI, "//"); i >= 0 && i < at {
		start = i + len("//")
	}

	return strconv.Quote(nodeURI[:start] + "xxxxx" + nodeURI[at:])
}

// newWSTarget maps a parsed node URI to its websocket endpoint the way the CometBFT
// client does: https and wss use wss, everything else ws; a unix socket is dialed
// directly, with its path standing in for the host; the URI's path prefix is kept;
// credentials in the URI, a user, a password or both, become a Basic Authorization
// header, as the node's HTTP client sends them.
func newWSTarget(u *url.URL) wsTarget {
	scheme := schemeWS
	if u.Scheme == schemeHTTPS || u.Scheme == schemeWSS {
		scheme = schemeWSS
	}

	t := wsTarget{network: networkTCP, address: u.Host}
	host, path := u.Host, strings.TrimSuffix(u.EscapedPath(), "/")
	if u.Scheme == schemeUnix {
		t.network, t.address = networkUnix, u.Path
		host, path = strings.ReplaceAll(u.EscapedPath(), "/", "."), ""
	}
	t.url = scheme + "://" + host + path + "/websocket"

	if u.User != nil {
		user := u.User.Username()
		pass, _ := u.User.Password()
		if user != "" || pass != "" {
			t.header = http.Header{"Authorization": {"Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+pass))}}
		}
	}

	return t
}

// queryURL is u as the HTTP client needs it. A node given by its websocket URI is
// queried over HTTP on the same address: ws becomes http and wss https, since the HTTP
// client can neither dial a ws network nor post to a ws URL.
func queryURL(u *url.URL) *url.URL {
	q := *u
	switch q.Scheme {
	case schemeWS:
		q.Scheme = schemeHTTP
	case schemeWSS:
		q.Scheme = schemeHTTPS
	}
	return &q
}

// httpQueryNode is the service's own HTTP client for the node.
type httpQueryNode struct {
	*rpchttp.HTTP
	hc *http.Client
}

// CloseIdleConnections closes the client's idle connections to the node. The service
// calls it once it no longer queries the node, since nothing else shares the pool.
func (n httpQueryNode) CloseIdleConnections() { n.hc.CloseIdleConnections() }

// newQueryNode returns an HTTP client for the node at u, with its own connection pool.
// It is never started: its websocket side goes unused.
func newQueryNode(ctx context.Context, u *url.URL) (queryNode, error) {
	remote := queryURL(u).String()

	hc, err := nodeclient.NewHTTPClient(ctx, remote)
	if err != nil {
		return nil, err
	}

	c, err := rpchttp.NewWithClient(remote, "/websocket", hc)
	if err != nil {
		return nil, err
	}

	return httpQueryNode{HTTP: c, hc: hc}, nil
}

// closeIdle releases the idle connections of a client the service owns. Clients that
// keep no pool of their own have nothing to release.
func closeIdle(client queryNode) {
	if c, ok := client.(interface{ CloseIdleConnections() }); ok {
		c.CloseIdleConnections()
	}
}
