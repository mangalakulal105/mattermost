// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

export type PluginConfiguration = {

    /** Plugin ID  */
    id: string;

    /** Name of the plugin to show in the UI. We recommend to use manifest.name */
    uiName: string;

    /** URL to the icon to show in the UI. No icon will show the plug outline icon. */
    icon?: string;
    sections: PluginConfigurationSection[];
}

export type PluginConfigurationSection = {
    settings: PluginConfigurationSetting[];

    /** The title of the section. All titles must be different. */
    title: string;

    /**
     * This function will be called whenever a section is saved.
     *
     * The configuration will be automatically saved in the user preferences,
     * so use this function only in case you want to add some side effect
     * to the change.
    */
    onSubmit?: (changes: {[name: string]: string}) => void;
}

export type BasePluginConfigurationSetting = {

    /** Name of the setting. This will be the name used to store in the preferences. */
    name: string;

    /** Optional header for this setting. */
    title?: string;

    /** Optional help text for this setting */
    helpText?: string;

    /** The default value to use */
    default?: string;
}

export type PluginConfigurationRadioSetting = BasePluginConfigurationSetting & {
    type: 'radio';

    /** The default value to use */
    default: string;
    options: PluginConfigurationRadioSettingOption[];
}

export type PluginConfigurationRadioSettingOption = {

    /** The value to store in the preferences */
    value: string;

    /** The text to show in the UI */
    text: string;

    /** Optional help text for this option */
    helpText?: string;
}

export type PluginConfigurationSetting = PluginConfigurationRadioSetting
