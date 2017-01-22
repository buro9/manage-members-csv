manage-members-csv
==================

Allows admins of sites to manage members using CSV files.

Requirements
------------

You will need a CSV file that contains two columns:
1. First column contains the email addresses of the users to add/remove
2. Second column contains a `1` (one) or `0` (zero) where 1 = grant access, and 0 = revoke access

This will look like this:

```csv
"someone@example.org",1
"person@example.org",0
```

Then you will need a configuration file called `manage-members-csv.json`, which looks like this:

```javascript
{
	"url": "https://site.microco.sm/",
	"access_token": "f4e466534b511b8d6f8771413acb6e0a"
}
```

The `url` needs to be for the subdomain of the Microcosm host. This is typically the site's URL, unless the site is using a custom domain in which case it is the original URL. i.e. LFGSS is using a custom domain but view-source on https://www.lfgss.com/ reveals a `<meta name="subdomain" content="https://lfgss.microco.sm"/>` that shows us that the URL is `https://lfgss.microco.sm/`.

The `access_token` needs to belong to the profile that owns the site.

Usage
-----

Ensure that the executable and the config file are in the same directory, and pass in the CSV file as an argument.

i.e.

`manage-members-csv members.csv`

This will then apply all of the changes and print a summary on screen of any issues encountered or the result of the program running.