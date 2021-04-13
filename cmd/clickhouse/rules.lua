ch = require("clickhouse")
log = require("log")
options = require("options")

q = "SELECT count(*) FROM platform_raw.tch WHERE link_id = " .. options.get() .. " AND date BETWEEN '2021-01-01' AND '2021-03-01'"
resAll, err = ch.query(q)
if err ~= nil then
    print("err =", err)
    return 0
end
log.info("ResultAll = " .. resAll)

q = "SELECT count(*) FROM platform_raw.tch WHERE link_id = " .. options.get() .. " AND date BETWEEN '2021-01-01' AND '2021-03-01' AND minus(time_interval_2, time_interval_1) >= 10"
resTimelog, err = ch.query(q)
if err ~= nil then
    print("err =", err)
    return 0
end
log.info("ResultTimelog = " .. resTimelog)

if resTimelog / resAll > 0.2 then
  bad = resTimelog / resAll
  statusTraff = 1 - ((bad - 0.2) / 0.8)
  traff = math.ceil(statusTraff * 100) / 100
  log.info("statusTraff = " .. traff)
end

return 1
