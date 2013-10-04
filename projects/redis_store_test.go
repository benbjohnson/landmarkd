package projects

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func newRedisStore() (*RedisStore, *mockRedisConn, *mockSkyClient, *mockSkyTable) {
    table := new(mockSkyTable)
    client := new(mockSkyClient)
    conn := new(mockRedisConn)
    s, err := NewStore(client, "redis://localhost:6379")
    if err != nil {
        panic(err)
    }
    store := s.(*RedisStore)
    store.conn = conn
    return store, conn, client, table
}

// Ensure that the store will find a project by API key and cache it.
func TestRedisStoreFindByApiKey(t *testing.T) {
    store, conn, client, table := newRedisStore()
    client.On("GetTable", "myTable").Return(table, nil)
    conn.On("Do", "HGET", []interface{}{"ldmk_projects", "XXX"}).Return("myTable", nil)

    // Search for a project by API key.
    p, err := store.FindByApiKey("XXX")
    assert.Nil(t, err)
    assert.Equal(t, p.ApiKey, "XXX")
    assert.Equal(t, p.table, table)

    // Search for the same project. Should receive cached copy.
    p2, err := store.FindByApiKey("XXX")
    assert.Nil(t, err)
    assert.Equal(t, p, p2)

    client.AssertExpectations(t)
    conn.AssertExpectations(t)
}
