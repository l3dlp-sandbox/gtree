package gtree

import (
	"io"
)

// Output outputs a tree to w with r as Markdown format input.
func Output(w io.Writer, r io.Reader, options ...Option) error {
	conf, err := newConfig(options...)
	if err != nil {
		return err
	}
	rg := newRootGenerator(r, conf.space)
	roots, err := rg.generate()
	if err != nil {
		return err
	}

	tree := initializeTree(conf, roots)
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.spread(w)
}

// Mkdir makes directories.
func Mkdir(r io.Reader, options ...Option) error {
	conf, err := newConfig(options...)
	if err != nil {
		return err
	}
	rg := newRootGenerator(r, conf.space)
	roots, err := rg.generate()
	if err != nil {
		return err
	}

	tree := initializeTree(conf, roots)
	if err := tree.grow(); err != nil {
		return err
	}
	return tree.mkdir()
}

func initializeTree(conf *config, rs []*Node) *tree {
	g := newGrower(conf.encode, conf.lastNodeFormat, conf.intermedialNodeFormat, conf.dryrun)
	s := newSpreader(conf.encode)
	m := newMkdirer(conf.fileExtensions)
	return newTree(rs, g, s, m)
}

type tree struct {
	roots    []*Node
	grower   grower
	spreader spreader
	mkdirer  mkdirer
}

func newTree(
	roots []*Node,
	grower grower,
	spreader spreader,
	mkdirer mkdirer,
) *tree {
	return &tree{
		roots:    roots,
		grower:   grower,
		spreader: spreader,
		mkdirer:  mkdirer,
	}
}

func (t *tree) grow() error {
	return t.grower.grow(t.roots)
}

func (t *tree) spread(w io.Writer) error {
	return t.spreader.spread(w, t.roots)
}

func (t *tree) mkdir() error {
	return t.mkdirer.mkdir(t.roots)
}
