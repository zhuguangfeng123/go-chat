local key = KEYS[1]
local cntKey = key..":cnt"
--用户输入的code
local expectedCode   = ARGV[1]
local code = redis.call("get", key)

--转成一个数字
local cnt = tonumber(redis.call("get", "phone_code:login:18860313695:cnt"))
if cnt <= 0 then
    --用户一直输错
    return -1
elseif expectedCode == code then
    -- 正确
    redis.call("del", cntKey)
    return 0
else
    redis.call("decr", cntKey)
    return -2
end