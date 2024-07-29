// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

export const adminConsoleNavigation = [
    {
        type: ['team', 'e20'],
        header: 'Edition and License',
        sidebar: 'Edition and License',
        url: 'admin_console/about/license',
    },
    {
        type: ['cloud_enterprise'],
        header: 'Subscription',
        sidebar: 'Subscription',
        url: 'admin_console/billing/subscription',
    },
    {
        type: ['cloud_enterprise'],
        header: 'Billing History',
        sidebar: 'Billing History',
        url: 'admin_console/billing/billing_history',
    },
    {
        type: ['cloud_enterprise'],
        header: 'Company Information',
        sidebar: 'Company Information',
        url: 'admin_console/billing/company_info',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'System Statistics',
        sidebar: 'Site Statistics',
        url: '/admin_console/reporting/system_analytics',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Team Statistics',
        sidebar: 'Team Statistics',
        url: '/admin_console/reporting/team_statistics',
        headerContains: true,
    },
    {
        type: ['team', 'e20'],
        header: 'Server Logs',
        sidebar: 'Server Logs',
        url: '/admin_console/reporting/server_logs',
        headerContains: true,
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Mattermost Users',
        headerSelector: '.admin-console__header #systemUsersTable-headerId',
        sidebar: 'Users',
        url: 'admin_console/user_management/users',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Groups',
        sidebar: 'Groups',
        url: 'admin_console/user_management/groups',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Mattermost Teams',
        sidebar: 'Teams',
        url: 'admin_console/user_management/teams',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Mattermost Channels',
        team_header: 'Channels',
        sidebar: 'Channels',
        url: 'admin_console/user_management/channels',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Permission Schemes',
        sidebar: 'Permissions',
        url: 'admin_console/user_management/permissions',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Delegated Granular Administration',
        sidebar: 'Delegated Granular Administration',
        url: 'admin_console/user_management/system_roles',
    },
    {
        type: ['team', 'e20'],
        header: 'Web Server',
        sidebar: 'Web Server',
        url: 'admin_console/environment/web_server',
    },
    {
        type: ['team', 'e20'],
        header: 'Database',
        sidebar: 'Database',
        url: 'admin_console/environment/database',
    },
    {
        type: ['e20'],
        section: 'Environment',
        header: 'Elasticsearch',
        sidebar: 'Elasticsearch',
        url: 'admin_console/environment/elasticsearch',
    },
    {
        type: ['team', 'e20'],
        section: 'Environment',
        header: 'File Storage',
        sidebar: 'File Storage',
        url: 'admin_console/environment/file_storage',
    },
    {
        type: ['team', 'e20'],
        section: 'Environment',
        header: 'Image Proxy',
        sidebar: 'Image Proxy',
        url: 'admin_console/environment/image_proxy',
    },
    {
        type: ['team', 'e20'],
        section: 'Environment',
        header: 'SMTP',
        sidebar: 'SMTP',
        url: 'admin_console/environment/smtp',
    },
    {
        type: ['team', 'e20'],
        section: 'Environment',
        header: 'Push Notification Server',
        sidebar: 'Push Notification Server',
        url: 'admin_console/environment/push_notification_server',
    },
    {
        type: ['e20'],
        section: 'Environment',
        header: 'High Availability',
        sidebar: 'High Availability',
        url: 'admin_console/environment/high_availability',
    },
    {
        type: ['team', 'e20'],
        section: 'Environment',
        header: 'Rate Limiting',
        sidebar: 'Rate Limiting',
        url: 'admin_console/environment/rate_limiting',
    },
    {
        type: ['team', 'e20'],
        section: 'Environment',
        header: 'Logging',
        sidebar: 'Logging',
        url: 'admin_console/environment/logging',
    },
    {
        type: ['team', 'e20'],
        section: 'Environment',
        header: 'Session Lengths',
        sidebar: 'Session Lengths',
        url: 'admin_console/environment/session_lengths',
    },
    {
        type: ['team', 'e20'],
        section: 'Environment',
        header: 'Performance Monitoring',
        sidebar: 'Performance Monitoring',
        url: 'admin_console/environment/performance_monitoring',
    },
    {
        type: ['team', 'e20'],
        section: 'Environment',
        header: 'Developer Settings',
        sidebar: 'Developer',
        url: 'admin_console/environment/developer',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Customization',
        sidebar: 'Customization',
        url: 'admin_console/site_config/customization',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Localization',
        sidebar: 'Localization',
        url: 'admin_console/site_config/localization',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Users and Teams',
        sidebar: 'Users and Teams',
        url: 'admin_console/site_config/users_and_teams',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Notifications',
        sidebar: 'Notifications',
        url: 'admin_console/environment/notifications',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Announcement Banner',
        sidebar: 'Announcement Banner',
        url: 'admin_console/site_config/announcement_banner',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Emoji',
        sidebar: 'Emoji',
        url: 'admin_console/site_config/emoji',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Posts',
        sidebar: 'Posts',
        url: 'admin_console/site_config/posts',
    },
    {
        type: ['team', 'e20'],
        header: 'File Sharing and Downloads',
        sidebar: 'File Sharing and Downloads',
        url: 'admin_console/site_config/file_sharing_downloads',
    },
    {
        type: ['team', 'e20'],
        header: 'Public Links',
        sidebar: 'Public Links',
        url: 'admin_console/site_config/public_links',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Notices',
        sidebar: 'Notices',
        url: 'admin_console/site_config/notices',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Signup',
        sidebar: 'Signup',
        url: 'admin_console/authentication/signup',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Email Authentication',
        sidebar: 'Email',
        url: 'admin_console/authentication/email',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Password',
        sidebar: 'Password',
        url: 'admin_console/authentication/password',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Multi-factor Authentication',
        sidebar: 'MFA',
        url: 'admin_console/authentication/mfa',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'AD/LDAP',
        sidebar: 'AD/LDAP',
        url: 'admin_console/authentication/ldap',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'SAML 2.0',
        sidebar: 'SAML 2.0',
        url: 'admin_console/authentication/saml',
    },
    {
        type: ['team'],
        header: 'GitLab',
        sidebar: 'GitLab',
        url: 'admin_console/authentication/gitlab',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'OpenID Connect',
        sidebar: 'OpenID Connect',
        url: 'admin_console/authentication/openid',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Guest Access',
        sidebar: 'Guest Access',
        url: 'admin_console/authentication/guest_access',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Plugin Management',
        sidebar: 'Plugin Management',
        url: 'admin_console/plugins/plugin_management',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Integration Management',
        sidebar: 'Integration Management',
        url: 'admin_console/integrations/integration_management',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Bot Accounts',
        sidebar: 'Bot Accounts',
        url: 'admin_console/integrations/bot_accounts',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'GIF (Beta)',
        sidebar: 'GIF (Beta)',
        url: 'admin_console/integrations/gif',
    },
    {
        type: ['team', 'e20'],
        header: 'CORS',
        sidebar: 'CORS',
        url: 'admin_console/integrations/cors',
    },
    {
        type: ['e20', 'cloud_enterprise'],
        header: 'Data Retention Policies',
        sidebar: 'Data Retention Policies',
        url: 'admin_console/compliance/data_retention_settings',
    },
    {
        type: ['team'],
        header: 'Data Retention Policy',
        sidebar: 'Data Retention Policy',
        url: 'admin_console/compliance/data_retention',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Compliance Export',
        sidebar: 'Compliance Export',
        url: 'admin_console/compliance/export',
    },
    {
        type: ['e20', 'cloud_enterprise'],
        header: 'Compliance Monitoring',
        sidebar: 'Compliance Monitoring',
        url: 'admin_console/compliance/monitoring',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Custom Terms of Service',
        sidebar: 'Custom Terms of Service',
        url: 'admin_console/compliance/custom_terms_of_service',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Experimental Features',
        sidebar: 'Features',
        url: 'admin_console/experimental/features',
    },
    {
        type: ['team', 'e20', 'cloud_enterprise'],
        header: 'Feature Flags',
        sidebar: 'Feature Flags',
        url: 'admin_console/experimental/feature_flags',
    },
    {
        type: ['team', 'e20'],
        header: 'Bleve',
        sidebar: 'Bleve',
        url: 'admin_console/experimental/blevesearch',
    },
];
