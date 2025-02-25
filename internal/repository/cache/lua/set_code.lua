--验证码在redis上的key
local key = KEYS[1]
--验证次数，一个验证码，最多重复三次，这个记录验证了几次
local cntKey = key..":cnt"
--验证码 123456
local val = ARGV[1]
--过期时间
local ttl = tonumber(redis.call("ttl", key))

if ttl == -1 then
    --key存在，但是没有过期时间
    return -2
elseif ttl == -2 or ttl < 540 then
    redis.call("set", key, val)
    redis.call("expire", key, 600)
    redis.call("set", cntKey, 3)
    redis.call("expire", cntKey, 600)
    return 0
else
    return -1
end
