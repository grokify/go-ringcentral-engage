#!ruby

require 'multi_json'

doc = 'email: To filter users on given email.
category_id: To filter users on given category id.
identity_id: To filter users on given identity id.
external_id: To filter users on given external id.
role_id: To filter users on given role id.
team_id: To filter users on given team id.'

doc = 'name: Team name.
leader_ids[]: List of user id as leaders | array:string.
user_ids[]: List of user id as team members. | array:string'

doc = 'label: Folder’s label (mandatory).
parent_id: ID of the parent folder.
position: position of the folder. | integer
query: query of the folder as described in ​Search API documentation.​\n\nExample: “​active_and_assigned_to_me:true”
render_threads_count: boolean describing display of the number of threads. | boolean
role_restriction[only][]: list of roles allowed to see this folder. This parameter has to be a hash otherwise it will raise a 400 error. The key should be "only". For example: `&role_restriction[only][]=4e5596cdae70f677b5000002` | array:string
team_restriction[only][]: list of teams allowed to see this folder. Same thing as role_restriction: team_restriction parameter has to be a hash with the key "only". | array:string'

parameters = []

doc.split("\n").each do |line|
  if line =~ /^\s*([\w\[\]]+)\s*:\s*(.+)\s*$/
    name = $1
    desc = $2.strip
    type = ''
    format = ''
    items = {}
    if desc =~ /^\s*(.+)\s*\|\s*(.+):(.+)\s*$/
      desc = $1.strip
      type = $2
      if type == 'array'
        items = {type: $3}
      end
    elsif desc =~ /^\s*(.+)\s*\|\s*(.+)\s*$/
      desc = $1
      type = $2
    end
    param = {
      name: name,
      description: desc,
      in: 'query',
      type: 'string',
      required: false
    }
    if type.length > 0
      param[:type] = type
    end
    if format.length > 0
      param[:format] = format
    end
    if items.length >0 
      param[:items] = items
    end
    parameters.push param
  else
    puts "E_NO_MATCH: " + line
  end
end

puts doc

puts MultiJson.encode parameters, pretty: true