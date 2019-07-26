package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func checkValidName(t *testing.T, name string, want bool) {
	got := ValidName(name)
	if got != want {
		t.Errorf("checking name %s, wanted %t but got %t",
			name, want, got)
	}
}

func TestValidName(t *testing.T) {
	checkValidName(t, "myname", true)
	checkValidName(t, "my name", true)
	checkValidName(t, "00112", true)
	checkValidName(t, "My_name 0001-0002", true)
	checkValidName(t, "Hunna Stoll Go", true)
	checkValidName(t, "Deusche telecom AG", true)
	checkValidName(t, "Sonoral S.A.", true)
	checkValidName(t, "UFGT Inc.", true)
	checkValidName(t, "Atlantic, Inc.", true)
	checkValidName(t, "Pillimo Go!", true)
	checkValidName(t, "", false)
	checkValidName(t, " name", false)
	checkValidName(t, "-name", false)
	checkValidName(t, "a;sldfj", false)
	checkValidName(t, "$fadf", false)
}

func checkValidIp(t *testing.T, ip []byte, want bool) {
	got := ValidIp(ip)
	if got != want {
		t.Errorf("checking %x, wanted %t but got %t",
			ip, want, got)
	}
}

func TestValidIp(t *testing.T) {
	checkValidIp(t, []byte{1, 2, 3, 4}, true)
	checkValidIp(t, []byte{1, 2, 3, 4, 5}, false)
	checkValidIp(t, []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16}, true)
	checkValidIp(t, []byte{1, 2, 3, 4, 5}, false)
	checkValidIp(t, nil, false)
}

func TestValidLDAPName(t *testing.T) {
	checkValidLDAPName(t, "myname", true)
	checkValidLDAPName(t, "my name", true)
	checkValidLDAPName(t, "00112", true)
	checkValidLDAPName(t, "My_name 0001-0002", true)
	checkValidLDAPName(t, "Hunna Stoll Go", true)
	checkValidLDAPName(t, "Deusche telecom AG", true)
	checkValidLDAPName(t, "Sonoral S.A.", true)
	checkValidLDAPName(t, "UFGT Inc.", true)
	checkValidLDAPName(t, "Atlantic, Inc.", true)
	checkValidLDAPName(t, "Pillimo Go!", true)
	checkValidLDAPName(t, "", false)
	checkValidLDAPName(t, " name", false)
	checkValidLDAPName(t, "name ", false)
	checkValidLDAPName(t, "name\\a", false)
	checkValidLDAPName(t, "name#a", false)
	checkValidLDAPName(t, "name+a", false)
	checkValidLDAPName(t, "name<a", false)
	checkValidLDAPName(t, "name>a", false)
	checkValidLDAPName(t, "name;a", false)
	checkValidLDAPName(t, "name\"a", false)

	name := EscapeLDAPName("foo, Inc.")
	require.Equal(t, "foo, Inc.", UnescapeLDAPName(name))

	user := EscapeLDAPName("jon,user")
	org := EscapeLDAPName("foo, Inc.")
	split := strings.Split(user+","+org, ",")
	require.Equal(t, "jon,user", UnescapeLDAPName(split[0]))
	require.Equal(t, "foo, Inc.", UnescapeLDAPName(split[1]))

	gname := GitlabGroupSanitize("atlantic, inc.")
	require.Equal(t, "atlantic--inc", gname)
}

func checkValidLDAPName(t *testing.T, name string, want bool) {
	got := ValidLDAPName(name)
	if got != want {
		t.Errorf("checking name %s, wanted %t but got %t",
			name, want, got)
	}
}

func TestValidOrgName(t *testing.T) {
	var err error

	err = ValidOrgName("orgname_123.dev")
	require.Nil(t, err, "valid org name")
	err = ValidOrgName(".orgname_123.dev")
	require.NotNil(t, err, "invalid org name")
	err = ValidOrgName("-orgname_123.dev")
	require.NotNil(t, err, "invalid org name")
	err = ValidOrgName("orgname_123.dev.")
	require.NotNil(t, err, "invalid org name")
	err = ValidOrgName("orgname_123$dev")
	require.NotNil(t, err, "invalid org name")
	err = ValidOrgName("orgname_123dev-cache")
	require.NotNil(t, err, "invalid org name")
	err = ValidOrgName("orgname_123dev.")
	require.NotNil(t, err, "invalid org name")
	err = ValidOrgName("orgname_123dev.git")
	require.NotNil(t, err, "invalid org name")
	err = ValidOrgName("orgname_123dev.atom")
	require.NotNil(t, err, "invalid org name")
	err = ValidOrgName("orgname_123dev test")
	require.NotNil(t, err, "invalid org name")
	err = ValidOrgName("orgname_123dev,test")
	require.NotNil(t, err, "invalid org name")
}
