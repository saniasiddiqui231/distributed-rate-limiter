local key = KEYS[1]

local limit = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local now = tonumber(ARGV[3])


redis.call(
    "ZREMRANGEBYSCORE",
    key,
    "-inf",
    now - window
)


local count = redis.call("ZCARD", key)

if count >= limit then
    return 0
end


redis.call(
    "ZADD",
    key,
    now,
    ARGV[4]
)

redis.call(
    "EXPIRE",
    key,
    window
)

return 1