// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

const values = {
    INVITE_USER: "invite_user",
    ADD_USER_TO_TEAM: "add_user_to_team",
    MANAGE_SLASH_COMMANDS: "manage_slash_commands",
    MANAGE_OTHERS_SLASH_COMMANDS: "manage_others_slash_commands",
    CREATE_PUBLIC_CHANNEL: "create_public_channel",
    CREATE_PRIVATE_CHANNEL: "create_private_channel",
    MANAGE_PUBLIC_CHANNEL_MEMBERS: "manage_public_channel_members",
    MANAGE_PRIVATE_CHANNEL_MEMBERS: "manage_private_channel_members",
    READ_PUBLIC_CHANNEL_GROUPS: "read_public_channel_groups",
    READ_PRIVATE_CHANNEL_GROUPS: "read_private_channel_groups",
    ASSIGN_SYSTEM_ADMIN_ROLE: "assign_system_admin_role",
    MANAGE_ROLES: "manage_roles",
    MANAGE_TEAM_ROLES: "manage_team_roles",
    MANAGE_CHANNEL_ROLES: "manage_channel_roles",
    MANAGE_SYSTEM: "manage_system",
    CREATE_DIRECT_CHANNEL: "create_direct_channel",
    CREATE_GROUP_CHANNEL: "create_group_channel",
    MANAGE_PUBLIC_CHANNEL_PROPERTIES: "manage_public_channel_properties",
    MANAGE_PRIVATE_CHANNEL_PROPERTIES: "manage_private_channel_properties",
    LIST_PUBLIC_TEAMS: "list_public_teams",
    JOIN_PUBLIC_TEAMS: "join_public_teams",
    LIST_PRIVATE_TEAMS: "list_private_teams",
    JOIN_PRIVATE_TEAMS: "join_private_teams",
    LIST_TEAM_CHANNELS: "list_team_channels",
    JOIN_PUBLIC_CHANNELS: "join_public_channels",
    DELETE_PUBLIC_CHANNEL: "delete_public_channel",
    CONVERT_PUBLIC_CHANNEL_TO_PRIVATE: "convert_public_channel_to_private",
    CONVERT_PRIVATE_CHANNEL_TO_PUBLIC: "convert_private_channel_to_public",
    DELETE_PRIVATE_CHANNEL: "delete_private_channel",
    EDIT_OTHER_USERS: "edit_other_users",
    READ_CHANNEL: "read_channel",
    READ_CHANNEL_CONTENT: "read_channel_content",
    READ_PUBLIC_CHANNEL: "read_public_channel",
    ADD_REACTION: "add_reaction",
    REMOVE_REACTION: "remove_reaction",
    REMOVE_OTHERS_REACTIONS: "remove_others_reactions",
    PERMANENT_DELETE_USER: "permanent_delete_user",
    UPLOAD_FILE: "upload_file",
    GET_PUBLIC_LINK: "get_public_link",
    MANAGE_WEBHOOKS: "manage_webhooks",
    MANAGE_OTHERS_WEBHOOKS: "manage_others_webhooks",
    MANAGE_INCOMING_WEBHOOKS: "manage_incoming_webhooks",
    MANAGE_OTHERS_INCOMING_WEBHOOKS: "manage_others_incoming_webhooks",
    MANAGE_OUTGOING_WEBHOOKS: "manage_outgoing_webhooks",
    MANAGE_OTHERS_OUTGOING_WEBHOOKS: "manage_others_outgoing_webhooks",
    MANAGE_OAUTH: "manage_oauth",
    MANAGE_SYSTEM_WIDE_OAUTH: "manage_system_wide_oauth",
    CREATE_POST: "create_post",
    CREATE_POST_PUBLIC: "create_post_public",
    EDIT_POST: "edit_post",
    EDIT_OTHERS_POSTS: "edit_others_posts",
    DELETE_POST: "delete_post",
    DELETE_OTHERS_POSTS: "delete_others_posts",
    REMOVE_USER_FROM_TEAM: "remove_user_from_team",
    CREATE_TEAM: "create_team",
    MANAGE_TEAM: "manage_team",
    IMPORT_TEAM: "import_team",
    VIEW_TEAM: "view_team",
    LIST_USERS_WITHOUT_TEAM: "list_users_without_team",
    CREATE_USER_ACCESS_TOKEN: "create_user_access_token",
    READ_USER_ACCESS_TOKEN: "read_user_access_token",
    REVOKE_USER_ACCESS_TOKEN: "revoke_user_access_token",
    MANAGE_JOBS: "manage_jobs",
    MANAGE_EMOJIS: "manage_emojis",
    MANAGE_OTHERS_EMOJIS: "manage_others_emojis",
    CREATE_EMOJIS: "create_emojis",
    DELETE_EMOJIS: "delete_emojis",
    DELETE_OTHERS_EMOJIS: "delete_others_emojis",
    VIEW_MEMBERS: "view_members",
    INVITE_GUEST: "invite_guest",
    PROMOTE_GUEST: "promote_guest",
    DEMOTE_TO_GUEST: "demote_to_guest",
    USE_CHANNEL_MENTIONS: "use_channel_mentions",
    USE_GROUP_MENTIONS: "use_group_mentions",
    READ_OTHER_USERS_TEAMS: "read_other_users_teams",
    EDIT_BRAND: "edit_brand",
    READ_JOBS: "read_jobs",
    DOWNLOAD_COMPLIANCE_EXPORT_RESULT: "download_compliance_export_result",
    CREATE_LDAP_SYNC_JOB: "create_ldap_sync_job",
    READ_LDAP_SYNC_JOB: "read_ldap_sync_job",
    TEST_LDAP: "test_ldap",
    GET_SAML_METADATA_FROM_IDP: "get_saml_metadata_from_idp",
    ADD_SAML_PUBLIC_CERT: "add_saml_public_cert",
    ADD_SAML_PRIVATE_CERT: "add_saml_private_cert",
    ADD_SAML_IDP_CERT: "add_saml_idp_cert",
    REMOVE_SAML_PUBLIC_CERT: "remove_saml_public_cert",
    REMOVE_SAML_PRIVATE_CERT: "remove_saml_private_cert",
    REMOVE_SAML_IDP_CERT: "remove_saml_idp_cert",
    GET_SAML_CERT_STATUS: "get_saml_cert_status",
    ADD_LDAP_PUBLIC_CERT: "add_ldap_public_cert",
    ADD_LDAP_PRIVATE_CERT: "add_ldap_private_cert",
    REMOVE_LDAP_PUBLIC_CERT: "remove_ldap_public_cert",
    REMOVE_LDAP_PRIVATE_CERT: "remove_ldap_private_cert",
    INVALIDATE_EMAIL_INVITE: "invalidate_email_invite",
    TEST_SITE_URL: "test_site_url",
    TEST_ELASTICSEARCH: "test_elasticsearch",
    TEST_S3: "test_s3",
    TEST_EMAIL: "test_email",
    RELOAD_CONFIG: "reload_config",
    INVALIDATE_CACHES: "invalidate_caches",
    PURGE_ELASTICSEARCH_INDEXES: "purge_elasticsearch_indexes",
    RECYCLE_DATABASE_CONNECTIONS: "recycle_database_connections",
    CREATE_ELASTICSEARCH_POST_INDEXING_JOB:
        "create_elasticsearch_post_indexing_job",
    CREATE_ELASTICSEARCH_POST_AGGREGATION_JOB:
        "create_elasticsearch_post_aggregation_job",
    READ_ELASTICSEARCH_POST_INDEXING_JOB:
        "read_elasticsearch_post_indexing_job",
    READ_ELASTICSEARCH_POST_AGGREGATION_JOB:
        "read_elasticsearch_post_aggregation_job",
    USE_SLASH_COMMANDS: "use_slash_commands",

    SYSCONSOLE_READ_ABOUT_EDITION_AND_LICENSE:
        "sysconsole_read_about_edition_and_license",
    SYSCONSOLE_WRITE_ABOUT_EDITION_AND_LICENSE:
        "sysconsole_write_about_edition_and_license",
    SYSCONSOLE_READ_BILLING: "sysconsole_read_billing",
    SYSCONSOLE_WRITE_BILLING: "sysconsole_write_billing",
    SYSCONSOLE_READ_REPORTING_SITE_STATISTICS:
        "sysconsole_read_reporting_site_statistics",
    SYSCONSOLE_WRITE_REPORTING_SITE_STATISTICS:
        "sysconsole_write_reporting_site_statistics",
    SYSCONSOLE_READ_REPORTING_TEAM_STATISTICS:
        "sysconsole_read_reporting_team_statistics",
    SYSCONSOLE_WRITE_REPORTING_TEAM_STATISTICS:
        "sysconsole_write_reporting_statistics",
    SYSCONSOLE_READ_REPORTING_SERVER_LOGS:
        "sysconsole_read_reporting_server_logs",
    SYSCONSOLE_WRITE_REPORTING_SERVER_LOGS:
        "sysconsole_write_reporting_server_logs",
    SYSCONSOLE_READ_USERMANAGEMENT_USERS:
        "sysconsole_read_user_management_users",
    SYSCONSOLE_WRITE_USERMANAGEMENT_USERS:
        "sysconsole_write_user_management_users",
    SYSCONSOLE_READ_USERMANAGEMENT_GROUPS:
        "sysconsole_read_user_management_groups",
    SYSCONSOLE_WRITE_USERMANAGEMENT_GROUPS:
        "sysconsole_write_user_management_groups",
    SYSCONSOLE_READ_USERMANAGEMENT_TEAMS:
        "sysconsole_read_user_management_teams",
    SYSCONSOLE_WRITE_USERMANAGEMENT_TEAMS:
        "sysconsole_write_user_management_teams",
    SYSCONSOLE_READ_USERMANAGEMENT_CHANNELS:
        "sysconsole_read_user_management_channels",
    SYSCONSOLE_WRITE_USERMANAGEMENT_CHANNELS:
        "sysconsole_write_user_management_channels",
    SYSCONSOLE_READ_USERMANAGEMENT_PERMISSIONS:
        "sysconsole_read_user_management_permissions",
    SYSCONSOLE_WRITE_USERMANAGEMENT_PERMISSIONS:
        "sysconsole_write_user_management_permissions",
    SYSCONSOLE_READ_USERMANAGEMENT_SYSTEM_ROLES:
        "sysconsole_read_user_management_system_roles",
    SYSCONSOLE_WRITE_USERMANAGEMENT_SYSTEM_ROLES:
        "sysconsole_write_user_management_system_roles",
    SYSCONSOLE_READ_SITE_CUSTOMIZATION: "sysconsole_read_site_customization",
    SYSCONSOLE_WRITE_SITE_CUSTOMIZATION: "sysconsole_write_site_customization",
    SYSCONSOLE_READ_SITE_LOCALIZATION: "sysconsole_read_site_localization",
    SYSCONSOLE_WRITE_SITE_LOCALIZATION: "sysconsole_write_site_localization",
    SYSCONSOLE_READ_SITE_USERS_AND_TEAMS:
        "sysconsole_read_site_users_and_teams",
    SYSCONSOLE_WRITE_SITE_USERS_AND_TEAMS:
        "sysconsole_write_site_users_and_teams",
    SYSCONSOLE_READ_SITE_NOTIFICATIONS: "sysconsole_read_site_notifications",
    SYSCONSOLE_WRITE_SITE_NOTIFICATIONS: "sysconsole_write_site_notifications",
    SYSCONSOLE_READ_SITE_ANNOUNCEMENT_BANNER:
        "sysconsole_read_site_announcement_banner",
    SYSCONSOLE_WRITE_SITE_ANNOUNCEMENT_BANNER:
        "sysconsole_write_site_announcement_banner",
    SYSCONSOLE_READ_SITE_EMOJI: "sysconsole_read_site_emoji",
    SYSCONSOLE_WRITE_SITE_EMOJI: "sysconsole_write_site_emoji",
    SYSCONSOLE_READ_SITE_POSTS: "sysconsole_read_site_posts",
    SYSCONSOLE_WRITE_SITE_POSTS: "sysconsole_write_site_posts",
    SYSCONSOLE_READ_SITE_FILE_SHARING_AND_DOWNLOADS:
        "sysconsole_read_site_file_sharing_and_downloads",
    SYSCONSOLE_WRITE_SITE_FILE_SHARING_AND_DOWNLOADS:
        "sysconsole_write_site_file_sharing_and_downloads",
    SYSCONSOLE_READ_SITE_PUBLIC_LINKS: "sysconsole_read_site_public_links",
    SYSCONSOLE_WRITE_SITE_PUBLIC_LINKS: "sysconsole_write_site_public_links",
    SYSCONSOLE_READ_SITE_NOTICES: "sysconsole_read_site_notices",
    SYSCONSOLE_WRITE_SITE_NOTICES: "sysconsole_write_site_notices",
    SYSCONSOLE_READ_ENVIRONMENT_WEB_SERVER:
        "sysconsole_read_environment_web_server",
    SYSCONSOLE_WRITE_ENVIRONMENT_WEB_SERVER:
        "sysconsole_write_environment_web_server",
    SYSCONSOLE_READ_ENVIRONMENT_DATABASE:
        "sysconsole_read_environment_database",
    SYSCONSOLE_WRITE_ENVIRONMENT_DATABASE:
        "sysconsole_write_environment_database",
    SYSCONSOLE_READ_ENVIRONMENT_ELASTICSEARCH:
        "sysconsole_read_environment_elasticsearch",
    SYSCONSOLE_WRITE_ENVIRONMENT_ELASTICSEARCH:
        "sysconsole_write_environment_elasticsearch",
    SYSCONSOLE_READ_ENVIRONMENT_FILE_STORAGE:
        "sysconsole_read_environment_file_storage",
    SYSCONSOLE_WRITE_ENVIRONMENT_FILE_STORAGE:
        "sysconsole_write_environment_file_storage",
    SYSCONSOLE_READ_ENVIRONMENT_IMAGE_PROXY:
        "sysconsole_read_environment_image_proxy",
    SYSCONSOLE_WRITE_ENVIRONMENT_IMAGE_PROXY:
        "sysconsole_write_environment_image_proxy",
    SYSCONSOLE_READ_ENVIRONMENT_SMTP: "sysconsole_read_environment_smtp",
    SYSCONSOLE_WRITE_ENVIRONMENT_SMTP: "sysconsole_write_environment_smtp",
    SYSCONSOLE_READ_ENVIRONMENT_PUSH_NOTIFICATION_SERVER:
        "sysconsole_read_environment_push_notification_server",
    SYSCONSOLE_WRITE_ENVIRONMENT_PUSH_NOTIFICATION_SERVER:
        "sysconsole_write_environment_push_notification_server",
    SYSCONSOLE_READ_ENVIRONMENT_HIGH_AVAILABILITY:
        "sysconsole_read_environment_high_availability",
    SYSCONSOLE_WRITE_ENVIRONMENT_HIGH_AVAILABILITY:
        "sysconsole_write_environment_high_availability",
    SYSCONSOLE_READ_ENVIRONMENT_RATE_LIMITING:
        "sysconsole_read_environment_rate_limiting",
    SYSCONSOLE_WRITE_ENVIRONMENT_RATE_LIMITING:
        "sysconsole_write_environment_rate_limiting",
    SYSCONSOLE_READ_ENVIRONMENT_LOGGING: "sysconsole_read_environment_logging",
    SYSCONSOLE_WRITE_ENVIRONMENT_LOGGING:
        "sysconsole_write_environment_logging",
    SYSCONSOLE_READ_ENVIRONMENT_SESSION_LENGTHS:
        "sysconsole_read_environment_session_lengths",
    SYSCONSOLE_WRITE_ENVIRONMENT_SESSION_LENGTHS:
        "sysconsole_write_environment_session_lengths",
    SYSCONSOLE_READ_ENVIRONMENT_PERFORMANCE_MONITORING:
        "sysconsole_read_environment_performance_monitoring",
    SYSCONSOLE_WRITE_ENVIRONMENT_PERFORMANCE_MONITORING:
        "sysconsole_write_environment_performance_monitoring",
    SYSCONSOLE_READ_ENVIRONMENT_DEVELOPER:
        "sysconsole_read_environment_developer",
    SYSCONSOLE_WRITE_ENVIRONMENT_DEVELOPER:
        "sysconsole_write_environment_developer",
    SYSCONSOLE_READ_AUTHENTICATION_SIGNUP:
        "sysconsole_read_authentication_signup",
    SYSCONSOLE_WRITE_AUTHENTICATION_SIGNUP:
        "sysconsole_write_authentication_signup",
    SYSCONSOLE_READ_AUTHENTICATION_EMAIL:
        "sysconsole_read_authentication_email",
    SYSCONSOLE_WRITE_AUTHENTICATION_EMAIL:
        "sysconsole_write_authentication_email",
    SYSCONSOLE_READ_AUTHENTICATION_PASSWORD:
        "sysconsole_read_authentication_password",
    SYSCONSOLE_WRITE_AUTHENTICATION_PASSWORD:
        "sysconsole_write_authentication_password",
    SYSCONSOLE_READ_AUTHENTICATION_MFA: "sysconsole_read_authentication_mfa",
    SYSCONSOLE_WRITE_AUTHENTICATION_MFA: "sysconsole_write_authentication_mfa",
    SYSCONSOLE_READ_AUTHENTICATION_LDAP: "sysconsole_read_authentication_ldap",
    SYSCONSOLE_WRITE_AUTHENTICATION_LDAP:
        "sysconsole_write_authentication_ldap",
    SYSCONSOLE_READ_AUTHENTICATION_SAML: "sysconsole_read_authentication_saml",
    SYSCONSOLE_WRITE_AUTHENTICATION_SAML:
        "sysconsole_write_authentication_saml",
    SYSCONSOLE_READ_AUTHENTICATION_OPENID:
        "sysconsole_read_authentication_openid",
    SYSCONSOLE_WRITE_AUTHENTICATION_OPENID:
        "sysconsole_write_authentication_openid",
    SYSCONSOLE_READ_AUTHENTICATION_GUEST_ACCESS:
        "sysconsole_read_authentication_guest_access",
    SYSCONSOLE_WRITE_AUTHENTICATION_GUEST_ACCESS:
        "sysconsole_write_authentication_guest_access",
    SYSCONSOLE_READ_PLUGINS: "sysconsole_read_plugins",
    SYSCONSOLE_WRITE_PLUGINS: "sysconsole_write_plugins",
    SYSCONSOLE_READ_INTEGRATIONS_INTEGRATION_MANAGEMENT:
        "sysconsole_read_integrations_integration_management",
    SYSCONSOLE_WRITE_INTEGRATIONS_INTEGRATION_MANAGEMENT:
        "sysconsole_write_integrations_integration_management",
    SYSCONSOLE_READ_INTEGRATIONS_BOT_ACCOUNTS:
        "sysconsole_read_integrations_bot_accounts",
    SYSCONSOLE_WRITE_INTEGRATIONS_BOT_ACCOUNTS:
        "sysconsole_write_integrations_bot_accounts",
    SYSCONSOLE_READ_INTEGRATIONS_GIF: "sysconsole_read_integrations_gif",
    SYSCONSOLE_WRITE_INTEGRATIONS_GIF: "sysconsole_write_integrations_gif",
    SYSCONSOLE_READ_INTEGRATIONS_CORS: "sysconsole_read_integrations_cors",
    SYSCONSOLE_WRITE_INTEGRATIONS_CORS: "sysconsole_write_integrations_cors",
    SYSCONSOLE_READ_COMPLIANCE_DATA_RETENTION_POLICY:
        "sysconsole_read_compliance_data_retention_policy",
    SYSCONSOLE_WRITE_COMPLIANCE_DATA_RETENTION_POLICY:
        "sysconsole_write_compliance_data_retention_policy",
    SYSCONSOLE_READ_COMPLIANCE_COMPLIANCE_EXPORT:
        "sysconsole_read_compliance_compliance_export",
    SYSCONSOLE_WRITE_COMPLIANCE_COMPLIANCE_EXPORT:
        "sysconsole_write_compliance_compliance_export",
    SYSCONSOLE_READ_COMPLIANCE_COMPLIANCE_MONITORING:
        "sysconsole_read_compliance_compliance_monitoring",
    SYSCONSOLE_WRITE_COMPLIANCE_COMPLIANCE_MONITORING:
        "sysconsole_write_compliance_compliance_monitoring",
    SYSCONSOLE_READ_COMPLIANCE_CUSTOM_TERMS_OF_SERVICE:
        "sysconsole_read_compliance_custom_terms_of_service",
    SYSCONSOLE_WRITE_COMPLIANCE_CUSTOM_TERMS_OF_SERVICE:
        "sysconsole_write_compliance_custom_terms_of_service",
    SYSCONSOLE_READ_EXPERIMENTAL_FEATURES:
        "sysconsole_read_experimental_features",
    SYSCONSOLE_WRITE_EXPERIMENTAL_FEATURES:
        "sysconsole_write_experimental_features",
    SYSCONSOLE_READ_EXPERIMENTAL_FEATURE_FLAGS:
        "sysconsole_read_experimental_feature_flags",
    SYSCONSOLE_WRITE_EXPERIMENTAL_FEATURE_FLAGS:
        "sysconsole_write_experimental_feature_flags",
    SYSCONSOLE_READ_EXPERIMENTAL_BLEVE: "sysconsole_read_experimental_bleve",
    SYSCONSOLE_WRITE_EXPERIMENTAL_BLEVE: "sysconsole_write_experimental_bleve",

    SYSCONSOLE_READ_PRODUCTS_BOARDS: "sysconsole_read_products_boards",
    SYSCONSOLE_WRITE_PRODUCTS_BOARDS: "sysconsole_write_products_boards",

    PLAYBOOK_PUBLIC_CREATE: "playbook_public_create",
    PLAYBOOK_PUBLIC_MANAGE_PROPERTIES: "playbook_public_manage_properties",
    PLAYBOOK_PUBLIC_MANAGE_MEMBERS: "playbook_public_manage_members",
    PLAYBOOK_PUBLIC_VIEW: "playbook_public_view",
    PLAYBOOK_PUBLIC_MAKE_PRIVATE: "playbook_public_make_private",

    PLAYBOOK_PRIVATE_CREATE: "playbook_private_create",
    PLAYBOOK_PRIVATE_MANAGE_PROPERTIES: "playbook_private_manage_properties",
    PLAYBOOK_PRIVATE_MANAGE_MEMBERS: "playbook_private_manage_members",
    PLAYBOOK_PRIVATE_VIEW: "playbook_private_view",
    PLAYBOOK_PRIVATE_MAKE_PUBLIC: "playbook_private_make_public",

    RUN_CREATE: "run_create",
    RUN_MANAGE_PROPERTIES: "run_manage_properties",
    RUN_MANAGE_MEMBERS: "run_manage_members",
    RUN_VIEW: "run_view",

    CHANNEL_MODERATED_PERMISSIONS: {
        CREATE_POST: "create_post",
        CREATE_REACTIONS: "create_reactions",
        MANAGE_MEMBERS: "manage_members",
        USE_CHANNEL_MENTIONS: "use_channel_mentions",
    },
    MANAGE_BOTS: "manage_bots",
    MANAGE_OTHERS_BOTS: "manage_others_bots",
    SYSCONSOLE_READ_PERMISSIONS: [] as string[],
    SYSCONSOLE_WRITE_PERMISSIONS: [] as string[],
    MANAGE_SHARED_CHANNELS: "manage_shared_channels",
    MANAGE_SECURE_CONNECTIONS: "manage_secure_connections",

    CREATE_CUSTOM_GROUP: "create_custom_group",
    MANAGE_CUSTOM_GROUP_MEMBERS: "manage_custom_group_members",
    EDIT_CUSTOM_GROUP: "edit_custom_group",
    DELETE_CUSTOM_GROUP: "delete_custom_group",
    RESTORE_CUSTOM_GROUP: "restore_custom_group",
};

values.SYSCONSOLE_READ_PERMISSIONS = [
    values.SYSCONSOLE_READ_ABOUT_EDITION_AND_LICENSE,
    values.SYSCONSOLE_READ_BILLING,
    values.SYSCONSOLE_READ_REPORTING_SITE_STATISTICS,
    values.SYSCONSOLE_READ_REPORTING_TEAM_STATISTICS,
    values.SYSCONSOLE_READ_REPORTING_SERVER_LOGS,
    values.SYSCONSOLE_READ_USERMANAGEMENT_USERS,
    values.SYSCONSOLE_READ_USERMANAGEMENT_GROUPS,
    values.SYSCONSOLE_READ_USERMANAGEMENT_TEAMS,
    values.SYSCONSOLE_READ_USERMANAGEMENT_CHANNELS,
    values.SYSCONSOLE_READ_USERMANAGEMENT_PERMISSIONS,
    values.SYSCONSOLE_READ_SITE_CUSTOMIZATION,
    values.SYSCONSOLE_READ_SITE_LOCALIZATION,
    values.SYSCONSOLE_READ_SITE_USERS_AND_TEAMS,
    values.SYSCONSOLE_READ_SITE_NOTIFICATIONS,
    values.SYSCONSOLE_READ_SITE_ANNOUNCEMENT_BANNER,
    values.SYSCONSOLE_READ_SITE_EMOJI,
    values.SYSCONSOLE_READ_SITE_POSTS,
    values.SYSCONSOLE_READ_SITE_FILE_SHARING_AND_DOWNLOADS,
    values.SYSCONSOLE_READ_SITE_PUBLIC_LINKS,
    values.SYSCONSOLE_READ_SITE_NOTICES,
    values.SYSCONSOLE_READ_ENVIRONMENT_WEB_SERVER,
    values.SYSCONSOLE_READ_ENVIRONMENT_DATABASE,
    values.SYSCONSOLE_READ_ENVIRONMENT_ELASTICSEARCH,
    values.SYSCONSOLE_READ_ENVIRONMENT_FILE_STORAGE,
    values.SYSCONSOLE_READ_ENVIRONMENT_IMAGE_PROXY,
    values.SYSCONSOLE_READ_ENVIRONMENT_SMTP,
    values.SYSCONSOLE_READ_ENVIRONMENT_PUSH_NOTIFICATION_SERVER,
    values.SYSCONSOLE_READ_ENVIRONMENT_HIGH_AVAILABILITY,
    values.SYSCONSOLE_READ_ENVIRONMENT_RATE_LIMITING,
    values.SYSCONSOLE_READ_ENVIRONMENT_LOGGING,
    values.SYSCONSOLE_READ_ENVIRONMENT_SESSION_LENGTHS,
    values.SYSCONSOLE_READ_ENVIRONMENT_PERFORMANCE_MONITORING,
    values.SYSCONSOLE_READ_ENVIRONMENT_DEVELOPER,
    values.SYSCONSOLE_READ_AUTHENTICATION_SIGNUP,
    values.SYSCONSOLE_READ_AUTHENTICATION_EMAIL,
    values.SYSCONSOLE_READ_AUTHENTICATION_PASSWORD,
    values.SYSCONSOLE_READ_AUTHENTICATION_MFA,
    values.SYSCONSOLE_READ_AUTHENTICATION_LDAP,
    values.SYSCONSOLE_READ_AUTHENTICATION_SAML,
    values.SYSCONSOLE_READ_AUTHENTICATION_OPENID,
    values.SYSCONSOLE_READ_AUTHENTICATION_GUEST_ACCESS,
    values.SYSCONSOLE_READ_PLUGINS,
    values.SYSCONSOLE_READ_INTEGRATIONS_INTEGRATION_MANAGEMENT,
    values.SYSCONSOLE_READ_INTEGRATIONS_BOT_ACCOUNTS,
    values.SYSCONSOLE_READ_INTEGRATIONS_GIF,
    values.SYSCONSOLE_READ_INTEGRATIONS_CORS,
    values.SYSCONSOLE_READ_COMPLIANCE_DATA_RETENTION_POLICY,
    values.SYSCONSOLE_READ_COMPLIANCE_COMPLIANCE_EXPORT,
    values.SYSCONSOLE_READ_COMPLIANCE_COMPLIANCE_MONITORING,
    values.SYSCONSOLE_READ_COMPLIANCE_CUSTOM_TERMS_OF_SERVICE,
    values.SYSCONSOLE_READ_EXPERIMENTAL_FEATURES,
    values.SYSCONSOLE_READ_EXPERIMENTAL_FEATURE_FLAGS,
    values.SYSCONSOLE_READ_EXPERIMENTAL_BLEVE,
    values.SYSCONSOLE_READ_PRODUCTS_BOARDS,
];

values.SYSCONSOLE_WRITE_PERMISSIONS = [
    values.SYSCONSOLE_WRITE_ABOUT_EDITION_AND_LICENSE,
    values.SYSCONSOLE_WRITE_BILLING,
    values.SYSCONSOLE_WRITE_REPORTING_SITE_STATISTICS,
    values.SYSCONSOLE_WRITE_REPORTING_TEAM_STATISTICS,
    values.SYSCONSOLE_WRITE_REPORTING_SERVER_LOGS,
    values.SYSCONSOLE_WRITE_USERMANAGEMENT_USERS,
    values.SYSCONSOLE_WRITE_USERMANAGEMENT_GROUPS,
    values.SYSCONSOLE_WRITE_USERMANAGEMENT_TEAMS,
    values.SYSCONSOLE_WRITE_USERMANAGEMENT_CHANNELS,
    values.SYSCONSOLE_WRITE_USERMANAGEMENT_PERMISSIONS,
    values.SYSCONSOLE_WRITE_SITE_CUSTOMIZATION,
    values.SYSCONSOLE_WRITE_SITE_LOCALIZATION,
    values.SYSCONSOLE_WRITE_SITE_USERS_AND_TEAMS,
    values.SYSCONSOLE_WRITE_SITE_NOTIFICATIONS,
    values.SYSCONSOLE_WRITE_SITE_ANNOUNCEMENT_BANNER,
    values.SYSCONSOLE_WRITE_SITE_EMOJI,
    values.SYSCONSOLE_WRITE_SITE_POSTS,
    values.SYSCONSOLE_WRITE_SITE_FILE_SHARING_AND_DOWNLOADS,
    values.SYSCONSOLE_WRITE_SITE_PUBLIC_LINKS,
    values.SYSCONSOLE_WRITE_SITE_NOTICES,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_WEB_SERVER,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_DATABASE,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_ELASTICSEARCH,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_FILE_STORAGE,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_IMAGE_PROXY,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_SMTP,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_PUSH_NOTIFICATION_SERVER,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_HIGH_AVAILABILITY,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_RATE_LIMITING,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_LOGGING,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_SESSION_LENGTHS,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_PERFORMANCE_MONITORING,
    values.SYSCONSOLE_WRITE_ENVIRONMENT_DEVELOPER,
    values.SYSCONSOLE_WRITE_AUTHENTICATION_SIGNUP,
    values.SYSCONSOLE_WRITE_AUTHENTICATION_EMAIL,
    values.SYSCONSOLE_WRITE_AUTHENTICATION_PASSWORD,
    values.SYSCONSOLE_WRITE_AUTHENTICATION_MFA,
    values.SYSCONSOLE_WRITE_AUTHENTICATION_LDAP,
    values.SYSCONSOLE_WRITE_AUTHENTICATION_SAML,
    values.SYSCONSOLE_WRITE_AUTHENTICATION_OPENID,
    values.SYSCONSOLE_WRITE_AUTHENTICATION_GUEST_ACCESS,
    values.SYSCONSOLE_WRITE_PLUGINS,
    values.SYSCONSOLE_WRITE_INTEGRATIONS_INTEGRATION_MANAGEMENT,
    values.SYSCONSOLE_WRITE_INTEGRATIONS_BOT_ACCOUNTS,
    values.SYSCONSOLE_WRITE_INTEGRATIONS_GIF,
    values.SYSCONSOLE_WRITE_INTEGRATIONS_CORS,
    values.SYSCONSOLE_WRITE_COMPLIANCE_DATA_RETENTION_POLICY,
    values.SYSCONSOLE_WRITE_COMPLIANCE_COMPLIANCE_EXPORT,
    values.SYSCONSOLE_WRITE_COMPLIANCE_COMPLIANCE_MONITORING,
    values.SYSCONSOLE_WRITE_COMPLIANCE_CUSTOM_TERMS_OF_SERVICE,
    values.SYSCONSOLE_WRITE_EXPERIMENTAL_FEATURES,
    values.SYSCONSOLE_WRITE_EXPERIMENTAL_FEATURE_FLAGS,
    values.SYSCONSOLE_WRITE_EXPERIMENTAL_BLEVE,
    values.SYSCONSOLE_WRITE_PRODUCTS_BOARDS,
];

export default values;
