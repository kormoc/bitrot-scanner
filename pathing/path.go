package pathing

import (
	"errors"
	"os"
	"strings"
	"sync"
)

type Path struct {
	parent   *Path
	name     string
	children []*Path
}

var pathSeparator = string(os.PathSeparator)
var root *Path
var rootMutex = &sync.Mutex{}

var ErrInvalidPath = errors.New("Invalid Path")

func init() {
	root = newPath()
}

func newPath() *Path {
	path := Path{}
	path.children = make([]*Path, 0, 0)
	return &path
}

func New(path string) (child *Path, err error) {
	if !strings.HasPrefix(path, pathSeparator) {
		err = ErrInvalidPath
		return
	}

	// Trim trailing and prepended path separators
	path = strings.Trim(path, pathSeparator)

	paths := strings.Split(path, pathSeparator)

	rootMutex.Lock()
	defer rootMutex.Unlock()

	child = root.GetPath(paths)

	return
}

func (path *Path) GetPath(searchPath []string) *Path {
	// End of the path? Just return
	if len(searchPath) == 0 {
		return path
	}

	subpath := searchPath[0]
	searchPath = searchPath[1:]

	// Search the children for a subpath match
	for _, child := range path.children {
		if child.name == subpath {
			return child.GetPath(searchPath)
		}
	}

	// Add a new child
	newChild := newPath()
	newChild.parent = path
	newChild.name = subpath
	path.children = append(path.children, newChild)
	return newChild.GetPath(searchPath)
}

func (path *Path) String() (pathStr string) {
	if path == nil {
		return
	}
	if path.parent != nil {
		pathStr = path.parent.String() + pathSeparator
	}
	pathStr += path.name
	return
}
