local key = KEYS[1]

local capacity = tonumber(ARGV[1])
local refillTokens = tonumber(ARGV[2])
local refillInterval = tonumber(ARGV[3])
local currentTime = tonumber(ARGV[4])
local ttl = tonumber(ARGV[5])

local bucket = redis.call(
    "HMGET",
    key,
    "tokens",
    "last_refill"
)

local tokens = tonumber(bucket[1])
local lastRefill = tonumber(bucket[2])

if tokens == nil then
    tokens = capacity
    lastRefill = currentTime
end

local elapsed = currentTime - lastRefill

local intervals = math.floor(elapsed / refillInterval)

if intervals > 0 then
    tokens = math.min(
        capacity,
        tokens + (intervals * refillTokens)
    )

    lastRefill = lastRefill + (intervals * refillInterval)
end

if tokens < 1 then
    return 0
end

tokens = tokens - 1

redis.call(
    "HMSET",
    key,
    "tokens",
    tokens,
    "last_refill",
    lastRefill
)

redis.call("EXPIRE", key, ttl)

return 1