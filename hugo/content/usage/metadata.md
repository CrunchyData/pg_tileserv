---
title: "Service Metadata"
date:
draft: false
weight: 100
---

You can explore the contents of the tile server using:

* an HTML web interface for humans; and
* a JSON API for computers.

The JSON API is useful for clients that auto-configure based on the service metadata. In fact, the HTML web interface itself is an example of such an auto-configuring interface: it reads the JSON and uses that to set up the web map visualization and interface elements.

## Web Interface

After start-up, you can connect to the server and explore the published tables and functions in the database via a web interface at:

* http://localhost:7800

Click the "preview" link of any of the layer entries to see a web map view of the layer. The "json" link provides a direct link to the JSON metadata for that layer.

## Layers List

A top-level list of layers is available in JSON at:

* http://localhost:7800/index.json

The index JSON returns the minimum information about each layer.
```json
{
    "public.ne_50m_admin_0_countries" : {
        "name" : "ne_50m_admin_0_countries",
        "schema" : "public",
        "type" : "table",
        "id" : "public.ne_50m_admin_0_countries",
        "description" : "Natural Earth country data",
        "detailurl" : "http://localhost:7800/public.ne_50m_admin_0_countries.json"
    }
}
```

* The `detailurl` provides more detailed metadata for table and function layers.
* The `description` field is read from the `comment` value of the table. To set a comment on a table, use the `COMMENT` command:
    ```sql
    COMMENT ON TABLE ne_50m_admin_0_countries IS 'This is my comment';
    ```
