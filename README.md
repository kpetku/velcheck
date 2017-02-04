# velcheck
velcheck is a command line utility that is capable of checking multiple sites on a shared hosting webserver against the Joomla! Vulnerable Extension List.

velcheck only scans the local machine where it is executed.  Remote scanning is beyond the scope of this utility.

# Installation
```go get github.com/kpetku/velcheck```

# Usage
```./velcheck --help``` for options.

# Information regarding the default feed
* At this time of writing, the remote GPL licensed JSON format feed containing the Joomla! Vulnerable Extension List comes courtesy of Phil Taylor.
* Unless otherwise specified with an alternative ```--feed=```, this program will make an outbound https connection by default to [vel.myjoomla.io](https://vel.myjoomla.io/).
* Please respect the terms of [docs.myjoomla.io](https://docs.myjoomla.io/) specifically the section that states: "This JSON is provided for free, don’t abuse this by hammering it with massive of calls please. Be nice. Cache."

To download a local copy of the feed into your current working directory, type:
```curl -o velcheck.json https://vel.myjoomla.io/``` and then run the program with the options of: ```./velcheck --feed=velcheck.json```

# Disclaimer
* velcheck and the author is not affiliated with or endorsed by The Joomla! Project™. Any products and services provided through this site are not supported or warrantied by The Joomla! Project or Open Source Matters, Inc. Use of the Joomla!® name, symbol, logo and related trademarks is permitted under a limited license granted by Open Source Matters, Inc.
* velcheck and the author of this utility is in no way affiliated with or endorsed by Phil Taylor, myJoomla.com, or WPScan.

# Thanks
The author would like to thank the JVEL team for all of their hard work.  If you happen to be a PHP developer, please consider volunteering over at the [Vulnerable Extensions List team](https://vel.joomla.org/articles/1868-the-vulnerable-extensions-list-team-is-looking-for-new-members).