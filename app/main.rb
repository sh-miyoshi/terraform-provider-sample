require "sinatra"
require "json"
require 'sinatra/reloader' if development?
require 'securerandom'

data = {
  vm: [],
  storage: []
}

#----------------------------------------
# For VM resource
#----------------------------------------
get '/vm' do
  content_type :json
  data[:vm].to_json
end

get '/vm/:id' do
  data[:vm].each do |vm|
    if vm[:id] == params["id"]
      content_type :json
      return vm.to_json
    end
  end

  404
end

post '/vm', provides: :json do
  params = JSON.parse request.body.read

  v = {
    id: SecureRandom.uuid,
    name: params["name"],
    spec: params["spec"]
  }

  data[:vm].push(v)
  content_type :json
  v.to_json
end

delete '/vm/:id' do
  data[:vm].delete_if{|vm| vm[:id] == params["id"]}
end

#----------------------------------------
# For Storage resource
#----------------------------------------
get '/storage' do
  content_type :json
  data[:storage].to_json
end

get '/storage/:id' do
  data[:storage].each do |storage|
    if storage[:id] == params["id"]
      content_type :json
      return storage.to_json
    end
  end

  404
end

post '/storage', provides: :json do
  params = JSON.parse request.body.read

  v = {
    id: SecureRandom.uuid,
    name: params["name"],
    spec: params["spec"]
  }

  data[:storage].push(v)
  content_type :json
  v.to_json
end

delete '/storage/:id' do
  data[:storage].delete_if{|storage| storage[:id] == params["id"]}
end
