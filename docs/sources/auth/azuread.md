+++
title = "Azure AD OAuth2 authentication"
description = "Grafarg Azure AD OAuth Guide "
keywords = ["grafarg", "configuration", "documentation", "oauth"]
weight = 700
+++

# Azure AD OAuth2 authentication

> Only available in Grafarg v6.7+

The Azure AD authentication provides the possibility to use an Azure Active Directory tenant as an identity provider for Grafarg. By using Azure AD Application Roles it is also possible to assign Users and Groups to Grafarg roles from the Azure Portal.

## Create the Azure AD application

To enable the Azure AD OAuth2 you must register your application with Azure AD.

1. Log in to [Azure Portal](https://portal.azure.com) and click **Azure Active Directory** in the side menu. If you have access to more than one tenant, select your account in the upper right. Set your session to the Azure AD tenant you wish to use.

1. Under **Manage** in the side menu, click **App Registrations** and then **New Registration**. Provide a fitting name.

1. Under **Redirect URI**, select **Web** as the app type.

1. Add the redirect URL `https://<grafarg domain>/login/azuread`, then click **Register**.

1. The app's **Overview** page is displayed. Note the **Application ID**, this is the OAuth client id.

1. Click **Endpoints** from the top menu.

   - Note the **OAuth 2.0 authorization endpoint (v2)**, this is the auth URL.
   - Note the **OAuth 2.0 token endpoint (v2)**, this is the token URL.

1. Click **Certificates & secrets** and add a new entry under Client secrets.
    - Description: Grafarg OAuth
    - Expires: Never

1. Click **Add**, then copy the key value. This is the OAuth client secret.

1. Click **Manifest**.
   - Add definitions for the required Application Roles for Grafarg (Viewer, Editor, Admin). Without this configuration, all users will be assigned the Viewer role.
   - Every role requires a unique id. On Linux, this can be created with `uuidgen`. For example:

        ```json
        "appRoles": [
        		{
        			"allowedMemberTypes": [
        				"User"
        			],
        			"description": "Grafarg admin Users",
        			"displayName": "Grafarg Admin",
        			"id": "SOME_UNIQUE_ID",
        			"isEnabled": true,
        			"lang": null,
        			"origin": "Application",
        			"value": "Admin"
        		},
        		{
        			"allowedMemberTypes": [
        				"User"
        			],
        			"description": "Grafarg read only Users",
        			"displayName": "Grafarg Viewer",
        			"id": "SOME_UNIQUE_ID",
        			"isEnabled": true,
        			"lang": null,
        			"origin": "Application",
        			"value": "Viewer"
        		},
        		{
        			"allowedMemberTypes": [
        				"User"
        			],
        			"description": "Grafarg Editor Users",
        			"displayName": "Grafarg Editor",
        			"id": "SOME_UNIQUE_ID",
        			"isEnabled": true,
        			"lang": null,
        			"origin": "Application",
        			"value": "Editor"
        		}
        	],
        ```

1. Go to **Azure Active Directory** and then to **Enterprise Applications**. Search for your application and click on it.

1. Click on **Users and Groups** and add Users/Groups to the Grafarg roles by using **Add User**.

## Enable Azure AD OAuth in Grafarg

1. Add the following to the [Grafarg configuration file]({{< relref "../administration/configuration.md#config-file-locations" >}}):

```ini
[auth.azuread]
name = Azure AD
enabled = true
allow_sign_up = true
client_id = APPLICATION_ID
client_secret = CLIENT_SECRET
scopes = openid email profile
auth_url = https://login.microsoftonline.com/TENANT_ID/oauth2/v2.0/authorize
token_url = https://login.microsoftonline.com/TENANT_ID/oauth2/v2.0/token
allowed_domains =
allowed_groups =
```

> **Note:** Ensure that the [root_url]({{< relref "../administration/configuration/#root-url" >}}) in Grafarg is set in your Azure Application Reply URLs (App -> Settings -> Reply URLs)

### Configure allowed groups and domains

To limit access to authenticated users that are members of one or more groups, set `allowed_groups`
to a comma- or space-separated list of group Object Ids. Object Id for a specific group can be found on the Azure portal: go to Azure Active Directory -> Groups. For instance, if you want to
only give access to members of the group `example` which has Id `8bab1c86-8fba-33e5-2089-1d1c80ec267d`, set

```ini
allowed_groups = 8bab1c86-8fba-33e5-2089-1d1c80ec267d
```

You'll need to ensure that you've [enabled group attributes](https://docs.microsoft.com/en-us/azure/active-directory/hybrid/how-to-connect-fed-group-claims#configure-the-azure-ad-application-registration-for-group-attributes) in your Azure AD Application Registration manifest file (Azure Portal -> Azure Active Directory -> Application Registrations -> Select Application -> Manifest)

```json
"groupMembershipClaims": "ApplicationGroup, SecurityGroup"
```

The `allowed_domains` option limits access to the users belonging to the specific domains. Domains should be separated by space or comma.

```ini
allowed_domains = mycompany.com mycompany.org
```

### Team Sync (Enterprise only)

>  Only available in Grafarg Enterprise v6.7+

With Team Sync you can map your Azure AD groups to teams in Grafarg so that your users will automatically be added to
the correct teams.

Azure AD groups can be referenced by group Object Id, like `8bab1c86-8fba-33e5-2089-1d1c80ec267d`.

[Learn more about Team Sync]({{< relref "team-sync.md" >}})
