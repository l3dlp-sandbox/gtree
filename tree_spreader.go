//go:build !wasm

package gtree

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"

	"github.com/fatih/color"
	toml "github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

func newSpreader(encode encode) spreader {
	switch encode {
	case encodeJSON:
		return &jsonSpreader{}
	case encodeYAML:
		return &yamlSpreader{}
	case encodeTOML:
		return &tomlSpreader{}
	default:
		return &defaultSpreader{}
	}
}

func newColorizeSpreader(fileExtensions []string) spreader {
	return &colorizeSpreader{
		fileConsiderer: newFileConsiderer(fileExtensions),
		fileColor:      color.New(color.Bold, color.FgHiCyan),
		fileCounter:    newCounter(),

		dirColor:   color.New(color.FgGreen),
		dirCounter: newCounter(),
	}
}

type encode int

const (
	encodeDefault encode = iota
	encodeJSON
	encodeYAML
	encodeTOML
)

type defaultSpreader struct {
	sync.Mutex
}

const workerSpreadNum = 10

func (ds *defaultSpreader) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer func() {
			close(errc)
		}()

		bw := bufio.NewWriter(w)
		wg := &sync.WaitGroup{}
		for i := 0; i < workerSpreadNum; i++ {
			wg.Add(1)
			go ds.worker(ctx, wg, bw, roots, errc)
		}
		wg.Wait()

		errc <- bw.Flush()
	}()

	return errc
}

func (ds *defaultSpreader) worker(ctx context.Context, wg *sync.WaitGroup, bw *bufio.Writer, roots <-chan *Node, errc chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case root, ok := <-roots:
			if !ok {
				return
			}
			ret := ds.spreadBranch(root)

			ds.Lock()
			_, err := bw.WriteString(ret)
			ds.Unlock()

			errc <- err
		}
	}
}

func (*defaultSpreader) spreadBranch(current *Node) string {
	ret := current.name + "\n"
	if !current.isRoot() {
		ret = current.branch() + " " + current.name + "\n"
	}

	for _, child := range current.children {
		ret += (*defaultSpreader)(nil).spreadBranch(child)
	}
	return ret
}

type colorizeSpreader struct {
	fileConsiderer *fileConsiderer
	fileColor      *color.Color
	fileCounter    *counter

	dirColor   *color.Color
	dirCounter *counter
}

func (cs *colorizeSpreader) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		bw := bufio.NewWriter(w)
	BREAK:
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					break BREAK
				}
				cs.fileCounter.reset()
				cs.dirCounter.reset()

				if _, err := bw.WriteString(
					fmt.Sprintf(
						"%s\n%s\n",
						cs.spreadBranch(root),
						cs.summary()),
				); err != nil {
					errc <- err
					return
				}
			}
			if err := bw.Flush(); err != nil {
				errc <- err
				return
			}
		}
	}()

	return errc
}

func (cs *colorizeSpreader) spreadBranch(current *Node) string {
	ret := ""
	if current.isRoot() {
		ret = cs.colorize(current) + "\n"
	} else {
		ret = current.branch() + " " + cs.colorize(current) + "\n"
	}

	for _, child := range current.children {
		ret += cs.spreadBranch(child)
	}
	return ret
}

func (cs *colorizeSpreader) colorize(current *Node) string {
	if cs.fileConsiderer.isFile(current) {
		_ = cs.fileCounter.next()
		return cs.fileColor.Sprint(current.name)
	} else {
		_ = cs.dirCounter.next()
		return cs.dirColor.Sprint(current.name)
	}
}

func (cs *colorizeSpreader) summary() string {
	return fmt.Sprintf(
		"%d directories, %d files",
		cs.dirCounter.current(),
		cs.fileCounter.current(),
	)
}

type jsonSpreader struct{}

func (*jsonSpreader) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		enc := json.NewEncoder(w)
	BREAK:
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					break BREAK
				}
				if err := enc.Encode(root.toJSONNode(nil)); err != nil {
					errc <- err
					return
				}
			}
		}
	}()

	return errc
}

type jsonNode struct {
	Name     string      `json:"value"`
	Children []*jsonNode `json:"children"`
}

func (parent *Node) toJSONNode(jParent *jsonNode) *jsonNode {
	if jParent == nil {
		jParent = &jsonNode{Name: parent.name}
	}
	if !parent.hasChild() {
		return jParent
	}

	jParent.Children = make([]*jsonNode, len(parent.children))
	for i := range parent.children {
		jParent.Children[i] = &jsonNode{Name: parent.children[i].name}
		_ = parent.children[i].toJSONNode(jParent.Children[i])
	}

	return jParent
}

type tomlSpreader struct{}

func (*tomlSpreader) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		enc := toml.NewEncoder(w)
	BREAK:
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					break BREAK
				}
				if err := enc.Encode(root.toTOMLNode(nil)); err != nil {
					errc <- err
					return
				}
			}
		}
	}()

	return errc
}

type tomlNode struct {
	Name     string      `toml:"value"`
	Children []*tomlNode `toml:"children"`
}

func (parent *Node) toTOMLNode(tParent *tomlNode) *tomlNode {
	if tParent == nil {
		tParent = &tomlNode{Name: parent.name}
	}
	if !parent.hasChild() {
		return tParent
	}

	tParent.Children = make([]*tomlNode, len(parent.children))
	for i := range parent.children {
		tParent.Children[i] = &tomlNode{Name: parent.children[i].name}
		_ = parent.children[i].toTOMLNode(tParent.Children[i])
	}

	return tParent
}

type yamlSpreader struct{}

func (*yamlSpreader) spread(ctx context.Context, w io.Writer, roots <-chan *Node) <-chan error {
	errc := make(chan error, 1)

	go func() {
		defer close(errc)

		enc := yaml.NewEncoder(w)
	BREAK:
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					break BREAK
				}
				if err := enc.Encode(root.toYAMLNode(nil)); err != nil {
					errc <- err
					return
				}
			}
		}
	}()

	return errc
}

type yamlNode struct {
	Name     string      `yaml:"value"`
	Children []*yamlNode `yaml:"children"`
}

func (parent *Node) toYAMLNode(yParent *yamlNode) *yamlNode {
	if yParent == nil {
		yParent = &yamlNode{Name: parent.name}
	}
	if !parent.hasChild() {
		return yParent
	}

	yParent.Children = make([]*yamlNode, len(parent.children))
	for i := range parent.children {
		yParent.Children[i] = &yamlNode{Name: parent.children[i].name}
		_ = parent.children[i].toYAMLNode(yParent.Children[i])
	}

	return yParent
}

var (
	_ spreader = (*defaultSpreader)(nil)
	_ spreader = (*colorizeSpreader)(nil)
	_ spreader = (*jsonSpreader)(nil)
	_ spreader = (*yamlSpreader)(nil)
	_ spreader = (*tomlSpreader)(nil)
)
