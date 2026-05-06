function envoy_on_request(handle)
    local headers = handle:headers()
  
    local request_id = headers:get("x-request-id")
    local client_id  = headers:get("x-client-id")
    
    handle:logInfo("LUA CALLED")

    if request_id == nil then
      handle:respond({[":status"]="400"}, "Missing x-request-id")
      return
    end
  
    if client_id == nil then
      handle:respond({[":status"]="400"}, "Missing x-client-id")
      return
    end
  
    if headers:get("x-session-id") == nil then
      headers:add("x-session-id", handle:streamInfo():requestID())
    end
  end
  