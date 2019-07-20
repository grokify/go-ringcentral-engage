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
doc = '  {
    "id":"5c8bf2c8d6cb0005214fab49",
    "created_at":"2019-03-15T18:45:28Z",
    "updated_at":"2019-05-06T07:26:22Z",
    "community_id":"5c8bf29b920026341f267d5a",
    "community_url":null,
    "company":null,
    "email":"test.rc-platform@email.dimelo.com",
    "firstname":null,
    "foreign_id":"test.rc-platform@email.dimelo.com",
    "gender":null,
    "home_phone":null,
    "identity_group_id":"5c8bf2c8d6cb0005214fab47",
    "lastname":null,
    "mobile_phone":null,
    "screenname":"Test email",
    "user_ids":[
      "5c8bf3ecd6cb009c9b4f9ae9",
      "5ccb970fd6cb005f918292e1"
    ],
    "uuid":"test.rc-platform@email.dimelo.com",
    "extra_values":{

    },
    "display_name":"Test email",
    "avatar_url":"https://rc-platform.engagement.dimelo.com/assets/default_avatar/thumb-25da1b581731c9bd0669fa2f2df6a4bc.png",
    "type":"emails"
  }'

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