package fs

type VirtualFileSystem struct {
    files map[string][]byte
    mu    sync.RWMutex
}

func NewVFS() *VirtualFileSystem {
    return &VirtualFileSystem{
        files: make(map[string][]byte),
    }
}

func (v *VirtualFileSystem) WriteFile(name string, data []byte) error {
    v.mu.Lock()
    defer v.mu.Unlock()
    v.files[name] = data
    return nil
}

func (v *VirtualFileSystem) ReadFile(name string) ([]byte, error) {
    v.mu.RLock()
    defer v.mu.RUnlock()
    data, exists := v.files[name]
    if !exists {
        return nil, fmt.Errorf("file not found")
    }
    return data, nil
}