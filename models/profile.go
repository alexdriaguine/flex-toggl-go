package models

import "time"

// Profile Type from toggl API response
type Profile struct {
	Since int `json:"since"`
	Data  struct {
		ID                    int       `json:"id"`
		APIToken              string    `json:"api_token"`
		DefaultWid            int       `json:"default_wid"`
		Email                 string    `json:"email"`
		Fullname              string    `json:"fullname"`
		JqueryTimeofdayFormat string    `json:"jquery_timeofday_format"`
		JqueryDateFormat      string    `json:"jquery_date_format"`
		TimeofdayFormat       string    `json:"timeofday_format"`
		DateFormat            string    `json:"date_format"`
		StoreStartAndStopTime bool      `json:"store_start_and_stop_time"`
		BeginningOfWeek       int       `json:"beginning_of_week"`
		Language              string    `json:"language"`
		ImageURL              string    `json:"image_url"`
		SidebarPiechart       bool      `json:"sidebar_piechart"`
		At                    time.Time `json:"at"`
		CreatedAt             time.Time `json:"created_at"`
		Retention             int       `json:"retention"`
		RecordTimeline        bool      `json:"record_timeline"`
		RenderTimeline        bool      `json:"render_timeline"`
		TimelineEnabled       bool      `json:"timeline_enabled"`
		TimelineExperiment    bool      `json:"timeline_experiment"`
		NewBlogPost           struct {
			Title    string    `json:"title"`
			URL      string    `json:"url"`
			Category string    `json:"category"`
			PubDate  time.Time `json:"pub_date"`
		} `json:"new_blog_post"`
		ShouldUpgrade          bool   `json:"should_upgrade"`
		AchievementsEnabled    bool   `json:"achievements_enabled"`
		Timezone               string `json:"timezone"`
		OpenidEnabled          bool   `json:"openid_enabled"`
		SendProductEmails      bool   `json:"send_product_emails"`
		SendWeeklyReport       bool   `json:"send_weekly_report"`
		SendTimerNotifications bool   `json:"send_timer_notifications"`
		LastBlogEntry          string `json:"last_blog_entry"`
		Invitation             struct {
		} `json:"invitation"`
		Workspaces []struct {
			ID                          int       `json:"id"`
			Name                        string    `json:"name"`
			Profile                     int       `json:"profile"`
			Premium                     bool      `json:"premium"`
			Admin                       bool      `json:"admin"`
			DefaultHourlyRate           int       `json:"default_hourly_rate"`
			DefaultCurrency             string    `json:"default_currency"`
			OnlyAdminsMayCreateProjects bool      `json:"only_admins_may_create_projects"`
			OnlyAdminsSeeBillableRates  bool      `json:"only_admins_see_billable_rates"`
			OnlyAdminsSeeTeamDashboard  bool      `json:"only_admins_see_team_dashboard"`
			ProjectsBillableByDefault   bool      `json:"projects_billable_by_default"`
			Rounding                    int       `json:"rounding"`
			RoundingMinutes             int       `json:"rounding_minutes"`
			At                          time.Time `json:"at"`
			LogoURL                     string    `json:"logo_url"`
			IcalURL                     string    `json:"ical_url"`
			IcalEnabled                 bool      `json:"ical_enabled"`
		} `json:"workspaces"`
		DurationFormat string `json:"duration_format"`
		Obm            struct {
			Included bool   `json:"included"`
			Nr       int    `json:"nr"`
			Actions  string `json:"actions"`
		} `json:"obm"`
	} `json:"data"`
}
