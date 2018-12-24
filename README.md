# Vergen

Generates version numbers for use with web service development.

# Usage

## Release Version `vergen release`

A release version is of the form `YY.MM.DD-$Revision` where `$Revision >= 0`. The date component is always computed from 
UTC. The purpose of `$Revision` is so in a given day multiple releases can occur with unique version numbers. In theory 
`$Revision` could be any integer such as seconds since midnight, however, in practice this program use an integer 
computed off the current number of release tags for the current date.

## Preview Version `vergen preview`

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

