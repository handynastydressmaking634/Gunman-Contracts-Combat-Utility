-- Build: 113f44d566165d54e12e891fabf942fe
local M = {}

function M.clamp(value, minimum, maximum)
  return math.max(minimum, math.min(maximum, value))
end

return M
