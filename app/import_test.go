// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package app

import (
	"path/filepath"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/store"
	"github.com/mattermost/mattermost-server/utils"
)

func ptrStr(s string) *string {
	return &s
}

func ptrInt64(i int64) *int64 {
	return &i
}

func ptrInt(i int) *int {
	return &i
}

func ptrBool(b bool) *bool {
	return &b
}

func checkPreference(t *testing.T, a *App, userId string, category string, name string, value string) {
	if res := <-a.Srv.Store.Preference().GetCategory(userId, category); res.Err != nil {
		debug.PrintStack()
		t.Fatalf("Failed to get preferences for user %v with category %v", userId, category)
	} else {
		preferences := res.Data.(model.Preferences)
		found := false
		for _, preference := range preferences {
			if preference.Name == name {
				found = true
				if preference.Value != value {
					debug.PrintStack()
					t.Fatalf("Preference for user %v in category %v with name %v has value %v, expected %v", userId, category, name, preference.Value, value)
				}
				break
			}
		}
		if !found {
			debug.PrintStack()
			t.Fatalf("Did not find preference for user %v in category %v with name %v", userId, category, name)
		}
	}
}

func checkNotifyProp(t *testing.T, user *model.User, key string, value string) {
	if actual, ok := user.NotifyProps[key]; !ok {
		debug.PrintStack()
		t.Fatalf("Notify prop %v not found. User: %v", key, user.Id)
	} else if actual != value {
		debug.PrintStack()
		t.Fatalf("Notify Prop %v was %v but expected %v. User: %v", key, actual, value, user.Id)
	}
}

func checkError(t *testing.T, err *model.AppError) {
	if err == nil {
		debug.PrintStack()
		t.Fatal("Should have returned an error.")
	}
}

func checkNoError(t *testing.T, err *model.AppError) {
	if err != nil {
		debug.PrintStack()
		t.Fatalf("Unexpected Error: %v", err.Error())
	}
}

func TestImportValidateSchemeImportData(t *testing.T) {
	// Test with minimum required valid properties and team scope.
	data := SchemeImportData{
		Name:        ptrStr("name"),
		DisplayName: ptrStr("display name"),
		Scope:       ptrStr("team"),
		DefaultTeamAdminRole: &RoleImportData{
			Name:        ptrStr("name"),
			DisplayName: ptrStr("display name"),
			Permissions: &[]string{"invite_user"},
		},
		DefaultTeamUserRole: &RoleImportData{
			Name:        ptrStr("name"),
			DisplayName: ptrStr("display name"),
			Permissions: &[]string{"invite_user"},
		},
		DefaultChannelAdminRole: &RoleImportData{
			Name:        ptrStr("name"),
			DisplayName: ptrStr("display name"),
			Permissions: &[]string{"invite_user"},
		},
		DefaultChannelUserRole: &RoleImportData{
			Name:        ptrStr("name"),
			DisplayName: ptrStr("display name"),
			Permissions: &[]string{"invite_user"},
		},
	}
	if err := validateSchemeImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.", err)
	}

	// Test with various invalid names.
	data.Name = nil
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid name.")
	}

	data.Name = ptrStr("")
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid name.")
	}

	data.Name = ptrStr(strings.Repeat("1234567890", 100))
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid name.")
	}

	data.Name = ptrStr("name")
	// Test with invalid display name.
	data.DisplayName = nil
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid display name.")
	}

	data.DisplayName = ptrStr("")
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid display name.")
	}

	data.DisplayName = ptrStr(strings.Repeat("1234567890", 100))
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid display name.")
	}

	data.DisplayName = ptrStr("display name")

	// Test with various missing roles.
	data.DefaultTeamAdminRole = nil
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to missing role.")
	}

	data.DefaultTeamAdminRole = &RoleImportData{
		Name:        ptrStr("name"),
		DisplayName: ptrStr("display name"),
		Permissions: &[]string{"invite_user"},
	}
	data.DefaultTeamUserRole = nil
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to missing role.")
	}

	data.DefaultTeamUserRole = &RoleImportData{
		Name:        ptrStr("name"),
		DisplayName: ptrStr("display name"),
		Permissions: &[]string{"invite_user"},
	}
	data.DefaultChannelAdminRole = nil
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to missing role.")
	}

	data.DefaultChannelAdminRole = &RoleImportData{
		Name:        ptrStr("name"),
		DisplayName: ptrStr("display name"),
		Permissions: &[]string{"invite_user"},
	}
	data.DefaultChannelUserRole = nil
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to missing role.")
	}

	data.DefaultChannelUserRole = &RoleImportData{
		Name:        ptrStr("name"),
		DisplayName: ptrStr("display name"),
		Permissions: &[]string{"invite_user"},
	}

	// Test with various invalid roles.
	data.DefaultTeamAdminRole.Name = nil
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid role.")
	}

	data.DefaultTeamAdminRole.Name = ptrStr("name")
	data.DefaultTeamUserRole.Name = nil
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid role.")
	}

	data.DefaultTeamUserRole.Name = ptrStr("name")
	data.DefaultChannelAdminRole.Name = nil
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid role.")
	}

	data.DefaultChannelAdminRole.Name = ptrStr("name")
	data.DefaultChannelUserRole.Name = nil
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid role.")
	}

	data.DefaultChannelUserRole.Name = ptrStr("name")

	// Change to a Channel scope role, and check with missing or extra roles again.
	data.Scope = ptrStr("channel")
	data.DefaultTeamAdminRole = nil
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to spurious role.")
	}

	data.DefaultTeamAdminRole = &RoleImportData{
		Name:        ptrStr("name"),
		DisplayName: ptrStr("display name"),
		Permissions: &[]string{"invite_user"},
	}
	data.DefaultTeamUserRole = nil
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to spurious role.")
	}

	data.DefaultTeamAdminRole = nil
	if err := validateSchemeImportData(&data); err != nil {
		t.Fatal("Should have succeeded.")
	}

	// Test with all combinations of optional parameters.
	data.Description = ptrStr(strings.Repeat("1234567890", 1024))
	if err := validateSchemeImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid description.")
	}

	data.Description = ptrStr("description")
	if err := validateSchemeImportData(&data); err != nil {
		t.Fatal("Should have succeeded.")
	}
}

func TestImportValidateRoleImportData(t *testing.T) {
	// Test with minimum required valid properties.
	data := RoleImportData{
		Name:        ptrStr("name"),
		DisplayName: ptrStr("display name"),
	}
	if err := validateRoleImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.", err)
	}

	// Test with various invalid names.
	data.Name = nil
	if err := validateRoleImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid name.")
	}

	data.Name = ptrStr("")
	if err := validateRoleImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid name.")
	}

	data.Name = ptrStr(strings.Repeat("1234567890", 100))
	if err := validateRoleImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid name.")
	}

	data.Name = ptrStr("name")
	// Test with invalid display name.
	data.DisplayName = nil
	if err := validateRoleImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid display name.")
	}

	data.DisplayName = ptrStr("")
	if err := validateRoleImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid display name.")
	}

	data.DisplayName = ptrStr(strings.Repeat("1234567890", 100))
	if err := validateRoleImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid display name.")
	}

	data.DisplayName = ptrStr("display name")

	// Test with various valid/invalid permissions.
	data.Permissions = &[]string{}
	if err := validateRoleImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.", err)
	}

	data.Permissions = &[]string{"invite_user", "add_user_to_team"}
	if err := validateRoleImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.", err)
	}

	data.Permissions = &[]string{"invite_user", "add_user_to_team", "derp"}
	if err := validateRoleImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to invalid permission.", err)
	}

	data.Permissions = &[]string{"invite_user", "add_user_to_team"}

	// Test with various valid/invalid descriptions.
	data.Description = ptrStr(strings.Repeat("1234567890", 1024))
	if err := validateRoleImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to invalid description.", err)
	}

	data.Description = ptrStr("description")
	if err := validateRoleImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.", err)
	}
}

func TestImportValidateTeamImportData(t *testing.T) {

	// Test with minimum required valid properties.
	data := TeamImportData{
		Name:        ptrStr("teamname"),
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
	}
	if err := validateTeamImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}

	// Test with various invalid names.
	data = TeamImportData{
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
	}
	if err := validateTeamImportData(&data); err == nil {
		t.Fatal("Should have failed due to missing name.")
	}

	data.Name = ptrStr(strings.Repeat("abcdefghij", 7))
	if err := validateTeamImportData(&data); err == nil {
		t.Fatal("Should have failed due to too long name.")
	}

	data.Name = ptrStr("login")
	if err := validateTeamImportData(&data); err == nil {
		t.Fatal("Should have failed due to reserved word in name.")
	}

	data.Name = ptrStr("Test::''ASD")
	if err := validateTeamImportData(&data); err == nil {
		t.Fatal("Should have failed due to non alphanum characters in name.")
	}

	data.Name = ptrStr("A")
	if err := validateTeamImportData(&data); err == nil {
		t.Fatal("Should have failed due to short name.")
	}

	// Test team various invalid display names.
	data = TeamImportData{
		Name: ptrStr("teamname"),
		Type: ptrStr("O"),
	}
	if err := validateTeamImportData(&data); err == nil {
		t.Fatal("Should have failed due to missing display_name.")
	}

	data.DisplayName = ptrStr("")
	if err := validateTeamImportData(&data); err == nil {
		t.Fatal("Should have failed due to empty display_name.")
	}

	data.DisplayName = ptrStr(strings.Repeat("abcdefghij", 7))
	if err := validateTeamImportData(&data); err == nil {
		t.Fatal("Should have failed due to too long display_name.")
	}

	// Test with various valid and invalid types.
	data = TeamImportData{
		Name:        ptrStr("teamname"),
		DisplayName: ptrStr("Display Name"),
	}
	if err := validateTeamImportData(&data); err == nil {
		t.Fatal("Should have failed due to missing type.")
	}

	data.Type = ptrStr("A")
	if err := validateTeamImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid type.")
	}

	data.Type = ptrStr("I")
	if err := validateTeamImportData(&data); err != nil {
		t.Fatal("Should have succeeded with valid type.")
	}

	// Test with all the combinations of optional parameters.
	data = TeamImportData{
		Name:            ptrStr("teamname"),
		DisplayName:     ptrStr("Display Name"),
		Type:            ptrStr("O"),
		Description:     ptrStr("The team description."),
		AllowOpenInvite: ptrBool(true),
	}
	if err := validateTeamImportData(&data); err != nil {
		t.Fatal("Should have succeeded with valid optional properties.")
	}

	data.AllowOpenInvite = ptrBool(false)
	if err := validateTeamImportData(&data); err != nil {
		t.Fatal("Should have succeeded with allow open invites false.")
	}

	data.Description = ptrStr(strings.Repeat("abcdefghij ", 26))
	if err := validateTeamImportData(&data); err == nil {
		t.Fatal("Should have failed due to too long description.")
	}

	// Test with an empty scheme name.
	data.Description = ptrStr("abcdefg")
	data.Scheme = ptrStr("")
	if err := validateTeamImportData(&data); err == nil {
		t.Fatal("Should have failed due to empty scheme name.")
	}

	// Test with a valid scheme name.
	data.Scheme = ptrStr("abcdefg")
	if err := validateTeamImportData(&data); err != nil {
		t.Fatal("Should have succeeded with valid scheme name.")
	}
}

func TestImportValidateChannelImportData(t *testing.T) {

	// Test with minimum required valid properties.
	data := ChannelImportData{
		Team:        ptrStr("teamname"),
		Name:        ptrStr("channelname"),
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
	}
	if err := validateChannelImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}

	// Test with missing team.
	data = ChannelImportData{
		Name:        ptrStr("channelname"),
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
	}
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to missing team.")
	}

	// Test with various invalid names.
	data = ChannelImportData{
		Team:        ptrStr("teamname"),
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
	}
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to missing name.")
	}

	data.Name = ptrStr(strings.Repeat("abcdefghij", 7))
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to too long name.")
	}

	data.Name = ptrStr("Test::''ASD")
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to non alphanum characters in name.")
	}

	data.Name = ptrStr("A")
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to short name.")
	}

	// Test team various invalid display names.
	data = ChannelImportData{
		Team: ptrStr("teamname"),
		Name: ptrStr("channelname"),
		Type: ptrStr("O"),
	}
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to missing display_name.")
	}

	data.DisplayName = ptrStr("")
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to empty display_name.")
	}

	data.DisplayName = ptrStr(strings.Repeat("abcdefghij", 7))
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to too long display_name.")
	}

	// Test with various valid and invalid types.
	data = ChannelImportData{
		Team:        ptrStr("teamname"),
		Name:        ptrStr("channelname"),
		DisplayName: ptrStr("Display Name"),
	}
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to missing type.")
	}

	data.Type = ptrStr("A")
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid type.")
	}

	data.Type = ptrStr("P")
	if err := validateChannelImportData(&data); err != nil {
		t.Fatal("Should have succeeded with valid type.")
	}

	// Test with all the combinations of optional parameters.
	data = ChannelImportData{
		Team:        ptrStr("teamname"),
		Name:        ptrStr("channelname"),
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
		Header:      ptrStr("Channel Header Here"),
		Purpose:     ptrStr("Channel Purpose Here"),
	}
	if err := validateChannelImportData(&data); err != nil {
		t.Fatal("Should have succeeded with valid optional properties.")
	}

	data.Header = ptrStr(strings.Repeat("abcdefghij ", 103))
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to too long header.")
	}

	data.Header = ptrStr("Channel Header Here")
	data.Purpose = ptrStr(strings.Repeat("abcdefghij ", 26))
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to too long purpose.")
	}

	// Test with an empty scheme name.
	data.Purpose = ptrStr("abcdefg")
	data.Scheme = ptrStr("")
	if err := validateChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to empty scheme name.")
	}

	// Test with a valid scheme name.
	data.Scheme = ptrStr("abcdefg")
	if err := validateChannelImportData(&data); err != nil {
		t.Fatal("Should have succeeded with valid scheme name.")
	}
}

func TestImportValidateUserImportData(t *testing.T) {

	// Test with minimum required valid properties.
	data := UserImportData{
		Username: ptrStr("bob"),
		Email:    ptrStr("bob@example.com"),
	}
	if err := validateUserImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}

	// Invalid Usernames.
	data.Username = nil
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to nil Username.")
	}

	data.Username = ptrStr("")
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to 0 length Username.")
	}

	data.Username = ptrStr(strings.Repeat("abcdefghij", 7))
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to too long Username.")
	}

	data.Username = ptrStr("i am a username with spaces and !!!")
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to invalid characters in Username.")
	}

	data.Username = ptrStr("bob")

	// Unexisting Picture Image
	data.ProfileImage = ptrStr("not-existing-file")
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to not existing profile image file.")
	}
	data.ProfileImage = nil

	// Invalid Emails
	data.Email = nil
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to nil Email.")
	}

	data.Email = ptrStr("")
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to 0 length Email.")
	}

	data.Email = ptrStr(strings.Repeat("abcdefghij", 13))
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to too long Email.")
	}

	data.Email = ptrStr("bob@example.com")

	data.AuthService = ptrStr("")
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to 0-length auth service.")
	}

	data.AuthService = ptrStr("saml")
	data.AuthData = ptrStr(strings.Repeat("abcdefghij", 15))
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to too long auth data.")
	}

	data.AuthData = ptrStr("bobbytables")
	if err := validateUserImportData(&data); err != nil {
		t.Fatal("Validation should have succeeded with valid auth service and auth data.")
	}

	// Test a valid User with all fields populated.
	testsDir, _ := utils.FindDir("tests")
	data = UserImportData{
		ProfileImage: ptrStr(filepath.Join(testsDir, "test.png")),
		Username:     ptrStr("bob"),
		Email:        ptrStr("bob@example.com"),
		AuthService:  ptrStr("ldap"),
		AuthData:     ptrStr("bob"),
		Nickname:     ptrStr("BobNick"),
		FirstName:    ptrStr("Bob"),
		LastName:     ptrStr("Blob"),
		Position:     ptrStr("The Boss"),
		Roles:        ptrStr("system_user"),
		Locale:       ptrStr("en"),
	}
	if err := validateUserImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}

	// Test various invalid optional field values.
	data.Nickname = ptrStr(strings.Repeat("abcdefghij", 7))
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to too long Nickname.")
	}
	data.Nickname = ptrStr("BobNick")

	data.FirstName = ptrStr(strings.Repeat("abcdefghij", 7))
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to too long First Name.")
	}
	data.FirstName = ptrStr("Bob")

	data.LastName = ptrStr(strings.Repeat("abcdefghij", 7))
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to too long Last name.")
	}
	data.LastName = ptrStr("Blob")

	data.Position = ptrStr(strings.Repeat("abcdefghij", 13))
	if err := validateUserImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to too long Position.")
	}
	data.Position = ptrStr("The Boss")

	data.Roles = nil
	if err := validateUserImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}

	data.Roles = ptrStr("")
	if err := validateUserImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}
	data.Roles = ptrStr("system_user")

	// Try various valid/invalid notify props.
	data.NotifyProps = &UserNotifyPropsImportData{}

	data.NotifyProps.Desktop = ptrStr("invalid")
	checkError(t, validateUserImportData(&data))

	data.NotifyProps.Desktop = ptrStr(model.USER_NOTIFY_ALL)
	data.NotifyProps.DesktopSound = ptrStr("invalid")
	checkError(t, validateUserImportData(&data))

	data.NotifyProps.DesktopSound = ptrStr("true")
	data.NotifyProps.Email = ptrStr("invalid")
	checkError(t, validateUserImportData(&data))

	data.NotifyProps.Email = ptrStr("true")
	data.NotifyProps.Mobile = ptrStr("invalid")
	checkError(t, validateUserImportData(&data))

	data.NotifyProps.Mobile = ptrStr(model.USER_NOTIFY_ALL)
	data.NotifyProps.MobilePushStatus = ptrStr("invalid")
	checkError(t, validateUserImportData(&data))

	data.NotifyProps.MobilePushStatus = ptrStr(model.STATUS_ONLINE)
	data.NotifyProps.ChannelTrigger = ptrStr("invalid")
	checkError(t, validateUserImportData(&data))

	data.NotifyProps.ChannelTrigger = ptrStr("true")
	data.NotifyProps.CommentsTrigger = ptrStr("invalid")
	checkError(t, validateUserImportData(&data))

	data.NotifyProps.CommentsTrigger = ptrStr(model.COMMENTS_NOTIFY_ROOT)
	data.NotifyProps.MentionKeys = ptrStr("valid")
	checkNoError(t, validateUserImportData(&data))
}

func TestImportValidateUserTeamsImportData(t *testing.T) {

	// Invalid Name.
	data := []UserTeamImportData{
		{
			Roles: ptrStr("team_admin team_user"),
		},
	}
	if err := validateUserTeamsImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid name.")
	}
	data[0].Name = ptrStr("teamname")

	// Valid (nil roles)
	data[0].Roles = nil
	if err := validateUserTeamsImportData(&data); err != nil {
		t.Fatal("Should have succeeded with empty roles.")
	}

	// Valid (empty roles)
	data[0].Roles = ptrStr("")
	if err := validateUserTeamsImportData(&data); err != nil {
		t.Fatal("Should have succeeded with empty roles.")
	}

	// Valid (with roles)
	data[0].Roles = ptrStr("team_admin team_user")
	if err := validateUserTeamsImportData(&data); err != nil {
		t.Fatal("Should have succeeded with valid roles.")
	}
}

func TestImportValidateUserChannelsImportData(t *testing.T) {

	// Invalid Name.
	data := []UserChannelImportData{
		{
			Roles: ptrStr("channel_admin channel_user"),
		},
	}
	if err := validateUserChannelsImportData(&data); err == nil {
		t.Fatal("Should have failed due to invalid name.")
	}
	data[0].Name = ptrStr("channelname")

	// Valid (nil roles)
	data[0].Roles = nil
	if err := validateUserChannelsImportData(&data); err != nil {
		t.Fatal("Should have succeeded with empty roles.")
	}

	// Valid (empty roles)
	data[0].Roles = ptrStr("")
	if err := validateUserChannelsImportData(&data); err != nil {
		t.Fatal("Should have succeeded with empty roles.")
	}

	// Valid (with roles)
	data[0].Roles = ptrStr("channel_admin channel_user")
	if err := validateUserChannelsImportData(&data); err != nil {
		t.Fatal("Should have succeeded with valid roles.")
	}

	// Empty notify props.
	data[0].NotifyProps = &UserChannelNotifyPropsImportData{}
	if err := validateUserChannelsImportData(&data); err != nil {
		t.Fatal("Should have succeeded with empty notify props.")
	}

	// Invalid desktop notify props.
	data[0].NotifyProps.Desktop = ptrStr("invalid")
	if err := validateUserChannelsImportData(&data); err == nil {
		t.Fatal("Should have failed with invalid desktop notify props.")
	}

	// Invalid mobile notify props.
	data[0].NotifyProps.Desktop = ptrStr("mention")
	data[0].NotifyProps.Mobile = ptrStr("invalid")
	if err := validateUserChannelsImportData(&data); err == nil {
		t.Fatal("Should have failed with invalid mobile notify props.")
	}

	// Invalid mark_unread notify props.
	data[0].NotifyProps.Mobile = ptrStr("mention")
	data[0].NotifyProps.MarkUnread = ptrStr("invalid")
	if err := validateUserChannelsImportData(&data); err == nil {
		t.Fatal("Should have failed with invalid mark_unread notify props.")
	}

	// Valid notify props.
	data[0].NotifyProps.MarkUnread = ptrStr("mention")
	if err := validateUserChannelsImportData(&data); err != nil {
		t.Fatal("Should have succeeded with valid notify props.")
	}
}

func TestImportValidateReactionImportData(t *testing.T) {
	// Test with minimum required valid properties.
	parentCreateAt := model.GetMillis() - 100
	data := ReactionImportData{
		User:      ptrStr("username"),
		EmojiName: ptrStr("emoji"),
		CreateAt:  ptrInt64(model.GetMillis()),
	}
	if err := validateReactionImportData(&data, parentCreateAt); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}

	// Test with missing required properties.
	data = ReactionImportData{
		EmojiName: ptrStr("emoji"),
		CreateAt:  ptrInt64(model.GetMillis()),
	}
	if err := validateReactionImportData(&data, parentCreateAt); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	data = ReactionImportData{
		User:     ptrStr("username"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateReactionImportData(&data, parentCreateAt); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	data = ReactionImportData{
		User:      ptrStr("username"),
		EmojiName: ptrStr("emoji"),
	}
	if err := validateReactionImportData(&data, parentCreateAt); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	// Test with invalid emoji name.
	data = ReactionImportData{
		User:      ptrStr("username"),
		EmojiName: ptrStr(strings.Repeat("1234567890", 500)),
		CreateAt:  ptrInt64(model.GetMillis()),
	}
	if err := validateReactionImportData(&data, parentCreateAt); err == nil {
		t.Fatal("Should have failed due to too long emoji name.")
	}

	// Test with invalid CreateAt
	data = ReactionImportData{
		User:      ptrStr("username"),
		EmojiName: ptrStr("emoji"),
		CreateAt:  ptrInt64(0),
	}
	if err := validateReactionImportData(&data, parentCreateAt); err == nil {
		t.Fatal("Should have failed due to 0 create-at value.")
	}

	data = ReactionImportData{
		User:      ptrStr("username"),
		EmojiName: ptrStr("emoji"),
		CreateAt:  ptrInt64(parentCreateAt - 100),
	}
	if err := validateReactionImportData(&data, parentCreateAt); err == nil {
		t.Fatal("Should have failed due parent with newer create-at value.")
	}
}

func TestImportValidateReplyImportData(t *testing.T) {
	// Test with minimum required valid properties.
	parentCreateAt := model.GetMillis() - 100
	maxPostSize := 10000
	data := ReplyImportData{
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateReplyImportData(&data, parentCreateAt, maxPostSize); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}

	// Test with missing required properties.
	data = ReplyImportData{
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateReplyImportData(&data, parentCreateAt, maxPostSize); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	data = ReplyImportData{
		User:     ptrStr("username"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateReplyImportData(&data, parentCreateAt, maxPostSize); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	data = ReplyImportData{
		User:    ptrStr("username"),
		Message: ptrStr("message"),
	}
	if err := validateReplyImportData(&data, parentCreateAt, maxPostSize); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	// Test with invalid message.
	data = ReplyImportData{
		User:     ptrStr("username"),
		Message:  ptrStr(strings.Repeat("0", maxPostSize+1)),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateReplyImportData(&data, parentCreateAt, maxPostSize); err == nil {
		t.Fatal("Should have failed due to too long message.")
	}

	// Test with invalid CreateAt
	data = ReplyImportData{
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(0),
	}
	if err := validateReplyImportData(&data, parentCreateAt, maxPostSize); err == nil {
		t.Fatal("Should have failed due to 0 create-at value.")
	}

	data = ReplyImportData{
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(parentCreateAt - 100),
	}
	if err := validateReplyImportData(&data, parentCreateAt, maxPostSize); err == nil {
		t.Fatal("Should have failed due parent with newer create-at value.")
	}
}

func TestImportValidatePostImportData(t *testing.T) {
	maxPostSize := 10000

	// Test with minimum required valid properties.
	data := PostImportData{
		Team:     ptrStr("teamname"),
		Channel:  ptrStr("channelname"),
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validatePostImportData(&data, maxPostSize); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}

	// Test with missing required properties.
	data = PostImportData{
		Channel:  ptrStr("channelname"),
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validatePostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	data = PostImportData{
		Team:     ptrStr("teamname"),
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validatePostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	data = PostImportData{
		Team:     ptrStr("teamname"),
		Channel:  ptrStr("channelname"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validatePostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	data = PostImportData{
		Team:     ptrStr("teamname"),
		Channel:  ptrStr("channelname"),
		User:     ptrStr("username"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validatePostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	data = PostImportData{
		Team:    ptrStr("teamname"),
		Channel: ptrStr("channelname"),
		User:    ptrStr("username"),
		Message: ptrStr("message"),
	}
	if err := validatePostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	// Test with invalid message.
	data = PostImportData{
		Team:     ptrStr("teamname"),
		Channel:  ptrStr("channelname"),
		User:     ptrStr("username"),
		Message:  ptrStr(strings.Repeat("0", maxPostSize+1)),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validatePostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to too long message.")
	}

	// Test with invalid CreateAt
	data = PostImportData{
		Team:     ptrStr("teamname"),
		Channel:  ptrStr("channelname"),
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(0),
	}
	if err := validatePostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to 0 create-at value.")
	}

	// Test with valid all optional parameters.
	reactions := []ReactionImportData{ReactionImportData{
		User:      ptrStr("username"),
		EmojiName: ptrStr("emoji"),
		CreateAt:  ptrInt64(model.GetMillis()),
	}}
	replies := []ReplyImportData{ReplyImportData{
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}}
	data = PostImportData{
		Team:      ptrStr("teamname"),
		Channel:   ptrStr("channelname"),
		User:      ptrStr("username"),
		Message:   ptrStr("message"),
		CreateAt:  ptrInt64(model.GetMillis()),
		Reactions: &reactions,
		Replies:   &replies,
	}
	if err := validatePostImportData(&data, maxPostSize); err != nil {
		t.Fatal("Should have succeeded.")
	}
}

func TestImportValidateDirectChannelImportData(t *testing.T) {

	// Test with valid number of members for direct message.
	data := DirectChannelImportData{
		Members: &[]string{
			model.NewId(),
			model.NewId(),
		},
	}
	if err := validateDirectChannelImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}

	// Test with valid number of members for group message.
	data = DirectChannelImportData{
		Members: &[]string{
			model.NewId(),
			model.NewId(),
			model.NewId(),
		},
	}
	if err := validateDirectChannelImportData(&data); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}

	// Test with all the combinations of optional parameters.
	data = DirectChannelImportData{
		Members: &[]string{
			model.NewId(),
			model.NewId(),
		},
		Header: ptrStr("Channel Header Here"),
	}
	if err := validateDirectChannelImportData(&data); err != nil {
		t.Fatal("Should have succeeded with valid optional properties.")
	}

	// Test with invalid Header.
	data.Header = ptrStr(strings.Repeat("abcdefghij ", 103))
	if err := validateDirectChannelImportData(&data); err == nil {
		t.Fatal("Should have failed due to too long header.")
	}

	// Test with different combinations of invalid member counts.
	data = DirectChannelImportData{
		Members: &[]string{},
	}
	if err := validateDirectChannelImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to invalid number of members.")
	}

	data = DirectChannelImportData{
		Members: &[]string{
			model.NewId(),
		},
	}
	if err := validateDirectChannelImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to invalid number of members.")
	}

	data = DirectChannelImportData{
		Members: &[]string{
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
		},
	}
	if err := validateDirectChannelImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to invalid number of members.")
	}

	// Test with invalid FavoritedBy
	member1 := model.NewId()
	member2 := model.NewId()
	data = DirectChannelImportData{
		Members: &[]string{
			member1,
			member2,
		},
		FavoritedBy: &[]string{
			member1,
			model.NewId(),
		},
	}
	if err := validateDirectChannelImportData(&data); err == nil {
		t.Fatal("Validation should have failed due to non-member favorited.")
	}

	// Test with valid FavoritedBy
	data = DirectChannelImportData{
		Members: &[]string{
			member1,
			member2,
		},
		FavoritedBy: &[]string{
			member1,
			member2,
		},
	}
	if err := validateDirectChannelImportData(&data); err != nil {
		t.Fatal(err)
	}
}

func TestImportValidateDirectPostImportData(t *testing.T) {
	maxPostSize := 10000

	// Test with minimum required valid properties.
	data := DirectPostImportData{
		ChannelMembers: &[]string{
			model.NewId(),
			model.NewId(),
		},
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}

	// Test with missing required properties.
	data = DirectPostImportData{
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	data = DirectPostImportData{
		ChannelMembers: &[]string{
			model.NewId(),
			model.NewId(),
		},
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	data = DirectPostImportData{
		ChannelMembers: &[]string{
			model.NewId(),
			model.NewId(),
		},
		User:     ptrStr("username"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	data = DirectPostImportData{
		ChannelMembers: &[]string{
			model.NewId(),
			model.NewId(),
		},
		User:    ptrStr("username"),
		Message: ptrStr("message"),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to missing required property.")
	}

	// Test with invalid numbers of channel members.
	data = DirectPostImportData{
		ChannelMembers: &[]string{},
		User:           ptrStr("username"),
		Message:        ptrStr("message"),
		CreateAt:       ptrInt64(model.GetMillis()),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to unsuitable number of members.")
	}

	data = DirectPostImportData{
		ChannelMembers: &[]string{
			model.NewId(),
		},
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to unsuitable number of members.")
	}

	data = DirectPostImportData{
		ChannelMembers: &[]string{
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
			model.NewId(),
		},
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to unsuitable number of members.")
	}

	// Test with group message number of members.
	data = DirectPostImportData{
		ChannelMembers: &[]string{
			model.NewId(),
			model.NewId(),
			model.NewId(),
		},
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err != nil {
		t.Fatal("Validation failed but should have been valid.")
	}

	// Test with invalid message.
	data = DirectPostImportData{
		ChannelMembers: &[]string{
			model.NewId(),
			model.NewId(),
		},
		User:     ptrStr("username"),
		Message:  ptrStr(strings.Repeat("0", maxPostSize+1)),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to too long message.")
	}

	// Test with invalid CreateAt
	data = DirectPostImportData{
		ChannelMembers: &[]string{
			model.NewId(),
			model.NewId(),
		},
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(0),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Should have failed due to 0 create-at value.")
	}

	// Test with invalid FlaggedBy
	member1 := model.NewId()
	member2 := model.NewId()
	data = DirectPostImportData{
		ChannelMembers: &[]string{
			member1,
			member2,
		},
		FlaggedBy: &[]string{
			member1,
			model.NewId(),
		},
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err == nil {
		t.Fatal("Validation should have failed due to non-member flagged.")
	}

	// Test with valid FlaggedBy
	data = DirectPostImportData{
		ChannelMembers: &[]string{
			member1,
			member2,
		},
		FlaggedBy: &[]string{
			member1,
			member2,
		},
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := validateDirectPostImportData(&data, maxPostSize); err != nil {
		t.Fatal(err)
	}

	// Test with valid all optional parameters.
	reactions := []ReactionImportData{ReactionImportData{
		User:      ptrStr("username"),
		EmojiName: ptrStr("emoji"),
		CreateAt:  ptrInt64(model.GetMillis()),
	}}
	replies := []ReplyImportData{ReplyImportData{
		User:     ptrStr("username"),
		Message:  ptrStr("message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}}
	data = DirectPostImportData{
		ChannelMembers: &[]string{
			member1,
			member2,
		},
		FlaggedBy: &[]string{
			member1,
			member2,
		},
		User:      ptrStr("username"),
		Message:   ptrStr("message"),
		CreateAt:  ptrInt64(model.GetMillis()),
		Reactions: &reactions,
		Replies:   &replies,
	}

	if err := validateDirectPostImportData(&data, maxPostSize); err != nil {
		t.Fatal(err)
	}
}

func TestImportImportScheme(t *testing.T) {
	th := Setup()
	defer th.TearDown()

	// Mark the phase 2 permissions migration as completed.
	<-th.App.Srv.Store.System().Save(&model.System{Name: model.MIGRATION_KEY_ADVANCED_PERMISSIONS_PHASE_2, Value: "true"})

	defer func() {
		<-th.App.Srv.Store.System().PermanentDeleteByName(model.MIGRATION_KEY_ADVANCED_PERMISSIONS_PHASE_2)
	}()

	// Try importing an invalid scheme in dryRun mode.
	data := SchemeImportData{
		Name:  ptrStr(model.NewId()),
		Scope: ptrStr("team"),
		DefaultTeamUserRole: &RoleImportData{
			Name:        ptrStr(model.NewId()),
			DisplayName: ptrStr(model.NewId()),
		},
		DefaultTeamAdminRole: &RoleImportData{
			Name:        ptrStr(model.NewId()),
			DisplayName: ptrStr(model.NewId()),
		},
		DefaultChannelUserRole: &RoleImportData{
			Name:        ptrStr(model.NewId()),
			DisplayName: ptrStr(model.NewId()),
		},
		DefaultChannelAdminRole: &RoleImportData{
			Name:        ptrStr(model.NewId()),
			DisplayName: ptrStr(model.NewId()),
		},
		Description: ptrStr("description"),
	}

	if err := th.App.ImportScheme(&data, true); err == nil {
		t.Fatalf("Should have failed to import.")
	}

	if res := <-th.App.Srv.Store.Scheme().GetByName(*data.Name); res.Err == nil {
		t.Fatalf("Scheme should not have imported.")
	}

	// Try importing a valid scheme in dryRun mode.
	data.DisplayName = ptrStr("display name")

	if err := th.App.ImportScheme(&data, true); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	if res := <-th.App.Srv.Store.Scheme().GetByName(*data.Name); res.Err == nil {
		t.Fatalf("Scheme should not have imported.")
	}

	// Try importing an invalid scheme.
	data.DisplayName = nil

	if err := th.App.ImportScheme(&data, false); err == nil {
		t.Fatalf("Should have failed to import.")
	}

	if res := <-th.App.Srv.Store.Scheme().GetByName(*data.Name); res.Err == nil {
		t.Fatalf("Scheme should not have imported.")
	}

	// Try importing a valid scheme with all params set.
	data.DisplayName = ptrStr("display name")

	if err := th.App.ImportScheme(&data, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	if res := <-th.App.Srv.Store.Scheme().GetByName(*data.Name); res.Err != nil {
		t.Fatalf("Failed to import scheme: %v", res.Err)
	} else {
		scheme := res.Data.(*model.Scheme)
		assert.Equal(t, *data.Name, scheme.Name)
		assert.Equal(t, *data.DisplayName, scheme.DisplayName)
		assert.Equal(t, *data.Description, scheme.Description)
		assert.Equal(t, *data.Scope, scheme.Scope)

		if res := <-th.App.Srv.Store.Role().GetByName(scheme.DefaultTeamAdminRole); res.Err != nil {
			t.Fatalf("Should have found the imported role.")
		} else {
			role := res.Data.(*model.Role)
			assert.Equal(t, *data.DefaultTeamAdminRole.DisplayName, role.DisplayName)
			assert.False(t, role.BuiltIn)
			assert.True(t, role.SchemeManaged)
		}

		if res := <-th.App.Srv.Store.Role().GetByName(scheme.DefaultTeamUserRole); res.Err != nil {
			t.Fatalf("Should have found the imported role.")
		} else {
			role := res.Data.(*model.Role)
			assert.Equal(t, *data.DefaultTeamUserRole.DisplayName, role.DisplayName)
			assert.False(t, role.BuiltIn)
			assert.True(t, role.SchemeManaged)
		}

		if res := <-th.App.Srv.Store.Role().GetByName(scheme.DefaultChannelAdminRole); res.Err != nil {
			t.Fatalf("Should have found the imported role.")
		} else {
			role := res.Data.(*model.Role)
			assert.Equal(t, *data.DefaultChannelAdminRole.DisplayName, role.DisplayName)
			assert.False(t, role.BuiltIn)
			assert.True(t, role.SchemeManaged)
		}

		if res := <-th.App.Srv.Store.Role().GetByName(scheme.DefaultChannelUserRole); res.Err != nil {
			t.Fatalf("Should have found the imported role.")
		} else {
			role := res.Data.(*model.Role)
			assert.Equal(t, *data.DefaultChannelUserRole.DisplayName, role.DisplayName)
			assert.False(t, role.BuiltIn)
			assert.True(t, role.SchemeManaged)
		}
	}

	// Try modifying all the fields and re-importing.
	data.DisplayName = ptrStr("new display name")
	data.Description = ptrStr("new description")

	if err := th.App.ImportScheme(&data, false); err != nil {
		t.Fatalf("Should have succeeded: %v", err)
	}

	if res := <-th.App.Srv.Store.Scheme().GetByName(*data.Name); res.Err != nil {
		t.Fatalf("Failed to import scheme: %v", res.Err)
	} else {
		scheme := res.Data.(*model.Scheme)
		assert.Equal(t, *data.Name, scheme.Name)
		assert.Equal(t, *data.DisplayName, scheme.DisplayName)
		assert.Equal(t, *data.Description, scheme.Description)
		assert.Equal(t, *data.Scope, scheme.Scope)

		if res := <-th.App.Srv.Store.Role().GetByName(scheme.DefaultTeamAdminRole); res.Err != nil {
			t.Fatalf("Should have found the imported role.")
		} else {
			role := res.Data.(*model.Role)
			assert.Equal(t, *data.DefaultTeamAdminRole.DisplayName, role.DisplayName)
			assert.False(t, role.BuiltIn)
			assert.True(t, role.SchemeManaged)
		}

		if res := <-th.App.Srv.Store.Role().GetByName(scheme.DefaultTeamUserRole); res.Err != nil {
			t.Fatalf("Should have found the imported role.")
		} else {
			role := res.Data.(*model.Role)
			assert.Equal(t, *data.DefaultTeamUserRole.DisplayName, role.DisplayName)
			assert.False(t, role.BuiltIn)
			assert.True(t, role.SchemeManaged)
		}

		if res := <-th.App.Srv.Store.Role().GetByName(scheme.DefaultChannelAdminRole); res.Err != nil {
			t.Fatalf("Should have found the imported role.")
		} else {
			role := res.Data.(*model.Role)
			assert.Equal(t, *data.DefaultChannelAdminRole.DisplayName, role.DisplayName)
			assert.False(t, role.BuiltIn)
			assert.True(t, role.SchemeManaged)
		}

		if res := <-th.App.Srv.Store.Role().GetByName(scheme.DefaultChannelUserRole); res.Err != nil {
			t.Fatalf("Should have found the imported role.")
		} else {
			role := res.Data.(*model.Role)
			assert.Equal(t, *data.DefaultChannelUserRole.DisplayName, role.DisplayName)
			assert.False(t, role.BuiltIn)
			assert.True(t, role.SchemeManaged)
		}
	}

	// Try changing the scope of the scheme and reimporting.
	data.Scope = ptrStr("channel")

	if err := th.App.ImportScheme(&data, false); err == nil {
		t.Fatalf("Should have failed to import.")
	}

	if res := <-th.App.Srv.Store.Scheme().GetByName(*data.Name); res.Err != nil {
		t.Fatalf("Failed to import scheme: %v", res.Err)
	} else {
		scheme := res.Data.(*model.Scheme)
		assert.Equal(t, *data.Name, scheme.Name)
		assert.Equal(t, *data.DisplayName, scheme.DisplayName)
		assert.Equal(t, *data.Description, scheme.Description)
		assert.Equal(t, "team", scheme.Scope)
	}
}

func TestImportImportRole(t *testing.T) {
	th := Setup()
	defer th.TearDown()

	// Try importing an invalid role in dryRun mode.
	rid1 := model.NewId()
	data := RoleImportData{
		Name: &rid1,
	}

	if err := th.App.ImportRole(&data, true, false); err == nil {
		t.Fatalf("Should have failed to import.")
	}

	if res := <-th.App.Srv.Store.Role().GetByName(rid1); res.Err == nil {
		t.Fatalf("Role should not have imported.")
	}

	// Try importing the valid role in dryRun mode.
	data.DisplayName = ptrStr("display name")

	if err := th.App.ImportRole(&data, true, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	if res := <-th.App.Srv.Store.Role().GetByName(rid1); res.Err == nil {
		t.Fatalf("Role should not have imported as we are in dry run mode.")
	}

	// Try importing an invalid role.
	data.DisplayName = nil

	if err := th.App.ImportRole(&data, false, false); err == nil {
		t.Fatalf("Should have failed to import.")
	}

	if res := <-th.App.Srv.Store.Role().GetByName(rid1); res.Err == nil {
		t.Fatalf("Role should not have imported.")
	}

	// Try importing a valid role with all params set.
	data.DisplayName = ptrStr("display name")
	data.Description = ptrStr("description")
	data.Permissions = &[]string{"invite_user", "add_user_to_team"}

	if err := th.App.ImportRole(&data, false, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	if res := <-th.App.Srv.Store.Role().GetByName(rid1); res.Err != nil {
		t.Fatalf("Should have found the imported role.")
	} else {
		role := res.Data.(*model.Role)
		assert.Equal(t, *data.Name, role.Name)
		assert.Equal(t, *data.DisplayName, role.DisplayName)
		assert.Equal(t, *data.Description, role.Description)
		assert.Equal(t, *data.Permissions, role.Permissions)
		assert.False(t, role.BuiltIn)
		assert.False(t, role.SchemeManaged)
	}

	// Try changing all the params and reimporting.
	data.DisplayName = ptrStr("new display name")
	data.Description = ptrStr("description")
	data.Permissions = &[]string{"use_slash_commands"}

	if err := th.App.ImportRole(&data, false, true); err != nil {
		t.Fatalf("Should have succeeded. %v", err)
	}

	if res := <-th.App.Srv.Store.Role().GetByName(rid1); res.Err != nil {
		t.Fatalf("Should have found the imported role.")
	} else {
		role := res.Data.(*model.Role)
		assert.Equal(t, *data.Name, role.Name)
		assert.Equal(t, *data.DisplayName, role.DisplayName)
		assert.Equal(t, *data.Description, role.Description)
		assert.Equal(t, *data.Permissions, role.Permissions)
		assert.False(t, role.BuiltIn)
		assert.True(t, role.SchemeManaged)
	}

	// Check that re-importing with only required fields doesn't update the others.
	data2 := RoleImportData{
		Name:        &rid1,
		DisplayName: ptrStr("new display name again"),
	}

	if err := th.App.ImportRole(&data2, false, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	if res := <-th.App.Srv.Store.Role().GetByName(rid1); res.Err != nil {
		t.Fatalf("Should have found the imported role.")
	} else {
		role := res.Data.(*model.Role)
		assert.Equal(t, *data2.Name, role.Name)
		assert.Equal(t, *data2.DisplayName, role.DisplayName)
		assert.Equal(t, *data.Description, role.Description)
		assert.Equal(t, *data.Permissions, role.Permissions)
		assert.False(t, role.BuiltIn)
		assert.False(t, role.SchemeManaged)
	}
}

func TestImportImportTeam(t *testing.T) {
	th := Setup()
	defer th.TearDown()

	// Mark the phase 2 permissions migration as completed.
	<-th.App.Srv.Store.System().Save(&model.System{Name: model.MIGRATION_KEY_ADVANCED_PERMISSIONS_PHASE_2, Value: "true"})

	defer func() {
		<-th.App.Srv.Store.System().PermanentDeleteByName(model.MIGRATION_KEY_ADVANCED_PERMISSIONS_PHASE_2)
	}()

	scheme1 := th.SetupTeamScheme()
	scheme2 := th.SetupTeamScheme()

	// Check how many teams are in the database.
	var teamsCount int64
	if r := <-th.App.Srv.Store.Team().AnalyticsTeamCount(); r.Err == nil {
		teamsCount = r.Data.(int64)
	} else {
		t.Fatalf("Failed to get team count.")
	}

	data := TeamImportData{
		Name:            ptrStr(model.NewId()),
		DisplayName:     ptrStr("Display Name"),
		Type:            ptrStr("XYZ"),
		Description:     ptrStr("The team description."),
		AllowOpenInvite: ptrBool(true),
		Scheme:          &scheme1.Name,
	}

	// Try importing an invalid team in dryRun mode.
	if err := th.App.ImportTeam(&data, true); err == nil {
		t.Fatalf("Should have received an error importing an invalid team.")
	}

	// Do a valid team in dry-run mode.
	data.Type = ptrStr("O")
	if err := th.App.ImportTeam(&data, true); err != nil {
		t.Fatalf("Received an error validating valid team.")
	}

	// Check that no more teams are in the DB.
	th.CheckTeamCount(t, teamsCount)

	// Do an invalid team in apply mode, check db changes.
	data.Type = ptrStr("XYZ")
	if err := th.App.ImportTeam(&data, false); err == nil {
		t.Fatalf("Import should have failed on invalid team.")
	}

	// Check that no more teams are in the DB.
	th.CheckTeamCount(t, teamsCount)

	// Do a valid team in apply mode, check db changes.
	data.Type = ptrStr("O")
	if err := th.App.ImportTeam(&data, false); err != nil {
		t.Fatalf("Received an error importing valid team: %v", err)
	}

	// Check that one more team is in the DB.
	th.CheckTeamCount(t, teamsCount+1)

	// Get the team and check that all the fields are correct.
	if team, err := th.App.GetTeamByName(*data.Name); err != nil {
		t.Fatalf("Failed to get team from database.")
	} else {
		assert.Equal(t, *data.DisplayName, team.DisplayName)
		assert.Equal(t, *data.Type, team.Type)
		assert.Equal(t, *data.Description, team.Description)
		assert.Equal(t, *data.AllowOpenInvite, team.AllowOpenInvite)
		assert.Equal(t, scheme1.Id, *team.SchemeId)
	}

	// Alter all the fields of that team (apart from unique identifier) and import again.
	data.DisplayName = ptrStr("Display Name 2")
	data.Type = ptrStr("P")
	data.Description = ptrStr("The new description")
	data.AllowOpenInvite = ptrBool(false)
	data.Scheme = &scheme2.Name

	// Check that the original number of teams are again in the DB (because this query doesn't include deleted).
	data.Type = ptrStr("O")
	if err := th.App.ImportTeam(&data, false); err != nil {
		t.Fatalf("Received an error importing updated valid team.")
	}

	th.CheckTeamCount(t, teamsCount+1)

	// Get the team and check that all fields are correct.
	if team, err := th.App.GetTeamByName(*data.Name); err != nil {
		t.Fatalf("Failed to get team from database.")
	} else {
		assert.Equal(t, *data.DisplayName, team.DisplayName)
		assert.Equal(t, *data.Type, team.Type)
		assert.Equal(t, *data.Description, team.Description)
		assert.Equal(t, *data.AllowOpenInvite, team.AllowOpenInvite)
		assert.Equal(t, scheme2.Id, *team.SchemeId)
	}
}

func TestImportImportChannel(t *testing.T) {
	th := Setup()
	defer th.TearDown()

	// Mark the phase 2 permissions migration as completed.
	<-th.App.Srv.Store.System().Save(&model.System{Name: model.MIGRATION_KEY_ADVANCED_PERMISSIONS_PHASE_2, Value: "true"})

	defer func() {
		<-th.App.Srv.Store.System().PermanentDeleteByName(model.MIGRATION_KEY_ADVANCED_PERMISSIONS_PHASE_2)
	}()

	scheme1 := th.SetupChannelScheme()
	scheme2 := th.SetupChannelScheme()

	// Import a Team.
	teamName := model.NewId()
	th.App.ImportTeam(&TeamImportData{
		Name:        &teamName,
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
	}, false)
	team, err := th.App.GetTeamByName(teamName)
	if err != nil {
		t.Fatalf("Failed to get team from database.")
	}

	// Check how many channels are in the database.
	var channelCount int64
	if r := <-th.App.Srv.Store.Channel().AnalyticsTypeCount("", model.CHANNEL_OPEN); r.Err == nil {
		channelCount = r.Data.(int64)
	} else {
		t.Fatalf("Failed to get team count.")
	}

	// Do an invalid channel in dry-run mode.
	data := ChannelImportData{
		Team:        &teamName,
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
		Header:      ptrStr("Channe Header"),
		Purpose:     ptrStr("Channel Purpose"),
		Scheme:      &scheme1.Name,
	}
	if err := th.App.ImportChannel(&data, true); err == nil {
		t.Fatalf("Expected error due to invalid name.")
	}

	// Check that no more channels are in the DB.
	th.CheckChannelsCount(t, channelCount)

	// Do a valid channel with a nonexistent team in dry-run mode.
	data.Name = ptrStr("channelname")
	data.Team = ptrStr(model.NewId())
	if err := th.App.ImportChannel(&data, true); err != nil {
		t.Fatalf("Expected success as cannot validate channel name in dry run mode.")
	}

	// Check that no more channels are in the DB.
	th.CheckChannelsCount(t, channelCount)

	// Do a valid channel in dry-run mode.
	data.Team = &teamName
	if err := th.App.ImportChannel(&data, true); err != nil {
		t.Fatalf("Expected success as valid team.")
	}

	// Check that no more channels are in the DB.
	th.CheckChannelsCount(t, channelCount)

	// Do an invalid channel in apply mode.
	data.Name = nil
	if err := th.App.ImportChannel(&data, false); err == nil {
		t.Fatalf("Expected error due to invalid name (apply mode).")
	}

	// Check that no more channels are in the DB.
	th.CheckChannelsCount(t, channelCount)

	// Do a valid channel in apply mode with a non-existent team.
	data.Name = ptrStr("channelname")
	data.Team = ptrStr(model.NewId())
	if err := th.App.ImportChannel(&data, false); err == nil {
		t.Fatalf("Expected error due to non-existent team (apply mode).")
	}

	// Check that no more channels are in the DB.
	th.CheckChannelsCount(t, channelCount)

	// Do a valid channel in apply mode.
	data.Team = &teamName
	if err := th.App.ImportChannel(&data, false); err != nil {
		t.Fatalf("Expected success in apply mode: %v", err.Error())
	}

	// Check that 1 more channel is in the DB.
	th.CheckChannelsCount(t, channelCount+1)

	// Get the Channel and check all the fields are correct.
	if channel, err := th.App.GetChannelByName(*data.Name, team.Id); err != nil {
		t.Fatalf("Failed to get channel from database.")
	} else {
		assert.Equal(t, *data.Name, channel.Name)
		assert.Equal(t, *data.DisplayName, channel.DisplayName)
		assert.Equal(t, *data.Type, channel.Type)
		assert.Equal(t, *data.Header, channel.Header)
		assert.Equal(t, *data.Purpose, channel.Purpose)
		assert.Equal(t, scheme1.Id, *channel.SchemeId)
	}

	// Alter all the fields of that channel.
	data.DisplayName = ptrStr("Chaned Disp Name")
	data.Type = ptrStr(model.CHANNEL_PRIVATE)
	data.Header = ptrStr("New Header")
	data.Purpose = ptrStr("New Purpose")
	data.Scheme = &scheme2.Name
	if err := th.App.ImportChannel(&data, false); err != nil {
		t.Fatalf("Expected success in apply mode: %v", err.Error())
	}

	// Check channel count the same.
	th.CheckChannelsCount(t, channelCount)

	// Get the Channel and check all the fields are correct.
	if channel, err := th.App.GetChannelByName(*data.Name, team.Id); err != nil {
		t.Fatalf("Failed to get channel from database.")
	} else {
		assert.Equal(t, *data.Name, channel.Name)
		assert.Equal(t, *data.DisplayName, channel.DisplayName)
		assert.Equal(t, *data.Type, channel.Type)
		assert.Equal(t, *data.Header, channel.Header)
		assert.Equal(t, *data.Purpose, channel.Purpose)
		assert.Equal(t, scheme2.Id, *channel.SchemeId)
	}

}

func TestImportImportUser(t *testing.T) {
	th := Setup()
	defer th.TearDown()

	// Check how many users are in the database.
	var userCount int64
	if r := <-th.App.Srv.Store.User().GetTotalUsersCount(); r.Err == nil {
		userCount = r.Data.(int64)
	} else {
		t.Fatalf("Failed to get user count.")
	}

	// Do an invalid user in dry-run mode.
	data := UserImportData{
		Username: ptrStr(model.NewId()),
	}
	if err := th.App.ImportUser(&data, true); err == nil {
		t.Fatalf("Should have failed to import invalid user.")
	}

	// Check that no more users are in the DB.
	if r := <-th.App.Srv.Store.User().GetTotalUsersCount(); r.Err == nil {
		if r.Data.(int64) != userCount {
			t.Fatalf("Unexpected number of users")
		}
	} else {
		t.Fatalf("Failed to get user count.")
	}

	// Do a valid user in dry-run mode.
	data = UserImportData{
		Username: ptrStr(model.NewId()),
		Email:    ptrStr(model.NewId() + "@example.com"),
	}
	if err := th.App.ImportUser(&data, true); err != nil {
		t.Fatalf("Should have succeeded to import valid user.")
	}

	// Check that no more users are in the DB.
	if r := <-th.App.Srv.Store.User().GetTotalUsersCount(); r.Err == nil {
		if r.Data.(int64) != userCount {
			t.Fatalf("Unexpected number of users")
		}
	} else {
		t.Fatalf("Failed to get user count.")
	}

	// Do an invalid user in apply mode.
	data = UserImportData{
		Username: ptrStr(model.NewId()),
	}
	if err := th.App.ImportUser(&data, false); err == nil {
		t.Fatalf("Should have failed to import invalid user.")
	}

	// Check that no more users are in the DB.
	if r := <-th.App.Srv.Store.User().GetTotalUsersCount(); r.Err == nil {
		if r.Data.(int64) != userCount {
			t.Fatalf("Unexpected number of users")
		}
	} else {
		t.Fatalf("Failed to get user count.")
	}

	// Do a valid user in apply mode.
	username := model.NewId()
	testsDir, _ := utils.FindDir("tests")
	data = UserImportData{
		ProfileImage: ptrStr(filepath.Join(testsDir, "test.png")),
		Username:     &username,
		Email:        ptrStr(model.NewId() + "@example.com"),
		Nickname:     ptrStr(model.NewId()),
		FirstName:    ptrStr(model.NewId()),
		LastName:     ptrStr(model.NewId()),
		Position:     ptrStr(model.NewId()),
	}
	if err := th.App.ImportUser(&data, false); err != nil {
		t.Fatalf("Should have succeeded to import valid user.")
	}

	// Check that one more user is in the DB.
	if r := <-th.App.Srv.Store.User().GetTotalUsersCount(); r.Err == nil {
		if r.Data.(int64) != userCount+1 {
			t.Fatalf("Unexpected number of users")
		}
	} else {
		t.Fatalf("Failed to get user count.")
	}

	// Get the user and check all the fields are correct.
	if user, err := th.App.GetUserByUsername(username); err != nil {
		t.Fatalf("Failed to get user from database.")
	} else {
		if user.Email != *data.Email || user.Nickname != *data.Nickname || user.FirstName != *data.FirstName || user.LastName != *data.LastName || user.Position != *data.Position {
			t.Fatalf("User properties do not match Import Data.")
		}
		// Check calculated properties.
		if user.AuthService != "" {
			t.Fatalf("Expected Auth Service to be empty.")
		}

		if !(user.AuthData == nil || *user.AuthData == "") {
			t.Fatalf("Expected AuthData to be empty.")
		}

		if len(user.Password) == 0 {
			t.Fatalf("Expected password to be set.")
		}

		if !user.EmailVerified {
			t.Fatalf("Expected EmailVerified to be true.")
		}

		if user.Locale != *th.App.Config().LocalizationSettings.DefaultClientLocale {
			t.Fatalf("Expected Locale to be the default.")
		}

		if user.Roles != "system_user" {
			t.Fatalf("Expected roles to be system_user")
		}
	}

	// Alter all the fields of that user.
	data.Email = ptrStr(model.NewId() + "@example.com")
	data.ProfileImage = ptrStr(filepath.Join(testsDir, "testgif.gif"))
	data.AuthService = ptrStr("ldap")
	data.AuthData = &username
	data.Nickname = ptrStr(model.NewId())
	data.FirstName = ptrStr(model.NewId())
	data.LastName = ptrStr(model.NewId())
	data.Position = ptrStr(model.NewId())
	data.Roles = ptrStr("system_admin system_user")
	data.Locale = ptrStr("zh_CN")
	if err := th.App.ImportUser(&data, false); err != nil {
		t.Fatalf("Should have succeeded to update valid user %v", err)
	}

	// Check user count the same.
	if r := <-th.App.Srv.Store.User().GetTotalUsersCount(); r.Err == nil {
		if r.Data.(int64) != userCount+1 {
			t.Fatalf("Unexpected number of users")
		}
	} else {
		t.Fatalf("Failed to get user count.")
	}

	// Get the user and check all the fields are correct.
	if user, err := th.App.GetUserByUsername(username); err != nil {
		t.Fatalf("Failed to get user from database.")
	} else {
		if user.Email != *data.Email || user.Nickname != *data.Nickname || user.FirstName != *data.FirstName || user.LastName != *data.LastName || user.Position != *data.Position {
			t.Fatalf("Updated User properties do not match Import Data.")
		}
		// Check calculated properties.
		if user.AuthService != "ldap" {
			t.Fatalf("Expected Auth Service to be ldap \"%v\"", user.AuthService)
		}

		if !(user.AuthData == data.AuthData || *user.AuthData == *data.AuthData) {
			t.Fatalf("Expected AuthData to be set.")
		}

		if len(user.Password) != 0 {
			t.Fatalf("Expected password to be empty.")
		}

		if !user.EmailVerified {
			t.Fatalf("Expected EmailVerified to be true.")
		}

		if user.Locale != *data.Locale {
			t.Fatalf("Expected Locale to be the set.")
		}

		if user.Roles != *data.Roles {
			t.Fatalf("Expected roles to be set: %v", user.Roles)
		}
	}

	// Check Password and AuthData together.
	data.Password = ptrStr("PasswordTest")
	if err := th.App.ImportUser(&data, false); err == nil {
		t.Fatalf("Should have failed to import invalid user.")
	}

	data.AuthData = nil
	if err := th.App.ImportUser(&data, false); err != nil {
		t.Fatalf("Should have succeeded to update valid user %v", err)
	}

	data.Password = ptrStr("")
	if err := th.App.ImportUser(&data, false); err == nil {
		t.Fatalf("Should have failed to import invalid user.")
	}

	data.Password = ptrStr(strings.Repeat("0123456789", 10))
	if err := th.App.ImportUser(&data, false); err == nil {
		t.Fatalf("Should have failed to import invalid user.")
	}

	data.Password = ptrStr("TestPassword")

	// Test team and channel memberships
	teamName := model.NewId()
	th.App.ImportTeam(&TeamImportData{
		Name:        &teamName,
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
	}, false)
	team, err := th.App.GetTeamByName(teamName)
	if err != nil {
		t.Fatalf("Failed to get team from database.")
	}

	channelName := model.NewId()
	th.App.ImportChannel(&ChannelImportData{
		Team:        &teamName,
		Name:        &channelName,
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
	}, false)
	channel, err := th.App.GetChannelByName(channelName, team.Id)
	if err != nil {
		t.Fatalf("Failed to get channel from database.")
	}

	username = model.NewId()
	data = UserImportData{
		Username:  &username,
		Email:     ptrStr(model.NewId() + "@example.com"),
		Nickname:  ptrStr(model.NewId()),
		FirstName: ptrStr(model.NewId()),
		LastName:  ptrStr(model.NewId()),
		Position:  ptrStr(model.NewId()),
	}

	teamMembers, err := th.App.GetTeamMembers(team.Id, 0, 1000)
	if err != nil {
		t.Fatalf("Failed to get team member count")
	}
	teamMemberCount := len(teamMembers)

	channelMemberCount, err := th.App.GetChannelMemberCount(channel.Id)
	if err != nil {
		t.Fatalf("Failed to get channel member count")
	}

	// Test with an invalid team & channel membership in dry-run mode.
	data.Teams = &[]UserTeamImportData{
		{
			Roles: ptrStr("invalid"),
			Channels: &[]UserChannelImportData{
				{
					Roles: ptrStr("invalid"),
				},
			},
		},
	}
	if err := th.App.ImportUser(&data, true); err == nil {
		t.Fatalf("Should have failed.")
	}

	// Test with an unknown team name & invalid channel membership in dry-run mode.
	data.Teams = &[]UserTeamImportData{
		{
			Name: ptrStr(model.NewId()),
			Channels: &[]UserChannelImportData{
				{
					Roles: ptrStr("invalid"),
				},
			},
		},
	}
	if err := th.App.ImportUser(&data, true); err == nil {
		t.Fatalf("Should have failed.")
	}

	// Test with a valid team & invalid channel membership in dry-run mode.
	data.Teams = &[]UserTeamImportData{
		{
			Name: &teamName,
			Channels: &[]UserChannelImportData{
				{
					Roles: ptrStr("invalid"),
				},
			},
		},
	}
	if err := th.App.ImportUser(&data, true); err == nil {
		t.Fatalf("Should have failed.")
	}

	// Test with a valid team & unknown channel name in dry-run mode.
	data.Teams = &[]UserTeamImportData{
		{
			Name: &teamName,
			Channels: &[]UserChannelImportData{
				{
					Name: ptrStr(model.NewId()),
				},
			},
		},
	}
	if err := th.App.ImportUser(&data, true); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	// Test with a valid team & valid channel name in dry-run mode.
	data.Teams = &[]UserTeamImportData{
		{
			Name: &teamName,
			Channels: &[]UserChannelImportData{
				{
					Name: &channelName,
				},
			},
		},
	}
	if err := th.App.ImportUser(&data, true); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	// Check no new member objects were created because dry run mode.
	if tmc, err := th.App.GetTeamMembers(team.Id, 0, 1000); err != nil {
		t.Fatalf("Failed to get Team Member Count")
	} else if len(tmc) != teamMemberCount {
		t.Fatalf("Number of team members not as expected")
	}

	if cmc, err := th.App.GetChannelMemberCount(channel.Id); err != nil {
		t.Fatalf("Failed to get Channel Member Count")
	} else if cmc != channelMemberCount {
		t.Fatalf("Number of channel members not as expected")
	}

	// Test with an invalid team & channel membership in apply mode.
	data.Teams = &[]UserTeamImportData{
		{
			Roles: ptrStr("invalid"),
			Channels: &[]UserChannelImportData{
				{
					Roles: ptrStr("invalid"),
				},
			},
		},
	}
	if err := th.App.ImportUser(&data, false); err == nil {
		t.Fatalf("Should have failed.")
	}

	// Test with an unknown team name & invalid channel membership in apply mode.
	data.Teams = &[]UserTeamImportData{
		{
			Name: ptrStr(model.NewId()),
			Channels: &[]UserChannelImportData{
				{
					Roles: ptrStr("invalid"),
				},
			},
		},
	}
	if err := th.App.ImportUser(&data, false); err == nil {
		t.Fatalf("Should have failed.")
	}

	// Test with a valid team & invalid channel membership in apply mode.
	data.Teams = &[]UserTeamImportData{
		{
			Name: &teamName,
			Channels: &[]UserChannelImportData{
				{
					Roles: ptrStr("invalid"),
				},
			},
		},
	}
	if err := th.App.ImportUser(&data, false); err == nil {
		t.Fatalf("Should have failed.")
	}

	// Check no new member objects were created because all tests should have failed so far.
	if tmc, err := th.App.GetTeamMembers(team.Id, 0, 1000); err != nil {
		t.Fatalf("Failed to get Team Member Count")
	} else if len(tmc) != teamMemberCount {
		t.Fatalf("Number of team members not as expected")
	}

	if cmc, err := th.App.GetChannelMemberCount(channel.Id); err != nil {
		t.Fatalf("Failed to get Channel Member Count")
	} else if cmc != channelMemberCount {
		t.Fatalf("Number of channel members not as expected")
	}

	// Test with a valid team & unknown channel name in apply mode.
	data.Teams = &[]UserTeamImportData{
		{
			Name: &teamName,
			Channels: &[]UserChannelImportData{
				{
					Name: ptrStr(model.NewId()),
				},
			},
		},
	}
	if err := th.App.ImportUser(&data, false); err == nil {
		t.Fatalf("Should have failed.")
	}

	// Check only new team member object created because dry run mode.
	if tmc, err := th.App.GetTeamMembers(team.Id, 0, 1000); err != nil {
		t.Fatalf("Failed to get Team Member Count")
	} else if len(tmc) != teamMemberCount+1 {
		t.Fatalf("Number of team members not as expected")
	}

	if cmc, err := th.App.GetChannelMemberCount(channel.Id); err != nil {
		t.Fatalf("Failed to get Channel Member Count")
	} else if cmc != channelMemberCount {
		t.Fatalf("Number of channel members not as expected")
	}

	// Check team member properties.
	user, err := th.App.GetUserByUsername(username)
	if err != nil {
		t.Fatalf("Failed to get user from database.")
	}
	if teamMember, err := th.App.GetTeamMember(team.Id, user.Id); err != nil {
		t.Fatalf("Failed to get team member from database.")
	} else if teamMember.Roles != "team_user" {
		t.Fatalf("Team member properties not as expected")
	}

	// Test with a valid team & valid channel name in apply mode.
	data.Teams = &[]UserTeamImportData{
		{
			Name: &teamName,
			Channels: &[]UserChannelImportData{
				{
					Name: &channelName,
				},
			},
		},
	}
	if err := th.App.ImportUser(&data, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	// Check only new channel member object created because dry run mode.
	if tmc, err := th.App.GetTeamMembers(team.Id, 0, 1000); err != nil {
		t.Fatalf("Failed to get Team Member Count")
	} else if len(tmc) != teamMemberCount+1 {
		t.Fatalf("Number of team members not as expected")
	}

	if cmc, err := th.App.GetChannelMemberCount(channel.Id); err != nil {
		t.Fatalf("Failed to get Channel Member Count")
	} else if cmc != channelMemberCount+1 {
		t.Fatalf("Number of channel members not as expected")
	}

	// Check channel member properties.
	if channelMember, err := th.App.GetChannelMember(channel.Id, user.Id); err != nil {
		t.Fatalf("Failed to get channel member from database.")
	} else if channelMember.Roles != "channel_user" || channelMember.NotifyProps[model.DESKTOP_NOTIFY_PROP] != "default" || channelMember.NotifyProps[model.PUSH_NOTIFY_PROP] != "default" || channelMember.NotifyProps[model.MARK_UNREAD_NOTIFY_PROP] != "all" {
		t.Fatalf("Channel member properties not as expected")
	}

	// Test with the properties of the team and channel membership changed.
	data.Teams = &[]UserTeamImportData{
		{
			Name:  &teamName,
			Roles: ptrStr("team_user team_admin"),
			Channels: &[]UserChannelImportData{
				{
					Name:  &channelName,
					Roles: ptrStr("channel_user channel_admin"),
					NotifyProps: &UserChannelNotifyPropsImportData{
						Desktop:    ptrStr(model.USER_NOTIFY_MENTION),
						Mobile:     ptrStr(model.USER_NOTIFY_MENTION),
						MarkUnread: ptrStr(model.USER_NOTIFY_MENTION),
					},
					Favorite: ptrBool(true),
				},
			},
		},
	}
	if err := th.App.ImportUser(&data, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	// Check both member properties.
	if teamMember, err := th.App.GetTeamMember(team.Id, user.Id); err != nil {
		t.Fatalf("Failed to get team member from database.")
	} else if teamMember.Roles != "team_user team_admin" {
		t.Fatalf("Team member properties not as expected: %v", teamMember.Roles)
	}

	if channelMember, err := th.App.GetChannelMember(channel.Id, user.Id); err != nil {
		t.Fatalf("Failed to get channel member Desktop from database.")
	} else if channelMember.Roles != "channel_user channel_admin" || channelMember.NotifyProps[model.DESKTOP_NOTIFY_PROP] != model.USER_NOTIFY_MENTION || channelMember.NotifyProps[model.PUSH_NOTIFY_PROP] != model.USER_NOTIFY_MENTION || channelMember.NotifyProps[model.MARK_UNREAD_NOTIFY_PROP] != model.USER_NOTIFY_MENTION {
		t.Fatalf("Channel member properties not as expected")
	}

	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_FAVORITE_CHANNEL, channel.Id, "true")

	// No more new member objects.
	if tmc, err := th.App.GetTeamMembers(team.Id, 0, 1000); err != nil {
		t.Fatalf("Failed to get Team Member Count")
	} else if len(tmc) != teamMemberCount+1 {
		t.Fatalf("Number of team members not as expected")
	}

	if cmc, err := th.App.GetChannelMemberCount(channel.Id); err != nil {
		t.Fatalf("Failed to get Channel Member Count")
	} else if cmc != channelMemberCount+1 {
		t.Fatalf("Number of channel members not as expected")
	}

	// Add a user with some preferences.
	username = model.NewId()
	data = UserImportData{
		Username:           &username,
		Email:              ptrStr(model.NewId() + "@example.com"),
		Theme:              ptrStr(`{"awayIndicator":"#DCBD4E","buttonBg":"#23A2FF","buttonColor":"#FFFFFF","centerChannelBg":"#ffffff","centerChannelColor":"#333333","codeTheme":"github","image":"/static/files/a4a388b38b32678e83823ef1b3e17766.png","linkColor":"#2389d7","mentionBg":"#2389d7","mentionColor":"#ffffff","mentionHighlightBg":"#fff2bb","mentionHighlightLink":"#2f81b7","newMessageSeparator":"#FF8800","onlineIndicator":"#7DBE00","sidebarBg":"#fafafa","sidebarHeaderBg":"#3481B9","sidebarHeaderTextColor":"#ffffff","sidebarText":"#333333","sidebarTextActiveBorder":"#378FD2","sidebarTextActiveColor":"#111111","sidebarTextHoverBg":"#e6f2fa","sidebarUnreadText":"#333333","type":"Mattermost"}`),
		UseMilitaryTime:    ptrStr("true"),
		CollapsePreviews:   ptrStr("true"),
		MessageDisplay:     ptrStr("compact"),
		ChannelDisplayMode: ptrStr("centered"),
		TutorialStep:       ptrStr("3"),
	}
	if err := th.App.ImportUser(&data, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	// Check their values.
	user, err = th.App.GetUserByUsername(username)
	if err != nil {
		t.Fatalf("Failed to get user from database.")
	}

	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_THEME, "", *data.Theme)
	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_DISPLAY_SETTINGS, "use_military_time", *data.UseMilitaryTime)
	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_DISPLAY_SETTINGS, "collapse_previews", *data.CollapsePreviews)
	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_DISPLAY_SETTINGS, "message_display", *data.MessageDisplay)
	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_DISPLAY_SETTINGS, "channel_display_mode", *data.ChannelDisplayMode)
	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_TUTORIAL_STEPS, user.Id, *data.TutorialStep)

	// Change those preferences.
	data = UserImportData{
		Username:           &username,
		Email:              ptrStr(model.NewId() + "@example.com"),
		Theme:              ptrStr(`{"awayIndicator":"#123456","buttonBg":"#23A2FF","buttonColor":"#FFFFFF","centerChannelBg":"#ffffff","centerChannelColor":"#333333","codeTheme":"github","image":"/static/files/a4a388b38b32678e83823ef1b3e17766.png","linkColor":"#2389d7","mentionBg":"#2389d7","mentionColor":"#ffffff","mentionHighlightBg":"#fff2bb","mentionHighlightLink":"#2f81b7","newMessageSeparator":"#FF8800","onlineIndicator":"#7DBE00","sidebarBg":"#fafafa","sidebarHeaderBg":"#3481B9","sidebarHeaderTextColor":"#ffffff","sidebarText":"#333333","sidebarTextActiveBorder":"#378FD2","sidebarTextActiveColor":"#111111","sidebarTextHoverBg":"#e6f2fa","sidebarUnreadText":"#333333","type":"Mattermost"}`),
		UseMilitaryTime:    ptrStr("false"),
		CollapsePreviews:   ptrStr("false"),
		MessageDisplay:     ptrStr("clean"),
		ChannelDisplayMode: ptrStr("full"),
		TutorialStep:       ptrStr("2"),
	}
	if err := th.App.ImportUser(&data, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	// Check their values again.
	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_THEME, "", *data.Theme)
	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_DISPLAY_SETTINGS, "use_military_time", *data.UseMilitaryTime)
	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_DISPLAY_SETTINGS, "collapse_previews", *data.CollapsePreviews)
	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_DISPLAY_SETTINGS, "message_display", *data.MessageDisplay)
	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_DISPLAY_SETTINGS, "channel_display_mode", *data.ChannelDisplayMode)
	checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_TUTORIAL_STEPS, user.Id, *data.TutorialStep)

	// Set Notify Props
	data.NotifyProps = &UserNotifyPropsImportData{
		Desktop:          ptrStr(model.USER_NOTIFY_ALL),
		DesktopSound:     ptrStr("true"),
		Email:            ptrStr("true"),
		Mobile:           ptrStr(model.USER_NOTIFY_ALL),
		MobilePushStatus: ptrStr(model.STATUS_ONLINE),
		ChannelTrigger:   ptrStr("true"),
		CommentsTrigger:  ptrStr(model.COMMENTS_NOTIFY_ROOT),
		MentionKeys:      ptrStr("valid,misc"),
	}
	if err := th.App.ImportUser(&data, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	user, err = th.App.GetUserByUsername(username)
	if err != nil {
		t.Fatalf("Failed to get user from database.")
	}

	checkNotifyProp(t, user, model.DESKTOP_NOTIFY_PROP, model.USER_NOTIFY_ALL)
	checkNotifyProp(t, user, model.DESKTOP_SOUND_NOTIFY_PROP, "true")
	checkNotifyProp(t, user, model.EMAIL_NOTIFY_PROP, "true")
	checkNotifyProp(t, user, model.PUSH_NOTIFY_PROP, model.USER_NOTIFY_ALL)
	checkNotifyProp(t, user, model.PUSH_STATUS_NOTIFY_PROP, model.STATUS_ONLINE)
	checkNotifyProp(t, user, model.CHANNEL_MENTIONS_NOTIFY_PROP, "true")
	checkNotifyProp(t, user, model.COMMENTS_NOTIFY_PROP, model.COMMENTS_NOTIFY_ROOT)
	checkNotifyProp(t, user, model.MENTION_KEYS_NOTIFY_PROP, "valid,misc")

	// Change Notify Props
	data.NotifyProps = &UserNotifyPropsImportData{
		Desktop:          ptrStr(model.USER_NOTIFY_MENTION),
		DesktopSound:     ptrStr("false"),
		Email:            ptrStr("false"),
		Mobile:           ptrStr(model.USER_NOTIFY_NONE),
		MobilePushStatus: ptrStr(model.STATUS_AWAY),
		ChannelTrigger:   ptrStr("false"),
		CommentsTrigger:  ptrStr(model.COMMENTS_NOTIFY_ANY),
		MentionKeys:      ptrStr("misc"),
	}
	if err := th.App.ImportUser(&data, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	user, err = th.App.GetUserByUsername(username)
	if err != nil {
		t.Fatalf("Failed to get user from database.")
	}

	checkNotifyProp(t, user, model.DESKTOP_NOTIFY_PROP, model.USER_NOTIFY_MENTION)
	checkNotifyProp(t, user, model.DESKTOP_SOUND_NOTIFY_PROP, "false")
	checkNotifyProp(t, user, model.EMAIL_NOTIFY_PROP, "false")
	checkNotifyProp(t, user, model.PUSH_NOTIFY_PROP, model.USER_NOTIFY_NONE)
	checkNotifyProp(t, user, model.PUSH_STATUS_NOTIFY_PROP, model.STATUS_AWAY)
	checkNotifyProp(t, user, model.CHANNEL_MENTIONS_NOTIFY_PROP, "false")
	checkNotifyProp(t, user, model.COMMENTS_NOTIFY_PROP, model.COMMENTS_NOTIFY_ANY)
	checkNotifyProp(t, user, model.MENTION_KEYS_NOTIFY_PROP, "misc")

	// Check Notify Props get set on *create* user.
	username = model.NewId()
	data = UserImportData{
		Username: &username,
		Email:    ptrStr(model.NewId() + "@example.com"),
	}
	data.NotifyProps = &UserNotifyPropsImportData{
		Desktop:          ptrStr(model.USER_NOTIFY_MENTION),
		DesktopSound:     ptrStr("false"),
		Email:            ptrStr("false"),
		Mobile:           ptrStr(model.USER_NOTIFY_NONE),
		MobilePushStatus: ptrStr(model.STATUS_AWAY),
		ChannelTrigger:   ptrStr("false"),
		CommentsTrigger:  ptrStr(model.COMMENTS_NOTIFY_ANY),
		MentionKeys:      ptrStr("misc"),
	}

	if err := th.App.ImportUser(&data, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	user, err = th.App.GetUserByUsername(username)
	if err != nil {
		t.Fatalf("Failed to get user from database.")
	}

	checkNotifyProp(t, user, model.DESKTOP_NOTIFY_PROP, model.USER_NOTIFY_MENTION)
	checkNotifyProp(t, user, model.DESKTOP_SOUND_NOTIFY_PROP, "false")
	checkNotifyProp(t, user, model.EMAIL_NOTIFY_PROP, "false")
	checkNotifyProp(t, user, model.PUSH_NOTIFY_PROP, model.USER_NOTIFY_NONE)
	checkNotifyProp(t, user, model.PUSH_STATUS_NOTIFY_PROP, model.STATUS_AWAY)
	checkNotifyProp(t, user, model.CHANNEL_MENTIONS_NOTIFY_PROP, "false")
	checkNotifyProp(t, user, model.COMMENTS_NOTIFY_PROP, model.COMMENTS_NOTIFY_ANY)
	checkNotifyProp(t, user, model.MENTION_KEYS_NOTIFY_PROP, "misc")

	// Test importing a user with roles set to a team and a channel which are affected by an override scheme.
	// The import subsystem should translate `channel_admin/channel_user/team_admin/team_user`
	// to the appropriate scheme-managed-role booleans.

	// Mark the phase 2 permissions migration as completed.
	<-th.App.Srv.Store.System().Save(&model.System{Name: model.MIGRATION_KEY_ADVANCED_PERMISSIONS_PHASE_2, Value: "true"})

	defer func() {
		<-th.App.Srv.Store.System().PermanentDeleteByName(model.MIGRATION_KEY_ADVANCED_PERMISSIONS_PHASE_2)
	}()

	teamSchemeData := &SchemeImportData{
		Name:        ptrStr(model.NewId()),
		DisplayName: ptrStr(model.NewId()),
		Scope:       ptrStr("team"),
		DefaultTeamUserRole: &RoleImportData{
			Name:        ptrStr(model.NewId()),
			DisplayName: ptrStr(model.NewId()),
		},
		DefaultTeamAdminRole: &RoleImportData{
			Name:        ptrStr(model.NewId()),
			DisplayName: ptrStr(model.NewId()),
		},
		DefaultChannelUserRole: &RoleImportData{
			Name:        ptrStr(model.NewId()),
			DisplayName: ptrStr(model.NewId()),
		},
		DefaultChannelAdminRole: &RoleImportData{
			Name:        ptrStr(model.NewId()),
			DisplayName: ptrStr(model.NewId()),
		},
		Description: ptrStr("description"),
	}

	if err := th.App.ImportScheme(teamSchemeData, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	var teamScheme *model.Scheme
	if res := <-th.App.Srv.Store.Scheme().GetByName(*teamSchemeData.Name); res.Err != nil {
		t.Fatalf("Failed to import scheme: %v", res.Err)
	} else {
		teamScheme = res.Data.(*model.Scheme)
	}

	teamData := &TeamImportData{
		Name:            ptrStr(model.NewId()),
		DisplayName:     ptrStr("Display Name"),
		Type:            ptrStr("O"),
		Description:     ptrStr("The team description."),
		AllowOpenInvite: ptrBool(true),
		Scheme:          &teamScheme.Name,
	}
	if err := th.App.ImportTeam(teamData, false); err != nil {
		t.Fatalf("Import should have succeeded: %v", err.Error())
	}
	team, err = th.App.GetTeamByName(teamName)
	if err != nil {
		t.Fatalf("Failed to get team from database.")
	}

	channelData := &ChannelImportData{
		Team:        &teamName,
		Name:        ptrStr(model.NewId()),
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
		Header:      ptrStr("Channe Header"),
		Purpose:     ptrStr("Channel Purpose"),
	}
	if err := th.App.ImportChannel(channelData, false); err != nil {
		t.Fatalf("Import should have succeeded.")
	}
	channel, err = th.App.GetChannelByName(*channelData.Name, team.Id)
	if err != nil {
		t.Fatalf("Failed to get channel from database: %v", err.Error())
	}

	// Test with a valid team & valid channel name in apply mode.
	userData := &UserImportData{
		Username: &username,
		Email:    ptrStr(model.NewId() + "@example.com"),
		Teams: &[]UserTeamImportData{
			{
				Name:  &team.Name,
				Roles: ptrStr("team_user team_admin"),
				Channels: &[]UserChannelImportData{
					{
						Name:  &channel.Name,
						Roles: ptrStr("channel_admin channel_user"),
					},
				},
			},
		},
	}
	if err := th.App.ImportUser(userData, false); err != nil {
		t.Fatalf("Should have succeeded.")
	}

	user, err = th.App.GetUserByUsername(*userData.Username)
	if err != nil {
		t.Fatalf("Failed to get user from database.")
	}

	teamMember, err := th.App.GetTeamMember(team.Id, user.Id)
	if err != nil {
		t.Fatalf("Failed to get the team member")
	}
	assert.True(t, teamMember.SchemeAdmin)
	assert.True(t, teamMember.SchemeUser)
	assert.Equal(t, "", teamMember.ExplicitRoles)

	channelMember, err := th.App.GetChannelMember(channel.Id, user.Id)
	if err != nil {
		t.Fatalf("Failed to get the channel member")
	}
	assert.True(t, channelMember.SchemeAdmin)
	assert.True(t, channelMember.SchemeUser)
	assert.Equal(t, "", channelMember.ExplicitRoles)

}

func AssertAllPostsCount(t *testing.T, a *App, initialCount int64, change int64, teamName string) {
	if result := <-a.Srv.Store.Post().AnalyticsPostCount(teamName, false, false); result.Err != nil {
		t.Fatal(result.Err)
	} else {
		if initialCount+change != result.Data.(int64) {
			debug.PrintStack()
			t.Fatalf("Did not find the expected number of posts.")
		}
	}
}

func TestImportImportPost(t *testing.T) {
	th := Setup()
	defer th.TearDown()

	// Create a Team.
	teamName := model.NewId()
	th.App.ImportTeam(&TeamImportData{
		Name:        &teamName,
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
	}, false)
	team, err := th.App.GetTeamByName(teamName)
	if err != nil {
		t.Fatalf("Failed to get team from database.")
	}

	// Create a Channel.
	channelName := model.NewId()
	th.App.ImportChannel(&ChannelImportData{
		Team:        &teamName,
		Name:        &channelName,
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
	}, false)
	channel, err := th.App.GetChannelByName(channelName, team.Id)
	if err != nil {
		t.Fatalf("Failed to get channel from database.")
	}

	// Create a user.
	username := model.NewId()
	th.App.ImportUser(&UserImportData{
		Username: &username,
		Email:    ptrStr(model.NewId() + "@example.com"),
	}, false)
	user, err := th.App.GetUserByUsername(username)
	if err != nil {
		t.Fatalf("Failed to get user from database.")
	}

	// Count the number of posts in the testing team.
	var initialPostCount int64
	if result := <-th.App.Srv.Store.Post().AnalyticsPostCount(team.Id, false, false); result.Err != nil {
		t.Fatal(result.Err)
	} else {
		initialPostCount = result.Data.(int64)
	}

	// Try adding an invalid post in dry run mode.
	data := &PostImportData{
		Team:    &teamName,
		Channel: &channelName,
		User:    &username,
	}
	if err := th.App.ImportPost(data, true); err == nil {
		t.Fatalf("Expected error.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 0, team.Id)

	// Try adding a valid post in dry run mode.
	data = &PostImportData{
		Team:     &teamName,
		Channel:  &channelName,
		User:     &username,
		Message:  ptrStr("Hello"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportPost(data, true); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 0, team.Id)

	// Try adding an invalid post in apply mode.
	data = &PostImportData{
		Team:     &teamName,
		Channel:  &channelName,
		User:     &username,
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportPost(data, false); err == nil {
		t.Fatalf("Expected error.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 0, team.Id)

	// Try adding a valid post with invalid team in apply mode.
	data = &PostImportData{
		Team:     ptrStr(model.NewId()),
		Channel:  &channelName,
		User:     &username,
		Message:  ptrStr("Message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportPost(data, false); err == nil {
		t.Fatalf("Expected error.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 0, team.Id)

	// Try adding a valid post with invalid channel in apply mode.
	data = &PostImportData{
		Team:     &teamName,
		Channel:  ptrStr(model.NewId()),
		User:     &username,
		Message:  ptrStr("Message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportPost(data, false); err == nil {
		t.Fatalf("Expected error.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 0, team.Id)

	// Try adding a valid post with invalid user in apply mode.
	data = &PostImportData{
		Team:     &teamName,
		Channel:  &channelName,
		User:     ptrStr(model.NewId()),
		Message:  ptrStr("Message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportPost(data, false); err == nil {
		t.Fatalf("Expected error.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 0, team.Id)

	// Try adding a valid post in apply mode.
	time := model.GetMillis()
	data = &PostImportData{
		Team:     &teamName,
		Channel:  &channelName,
		User:     &username,
		Message:  ptrStr("Message"),
		CreateAt: &time,
	}
	if err := th.App.ImportPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 1, team.Id)

	// Check the post values.
	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(channel.Id, time); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		if post.Message != *data.Message || post.CreateAt != *data.CreateAt || post.UserId != user.Id {
			t.Fatal("Post properties not as expected")
		}
	}

	// Update the post.
	data = &PostImportData{
		Team:     &teamName,
		Channel:  &channelName,
		User:     &username,
		Message:  ptrStr("Message"),
		CreateAt: &time,
	}
	if err := th.App.ImportPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 1, team.Id)

	// Check the post values.
	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(channel.Id, time); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		if post.Message != *data.Message || post.CreateAt != *data.CreateAt || post.UserId != user.Id {
			t.Fatal("Post properties not as expected")
		}
	}

	// Save the post with a different time.
	newTime := time + 1
	data = &PostImportData{
		Team:     &teamName,
		Channel:  &channelName,
		User:     &username,
		Message:  ptrStr("Message"),
		CreateAt: &newTime,
	}
	if err := th.App.ImportPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 2, team.Id)

	// Save the post with a different message.
	data = &PostImportData{
		Team:     &teamName,
		Channel:  &channelName,
		User:     &username,
		Message:  ptrStr("Message 2"),
		CreateAt: &time,
	}
	if err := th.App.ImportPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 3, team.Id)

	// Test with hashtags
	hashtagTime := time + 2
	data = &PostImportData{
		Team:     &teamName,
		Channel:  &channelName,
		User:     &username,
		Message:  ptrStr("Message 2 #hashtagmashupcity"),
		CreateAt: &hashtagTime,
	}
	if err := th.App.ImportPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 4, team.Id)

	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(channel.Id, hashtagTime); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		if post.Message != *data.Message || post.CreateAt != *data.CreateAt || post.UserId != user.Id {
			t.Fatal("Post properties not as expected")
		}
		if post.Hashtags != "#hashtagmashupcity" {
			t.Fatalf("Hashtags not as expected: %s", post.Hashtags)
		}
	}

	// Post with flags.
	username2 := model.NewId()
	th.App.ImportUser(&UserImportData{
		Username: &username2,
		Email:    ptrStr(model.NewId() + "@example.com"),
	}, false)
	user2, err := th.App.GetUserByUsername(username2)
	if err != nil {
		t.Fatalf("Failed to get user from database.")
	}

	flagsTime := hashtagTime + 1
	data = &PostImportData{
		Team:     &teamName,
		Channel:  &channelName,
		User:     &username,
		Message:  ptrStr("Message with Favorites"),
		CreateAt: &flagsTime,
		FlaggedBy: &[]string{
			username,
			username2,
		},
	}
	if err := th.App.ImportPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 5, team.Id)

	// Check the post values.
	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(channel.Id, flagsTime); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		if post.Message != *data.Message || post.CreateAt != *data.CreateAt || post.UserId != user.Id {
			t.Fatal("Post properties not as expected")
		}

		checkPreference(t, th.App, user.Id, model.PREFERENCE_CATEGORY_FLAGGED_POST, post.Id, "true")
		checkPreference(t, th.App, user2.Id, model.PREFERENCE_CATEGORY_FLAGGED_POST, post.Id, "true")
	}

	// Post with reaction.
	reactionPostTime := hashtagTime + 2
	reactionTime := hashtagTime + 3
	data = &PostImportData{
		Team:     &teamName,
		Channel:  &channelName,
		User:     &username,
		Message:  ptrStr("Message with reaction"),
		CreateAt: &reactionPostTime,
		Reactions: &[]ReactionImportData{{
			User:      &user2.Username,
			EmojiName: ptrStr("+1"),
			CreateAt:  &reactionTime,
		}},
	}
	if err := th.App.ImportPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 6, team.Id)

	// Check the post values.
	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(channel.Id, reactionPostTime); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		if post.Message != *data.Message || post.CreateAt != *data.CreateAt || post.UserId != user.Id || !post.HasReactions {
			t.Fatal("Post properties not as expected")
		}

		if result := <-th.App.Srv.Store.Reaction().GetForPost(post.Id, false); result.Err != nil {
			t.Fatal("Can't get reaction")
		} else if len(result.Data.([]*model.Reaction)) != 1 {
			t.Fatal("Invalid number of reactions")
		}
	}

	// Post with reply.
	replyPostTime := hashtagTime + 4
	replyTime := hashtagTime + 5
	data = &PostImportData{
		Team:     &teamName,
		Channel:  &channelName,
		User:     &username,
		Message:  ptrStr("Message with reply"),
		CreateAt: &replyPostTime,
		Replies: &[]ReplyImportData{{
			User:     &user2.Username,
			Message:  ptrStr("Message reply"),
			CreateAt: &replyTime,
		}},
	}
	if err := th.App.ImportPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 8, team.Id)

	// Check the post values.
	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(channel.Id, replyPostTime); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		if post.Message != *data.Message || post.CreateAt != *data.CreateAt || post.UserId != user.Id {
			t.Fatal("Post properties not as expected")
		}

		// Check the reply values.
		if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(channel.Id, replyTime); result.Err != nil {
			t.Fatal(result.Err.Error())
		} else {
			replies := result.Data.([]*model.Post)
			if len(replies) != 1 {
				t.Fatal("Unexpected number of posts found.")
			}
			reply := replies[0]
			if reply.Message != *(*data.Replies)[0].Message || reply.CreateAt != *(*data.Replies)[0].CreateAt || reply.UserId != user2.Id {
				t.Fatal("Post properties not as expected")
			}

			if reply.RootId != post.Id {
				t.Fatal("Unexpected reply RootId")
			}
		}
	}
}

func TestImportImportDirectChannel(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	// Check how many channels are in the database.
	var directChannelCount int64
	if r := <-th.App.Srv.Store.Channel().AnalyticsTypeCount("", model.CHANNEL_DIRECT); r.Err == nil {
		directChannelCount = r.Data.(int64)
	} else {
		t.Fatalf("Failed to get direct channel count.")
	}

	var groupChannelCount int64
	if r := <-th.App.Srv.Store.Channel().AnalyticsTypeCount("", model.CHANNEL_GROUP); r.Err == nil {
		groupChannelCount = r.Data.(int64)
	} else {
		t.Fatalf("Failed to get group channel count.")
	}

	// Do an invalid channel in dry-run mode.
	data := DirectChannelImportData{
		Members: &[]string{
			model.NewId(),
		},
		Header: ptrStr("Channel Header"),
	}
	if err := th.App.ImportDirectChannel(&data, true); err == nil {
		t.Fatalf("Expected error due to invalid name.")
	}

	// Check that no more channels are in the DB.
	AssertChannelCount(t, th.App, model.CHANNEL_DIRECT, directChannelCount)
	AssertChannelCount(t, th.App, model.CHANNEL_GROUP, groupChannelCount)

	// Do a valid DIRECT channel with a nonexistent member in dry-run mode.
	data.Members = &[]string{
		model.NewId(),
		model.NewId(),
	}
	if err := th.App.ImportDirectChannel(&data, true); err != nil {
		t.Fatalf("Expected success as cannot validate existence of channel members in dry run mode.")
	}

	// Check that no more channels are in the DB.
	AssertChannelCount(t, th.App, model.CHANNEL_DIRECT, directChannelCount)
	AssertChannelCount(t, th.App, model.CHANNEL_GROUP, groupChannelCount)

	// Do a valid GROUP channel with a nonexistent member in dry-run mode.
	data.Members = &[]string{
		model.NewId(),
		model.NewId(),
		model.NewId(),
	}
	if err := th.App.ImportDirectChannel(&data, true); err != nil {
		t.Fatalf("Expected success as cannot validate existence of channel members in dry run mode.")
	}

	// Check that no more channels are in the DB.
	AssertChannelCount(t, th.App, model.CHANNEL_DIRECT, directChannelCount)
	AssertChannelCount(t, th.App, model.CHANNEL_GROUP, groupChannelCount)

	// Do an invalid channel in apply mode.
	data.Members = &[]string{
		model.NewId(),
	}
	if err := th.App.ImportDirectChannel(&data, false); err == nil {
		t.Fatalf("Expected error due to invalid member (apply mode).")
	}

	// Check that no more channels are in the DB.
	AssertChannelCount(t, th.App, model.CHANNEL_DIRECT, directChannelCount)
	AssertChannelCount(t, th.App, model.CHANNEL_GROUP, groupChannelCount)

	// Do a valid DIRECT channel.
	data.Members = &[]string{
		th.BasicUser.Username,
		th.BasicUser2.Username,
	}
	if err := th.App.ImportDirectChannel(&data, false); err != nil {
		t.Fatalf("Expected success: %v", err.Error())
	}

	// Check that one more DIRECT channel is in the DB.
	AssertChannelCount(t, th.App, model.CHANNEL_DIRECT, directChannelCount+1)
	AssertChannelCount(t, th.App, model.CHANNEL_GROUP, groupChannelCount)

	// Do the same DIRECT channel again.
	if err := th.App.ImportDirectChannel(&data, false); err != nil {
		t.Fatalf("Expected success.")
	}

	// Check that no more channels are in the DB.
	AssertChannelCount(t, th.App, model.CHANNEL_DIRECT, directChannelCount+1)
	AssertChannelCount(t, th.App, model.CHANNEL_GROUP, groupChannelCount)

	// Update the channel's HEADER
	data.Header = ptrStr("New Channel Header 2")
	if err := th.App.ImportDirectChannel(&data, false); err != nil {
		t.Fatalf("Expected success.")
	}

	// Check that no more channels are in the DB.
	AssertChannelCount(t, th.App, model.CHANNEL_DIRECT, directChannelCount+1)
	AssertChannelCount(t, th.App, model.CHANNEL_GROUP, groupChannelCount)

	// Get the channel to check that the header was updated.
	if channel, err := th.App.createDirectChannel(th.BasicUser.Id, th.BasicUser2.Id); err == nil || err.Id != store.CHANNEL_EXISTS_ERROR {
		t.Fatal("Should have got store.CHANNEL_EXISTS_ERROR")
	} else {
		if channel.Header != *data.Header {
			t.Fatal("Channel header has not been updated successfully.")
		}
	}

	// Do a GROUP channel with an extra invalid member.
	user3 := th.CreateUser()
	data.Members = &[]string{
		th.BasicUser.Username,
		th.BasicUser2.Username,
		user3.Username,
		model.NewId(),
	}
	if err := th.App.ImportDirectChannel(&data, false); err == nil {
		t.Fatalf("Should have failed due to invalid member in list.")
	}

	// Check that no more channels are in the DB.
	AssertChannelCount(t, th.App, model.CHANNEL_DIRECT, directChannelCount+1)
	AssertChannelCount(t, th.App, model.CHANNEL_GROUP, groupChannelCount)

	// Do a valid GROUP channel.
	data.Members = &[]string{
		th.BasicUser.Username,
		th.BasicUser2.Username,
		user3.Username,
	}
	if err := th.App.ImportDirectChannel(&data, false); err != nil {
		t.Fatalf("Expected success.")
	}

	// Check that one more GROUP channel is in the DB.
	AssertChannelCount(t, th.App, model.CHANNEL_DIRECT, directChannelCount+1)
	AssertChannelCount(t, th.App, model.CHANNEL_GROUP, groupChannelCount+1)

	// Do the same DIRECT channel again.
	if err := th.App.ImportDirectChannel(&data, false); err != nil {
		t.Fatalf("Expected success.")
	}

	// Check that no more channels are in the DB.
	AssertChannelCount(t, th.App, model.CHANNEL_DIRECT, directChannelCount+1)
	AssertChannelCount(t, th.App, model.CHANNEL_GROUP, groupChannelCount+1)

	// Update the channel's HEADER
	data.Header = ptrStr("New Channel Header 3")
	if err := th.App.ImportDirectChannel(&data, false); err != nil {
		t.Fatalf("Expected success.")
	}

	// Check that no more channels are in the DB.
	AssertChannelCount(t, th.App, model.CHANNEL_DIRECT, directChannelCount+1)
	AssertChannelCount(t, th.App, model.CHANNEL_GROUP, groupChannelCount+1)

	// Get the channel to check that the header was updated.
	userIds := []string{
		th.BasicUser.Id,
		th.BasicUser2.Id,
		user3.Id,
	}
	if channel, err := th.App.createGroupChannel(userIds, th.BasicUser.Id); err.Id != store.CHANNEL_EXISTS_ERROR {
		t.Fatal("Should have got store.CHANNEL_EXISTS_ERROR")
	} else {
		if channel.Header != *data.Header {
			t.Fatal("Channel header has not been updated successfully.")
		}
	}

	// Import a channel with some favorites.
	data.Members = &[]string{
		th.BasicUser.Username,
		th.BasicUser2.Username,
	}
	data.FavoritedBy = &[]string{
		th.BasicUser.Username,
		th.BasicUser2.Username,
	}
	if err := th.App.ImportDirectChannel(&data, false); err != nil {
		t.Fatal(err)
	}

	if channel, err := th.App.createDirectChannel(th.BasicUser.Id, th.BasicUser2.Id); err == nil || err.Id != store.CHANNEL_EXISTS_ERROR {
		t.Fatal("Should have got store.CHANNEL_EXISTS_ERROR")
	} else {
		checkPreference(t, th.App, th.BasicUser.Id, model.PREFERENCE_CATEGORY_FAVORITE_CHANNEL, channel.Id, "true")
		checkPreference(t, th.App, th.BasicUser2.Id, model.PREFERENCE_CATEGORY_FAVORITE_CHANNEL, channel.Id, "true")
	}
}

func AssertChannelCount(t *testing.T, a *App, channelType string, expectedCount int64) {
	if r := <-a.Srv.Store.Channel().AnalyticsTypeCount("", channelType); r.Err == nil {
		count := r.Data.(int64)
		if count != expectedCount {
			debug.PrintStack()
			t.Fatalf("Channel count of type: %v. Expected: %v, Got: %v", channelType, expectedCount, count)
		}
	} else {
		debug.PrintStack()
		t.Fatalf("Failed to get channel count.")
	}
}

func TestImportImportDirectPost(t *testing.T) {
	th := Setup().InitBasic()
	defer th.TearDown()

	// Create the DIRECT channel.
	channelData := DirectChannelImportData{
		Members: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
		},
	}
	if err := th.App.ImportDirectChannel(&channelData, false); err != nil {
		t.Fatalf("Expected success: %v", err.Error())
	}

	// Get the channel.
	var directChannel *model.Channel
	if channel, err := th.App.createDirectChannel(th.BasicUser.Id, th.BasicUser2.Id); err.Id != store.CHANNEL_EXISTS_ERROR {
		t.Fatal("Should have got store.CHANNEL_EXISTS_ERROR")
	} else {
		directChannel = channel
	}

	// Get the number of posts in the system.
	var initialPostCount int64
	if result := <-th.App.Srv.Store.Post().AnalyticsPostCount("", false, false); result.Err != nil {
		t.Fatal(result.Err)
	} else {
		initialPostCount = result.Data.(int64)
	}

	// Try adding an invalid post in dry run mode.
	data := &DirectPostImportData{
		ChannelMembers: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
		},
		User:     ptrStr(th.BasicUser.Username),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportDirectPost(data, true); err == nil {
		t.Fatalf("Expected error.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 0, "")

	// Try adding a valid post in dry run mode.
	data = &DirectPostImportData{
		ChannelMembers: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
		},
		User:     ptrStr(th.BasicUser.Username),
		Message:  ptrStr("Message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportDirectPost(data, true); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 0, "")

	// Try adding an invalid post in apply mode.
	data = &DirectPostImportData{
		ChannelMembers: &[]string{
			th.BasicUser.Username,
			model.NewId(),
		},
		User:     ptrStr(th.BasicUser.Username),
		Message:  ptrStr("Message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportDirectPost(data, false); err == nil {
		t.Fatalf("Expected error.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 0, "")

	// Try adding a valid post in apply mode.
	data = &DirectPostImportData{
		ChannelMembers: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
		},
		User:     ptrStr(th.BasicUser.Username),
		Message:  ptrStr("Message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportDirectPost(data, false); err != nil {
		t.Fatalf("Expected success: %v", err.Error())
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 1, "")

	// Check the post values.
	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(directChannel.Id, *data.CreateAt); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		if post.Message != *data.Message || post.CreateAt != *data.CreateAt || post.UserId != th.BasicUser.Id {
			t.Fatal("Post properties not as expected")
		}
	}

	// Import the post again.
	if err := th.App.ImportDirectPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 1, "")

	// Check the post values.
	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(directChannel.Id, *data.CreateAt); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		if post.Message != *data.Message || post.CreateAt != *data.CreateAt || post.UserId != th.BasicUser.Id {
			t.Fatal("Post properties not as expected")
		}
	}

	// Save the post with a different time.
	data.CreateAt = ptrInt64(*data.CreateAt + 1)
	if err := th.App.ImportDirectPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 2, "")

	// Save the post with a different message.
	data.Message = ptrStr("Message 2")
	if err := th.App.ImportDirectPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 3, "")

	// Test with hashtags
	data.Message = ptrStr("Message 2 #hashtagmashupcity")
	data.CreateAt = ptrInt64(*data.CreateAt + 1)
	if err := th.App.ImportDirectPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 4, "")

	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(directChannel.Id, *data.CreateAt); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		if post.Message != *data.Message || post.CreateAt != *data.CreateAt || post.UserId != th.BasicUser.Id {
			t.Fatal("Post properties not as expected")
		}
		if post.Hashtags != "#hashtagmashupcity" {
			t.Fatalf("Hashtags not as expected: %s", post.Hashtags)
		}
	}

	// Test with some flags.
	data = &DirectPostImportData{
		ChannelMembers: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
		},
		FlaggedBy: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
		},
		User:     ptrStr(th.BasicUser.Username),
		Message:  ptrStr("Message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}

	if err := th.App.ImportDirectPost(data, false); err != nil {
		t.Fatalf("Expected success: %v", err.Error())
	}

	// Check the post values.
	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(directChannel.Id, *data.CreateAt); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		checkPreference(t, th.App, th.BasicUser.Id, model.PREFERENCE_CATEGORY_FLAGGED_POST, post.Id, "true")
		checkPreference(t, th.App, th.BasicUser2.Id, model.PREFERENCE_CATEGORY_FLAGGED_POST, post.Id, "true")
	}

	// ------------------ Group Channel -------------------------

	// Create the GROUP channel.
	user3 := th.CreateUser()
	channelData = DirectChannelImportData{
		Members: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
			user3.Username,
		},
	}
	if err := th.App.ImportDirectChannel(&channelData, false); err != nil {
		t.Fatalf("Expected success: %v", err.Error())
	}

	// Get the channel.
	var groupChannel *model.Channel
	userIds := []string{
		th.BasicUser.Id,
		th.BasicUser2.Id,
		user3.Id,
	}
	if channel, err := th.App.createGroupChannel(userIds, th.BasicUser.Id); err.Id != store.CHANNEL_EXISTS_ERROR {
		t.Fatal("Should have got store.CHANNEL_EXISTS_ERROR")
	} else {
		groupChannel = channel
	}

	// Get the number of posts in the system.
	if result := <-th.App.Srv.Store.Post().AnalyticsPostCount("", false, false); result.Err != nil {
		t.Fatal(result.Err)
	} else {
		initialPostCount = result.Data.(int64)
	}

	// Try adding an invalid post in dry run mode.
	data = &DirectPostImportData{
		ChannelMembers: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
			user3.Username,
		},
		User:     ptrStr(th.BasicUser.Username),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportDirectPost(data, true); err == nil {
		t.Fatalf("Expected error.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 0, "")

	// Try adding a valid post in dry run mode.
	data = &DirectPostImportData{
		ChannelMembers: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
			user3.Username,
		},
		User:     ptrStr(th.BasicUser.Username),
		Message:  ptrStr("Message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportDirectPost(data, true); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 0, "")

	// Try adding an invalid post in apply mode.
	data = &DirectPostImportData{
		ChannelMembers: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
			user3.Username,
			model.NewId(),
		},
		User:     ptrStr(th.BasicUser.Username),
		Message:  ptrStr("Message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportDirectPost(data, false); err == nil {
		t.Fatalf("Expected error.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 0, "")

	// Try adding a valid post in apply mode.
	data = &DirectPostImportData{
		ChannelMembers: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
			user3.Username,
		},
		User:     ptrStr(th.BasicUser.Username),
		Message:  ptrStr("Message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}
	if err := th.App.ImportDirectPost(data, false); err != nil {
		t.Fatalf("Expected success: %v", err.Error())
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 1, "")

	// Check the post values.
	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(groupChannel.Id, *data.CreateAt); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		if post.Message != *data.Message || post.CreateAt != *data.CreateAt || post.UserId != th.BasicUser.Id {
			t.Fatal("Post properties not as expected")
		}
	}

	// Import the post again.
	if err := th.App.ImportDirectPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 1, "")

	// Check the post values.
	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(groupChannel.Id, *data.CreateAt); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		if post.Message != *data.Message || post.CreateAt != *data.CreateAt || post.UserId != th.BasicUser.Id {
			t.Fatal("Post properties not as expected")
		}
	}

	// Save the post with a different time.
	data.CreateAt = ptrInt64(*data.CreateAt + 1)
	if err := th.App.ImportDirectPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 2, "")

	// Save the post with a different message.
	data.Message = ptrStr("Message 2")
	if err := th.App.ImportDirectPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 3, "")

	// Test with hashtags
	data.Message = ptrStr("Message 2 #hashtagmashupcity")
	data.CreateAt = ptrInt64(*data.CreateAt + 1)
	if err := th.App.ImportDirectPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}
	AssertAllPostsCount(t, th.App, initialPostCount, 4, "")

	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(groupChannel.Id, *data.CreateAt); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		if post.Message != *data.Message || post.CreateAt != *data.CreateAt || post.UserId != th.BasicUser.Id {
			t.Fatal("Post properties not as expected")
		}
		if post.Hashtags != "#hashtagmashupcity" {
			t.Fatalf("Hashtags not as expected: %s", post.Hashtags)
		}
	}

	// Test with some flags.
	data = &DirectPostImportData{
		ChannelMembers: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
			user3.Username,
		},
		FlaggedBy: &[]string{
			th.BasicUser.Username,
			th.BasicUser2.Username,
		},
		User:     ptrStr(th.BasicUser.Username),
		Message:  ptrStr("Message"),
		CreateAt: ptrInt64(model.GetMillis()),
	}

	if err := th.App.ImportDirectPost(data, false); err != nil {
		t.Fatalf("Expected success: %v", err.Error())
	}

	// Check the post values.
	if result := <-th.App.Srv.Store.Post().GetPostsCreatedAt(groupChannel.Id, *data.CreateAt); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		if len(posts) != 1 {
			t.Fatal("Unexpected number of posts found.")
		}
		post := posts[0]
		checkPreference(t, th.App, th.BasicUser.Id, model.PREFERENCE_CATEGORY_FLAGGED_POST, post.Id, "true")
		checkPreference(t, th.App, th.BasicUser2.Id, model.PREFERENCE_CATEGORY_FLAGGED_POST, post.Id, "true")
	}
}

func TestImportImportLine(t *testing.T) {
	th := Setup()
	defer th.TearDown()

	// Try import line with an invalid type.
	line := LineImportData{
		Type: "gibberish",
	}

	if err := th.App.ImportLine(line, false); err == nil {
		t.Fatalf("Expected an error when importing a line with invalid type.")
	}

	// Try import line with team type but nil team.
	line.Type = "team"
	if err := th.App.ImportLine(line, false); err == nil {
		t.Fatalf("Expected an error when importing a line of type team with a nil team.")
	}

	// Try import line with channel type but nil channel.
	line.Type = "channel"
	if err := th.App.ImportLine(line, false); err == nil {
		t.Fatalf("Expected an error when importing a line with type channel with a nil channel.")
	}

	// Try import line with user type but nil user.
	line.Type = "user"
	if err := th.App.ImportLine(line, false); err == nil {
		t.Fatalf("Expected an error when importing a line with type uesr with a nil user.")
	}

	// Try import line with post type but nil post.
	line.Type = "post"
	if err := th.App.ImportLine(line, false); err == nil {
		t.Fatalf("Expected an error when importing a line with type post with a nil post.")
	}

	// Try import line with direct_channel type but nil direct_channel.
	line.Type = "direct_channel"
	if err := th.App.ImportLine(line, false); err == nil {
		t.Fatalf("Expected an error when importing a line with type direct_channel with a nil direct_channel.")
	}

	// Try import line with direct_post type but nil direct_post.
	line.Type = "direct_post"
	if err := th.App.ImportLine(line, false); err == nil {
		t.Fatalf("Expected an error when importing a line with type direct_post with a nil direct_post.")
	}

	// Try import line with scheme type but nil scheme.
	line.Type = "scheme"
	if err := th.App.ImportLine(line, false); err == nil {
		t.Fatalf("Expected an error when importing a line with type scheme with a nil scheme.")
	}
}

func TestImportBulkImport(t *testing.T) {
	th := Setup()
	defer th.TearDown()

	th.App.UpdateConfig(func(cfg *model.Config) { *cfg.ServiceSettings.EnableCustomEmoji = true })

	teamName := model.NewId()
	channelName := model.NewId()
	username := model.NewId()
	username2 := model.NewId()
	username3 := model.NewId()
	emojiName := model.NewId()
	testsDir, _ := utils.FindDir("tests")
	testImage := filepath.Join(testsDir, "test.png")

	// Run bulk import with a valid 1 of everything.
	data1 := `{"type": "version", "version": 1}
{"type": "team", "team": {"type": "O", "display_name": "lskmw2d7a5ao7ppwqh5ljchvr4", "name": "` + teamName + `"}}
{"type": "channel", "channel": {"type": "O", "display_name": "xr6m6udffngark2uekvr3hoeny", "team": "` + teamName + `", "name": "` + channelName + `"}}
{"type": "user", "user": {"username": "` + username + `", "email": "` + username + `@example.com", "teams": [{"name": "` + teamName + `", "channels": [{"name": "` + channelName + `"}]}]}}
{"type": "user", "user": {"username": "` + username2 + `", "email": "` + username2 + `@example.com", "teams": [{"name": "` + teamName + `", "channels": [{"name": "` + channelName + `"}]}]}}
{"type": "user", "user": {"username": "` + username3 + `", "email": "` + username3 + `@example.com", "teams": [{"name": "` + teamName + `", "channels": [{"name": "` + channelName + `"}]}]}}
{"type": "post", "post": {"team": "` + teamName + `", "channel": "` + channelName + `", "user": "` + username + `", "message": "Hello World", "create_at": 123456789012, "attachements":[{"path": "` + testImage + `"}]}}
{"type": "direct_channel", "direct_channel": {"members": ["` + username + `", "` + username2 + `"]}}
{"type": "direct_channel", "direct_channel": {"members": ["` + username + `", "` + username2 + `", "` + username3 + `"]}}
{"type": "direct_post", "direct_post": {"channel_members": ["` + username + `", "` + username2 + `"], "user": "` + username + `", "message": "Hello Direct Channel", "create_at": 123456789013}}
{"type": "direct_post", "direct_post": {"channel_members": ["` + username + `", "` + username2 + `", "` + username3 + `"], "user": "` + username + `", "message": "Hello Group Channel", "create_at": 123456789014}}
{"type": "emoji", "emoji": {"name": "` + emojiName + `", "image": "` + testImage + `"}}`

	if err, line := th.App.BulkImport(strings.NewReader(data1), false, 2); err != nil || line != 0 {
		t.Fatalf("BulkImport should have succeeded: %v, %v", err.Error(), line)
	}

	// Run bulk import using a string that contains a line with invalid json.
	data2 := `{"type": "version", "version": 1`
	if err, line := th.App.BulkImport(strings.NewReader(data2), false, 2); err == nil || line != 1 {
		t.Fatalf("Should have failed due to invalid JSON on line 1.")
	}

	// Run bulk import using valid JSON but missing version line at the start.
	data3 := `{"type": "team", "team": {"type": "O", "display_name": "lskmw2d7a5ao7ppwqh5ljchvr4", "name": "` + teamName + `"}}
{"type": "channel", "channel": {"type": "O", "display_name": "xr6m6udffngark2uekvr3hoeny", "team": "` + teamName + `", "name": "` + channelName + `"}}
{"type": "user", "user": {"username": "kufjgnkxkrhhfgbrip6qxkfsaa", "email": "kufjgnkxkrhhfgbrip6qxkfsaa@example.com"}}
{"type": "user", "user": {"username": "bwshaim6qnc2ne7oqkd5b2s2rq", "email": "bwshaim6qnc2ne7oqkd5b2s2rq@example.com", "teams": [{"name": "` + teamName + `", "channels": [{"name": "` + channelName + `"}]}]}}`
	if err, line := th.App.BulkImport(strings.NewReader(data3), false, 2); err == nil || line != 1 {
		t.Fatalf("Should have failed due to missing version line on line 1.")
	}
}

func TestImportProcessImportDataFileVersionLine(t *testing.T) {
	data := LineImportData{
		Type:    "version",
		Version: ptrInt(1),
	}
	if version, err := processImportDataFileVersionLine(data); err != nil || version != 1 {
		t.Fatalf("Expected no error and version 1.")
	}

	data.Type = "NotVersion"
	if _, err := processImportDataFileVersionLine(data); err == nil {
		t.Fatalf("Expected error on invalid version line.")
	}

	data.Type = "version"
	data.Version = nil
	if _, err := processImportDataFileVersionLine(data); err == nil {
		t.Fatalf("Expected error on invalid version line.")
	}
}

func TestImportValidateEmojiImportData(t *testing.T) {
	data := EmojiImportData{
		Name:  ptrStr("parrot"),
		Image: ptrStr("/path/to/image"),
	}

	err := validateEmojiImportData(&data)
	assert.Nil(t, err, "Validation should succeed")

	*data.Name = "smiley"
	err = validateEmojiImportData(&data)
	assert.NotNil(t, err)

	*data.Name = ""
	err = validateEmojiImportData(&data)
	assert.NotNil(t, err)

	*data.Name = ""
	*data.Image = ""
	err = validateEmojiImportData(&data)
	assert.NotNil(t, err)

	*data.Image = "/path/to/image"
	data.Name = nil
	err = validateEmojiImportData(&data)
	assert.NotNil(t, err)

	data.Name = ptrStr("parrot")
	data.Image = nil
	err = validateEmojiImportData(&data)
	assert.NotNil(t, err)
}

func TestImportImportEmoji(t *testing.T) {
	th := Setup()
	defer th.TearDown()

	th.App.UpdateConfig(func(cfg *model.Config) { *cfg.ServiceSettings.EnableCustomEmoji = true })

	testsDir, _ := utils.FindDir("tests")
	testImage := filepath.Join(testsDir, "test.png")

	data := EmojiImportData{Name: ptrStr(model.NewId())}
	err := th.App.ImportEmoji(&data, true)
	assert.NotNil(t, err, "Invalid emoji should have failed dry run")

	result := <-th.App.Srv.Store.Emoji().GetByName(*data.Name)
	assert.Nil(t, result.Data, "Emoji should not have been imported")

	data.Image = ptrStr(testImage)
	err = th.App.ImportEmoji(&data, true)
	assert.Nil(t, err, "Valid emoji should have passed dry run")

	data = EmojiImportData{Name: ptrStr(model.NewId())}
	err = th.App.ImportEmoji(&data, false)
	assert.NotNil(t, err, "Invalid emoji should have failed apply mode")

	data.Image = ptrStr("non-existent-file")
	err = th.App.ImportEmoji(&data, false)
	assert.NotNil(t, err, "Emoji with bad image file should have failed apply mode")

	data.Image = ptrStr(testImage)
	err = th.App.ImportEmoji(&data, false)
	assert.Nil(t, err, "Valid emoji should have succeeded apply mode")

	result = <-th.App.Srv.Store.Emoji().GetByName(*data.Name)
	assert.NotNil(t, result.Data, "Emoji should have been imported")

	err = th.App.ImportEmoji(&data, false)
	assert.Nil(t, err, "Second run should have succeeded apply mode")
}

func TestImportAttachment(t *testing.T) {
	th := Setup()
	defer th.TearDown()

	testsDir, _ := utils.FindDir("tests")
	testImage := filepath.Join(testsDir, "test.png")
	invalidPath := "some-invalid-path"

	userId := model.NewId()
	data := AttachmentImportData{Path: &testImage}
	_, err := th.App.ImportAttachment(&data, &model.Post{UserId: userId, ChannelId: "some-channel"}, "some-team", true)
	assert.Nil(t, err, "sample run without errors")

	attachments := GetAttachments(userId, th, t)
	assert.Equal(t, len(attachments), 1)

	data = AttachmentImportData{Path: &invalidPath}
	_, err = th.App.ImportAttachment(&data, &model.Post{UserId: model.NewId(), ChannelId: "some-channel"}, "some-team", true)
	assert.NotNil(t, err, "should have failed when opening the file")
	assert.Equal(t, err.Id, "app.import.attachment.bad_file.error")
}

func TestImportPostAndRepliesWithAttachments(t *testing.T) {

	th := Setup()
	defer th.TearDown()

	// Create a Team.
	teamName := model.NewId()
	th.App.ImportTeam(&TeamImportData{
		Name:        &teamName,
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
	}, false)
	team, err := th.App.GetTeamByName(teamName)
	if err != nil {
		t.Fatalf("Failed to get team from database.")
	}

	// Create a Channel.
	channelName := model.NewId()
	th.App.ImportChannel(&ChannelImportData{
		Team:        &teamName,
		Name:        &channelName,
		DisplayName: ptrStr("Display Name"),
		Type:        ptrStr("O"),
	}, false)
	_, err = th.App.GetChannelByName(channelName, team.Id)
	if err != nil {
		t.Fatalf("Failed to get channel from database.")
	}

	// Create a user3.
	username := model.NewId()
	th.App.ImportUser(&UserImportData{
		Username: &username,
		Email:    ptrStr(model.NewId() + "@example.com"),
	}, false)
	user3, err := th.App.GetUserByUsername(username)
	if err != nil {
		t.Fatalf("Failed to get user3 from database.")
	}

	username2 := model.NewId()
	th.App.ImportUser(&UserImportData{
		Username: &username2,
		Email:    ptrStr(model.NewId() + "@example.com"),
	}, false)
	user4, err := th.App.GetUserByUsername(username2)
	if err != nil {
		t.Fatalf("Failed to get user3 from database.")
	}


	// Post with attachments.
	time := model.GetMillis()
	attachmentsPostTime := time
	attachmentsReplyTime := time + 1
	testsDir, _ := utils.FindDir("tests")
	testImage := filepath.Join(testsDir, "test.png")
	testMarkDown := filepath.Join(testsDir, "test-attachments.md")
	data := &PostImportData{
		Team:     &teamName,
		Channel:  &channelName,
		User:     &username,
		Message:  ptrStr("Message with reply"),
		CreateAt: &attachmentsPostTime,
		Attachments: &[]AttachmentImportData{{Path: &testImage}, {Path: &testMarkDown}},
		Replies: &[]ReplyImportData{{
			User:     &user4.Username,
			Message:  ptrStr("Message reply"),
			CreateAt: &attachmentsReplyTime,
			Attachments: &[]AttachmentImportData{{Path: &testImage}},
		}},
	}

	if err := th.App.ImportPost(data, false); err != nil {
		t.Fatalf("Expected success.")
	}

	attachments := GetAttachments(user3.Id, th, t)
	assert.Equal(t, len(attachments), 2)
	assert.Contains(t, attachments[0].Path, team.Id)
	assert.Contains(t, attachments[1].Path, team.Id)
	AssertFileIdsInPost(attachments, th, t)

	attachments = GetAttachments(user4.Id, th, t)
	assert.Equal(t, len(attachments), 1)
	assert.Contains(t, attachments[0].Path, team.Id)
	AssertFileIdsInPost(attachments, th, t)

	// Reply with Attachments in Direct Post

	// Create direct post users.

	username3 := model.NewId()
	th.App.ImportUser(&UserImportData{
		Username: &username3,
		Email:    ptrStr(model.NewId() + "@example.com"),
	}, false)
	user3, err = th.App.GetUserByUsername(username3)
	if err != nil {
		t.Fatalf("Failed to get user3 from database.")
	}

	username4 := model.NewId()
	th.App.ImportUser(&UserImportData{
		Username: &username4,
		Email:    ptrStr(model.NewId() + "@example.com"),
	}, false)

	user4, err = th.App.GetUserByUsername(username4)
	if err != nil {
		t.Fatalf("Failed to get user3 from database.")
	}

	directImportData := &DirectPostImportData{
		ChannelMembers: &[]string{
			user3.Username,
			user4.Username,
		},
		User:     &user3.Username,
		Message:  ptrStr("Message with Replies"),
		CreateAt: ptrInt64(model.GetMillis()),
		Replies:  &[]ReplyImportData{{
			User:     &user4.Username,
			Message:  ptrStr("Message reply with attachment"),
			CreateAt: ptrInt64(model.GetMillis()),
			Attachments: &[]AttachmentImportData{{Path: &testImage}},
		}},
	}

	if err := th.App.ImportDirectPost(directImportData, false); err != nil {
		t.Fatalf("Expected success.")
	}

	attachments = GetAttachments(user4.Id, th, t)
	assert.Equal(t, len(attachments), 1)
	assert.Contains(t, attachments[0].Path, "noteam")
	AssertFileIdsInPost(attachments, th, t)

}


func GetAttachments(userId string, th *TestHelper, t *testing.T) []*model.FileInfo {
	if result := <-th.App.Srv.Store.FileInfo().GetForUser(userId); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		return result.Data.([]*model.FileInfo)
	}
	return nil
}

func AssertFileIdsInPost(files []*model.FileInfo, th *TestHelper, t *testing.T) {
	postId := files[0].PostId
	assert.NotNil(t, postId)

	if result := <-th.App.Srv.Store.Post().GetPostsByIds([]string{postId}); result.Err != nil {
		t.Fatal(result.Err.Error())
	} else {
		posts := result.Data.([]*model.Post)
		assert.Equal(t, len(posts), 1)
		for _, file := range files {
			assert.Contains(t, posts[0].FileIds, file.Id)
		}
	}
}
