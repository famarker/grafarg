+++
title = "Standard field options"
keywords = ["grafarg", "table options", "documentation", "format tables"]
weight = 400
+++

# Standard field options

This section explains all available field options. They are listed in alphabetical order.

You can apply standard field options to most built-in Grafarg panels. Some older panels and community panels that have not updated to the new panel and data model will be missing either all or some of these field options.

Most field options will not affect the visualization until you click outside of the field option box you are editing or press Enter.

For more information about applying these options, refer to:

- [Configure all fields]({{< relref "configure-all-fields.md" >}})
- [Configure specific fields]({{< relref "configure-specific-fields.md" >}})

> **Note:** We are constantly working to add and expand options for all visualization, so all options might not be available for all visualizations.

## Decimals

Number of decimals to render value with. Leave empty for Grafarg to use the number of decimals provided by the data source.

To change this setting, type a number in the field and then click outside the field or press Enter.

## Data links

Lets you control the URL to which a value or visualization link.

For more information and instructions, refer to [Data links]({{< relref "../../linking/data-links.md" >}}).

## Display name

Lets you set the display title of all fields. You can use [variables]({{< relref "../../variables/_index.md" >}}) in the field title.

When multiple stats, fields, or series are shown, this field controls the title in each stat. You can use expressions like `${__field.name}` to use only the series name or the field name in title.

Given a field with a name of Temp, and labels of {"Loc"="PBI", "Sensor"="3"}

| Expression syntax            | Example                 | Renders to                     | Explanation                                                                                                                                                                                                        |
| ---------------------------- | ----------------------- | ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `${__field.displayName}`     | Same as syntax          | `Temp {Loc="PBI", Sensor="3"}` | Displays the field name, and labels in `{}` if they are present. If there is only one label key in the response, then for the label portion, Grafarg displays the value of the label without the enclosing braces. |
| `${__field.name}`            | Same as syntax          | `Temp`                         | Displays the name of the field (without labels).                                                                                                                                                                   |
| `${__field.labels}`          | Same as syntax          | `Loc="PBI", Sensor="3"`        | Displays the labels without the name.                                                                                                                                                                              |
| `${__field.labels.X}`        | `${__field.labels.Loc}` | `PBI`                          | Displays the value of the specified label key.                                                                                                                                                                     |
| `${__field.labels.__values}` | Same as Syntax          | `PBI, 3`                       | Displays the values of the labels separated by a comma (without label keys).                                                                                                                                       |

If the value is an empty string after rendering the expression for a particular field, then the default display method is used.

## Max

Lets you set the maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields

## Min

Lets you set the minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields

## No value

Enter what Grafarg should display if the field value is empty or null.

## Unit

Lets you choose what unit a field should use. Click in the **Unit** field, then drill down until you find the unit you want. The unit you select is applied to all fields except time.

### Custom units

You can use the unit dropdown to also specify custom units, custom prefix or suffix and date time formats.

To select a custom unit enter the unit and select the last `Custom: xxx` option in the dropdown.

- If y u want a space -> If you want a space
- `suffix:<suffix>` for custom unit that should go after value.
- `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`. See [formats](https://momentjs.com/docs/#/displaying/) for the format syntax and options.
- `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the
  source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that
  SI scale character.
- `count:<unit>` for a custom count unit.
- `currency:<unit>` for custom a currency unit.

You can also paste a native emoji in the unit picker and pick it as a custom unit:

{{< figure src="/static/img/docs/v66/custom_unit_burger2.png" max-width="600px" caption="Custom unit emoji" >}}

### String units

Grafarg can sometime be too aggressive in parsing strings and displaying them as numbers. To make Grafarg show the original string create a field override and add a unit property with the `string` unit.

## Color scheme

> **Note:** Only available in Grafarg 7.3+.

{{< figure src="/static/img/docs/v73/color_scheme_dropdown.png" max-width="350px" caption="Color scheme" class="pull-right" >}}

The color scheme option defines how Grafarg colors series or fields. There are multiple modes here that work very differently and their utility depends largely on what visualization you currently have selected.

**Color by value**
In addition to deriving color from thresholds there are also continuous (gradient) color schemes. Useful for visualizations that color individual values. For example stat panels and the table panel. Continuous color modes use the percentage of a value relative to min and max to interpolate a color.

<div class="clearfix"></div>

| Color mode  | Description                    |
| ------------------ | ----------------------- |
| **Single color**   | Specify a single color, useful in an override rule  |
| **From thresholds**   | Informs Grafarg to take the color from the matching threshold  |
| **Classic palette**   | Grafarg will assign color by looking up a color in a palette by series index. Useful for Graphs and pie charts and other categorical data visualizations |
| **Green-Yellow-Red (by value)** | Continuous color scheme |
| **Blue-Yellow-Red (by value)** | Continuous color scheme |
| **Blues (by value)**  | Continuous color scheme (panel background to blue) |
| **Reds (by value)**   |  Continuous color scheme (panel background color to blue) |
| **Greens (by value)** | Continuous color scheme (panel background color to blue) |
| **Purple (by value)** |  Continuous color scheme (panel background color to blue) |

## Thresholds

Thresholds allow you to change the color of a field based on the value.

For more information and instructions, refer to [Thresholds]({{< relref "../thresholds.md" >}}).

## Value mapping

Lets you set rules that translate a field value or range of values into explicit text. You can add more than one value mapping.

- **Mapping type -** Click an option.
  - **Value -** Enter a value. If the field value is greater than or equal to the value, then the **Text** is displayed.
  - **From** and **To -** Enter a range. If the field value is between or equal to the values in the range, then the **Text** is displayed.
- **Text -** Text that is displayed if the conditions are met in a field. This field accepts variables.
