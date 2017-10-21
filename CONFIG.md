# Configuration of xml files

In this file the contents of the `company.xml` and a `game.xml` are explained. For example data checkout the [example folder](/example).

All data is stored as the value of an xml tag. See following example:

```xml
<title/> => <title>My Title</title>
```

- [Company](#company)
- [Game](#game)

## Company

```xml
<?xml version="1.0" encoding="UTF-8"?>
<company>
    <title/>
    <founding-date/>
    <website>
        <label/>
        <link/>
    </website>
    <based-in/>
    <press-contact/>
    <phone/>
    <description/>
    <history>
        <paragraph>
            <header/>
            <text/>
        </paragraph>
    </history>
    <address>
        <line/>
    </address>
    <social>
        <link>
            <label/>
            <link/>
        </link>
    </social>
    <videos>
        <video>
            <title/>
            <source/> <!-- e.g. youtube -->
            <id/> <!-- e.g. youtube watch id -->
        </video>
    </videos>
    <awards>
        <award>
            <description/>
            <info/>
        </award>
    </awards>
    <quotes>
        <quote>
            <text/>
            <from/>
            <link>
                <label/>
                <link/>
            </link>
        </quote>
    </quotes>
    <additionals>
        <additional>
            <title/>
            <description/>
            <link>
                <label/>
                <link/>
            </link>
        </additional>
    </additionals>
    <credits>
        <credit>
            <person/>
            <role/>
            <link/>
        </credit>
    </credits>
    <contacts>
        <contact>
            <name/>
            <link>
                <label/>
                <link/>
            </link>
        </contact>
    </contacts>
</company>
```

## Game

```xml
<?xml version="1.0" encoding="UTF-8"?>
<game>
    <title/>
    <release-dates>
        <release-date/>
    </release-dates>
    <website>
        <label/>
        <link/>
    </website>
    <description/>
    <history>
        <paragraph>
            <header/>
            <text/>
        </paragraph>
    </history>
    <press-can-request-copy/> <!-- true / false -->
    <monetization-permission/> <!-- false, ask, non-commercial, monetize -->
    <platforms>
        <platform>
            <label/>v
            <link/>
        </platform>
    </platforms>
    <prices>
        <price/>
    </prices>
    <features>
        <feature/>
    </features>
    <videos>
        <video>
            <title/>
            <source/> <!-- e.g. youtube -->
            <id/> <!-- e.g. youtube watch id -->
        </video>
    </videos>
    <awards>
        <award>
            <description/>
            <info/>
        </award>
    </awards>
    <quotes>
        <quote>
            <text/>
            <from/>
            <link>
                <label/>
                <link/>
            </link>
        </quote>
    </quotes>
    <additionals>
        <additional>
            <title/>
            <description/>
            <link>
                <label/>
                <link/>
            </link>
        </additional>
    </additionals>
    <credits>
        <credit>
            <person/>
            <role/>
            <link/>
        </credit>
    </credits>
</game>
```
