---
organization: matteoredaelli
category: ["saas"]
icon_url: "/images/plugins/matteoredaelli/talendcloud.svg"
brand_color: "#03363D"
display_name: "Talend"
name: "talend"
description: "Steampipe plugin for querying tickets, users, groups and more from Talend."
og_description: "Use SQL to query tickets, users and more from Talend. Open source CLI. No DB required." 
og_image: "/images/plugins/turbot/talend-social-graphic.png"
---

# Talend + Turbot

[Talend](https://www.talend.com/) is a customer service SaaS platform with 200,000+ customers. It enables organizations to provide customer service via text, mobile, phone, email, live chat, social media.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  id,
  created_at,
  assignee_id,
  subject
from
  talend_ticket
where
  status = 'open';
```
```
+------+---------------------+--------------+----------------------------------+
| id   | created_at          | assignee_id  | subject                          |
+------+---------------------+--------------+----------------------------------+
| 4582 | 2021-04-09 14:53:25 | 383110186421 | Need help with Export            |
| 4579 | 2021-04-08 21:19:23 | 383110186421 | DB and Workspace Scaling Options |
| 4577 | 2021-04-07 23:27:21 | 383110186421 | How do i create a Report?        |
+------+---------------------+--------------+----------------------------------+
```


## Documentation

- **[Table definitions & examples →](/plugins/turbot/talend/tables)**

## Get started

### Install

Download and install the latest Talend plugin:

```bash
steampipe plugin install talend
```

### Credentials

| Item | Description |
| - | - |
| Credentials | Talend requires an [API token](https://support.talend.com/hc/en-us/articles/226022787-Generating-a-new-API-token-), subdomain and email for all requests. |
| Permissions | You must be an administrator of your domain to create an API token. |
| Radius | A Talend connection is scoped to a single Talend account, with a single set of credentials. |
| Resolution |  1. Credentials specified in environment variables e.g. `ZENDESK_TOKEN`.<br />2. Credentials in the Steampipe configuration file (`~/.steampipe/config/talend.spc`) |

### Configuration

Installing the latest aws plugin will create a config file (`~/.steampipe/config/talend.spc`) with a single connection named `talend`:

```hcl
connection "talend" {
  plugin = "talend"
  base_url  = "https://eu...."
  api_key      = "b23kj4"
}
```

- `base_url` - The base_url name of your Talend account.
- `api_key` - api_key address of agent user who have permission to access the API.

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-talend
* Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
