package chrome

type ProfileData struct {
	ActiveTime                           float64 `json:"active_time"`
	AvatarIcon                           string  `json:"avatar_icon"`
	BackgroundApps                       bool    `json:"background_apps"`
	DefaultAvatarFillColor               int     `json:"default_avatar_fill_color"`
	DefaultAvatarStrokeColor             int     `json:"default_avatar_stroke_color"`
	FirstAccountNameHash                 int     `json:"first_account_name_hash"`
	ForceSigninProfileLocked             bool    `json:"force_signin_profile_locked"`
	GaiaGivenName                        string  `json:"gaia_given_name"`
	GaiaId                               string  `json:"gaia_id"`
	GaiaName                             string  `json:"gaia_name"`
	GaiaPictureFileName                  string  `json:"gaia_picture_file_name"`
	HostedDomain                         string  `json:"hosted_domain"`
	IsConsentedPrimaryAccount            bool    `json:"is_consented_primary_account"`
	IsEphemeral                          bool    `json:"is_ephemeral"`
	IsUsingDefaultAvatar                 bool    `json:"is_using_default_avatar"`
	IsUsingDefaultName                   bool    `json:"is_using_default_name"`
	LastDownloadedGaiaPictureUrlWithSize string  `json:"last_downloaded_gaia_picture_url_with_size"`
	ManagedUserId                        string  `json:"managed_user_id"`
	MetricsBucketIndex                   int     `json:"metrics_bucket_index"`
	Name                                 string  `json:"name"`
	ProfileColorSeed                     int     `json:"profile_color_seed"`
	ProfileHighlightColor                int     `json:"profile_highlight_color"`
	SigninWithCredentialProvider         bool    `json:"signin.with_credential_provider"`
	UserAcceptedAccountManagement        bool    `json:"user_accepted_account_management"`
	UserName                             string  `json:"user_name"`
}

type LocalState struct {
	Profile struct {
		InfoCache map[string]ProfileData `json:"info_cache"`
	} `json:"profile"`
}
