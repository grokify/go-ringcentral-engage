#!ruby

require 'multi_json'

doc = 'email: To filter users on given email.
category_id: To filter users on given category id.
identity_id: To filter users on given identity id.
external_id: To filter users on given external id.
role_id: To filter users on given role id.
team_id: To filter users on given team id.'

parameters = []

doc.split("\n").each do |line|
  if line =~ /^\s*(\w+)\s*:\s*(.+)\s*$/
    name = $1
    desc = $2
    param = {
      name: name,
      description: desc,
      in: 'query',
      type: 'string',
      required: false
    }
    parameters.push param
  else
    puts "E_NO_MATCH: " + line
  end
end

puts doc

puts MultiJson.encode parameters, pretty: true