# Pretty Pocket

When exporting data from [Pocket](https://getpocket.com/export), the output file is an HTML file containing the saved URL's without any tags or other metadata visible

[Pretty-Pocket](https://github.com/n89nanda/pretty-pocket) parses the html file and retrieves the metadata like tags, time-added and writes to new Json file which is more clear!

# Usage

`pretty-pocket ril_export.html`


# How to get the html export file

- Go to [Pocket Export Link](https://getpocket.com/export)
- Click on `Export HTML file` which will download a `html` file
- Move the export file to the same directory as `pretty-pocket`


# Install 

- Download the correct executable from the [releases](https://github.com/n89nanda/pretty-pocket/releases) for your CPU/OS type
- Unzip the archive and navigate to the directory in the terminal
- Run `./pretty-pocket ril_export.html`
- Note: In MacOS, There may be a systerm preferences pop-up to enable running apps from third-party developers. 