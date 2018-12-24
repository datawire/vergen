# Vergen

Generates version numbers automatically based of calendar dates and tags. This makes it suitable to for use in a 
continuous delivery or deployment context where a person manually generating version numbers is either unlikely or not
useful.

# Usage

# Latest Version `vergen latest`

Examines the current tags on the Git repository and returns the latest tagged version.

# Next Version `vergen next`

Examines the current tags on the Git repository and returns the next version.

# Preview Version `vergen preview`

Creates a preview version

A preview version has the form `$CommitID.$Branch.$Authority[-$Revision]`. The `$Revision` in Preview Version is 
optional and useful in cases where a nested resource needs a unique version. For example, if you are using `vergen` to
create tags for Docker images without committing during a development loop. Because the `$Commit` does not change your
tag names would be non-unique across subsequent runs. The `$Revision` mechanism allows some unique "data" to be appended
that would result in a unique version.

| Component | Description |
| --------- | ----------- |
| CommitID  | The Git SHA in long format |
| Branch    | The Git branch name with non-alphanum chars converted to `-` |
| Authority | The person or system that generated the version |
