## [1.0.1](https://github.com/auto-staging/stagectl/compare/1.0.0...1.0.1) (2019-01-09)


### Bug Fixes

* replaced exit 1 with log.Fatal() to solve missing error logging ([fcea664](https://github.com/auto-staging/stagectl/commit/fcea664))

# 1.0.0 (2019-01-09)


### Bug Fixes

* fixed args check (requires 2 args) ([6e5bd8f](https://github.com/auto-staging/stagectl/commit/6e5bd8f))
* fixed wrong function names ([849d342](https://github.com/auto-staging/stagectl/commit/849d342))
* fixed wrong naming for add repository ([5580abd](https://github.com/auto-staging/stagectl/commit/5580abd))
* stop after an error getting the environment ([0194c7f](https://github.com/auto-staging/stagectl/commit/0194c7f))


### Features

* added add subcommand to create new repository based on definition json ([49aa175](https://github.com/auto-staging/stagectl/commit/49aa175))
* added flag to only get single repository ([0195507](https://github.com/auto-staging/stagectl/commit/0195507))
* added function to ask user for environment update input ([1f59268](https://github.com/auto-staging/stagectl/commit/1f59268))
* added get environments for repo command ([8256820](https://github.com/auto-staging/stagectl/commit/8256820))
* added get repositories command ([59cf325](https://github.com/auto-staging/stagectl/commit/59cf325))
* added get subcommand to get current tower config ([809dd6b](https://github.com/auto-staging/stagectl/commit/809dd6b))
* added output flag and flags to filter by environment for status ([c754081](https://github.com/auto-staging/stagectl/commit/c754081))
* added remove subcommand to delete environment ([e2d28ce](https://github.com/auto-staging/stagectl/commit/e2d28ce))
* added remove subcommand to delete repository ([659f8e5](https://github.com/auto-staging/stagectl/commit/659f8e5))
* added start environment and stop environment commands ([30d7c5b](https://github.com/auto-staging/stagectl/commit/30d7c5b))
* added status command ([fd287d7](https://github.com/auto-staging/stagectl/commit/fd287d7))
* added subcommand to add environment for repository ([40e6a1f](https://github.com/auto-staging/stagectl/commit/40e6a1f))
* added subcommand to get gerneral config ([afd6ed9](https://github.com/auto-staging/stagectl/commit/afd6ed9))
* added subcommand to update general configuration ([abd749b](https://github.com/auto-staging/stagectl/commit/abd749b))
* added subcommand to update tower configuration ([160c274](https://github.com/auto-staging/stagectl/commit/160c274))
* added update subcommand for environments ([3492446](https://github.com/auto-staging/stagectl/commit/3492446))
* added update subcommand for repository ([2c4f0ec](https://github.com/auto-staging/stagectl/commit/2c4f0ec))
* added version flag ([d353f81](https://github.com/auto-staging/stagectl/commit/d353f81))
* cobra init ([8da01c5](https://github.com/auto-staging/stagectl/commit/8da01c5))
* improved help output with command examples ([6e09686](https://github.com/auto-staging/stagectl/commit/6e09686))
* improved output for api error message ([1f140bf](https://github.com/auto-staging/stagectl/commit/1f140bf))
