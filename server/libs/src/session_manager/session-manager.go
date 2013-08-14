/*
Session Management and User Login

- Get the session id from the cookie
- if the session exists
	- check that it's valid in the sessio manager
	- from session manager update the time stamp
	- retrieve the user data
- else the session DOESN'T exist
	- redirct the user to the login page

- Login
	- Check the username + password are valid in the user DB
	- if credentials valid
		- read user data from the db
		- have session manager store user data and create a session ID for the user
		- session manager returns user data and session token.
	- else provided invalid credentials
		- redirect to login "Invalid username or password."
*/

package session_manager

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"session_manager/hashes"
	"time"
	"uuid"
)

const (
	KEY_USERS    = "users:"
	KEY_SESSIONS = "sessions:"
	FLUSHALL     = "FLUSHALL"
	UUID_LEN     = 36
)

func SessionCount() (int, error) {
	res, err := GetConn().Do(hashes.HLEN, KEY_SESSIONS)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return redis.Int(res, err)
}

func ReadSessions() (s map[string]string, err error) {

	res, err := GetConn().Do(hashes.HGETALL, KEY_USERS)

	if err != nil {
		fmt.Println(err)
		return s, err
	}

	pairs, _, err := ReadPairs(res)

	return pairs, err
}

func SessionExists(sessionId string) (bool, error) {
	reply, err := GetConn().Do(hashes.HEXISTS, KEY_SESSIONS, sessionId)

	if err != nil {
		fmt.Println(err, reply)
		return false, err
	}

	return redis.Bool(reply, err)
}

func SessionValue(sessionId string) (userId, timestamp string, err error) {

	reply, err := GetConn().Do(hashes.HGET, KEY_SESSIONS, sessionId)

	if err != nil {
		fmt.Println(err, reply)
		return userId, timestamp, err
	}

	val, err := redis.String(reply, err)

	if err != nil {
		return userId, timestamp, err
	}

	if len(val) < UUID_LEN {
		return userId, timestamp, NewSessionError(
			"Session value %v isn't in the proper form.", val)
	}

	uuid_start := len(val) - UUID_LEN

	return val[uuid_start:], val[:uuid_start-1], nil
}

/**
 * session: =>
 *		sessionId => time-stamp/user_guid
 */
func AddSession(internalId string) (string, string, error) {
	return UpdateSession(uuid.New(), internalId)
}

func UpdateSession(sessionId, internalId string) (string, string, error) {

	timestamp := time.Now().String()

	// Setting up Timestamp first so that the values can
	// be sorted if that's a consideration.
	_, err := GetConn().Do(hashes.HMSET,
		KEY_SESSIONS,
		sessionId,
		internalId+"/"+timestamp)

	if err != nil {
		return sessionId, timestamp, err
	}

	return sessionId, timestamp, err
}
