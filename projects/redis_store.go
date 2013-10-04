package projects

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
    "github.com/skydb/sky.go"
    "net/url"
    "sync"
)

type RedisStore struct {
    sync.Mutex
    client sky.Client
    uri *url.URL
    conn redis.Conn
    projects map[string]*Project
}

// Creates a new Redis-backed project store.
func NewRedisStore(client sky.Client, uri *url.URL) *RedisStore {
    return &RedisStore{
        client: client,
        uri: uri,
        projects: make(map[string]*Project),
    }
}

// Retrieves the host and port used to connect to Redis.
func (s *RedisStore) Host() string {
    return s.uri.Host
}

// Retrieves the hash key used to lookup projects.
func (s *RedisStore) HashKey() string {
    if s.uri.Path == "" {
        return "ldmk_projects"
    }

    return s.uri.Path[1:]
}


// Opens a connection to Redis.
func (s *RedisStore) Open() error {
    s.Lock()
    defer s.Unlock()

    // Close the store if it's already open.
    s.close()

    // Open a new connection to Redis.
    var err error
    if s.conn, err = redis.Dial("tcp", s.Host()); err != nil {
        return err
    }

    return nil
}

// Cleans up any remote connections.
func (s *RedisStore) Close() {
    s.Lock()
    defer s.Unlock()
    s.close()
}

func (s *RedisStore) close() {
    if s.conn != nil {
        s.conn.Close()
        s.conn = nil
    }
}

// Looks up a project by API key. The store will look for a cached copy
// first and then check Redis.
func (s *RedisStore) FindByApiKey(apiKey string) (*Project, error) {
    s.Lock()
    defer s.Unlock()

    // Find local reference to project.
    p := s.projects[apiKey]

    // If project doesn't exist then find it from Redis.
    if p == nil {
        // Find project from Redis.
        reply, err := s.conn.Do("HGET", s.HashKey(), apiKey)
        if err != nil {
            return nil, err
        }

        // Create project based on reply.
        if reply, ok := reply.(string); ok {
            table, err := s.client.GetTable(reply)
            if err != nil {
                return nil, err
            }
            p = New(apiKey, table)
            s.projects[apiKey] = p
        } else {
            return nil, fmt.Errorf("Project not found: %s (%v)", apiKey, reply)
        }
    }

    return p, nil
}
