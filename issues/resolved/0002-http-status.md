# HTTP Status Code

Our code currently treats any status code other than `http.StatusOK` as an error
-- but this may not be the case.

It makes more sense to pass the status code back to the user and let them deal
with it appropriately.

Possibly there is some middle ground here -- is a 404 an error? Probably. But
again, passing the body and status code back to the user so they can handle it
in their own way without having to parse out an error message is probably the
better approach.
