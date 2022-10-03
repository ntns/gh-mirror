# gh-mirror
Tool to mirror Github repositories locally

# use cases
Keeping a local backup of Github repositories

# installation
`go install github.com/ntns/gh-mirror`

# usage
```
$ gh-mirror -init                # init gh-mirror directory and config
$ gh-mirror -add username/repo   # add repo to gh-mirror
$ gh-mirror -list                # list repos added to gh-mirror
$ gh-mirror                      # mirror and update repos
```

You can use gh-mirror from any directory. Repos will always be saved at ~/gh-mirror/username/repo.

# how it works
New repositories are cloned using `git clone --mirror`

Existing repositories are kept up to date using `git remote update`

That's it!


