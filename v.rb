#!/usr/bin/env ruby

require 'YAML'
require 'pp'

f = File.open('w.yml') 
d = f.read()
y = YAML.load(d)
pp y
