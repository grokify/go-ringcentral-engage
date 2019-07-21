#!ruby

require 'multi_json'

doc = 'name: Team name.
leader_ids[]: List of user id as leaders | array:string
user_ids[]: List of user id as team members. | array:string'

doc = 'label: Folder’s label (mandatory).
parent_id: ID of the parent folder.
position: position of the folder. | integer
query: query of the folder as described in ​Search API documentation.​\n\nExample: “​active_and_assigned_to_me:true”
render_threads_count: boolean describing display of the number of threads. | boolean
role_restriction[only][]: list of roles allowed to see this folder. This parameter has to be a hash otherwise it will raise a 400 error. The key should be "only". For example: `&role_restriction[only][]=4e5596cdae70f677b5000002` | array:string
team_restriction[only][]: list of teams allowed to see this folder. Same thing as role_restriction: team_restriction parameter has to be a hash with the key "only". | array:string'

doc = '● name: Source name
● active: Activate/deactivate the source (Boolean)
● channel_id: Channel
● color: Color of the icon (see S​ ource colors​) (Integer)
● sla_response: Response time (seconds)
● sla_expired_strategy: SLA expired strategy ("max", "half" or "base")
● intervention_messages_boost: Priority boost of messages with intervention (Integer)
● transferred_tasks_boost: Priority boost of transferred tasks (Integer)
● hidden_from_stats: Hide from statistics (Boolean)
● default_category_ids[]: Default categories | array:string
● user_thread_default_category_ids[]: Default categories (agent messages) | array:string
● default_content_language: Default content language
● auto_detect_content_language: Auto-detect content language (Boolean)
● content_archiving: Automatic archiving of old contents (Boolean)
● content_archiving_period: Archive contents older than (seconds)'

doc = '● community_id: To filter identities on given community id.
● identity_group_id: To filter on given group id.
● user_id: To filter identities on given user id.
● sort: To change the criteria chosen to sort the identities. The value can be “created_at” or
“updated_at”.
● foreign_id: To filter identities on given user id
● uuid: To filter identities on given uuid'

doc = '● category_ids: User list of category ids (multiple).
● email: User email (mandatory).
● enabled: Whether the user is enabled or not (boolean).
● external_id: User external id.
● firstname: User firstname (mandatory).
● gender: User gender ("man" or "woman").
● identity_ids: User list of identity ids (multiple).
● lastname: User lastname (mandatory).
● locale: Language for the user interface.
● nickname: User nickname.
● role_id: User role id (mandatory).
● team_ids: User list of team ids (multiple).
● timezone: Use the timezone endpoint to get the timezone name (String), default is empty for
domain timezone.
● spoken_languages: List of locales corresponding to the languages spoken by the user (multiple).'

doc = '
● access_help_center
● access_previous_messages
● access_pull_mode
● admin_stamp_answer
● approve_content
● assign_intervention
● author_block_content
● close_content_thread
● create_and_destroy_extension
● create_community
● create_content_source
● create_user
● delay_export_content
● delete_content_thread
● impersonate_user
● invite_user
● manage_api_access_tokens
● manage_app_sdk_applications
● manage_automatic_exports_tasks *
● manage_categories
● manage_chat
● manage_custom_fields
● manage_custom_notifications
● manage_emails_templates
● manage_folders
● manage_ice
● manage_identities
● manage_own_notifications
● manage_reply_assistant *
● manage_roles
● manage_rules_engine_rules *
● manage_surveys *
● manage_tags
● manage_teams
● manage_topologies
● manage_users_of_my_teams
● monitor_tasks
● monitor_team_tasks
● mute_content
● open_content_thread
● publish_content
● read_community
● read_content_source
● read_event
● read_export
● read_identity
● read_own_stats
● read_presence
● read_stats
● read_surveys *
● read_user
● receive_tasks
● reply_with_assistant *
● search_contents
● search_event
● update_community
● update_content_source
● update_extension
● update_identity
● update_intervention
● update_own_intervention
● update_settings
● update_time_sheet
● update_user
● use_emoji
'

doc = '
● category_ids[]: User list of category ids (multiple).
● email: User email (mandatory).
● enabled: Whether the user is enabled or not (boolean).
● external_id: User external id, used for SSO.
● firstname: User firstname (mandatory).
● gender: User gender ("man" or "woman").
● identity_ids[]: User list of identity ids (multiple).
● lastname: User lastname (mandatory).
● locale: Language for the user interface.
● nickname: User nickname.
● password: User plain password (mandatory).
● role_id: User role id (mandatory).
● team_ids[]: User list of team ids (multiple).
● timezone: Use the timezone endpoint to get the timezone name (String), default is empty for domain timezone.
● spoken_languages[]: List of locales corresponding to the languages spoken by the user (multiple).
'

doc = '
● name: Category name.
● parent_id: ID of parent category.
● color: displayed color for the category, see Category colors. | integer
● mandatory: mandatory categorization (Boolean).
● multiple: allow to assign multiple child categories (Boolean).
● post_qualification: post qualification (Boolean).
● unselectable: root category is unselectable (Boolean).
● source_ids[]: List of source id. | array:string
'
parameters = []

doc.split("\n").each_with_index do |line,i|
  puts "LINE[#{i}] " + line
  if line =~ /^●?\s*([\w\[\]]+)\s*:\s*(.+)\s*$/
    name = $1
    desc = $2.strip
    type = ''
    format = ''
    items = {}
    param = {name: name, in: 'query', required: false}
    if desc =~ /^\s*(.+)\s*\|\s*(.+):(.+)\s*$/
      desc = $1.strip
      type = $2
      if type == 'array'
        items = {type: $3}
      end
    elsif desc =~ /^\s*(.+)\s*\|\s*(.+)\s*\.?\s*$/
      desc = $1
      type = $2
    elsif desc =~/\((\w+)\)/
      type_raw = $1.strip.downcase
      if type_raw == 'seconds'
        type = 'integer'
      elsif type_raw == 'mandatory'
        type = 'string'
        param[:required] = true
      elsif type_raw == 'multiple'
        type = 'array'
        items = {type: 'string'}
      else 
        type = type_raw
      end
    end
    param[:required] = false
    desc.gsub!(/\s*\(mandatory\)/,'')
    param[:description] = desc
    if type.length > 0
      param[:schema] = {} unless param.key? :schema
      param[:schema][:type] = type
    else
        param[:schema] = {} unless param.key? :schema
        param[:schema][:type] = 'string'
    end
    if format.length > 0
      param[:format] = format
    end
    if type == 'array'
      if items.length >0 
        param[:schema][:items] = items
      end
    end
    parameters.push param
  elsif line =~ /^●?\s*(\w+)(\s*\*\s*)?$/
    name = $1.strip
    desc = $2 || ''
    desc.strip!
    param = {name: name, in: 'query', required: false}
    if desc == "*"
      param[:description] = "permission only available with the corresponding extension enabled"
    end
    param[:schema] = {} unless param.key? :schema
    param[:schema][:type] = 'boolean'
    parameters.push param
  else
    puts "E_NO_MATCH: " + line
  end
end

puts doc

puts MultiJson.encode parameters, pretty: true