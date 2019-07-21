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
doc = '
  {
    "id":"5c8bf22b14bf8a84c44c7753",
    "created_at":"2019-03-15T18:42:51Z",
    "updated_at":"2019-03-15T18:42:51Z",
    "label":"Agent",
    "approve_content":false,
    "assign_intervention":false,
    "author_block_content":false,
    "admin_stamp_answer":true,
    "close_content_thread":true,
    "mute_content":false,
    "update_own_intervention":true,
    "update_intervention":false,
    "publish_content":true,
    "delay_export_content":true,
    "receive_tasks":true,
    "search_contents":true,
    "open_content_thread":true,
    "impersonate_user":false,
    "delete_content_thread":false,
    "use_emoji":true,
    "access_pull_mode":true,
    "read_event":true,
    "read_presence":true,
    "read_identity":true,
    "access_previous_messages":true,
    "use_cobrowsing":true,
    "invite_user":false,
    "create_user":false,
    "read_user":true,
    "update_user":false,
    "manage_users_of_my_teams":false,
    "manage_identities":false,
    "update_identity":true,
    "manage_teams":false,
    "manage_roles":false,
    "manage_own_notifications":false,
    "manage_categories":false,
    "manage_folders":false,
    "manage_custom_notifications":false,
    "read_community":false,
    "create_community":false,
    "update_community":false,
    "read_content_source":false,
    "create_content_source":false,
    "update_content_source":false,
    "manage_chat":false,
    "manage_messaging":false,
    "manage_topologies":false,
    "read_export":false,
    "search_event":false,
    "update_settings":false,
    "manage_tags":false,
    "manage_custom_fields":false,
    "manage_emails_templates":false,
    "manage_api_access_tokens":false,
    "create_and_destroy_extension":false,
    "update_extension":false,
    "update_time_sheet":false,
    "manage_app_sdk_applications":false,
    "access_help_center":true,
    "manage_ice":false,
    "export_identity":false,
    "anonymize_identity":false,
    "lock_identity":false,
    "read_stats":false,
    "read_own_stats":true,
    "monitor_tasks":false,
    "monitor_team_tasks":false,
    "manage_reply_assistant":false,
    "reply_with_assistant":true,
    "manage_rules_engine_rules":false
  }
'

doc = '
{  "id": "5b87a6b2f042de5f94fabf8d",

  "created_at": "2018-08-30T08:11:30Z",
  "updated_at": "2018-08-30T08:11:30Z",
  "content_type": "image/jpeg",
  "filename": "cat.jpeg",
  "foreign_id": null,
  

  "size": 700754,
  "url": "http://domain-test.engagement.dimelo.dev/attachments/5b87a6b2f042de5f94fabf8d",
 

  "video_metadata": [],
  "embed": false,
  "public?": true

}
'
doc = '
{
  "id": "60944e5702bdafb74ec96142",
  "parent_id": "eb3c62690ec9025845fd3495",
  "name": "Technical",
  "created_at": "2012-05-23T01:12:49Z",
  "updated_at": "2012-05-23T01:12:49Z",
  "color": 0,
  "mandatory": false,
  "multiple": true,
  "post_qualification": false,
  "source_ids": [],
  "unselectable": false  }'

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
  elsif v.is_a? Integer
    prop[:type] = 'integer'
  end
  if v.kind_of?(Array)
    prop[:type] = 'array'
    prop[:items] = { type: "string" }
  end

  properties[name] = prop
end

puts doc

puts MultiJson.dump properties, pretty: true