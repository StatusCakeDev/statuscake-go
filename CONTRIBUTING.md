# How To Contribute

Contributions are welcome as we strive to make this package as useful as
possible for everyone. However time is not always on our side, and changes may
not be reviewed or merged in a timely manner.

If this package is found to be missing in functionality, please open an issue
describing the proposed change - discussing changes ahead of time reduces
friction within pull requests.

## Installation

* `git clone <repository-url>` this repository
* `cd statuscake-go`

## Linting

* `golint ./...`

## Running tests

* `go test ./...`

## Making Changes

Begin by creating a new branch. It is appreciated if branch names are written
using kebab-case.

```bash
git checkout master
git pull --rebase
git checkout -b my-new-feature
```

Make the desired change, and ensure both the linter and test suite continue to
pass. Once this requirement is met push the change back to a fork of this
repository.

```bash
git push -u origin my-new-feature
```

Finally open a pull request through the GitHub UI. Upon doing this the CI suite
will be run to ensure changes do not break current functionality.

Changes are more likely to be approve if they:

- Include tests for new functionality,
- Are accompanied with a [good commit message](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html),
- Contain few commits (preferably a single commit),
- Do not contain merge commits,
- Maintain backward compatibility.
