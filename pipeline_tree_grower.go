//go:build !wasm

package gtree

import (
	"context"
	"sync"
)

func newGrowerPipeline(
	lastNodeFormat, intermedialNodeFormat branchFormat,
	enabledValidation bool,
) growerPipeline {
	return &defaultGrowerPipeline{
		lastNodeFormat:        lastNodeFormat,
		intermedialNodeFormat: intermedialNodeFormat,
		enabledValidation:     enabledValidation,
	}
}

func newNopGrowerPipeline() growerPipeline {
	return &nopGrowerPipeline{}
}

type defaultGrowerPipeline struct {
	lastNodeFormat        branchFormat
	intermedialNodeFormat branchFormat
	enabledValidation     bool
}

const workerGrowNum = 10

func (dg *defaultGrowerPipeline) grow(ctx context.Context, roots <-chan *Node) (<-chan *Node, <-chan error) {
	nodes := make(chan *Node)
	errc := make(chan error, 1)

	go func() {
		defer func() {
			close(nodes)
			close(errc)
		}()

		wg := &sync.WaitGroup{}
		for i := 0; i < workerGrowNum; i++ {
			wg.Add(1)
			go dg.worker(ctx, wg, roots, nodes, errc)
		}
		wg.Wait()
	}()

	return nodes, errc
}

func (dg *defaultGrowerPipeline) worker(ctx context.Context, wg *sync.WaitGroup, roots <-chan *Node, nodes chan<- *Node, errc chan<- error) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case root, ok := <-roots:
			if !ok {
				return
			}
			if err := dg.assemble(root); err != nil {
				errc <- err
				return
			}
			nodes <- root
		}
	}
}

func (dg *defaultGrowerPipeline) assemble(current *Node) error {
	if err := dg.assembleBranch(current); err != nil {
		return err
	}

	for _, child := range current.children {
		if err := dg.assemble(child); err != nil {
			return err
		}
	}
	return nil
}

func (dg *defaultGrowerPipeline) assembleBranch(current *Node) error {
	current.clean() // 例えば、MkdirProgrammably funcでrootノードを使いまわすと、前回func実行時に形成されたノードの枝が残ったまま追記されてしまうため。

	dg.assembleBranchDirectly(current)

	// go back to the root to form a branch.
	tmpParent := current.parent
	if tmpParent != nil {
		for ; !tmpParent.isRoot(); tmpParent = tmpParent.parent {
			dg.assembleBranchIndirectly(current, tmpParent)
		}
	}

	dg.assembleBranchFinally(current, tmpParent)

	if dg.enabledValidation {
		return current.validatePath()
	}
	return nil
}

func (dg *defaultGrowerPipeline) assembleBranchDirectly(current *Node) {
	if current == nil || current.isRoot() {
		return
	}

	current.setPath(current.name)

	if current.isLastOfHierarchy() {
		current.setBranch(current.branch(), dg.lastNodeFormat.directly)
	} else {
		current.setBranch(current.branch(), dg.intermedialNodeFormat.directly)
	}
}

func (dg *defaultGrowerPipeline) assembleBranchIndirectly(current, parent *Node) {
	if current == nil || parent == nil || current.isRoot() {
		return
	}

	current.setPath(parent.name, current.path())

	if parent.isLastOfHierarchy() {
		current.setBranch(dg.lastNodeFormat.indirectly, current.branch())
	} else {
		current.setBranch(dg.intermedialNodeFormat.indirectly, current.branch())
	}
}

func (*defaultGrowerPipeline) assembleBranchFinally(current, root *Node) {
	if current == nil {
		return
	}

	if root != nil {
		current.setPath(root.path(), current.path())
	}
}

func (dg *defaultGrowerPipeline) enableValidation() {
	dg.enabledValidation = true
}

type nopGrowerPipeline struct{}

func (*nopGrowerPipeline) grow(ctx context.Context, roots <-chan *Node) (<-chan *Node, <-chan error) {
	nodes := make(chan *Node)
	errc := make(chan error, 1)

	go func() {
		defer func() {
			close(nodes)
			close(errc)
		}()

	BREAK:
		for {
			select {
			case <-ctx.Done():
				return
			case root, ok := <-roots:
				if !ok {
					break BREAK
				}
				nodes <- root
			}
		}
	}()

	return nodes, errc
}

func (*nopGrowerPipeline) enableValidation() {}

var (
	_ growerPipeline = (*defaultGrowerPipeline)(nil)
	_ growerPipeline = (*nopGrowerPipeline)(nil)
)