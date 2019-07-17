#!ruby

require 'multi_json'

doc = '
    {
      "id":"5cde163bd6cb00d3034e2332",
      "created_at":"2019-05-17T02:02:35Z",
      "updated_at":"2019-05-17T02:31:03Z",
      "category_ids":[

      ],
      "email":"embbnux.ji@ringcentral.com",
      "enabled":true,
      "external_id":null,
      "firstname":"Nux",
      "gender":null,
      "identity_ids":[

      ],
      "lastname":"Ji",
      "locale":"en",
      "nickname":null,
      "rc_user_id":null,
      "role_id":"5c8bf22b14bf8a84c44c7750",
      "spoken_languages":[
        "en"
      ],
      "team_ids":[

      ],
      "timezone":"",
      "invitation_pending":false
    }
'

properties = {}

object = MultiJson.load doc

object.each do |name,v|
  prop = {
    type: 'string'
  }
  if v =~ /^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z$/
    prop[:format] = 'date-time'
  end

  if !!v == v
    prop[:type] = 'boolean'
  end
  if v.kind_of?(Array)
    prop[:type] = 'array'
    prop[:items] = { type: "string" }
  end

  properties[name] = prop
end

puts doc

puts MultiJson.dump properties, pretty: true