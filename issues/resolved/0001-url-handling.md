# URL Handling

Currently, correct functioning relies on then end user specifying both the root URL and the endpoint URLs in a consistent manner (with regard to the trailing and/or leading slash) such that when they are concatenated, the resulting URL works.

For example:

- Root: `http://root.url`
- EP: `/endpoint`

OR:

- Root: `http://root.url/`
- EP: `endpoint`

If these approaches are not used consistently, the resulting URL may be `http://root.url//endpoint` or (even worse) `http://root.urlendpoint`.

The argument could be made that this is up to the user to manage -- but there is an easy fix which should not cause any problems going forward.
