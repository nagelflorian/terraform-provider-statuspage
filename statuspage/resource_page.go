package statuspage

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	statuspageAPI "github.com/nagelflorian/statuspage-go"
	log "github.com/sirupsen/logrus"
)

func resourcePageCreate(d *schema.ResourceData, metaRaw interface{}) error {
	// XXX the public Statuspage API is missing an endpoint to create a new page
	return resourcePageRead(d, metaRaw)
}
func resourcePageRead(d *schema.ResourceData, metaRaw interface{}) error {
	log.WithFields(log.Fields{"id": d.Id()}).Debug("Read resource: Incident")

	meta := metaRaw.(*Meta)

	pageID := d.Get("page_id").(string)
	page, err := meta.client.Page.GetPage(context.TODO(), pageID)
	if err != nil {
		return err
	}

	d.SetId(page.ID)
	d.Set("page_id", page.ID)
	d.Set("name", page.Name)
	d.Set("domain", page.Domain)
	d.Set("subdomain", page.Subdomain)
	d.Set("url", page.URL)
	d.Set("branding", page.Branding)
	d.Set("css_body_background_color", page.CSSBodyBackgroundColor)
	d.Set("css_font_color", page.CSSFontColor)
	d.Set("css_light_font_color", page.CSSLightFontColor)
	d.Set("css_greens", page.CSSGreens)
	d.Set("css_yellows", page.CSSYellows)
	d.Set("css_oranges", page.CSSOranges)
	d.Set("css_reds", page.CSSReds)
	d.Set("css_blues", page.CSSBlues)
	d.Set("css_border_color", page.CSSBorderColor)
	d.Set("css_graph_color", page.CSSGraphColor)
	d.Set("css_link_color", page.CSSLinkColor)
	d.Set("hidden_from_search", page.HiddenFromSearch)
	d.Set("viewers_must_be_team_members", page.ViewersMustBeTeamMembers)
	d.Set("allow_page_subscribers", page.AllowPageSubscribers)
	d.Set("allow_incident_subscribers", page.AllowIncidentSubscribers)
	d.Set("allow_email_subscribers", page.AllowEmailSubscribers)
	d.Set("allow_sms_subscribers", page.AllowSmsSubscribers)
	d.Set("allow_rss_atom_feeds", page.AllowRssAtomFeeds)
	d.Set("allow_webhook_subscribers", page.AllowWebhookSubscribers)
	d.Set("notifications_from_email", page.NotificationsFromEmail)
	d.Set("time_zone", page.TimeZone)
	d.Set("notifications_email_footer", page.NotificationsEmailFooter)

	return nil
}
func resourcePageDelete(d *schema.ResourceData, metaRaw interface{}) error {
	// XXX the public Statuspage API is missing an endpoint to delete a new page
	d.SetId("")
	return nil
}
func resourcePageUpdate(d *schema.ResourceData, metaRaw interface{}) error {
	log.WithFields(log.Fields{"id": d.Id()}).Debug("Update resource: Page")

	meta := metaRaw.(*Meta)

	pageID := d.Id()
	params := statuspageAPI.UpdatePageParams{
		Name:                     d.Get("name").(string),
		Domain:                   d.Get("domain").(string),
		Subdomain:                d.Get("subdomain").(string),
		URL:                      d.Get("url").(string),
		Branding:                 d.Get("branding").(string),
		CSSBodyBackgroundColor:   d.Get("css_body_background_color").(string),
		CSSFontColor:             d.Get("css_font_color").(string),
		CSSLightFontColor:        d.Get("css_light_font_color").(string),
		CSSGreens:                d.Get("css_greens").(string),
		CSSYellows:               d.Get("css_yellows").(string),
		CSSOranges:               d.Get("css_oranges").(string),
		CSSReds:                  d.Get("css_reds").(string),
		CSSBlues:                 d.Get("css_blues").(string),
		CSSBorderColor:           d.Get("css_border_color").(string),
		CSSGraphColor:            d.Get("css_graph_color").(string),
		CSSLinkColor:             d.Get("css_link_color").(string),
		HiddenFromSearch:         d.Get("hidden_from_search").(*bool),
		ViewersMustBeTeamMembers: d.Get("viewers_must_be_team_members").(*bool),
		AllowPageSubscribers:     d.Get("allow_page_subscribers").(*bool),
		AllowIncidentSubscribers: d.Get("allow_incident_subscribers").(*bool),
		AllowEmailSubscribers:    d.Get("allow_email_subscribers").(*bool),
		AllowSmsSubscribers:      d.Get("allow_sms_subscribers").(*bool),
		AllowRssAtomFeeds:        d.Get("allow_rss_atom_feeds").(*bool),
		AllowWebhookSubscribers:  d.Get("allow_webhook_subscribers").(*bool),
		NotificationsFromEmail:   d.Get("notifications_from_email").(string),
		TimeZone:                 d.Get("time_zone").(string),
		NotificationsEmailFooter: d.Get("notifications_email_footer").(string),
	}

	_, err := meta.client.Page.UpdatePage(context.TODO(), pageID, params)
	if err != nil {
		return fmt.Errorf("Failed to update page %q: %s", pageID, err)
	}

	return resourcePageRead(d, metaRaw)
}

func resourcePage() *schema.Resource {
	return &schema.Resource{
		Create: resourcePageCreate,
		Read:   resourcePageRead,
		Delete: resourcePageDelete,
		Update: resourcePageUpdate,
		Schema: map[string]*schema.Schema{
			"page_id":                      {Type: schema.TypeString, Required: true},
			"name":                         {Type: schema.TypeString, Optional: true},
			"domain":                       {Type: schema.TypeString, Optional: true},
			"subdomain":                    {Type: schema.TypeString, Optional: true},
			"url":                          {Type: schema.TypeString, Optional: true},
			"branding":                     {Type: schema.TypeString, Optional: true},
			"css_body_background_color":    {Type: schema.TypeString, Optional: true},
			"css_font_color":               {Type: schema.TypeString, Optional: true},
			"css_light_font_color":         {Type: schema.TypeString, Optional: true},
			"css_greens":                   {Type: schema.TypeString, Optional: true},
			"css_yellows":                  {Type: schema.TypeString, Optional: true},
			"css_oranges":                  {Type: schema.TypeString, Optional: true},
			"css_reds":                     {Type: schema.TypeString, Optional: true},
			"css_blues":                    {Type: schema.TypeString, Optional: true},
			"css_border_color":             {Type: schema.TypeString, Optional: true},
			"css_graph_color":              {Type: schema.TypeString, Optional: true},
			"css_link_color":               {Type: schema.TypeString, Optional: true},
			"hidden_from_search":           {Type: schema.TypeBool, Optional: true},
			"viewers_must_be_team_members": {Type: schema.TypeBool, Optional: true},
			"allow_page_subscribers":       {Type: schema.TypeBool, Optional: true},
			"allow_incident_subscribers":   {Type: schema.TypeBool, Optional: true},
			"allow_email_subscribers":      {Type: schema.TypeBool, Optional: true},
			"allow_sms_subscribers":        {Type: schema.TypeBool, Optional: true},
			"allow_rss_atom_feeds":         {Type: schema.TypeBool, Optional: true},
			"allow_webhook_subscribers":    {Type: schema.TypeBool, Optional: true},
			"notifications_from_email":     {Type: schema.TypeString, Optional: true},
			"time_zone":                    {Type: schema.TypeString, Optional: true},
			"notifications_email_footer":   {Type: schema.TypeString, Optional: true},
		},
	}
}
