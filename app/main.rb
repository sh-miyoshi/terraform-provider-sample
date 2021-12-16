require "sinatra"
require "json"
require 'sinatra/reloader' if development?
require 'securerandom'

#----------------------------------------
# For VM resource
#----------------------------------------
get '/vm' do
  content_type :json
  data = load()
  data["vm"].to_json
end

get '/vm/:id' do
  data = load()
  data["vm"].each do |vm|
    if vm["id"] == params[:id]
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

  data = load()
  data["vm"].push(v)
  save!(data)
  content_type :json
  v.to_json
end

delete '/vm/:id' do
  data = load()
  data["vm"].delete_if{|vm| vm["id"] == params[:id]}
  save!(data)

  201
end

put '/vm/:id' do
  data = load()
  body = JSON.parse request.body.read
  updated = false
  data["vm"].map! do |vm|
    if vm["id"] == params[:id]
      updated = true
      {
        id:   vm["id"],
        name: body["name"],
        spec: body["spec"]
      }
    else
      vm
    end
  end
  data["vm"].push({
    id:   params[:id],
    name: body["name"],
    spec: body["spec"]
  }) if !updated

  save!(data)
  201
end

#----------------------------------------
# For Storage resource
#----------------------------------------
get '/storage' do
  content_type :json
  data = load()
  data["storage"].to_json
end

get '/storage/:id' do
  data = load()
  data["storage"].each do |storage|
    if storage["id"] == params[:id]
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

  data = load()
  data["storage"].push(v)
  save!(data)
  content_type :json
  v.to_json
end

delete '/storage/:id' do
  data = load()
  data["storage"].delete_if{|storage| storage["id"] == params[:id]}
  save!(data)

  201
end

put '/storage/:id' do
  data = load()
  body = JSON.parse request.body.read
  updated = false
  data["storage"].map! do |storage|
    if storage["id"] == params[:id]
      updated = true
      {
        id:   storage["id"],
        name: body["name"],
        spec: body["spec"]
      }
    else
      storage
    end
  end
  data["storage"].push({
    id:   params[:id],
    name: body["name"],
    spec: body["spec"]
  }) if !updated

  save!(data)
  201
end

def load()
  begin
    File.open("data.json") do |fp|
      fp.flock(File::LOCK_SH)
      JSON.load(fp)
    end
  rescue
    data = {
      "vm" => [],
      "storage" => []
    }
  end
end

def save!(data)
  File.open("data.json","w") do |fp|
    fp.flock(File::LOCK_EX)
    fp.write(JSON.pretty_generate(data))
  end
end
