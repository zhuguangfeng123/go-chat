local key = KEYS[1]
local cntKey = key..":cnt"
redis.call("del", key)
redis.call("del", cntKey)


