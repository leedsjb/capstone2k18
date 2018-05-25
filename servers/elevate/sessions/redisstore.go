package sessions

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

//RedisStore represents a session.Store backed by redis.
type RedisStore struct {
	//Redis client used to talk to redis server.
	Client *redis.Client
	//Used for key expiry time on redis.
	SessionDuration time.Duration
}

//NewRedisStore constructs a new RedisStore
func NewRedisStore(client *redis.Client, sessionDuration time.Duration) *RedisStore {
	//initialize and return a new RedisStore struct
	return &RedisStore{client, sessionDuration}
}

//Store implementation

//Save saves the provided `sessionState` and associated SessionID to the store.
//The `sessionState` parameter is typically a pointer to a struct containing
//all the data you want to associated with the given SessionID.
func (rs *RedisStore) Save(sid SessionID, sessionState interface{}) error {
	//marshal the `sessionState` to JSON
	state, err := json.Marshal(sessionState)
	if err != nil {
		return fmt.Errorf("Error marshaling JSON: %v", err)
	}
	//and save it in the redis database, using `sid.getRedisKey()` for the key.
	if err = rs.Client.Set(sid.getRedisKey(), state, rs.SessionDuration).Err(); err != nil {
		return fmt.Errorf("Error saving state: %v", err)
	}
	return nil
}

//Get populates `sessionState` with the data previously saved
//for the given SessionID
func (rs *RedisStore) Get(sid SessionID, sessionState interface{}) error {
	//get the previously-saved session state data from redis
	state, err := rs.Client.Get(sid.getRedisKey()).Result()
	if err != nil {
		return ErrStateNotFound
	}
	//unmarshal it back into the `sessionState` parameter
	if err := json.Unmarshal([]byte(state), sessionState); err != nil {
		return fmt.Errorf("Error unmarshaling state: %v", err)
	}
	//and reset the expiry time, so that it doesn't get deleted until
	//the SessionDuration has elapsed.
	if err := rs.Client.Expire(sid.getRedisKey(), rs.SessionDuration).Err(); err != nil {
		return fmt.Errorf("Error resetting expiration time: %v", err)
	}
	return nil
}

//Delete deletes all state data associated with the SessionID from the store.
func (rs *RedisStore) Delete(sid SessionID) error {
	//delete the data stored in redis for the provided SessionID
	err := rs.Client.Del(sid.getRedisKey()).Err()
	if err != nil {
		return fmt.Errorf("Error deleting session: %v", err)
	}
	return nil
}

//getRedisKey() returns the redis key to use for the SessionID
func (sid SessionID) getRedisKey() string {
	//convert the SessionID to a string and add the prefix "sid:" to keep
	//SessionID keys separate from other keys that might end up in this
	//redis instance
	return "sid:" + sid.String()
}
