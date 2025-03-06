package syscall

type TaskOSCalls interface {
    Exec(command string) (taskID string)
    FileWrite(name string, data []byte) error
    FileRead(name string) ([]byte, error)
}

type RealCalls struct {
    tm  *process.TaskManager
    vfs *fs.VirtualFileSystem
}

func (rc *RealCalls) Exec(command string) string {
    task := rc.tm.StartTask(command)
    return task.ID.String()
}