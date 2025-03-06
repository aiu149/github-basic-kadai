package process

type Task struct {
    ID        uuid.UUID
    Command   string
    Status    string
    CreatedAt time.Time
}

type TaskManager struct {
    tasks map[uuid.UUID]*Task
    mu    sync.RWMutex
}

func NewTaskManager() *TaskManager {
    return &TaskManager{
        tasks: make(map[uuid.UUID]*Task),
    }
}

func (tm *TaskManager) StartTask(cmd string) *Task {
    task := &Task{
        ID:        uuid.New(),
        Command:   cmd,
        Status:    "running",
        CreatedAt: time.Now(),
    }
    
    go tm.executeTask(task)
    return task
}

func (tm *TaskManager) executeTask(task *Task) {
    // タスク実行ロジック
    defer func() {
        tm.mu.Lock()
        task.Status = "completed"
        tm.mu.Unlock()
    }()
}