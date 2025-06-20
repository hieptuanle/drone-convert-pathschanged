# Drone Conversion Extension: Paths Changed

A [Drone](https://drone.io/) [conversion extension](https://docs.drone.io/extensions/conversion/) to include/exclude pipelines and steps based on paths changed.

_Please note this project requires Drone server version 1.4 or higher._

## Fork of meltwater/drone-convert-pathschanged

This is a fork of the meltwater/drone-convert-pathschanged project. This fork adds support for Gitea. The original project is no longer maintained and has been archived.

## Installation

## Github Cloud

1. Create a github token via https://github.com/settings/tokens with the scope of`repo` (see [issue 13](https://github.com/meltwater/drone-convert-pathschanged/issues/13) for background).

2. Create a shared secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

3. Download and run the plugin:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --env=TOKEN=9e6eij3ckzvpe9mrhnqcis6zf8dhopmm46e3pi96 \
  --env=PROVIDER=github \
  --restart=always \
  --name=converter meltwater/drone-convert-pathschanged
```

If you wish to use an enviroment file you can pass it when starting the container :

```console
$ docker run -d \
...
  --name=converter meltwater/drone-convert-pathschanged --envfile drone.env
```

4. Update your Drone server configuration to include the plugin address and the shared secret.

```text
DRONE_CONVERT_PLUGIN_ENDPOINT=http://1.2.3.4:3000
DRONE_CONVERT_PLUGIN_SECRET=bea26a2221fd8090ea38720fc445eca6
```

## Github Server

1. Create a github token via https://your-github-server-address/settings/token with the scope of `repo`

2. Create a shares secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

3. Download ran run the plugin:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --env=TOKEN=9e6eij3ckzvpe9mrhnqcis6zf8dhopmm46e3pi96 \
  --env=PROVIDER=github \
  --env=GITHUB_SERVER=https://your-github-server-address
  --restart=always \
  --name=converter meltwater/drone-convert-pathschanged
```

4. Update your Drone server configuration to include the plugin address and the shared secret.

```text
DRONE_CONVERT_PLUGIN_ENDPOINT=http://1.2.3.4:3000
DRONE_CONVERT_PLUGIN_SECRET=bea26a2221fd8090ea38720fc445eca6
```

## Bitbucket Cloud

1.  Create an "App password" via https://bitbucket.org/account/settings/app-passwords and select only "Read" under "Repositories"

2.  Create a shared secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

3. Download and run the plugin:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --env=BITBUCKET_USER=youruser \
  --env=BITBUCKET_PASSWORD='yourpassword' \
  --env=PROVIDER=bitbucket \
  --restart=always \
  --name=converter meltwater/drone-convert-pathschanged
```

4. Update your Drone server configuration to include the plugin address and the shared secret.

```text
DRONE_CONVERT_PLUGIN_ENDPOINT=http://1.2.3.4:3000
DRONE_CONVERT_PLUGIN_SECRET=bea26a2221fd8090ea38720fc445eca6
```

## Stash (Bitbucket Server)

1. Create a Stash access token via https://your-bitbucket-address/plugins/servlet/access-tokens/manage with read-only rights

2. Create a shared secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

3. Download and run the plugin:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --env=TOKEN=9e6eij3ckzvpe9mrhnqcis6zf8dhopmm46e3pi96 \
  --env=PROVIDER=bitbucket-server \
  --env=STASH_SERVER=https://your-bitbucket-server-address
  --restart=always \
  --name=converter meltwater/drone-convert-pathschanged
```

4. Update your Drone server configuration to include the plugin address and the shared secret.

```text
DRONE_CONVERT_PLUGIN_ENDPOINT=http://1.2.3.4:3000
DRONE_CONVERT_PLUGIN_SECRET=bea26a2221fd8090ea38720fc445eca6
```

## Gitee Cloud

1. Create a gitee token via https://gitee.com/personal_access_tokens with the scope of`repo` (see [issue 13](https://github.com/meltwater/drone-convert-pathschanged/issues/13) for background).

2. Create a shared secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

3. Download and run the plugin:

```console
$   docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --env=TOKEN=b4af6fc778a7aba5f2a133d155f5b4a9cbe7becb255557e67597a4967eb50a88 \
  --env=PROVIDER=gitee \
  --restart=always \
  --name=converter meltwater/drone-convert-pathschanged
```

4. Update your Drone server configuration to include the plugin address and the shared secret.

```text
DRONE_CONVERT_PLUGIN_ENDPOINT=http://1.2.3.4:3000
DRONE_CONVERT_PLUGIN_SECRET=bea26a2221fd8090ea38720fc445eca6
```

## Gitea

1. Create a Gitea access token via https://your-gitea-server/user/settings/applications with the scope of `repo`.

2. Create a shared secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

3. Download and run the plugin:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --env=TOKEN=9e6eij3ckzvpe9mrhnqcis6zf8dhopmm46e3pi96 \
  --env=PROVIDER=gitea \
  --env=GITEA_SERVER=https://your-gitea-server \
  --restart=always \
  --name=converter meltwater/drone-convert-pathschanged
```

4. Update your Drone server configuration to include the plugin address and the shared secret.

```text
DRONE_CONVERT_PLUGIN_ENDPOINT=http://1.2.3.4:3000
DRONE_CONVERT_PLUGIN_SECRET=bea26a2221fd8090ea38720fc445eca6
```

## Examples

This extension uses [doublestar](https://github.com/bmatcuk/doublestar) for matching paths changed in your commit range, refer to their documentation for all supported patterns.

### `include`

Only run a pipeline when `README.md` is changed:

```yaml
---
kind: pipeline
name: readme

trigger:
  paths:
    include:
      - README.md

steps:
  - name: message
    image: busybox
    commands:
      - echo "README.md was changed”
```

Only run a pipeline step when `README.md` is changed:

```yaml
---
kind: pipeline
name: readme

steps:
  - name: message
    image: busybox
    commands:
      - echo "README.md was changed”
    when:
      paths:
        include:
          - README.md
```

Same as above, but with an implicit `include`:

```yaml
---
kind: pipeline
name: readme

steps:
  - name: message
    image: busybox
    commands:
      - echo "README.md was changed”
    when:
      paths:
        - README.md
```

### `include` and `exclude`

Run a pipeline step when `.yml` files are changed in the root, except for `.drone.yml`:

```yaml
---
kind: pipeline
name: yaml

steps:
  - name: message
    image: busybox
    commands:
      - echo "A .yml file in the root of the repo other than .drone.yml was changed"
    when:
      paths:
        include:
          - "*.yml"
        exclude:
          - .drone.yml
```

### `depends_on`

When using [`depends_on`](https://docker-runner.docs.drone.io/configuration/parallelism/) in a pipeline step, ensure the `paths` rules match, otherwise your steps may run out of order.

Only run two steps when `README.md` is changed, one after the other:

```yaml
---
kind: pipeline
name: depends_on

steps:
  - name: message
    image: busybox
    commands:
      - echo "README.md was changed”
    when:
      paths:
        include:
          - README.md

  - name: depends_on_message
    depends_on:
      - message
    image: busybox
    commands:
      - echo "This step runs after the message step"
    when:
      paths:
        include:
          - README.md
```

## Changesets

The changeset is generated by comparing the list of files changed between the commit [before] the patch is applied and the commit [after] the patch is applied.
As a result, the changeset for a commit may be different depending on which type of event triggered the build.

For example, the `push` and `tag` events may generate a changeset against the previous commit, where as the `pull_request` event may generate a changeset against the source branch. For more specifics on how before and after are set, review the webhook [parser].

[before]: https://docs.drone.io/pipeline/environment/reference/drone-commit-before/
[after]: https://docs.drone.io/pipeline/environment/reference/drone-commit-after/
[parser]: https://github.com/drone/drone/blob/ca454594021099909fb4ee9471720cacfe3207bd/service/hook/parser/parse.go

## Known issues

### Empty commits

Be careful when making empty commits with [`git commit --allow-empty`](https://git-scm.com/docs/git-commit#Documentation/git-commit.txt---allow-empty). When an empty commit is made, no files have changed, so this plugin will return the unmodified `.drone.yml` back to the drone server process.

This can lead to potentially unexpected behavior, since any `include` or `exclude` rules will effectively be ignored.

### YAML anchors

There is a problem in the YAML library where ordering matters during unmarshaling, see https://github.com/meltwater/drone-convert-pathschanged/issues/18

This syntax will fail:

```yaml
anchor: &anchor
  image: busybox
  settings:
    foo: bar

- name: test
  <<: *anchor
  when:
    event: push
    branch: master
```

But this will succeed:

```yaml
anchor: &anchor
  image: busybox
  settings:
    foo: bar

- <<: *anchor
  name: test
  when:
    event: push
    branch: master
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) to understand how to submit pull requests to us, and also see our [code of conduct](CODE_OF_CONDUCT.md).

### Protected Repos

When this plugin is used in conjunction with [protected repos](https://docs.drone.io/signature/),
signature validation will frequently fail.

This occurs due to Drone's order of operations, in that the Drone file's
signature is checked after the this plugin has rewritten sections based on
the paths-changed triggers, resulting in a different signature for the file.
