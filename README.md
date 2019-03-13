# Serverless Go CRUD

This is an example Go CRUD application deployed to AWS using the
[serverless framework][1] based on [Yos Riady's example][3] in his book
[*A Practical Guide to Serverless Go.*][2]

## Build and Deploy

```bash
$ make deploy
```

## Changes from Yos Riady's Example

I made a few changes/deviations from [Yos Riady's example app][3]:

- Migrated from `dep` to `mod` using [these instructions][dep2mod].
- Replaced `./scripts/build.sh` and `./scripts/deploy.sh` with a
  `Makefile`.
- Moved handlers into separate directories so that multiple, different
  `main` package files weren't in the same directory.
- Instead of using individual package includes in `serverless.yml`,
  include `./bin/**`.

[1]: https://serverless.com
[2]: https://leanpub.com/serverless-go
[3]: https://github.com/yosriady/serverless-crud-go
[dep2mod]:
https://blog.callr.tech/migrating-from-dep-to-go-1.11-modules/
