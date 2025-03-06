package api

type APIServer struct {
    taskManager *process.TaskManager
    vfs         *fs.VirtualFileSystem
    router      *gin.Engine
}

func NewAPIServer(tm *process.TaskManager, vfs *fs.VirtualFileSystem) *APIServer {
    server := &APIServer{
        taskManager: tm,
        vfs:         vfs,
        router:      gin.Default(),
    }
    
    server.router.POST("/tasks", server.createTask)
    server.router.GET("/files/:name", server.readFile)
    return server
}

func (s *APIServer) Run(addr string) error {
    return s.router.Run(addr)
}