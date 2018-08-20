#!perl
use strict;
use warnings;

use lib 't';
use helper;

my $tmp = workdir;
ok -d "t/repos/pipeline-test", "pipeline-test repo exists" or die;
chdir "t/repos/pipeline-test" or die;

runs_ok "genesis repipe --dry-run --config ci/aws/pipeline" and # {{{
runs_ok "genesis repipe --dry-run --config ci/aws/pipeline >$tmp/pipeline.yml" and
yaml_is get_file("$tmp/pipeline.yml"), <<'EOF', "pipeline generated for aws/pipeline (no smoke-tests, untagged)";
groups:
- jobs:
  - client-aws-1-preprod
  - client-aws-1-prod
  - client-aws-1-sandbox
  name: '*'
jobs:
- name: client-aws-1-preprod
  public: true
  serial: true
  plan:
  - aggregate:
    - get: code-changes
    - get: client-aws-1-preprod-changes
      passed:
      - client-aws-1-sandbox
  - config:
      image_resource:
        source:
          repository: starkandwayne/concourse
          tag: latest
        type: docker-image
      inputs:
      - name: code-changes
      - name: config-changes
      outputs:
      - name: out
      params:
        BOSH_CONFIG: |
          auth:
            https://preprod.example.com:25555:
              username: pp-admin
              password: Ahti2eeth3aewohnee1Phaec
          alias:
            target:
              default: https://preprod.example.com:25555
        BOSH_TARGET: default
        CURRENT_ENV: client-aws-1-preprod
        GIT_BRANCH: master
        GIT_PRIVATE_KEY: |
          -----BEGIN RSA PRIVATE KEY-----
          lol. you didn't really think that
          we'd put the key here, in a test,
          did you?!
          -----END RSA PRIVATE KEY-----
        VAULT_ADDR: https://127.0.0.1:8200
        VAULT_APP_ID: concourse
        VAULT_SKIP_VERIFY: null
        VAULT_USER_ID: aws-1
        WORKING_DIR: out/git
      platform: linux
      run:
        args:
        - ci
        - pipeline
        - stage1
        path: code/bin/genesis
    task: client-aws-1-preprod
  - params:
      rebase: true
      repository: out/git
    put: git

- name: client-aws-1-prod
  public: true
  serial: true
  plan:
  - aggregate:
    - get: code-changes
    - get: client-aws-1-prod-changes
      passed:
      - client-aws-1-preprod
  - config:
      image_resource:
        source:
          repository: starkandwayne/concourse
          tag: latest
        type: docker-image
      inputs:
      - name: code-changes
      - name: config-changes
      outputs:
      - name: out
      params:
        BOSH_CONFIG: |
          auth:
            https://prod.example.com:25555:
              username: pr-admin
              password: eeheelod3veepaepiepee8ahc3rukaefo6equiezuapohS2u
          alias:
            target:
              default: https://prod.example.com:25555
        BOSH_TARGET: default
        CURRENT_ENV: client-aws-1-prod
        GIT_BRANCH: master
        GIT_PRIVATE_KEY: |
          -----BEGIN RSA PRIVATE KEY-----
          lol. you didn't really think that
          we'd put the key here, in a test,
          did you?!
          -----END RSA PRIVATE KEY-----
        VAULT_ADDR: https://127.0.0.1:8200
        VAULT_APP_ID: concourse
        VAULT_SKIP_VERIFY: null
        VAULT_USER_ID: aws-1
        WORKING_DIR: out/git
      platform: linux
      run:
        args:
        - ci
        - pipeline
        - stage1
        path: code/bin/genesis
    task: client-aws-1-prod
  - params:
      rebase: true
      repository: out/git
    put: git

- name: client-aws-1-sandbox
  public: true
  serial: true
  plan:
  - aggregate:
    - get: code-changes
    - get: client-aws-1-sandbox-changes
      passed: []
  - config:
      image_resource:
        source:
          repository: starkandwayne/concourse
          tag: latest
        type: docker-image
      inputs:
      - name: code-changes
      - name: config-changes
      outputs:
      - name: out
      params:
        BOSH_CONFIG: |
          auth:
            https://sandbox.example.com:25555:
              username: sb-admin
              password: PaeM2Eip
          alias:
            target:
              default: https://sandbox.example.com:25555
        BOSH_TARGET: default
        CURRENT_ENV: client-aws-1-sandbox
        GIT_BRANCH: master
        GIT_PRIVATE_KEY: |
          -----BEGIN RSA PRIVATE KEY-----
          lol. you didn't really think that
          we'd put the key here, in a test,
          did you?!
          -----END RSA PRIVATE KEY-----
        VAULT_ADDR: https://127.0.0.1:8200
        VAULT_APP_ID: concourse
        VAULT_SKIP_VERIFY: null
        VAULT_USER_ID: aws-1
        WORKING_DIR: out/git
      platform: linux
      run:
        args:
        - ci
        - pipeline
        - stage1
        path: code/bin/genesis
    task: client-aws-1-sandbox
  - params:
      rebase: true
      repository: out/git
    put: git

resource_types:
- name: script
  source:
    repository: cfcommunity/script-resource
  type: docker-image
- name: email
  source:
    repository: pcfseceng/email-resource
  type: docker-image
- name: slack-notification
  source:
    repository: cfcommunity/slack-notification-resource
  type: docker-image
resources:
- name: git
  source:
    branch: master
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments
  type: git
- name: code-changes
  source:
    branch: master
    paths:
    - bin/genesis
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments
  type: git
- name: client-aws-1-preprod-changes
  source:
    branch: master
    paths:
    - .genesis/cached/client-aws-1-sandbox/client.yml
    - .genesis/cached/client-aws-1-sandbox/client-aws.yml
    - .genesis/cached/client-aws-1-sandbox/client-aws-1.yml
    - client-aws-1-preprod.yml
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments
  type: git
- name: client-aws-1-prod-changes
  source:
    branch: master
    paths:
    - .genesis/cached/client-aws-1-preprod/client.yml
    - .genesis/cached/client-aws-1-preprod/client-aws.yml
    - .genesis/cached/client-aws-1-preprod/client-aws-1.yml
    - client-aws-1-prod.yml
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments
  type: git
- name: client-aws-1-sandbox-changes
  source:
    branch: master
    paths:
    - client.yml
    - client-aws.yml
    - client-aws-1.yml
    - client-aws-1-sandbox.yml
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments
  type: git
- name: slack
  source:
    url: http://127.0.0.1:1337
  type: slack-notification

EOF
# }}}
runs_ok "genesis repipe --dry-run --config ci/aws/pipeline.tagged" and # {{{
runs_ok "genesis repipe --dry-run --config ci/aws/pipeline.tagged >$tmp/pipeline.yml" and
yaml_is get_file("$tmp/pipeline.yml"), <<'EOF', "pipeline generated for aws/pipeline (no smoke-tests, tagged)";
groups:
- jobs:
  - client-aws-1-preprod
  - client-aws-1-prod
  - client-aws-1-sandbox
  name: '*'
jobs:
- name: client-aws-1-preprod
  plan:
  - aggregate:
    - get: code-changes
    - get: client-aws-1-preprod-changes
      passed:
      - client-aws-1-sandbox
  - config:
      image_resource:
        source:
          repository: starkandwayne/concourse
          tag: latest
        type: docker-image
      inputs:
      - name: code-changes
      - name: config-changes
      outputs:
      - name: out
      params:
        BOSH_CONFIG: |
          auth:
            https://preprod.example.com:25555:
              username: pp-admin
              password: Ahti2eeth3aewohnee1Phaec
          alias:
            target:
              default: https://preprod.example.com:25555
        BOSH_TARGET: default
        CURRENT_ENV: client-aws-1-preprod
        GIT_BRANCH: master
        GIT_PRIVATE_KEY: |
          -----BEGIN RSA PRIVATE KEY-----
          lol. you didn't really think that
          we'd put the key here, in a test,
          did you?!
          -----END RSA PRIVATE KEY-----
        VAULT_ADDR: https://127.0.0.1:8200
        VAULT_APP_ID: concourse
        VAULT_SKIP_VERIFY: null
        VAULT_USER_ID: aws-1
        WORKING_DIR: out/git
      platform: linux
      run:
        args:
        - ci
        - pipeline
        - stage1
        path: code/bin/genesis
    task: client-aws-1-preprod
    tags: [client-aws-1-preprod]
  - params:
      rebase: true
      repository: out/git
    put: git
  public: true
  serial: true
- name: client-aws-1-prod
  plan:
  - aggregate:
    - get: code-changes
    - get: client-aws-1-prod-changes
      passed:
      - client-aws-1-preprod
  - config:
      image_resource:
        source:
          repository: starkandwayne/concourse
          tag: latest
        type: docker-image
      inputs:
      - name: code-changes
      - name: config-changes
      outputs:
      - name: out
      params:
        BOSH_CONFIG: |
          auth:
            https://prod.example.com:25555:
              username: pr-admin
              password: eeheelod3veepaepiepee8ahc3rukaefo6equiezuapohS2u
          alias:
            target:
              default: https://prod.example.com:25555
        BOSH_TARGET: default
        CURRENT_ENV: client-aws-1-prod
        GIT_BRANCH: master
        GIT_PRIVATE_KEY: |
          -----BEGIN RSA PRIVATE KEY-----
          lol. you didn't really think that
          we'd put the key here, in a test,
          did you?!
          -----END RSA PRIVATE KEY-----
        VAULT_ADDR: https://127.0.0.1:8200
        VAULT_APP_ID: concourse
        VAULT_SKIP_VERIFY: null
        VAULT_USER_ID: aws-1
        WORKING_DIR: out/git
      platform: linux
      run:
        args:
        - ci
        - pipeline
        - stage1
        path: code/bin/genesis
    task: client-aws-1-prod
    tags: [client-aws-1-prod]
  - params:
      rebase: true
      repository: out/git
    put: git
  public: true
  serial: true
- name: client-aws-1-sandbox
  plan:
  - aggregate:
    - get: code-changes
    - get: client-aws-1-sandbox-changes
      passed: []
  - config:
      image_resource:
        source:
          repository: starkandwayne/concourse
          tag: latest
        type: docker-image
      inputs:
      - name: code-changes
      - name: config-changes
      outputs:
      - name: out
      params:
        BOSH_CONFIG: |
          auth:
            https://sandbox.example.com:25555:
              username: sb-admin
              password: PaeM2Eip
          alias:
            target:
              default: https://sandbox.example.com:25555
        BOSH_TARGET: default
        CURRENT_ENV: client-aws-1-sandbox
        GIT_BRANCH: master
        GIT_PRIVATE_KEY: |
          -----BEGIN RSA PRIVATE KEY-----
          lol. you didn't really think that
          we'd put the key here, in a test,
          did you?!
          -----END RSA PRIVATE KEY-----
        VAULT_ADDR: https://127.0.0.1:8200
        VAULT_APP_ID: concourse
        VAULT_SKIP_VERIFY: null
        VAULT_USER_ID: aws-1
        WORKING_DIR: out/git
      platform: linux
      run:
        args:
        - ci
        - pipeline
        - stage1
        path: code/bin/genesis
    task: client-aws-1-sandbox
    tags: [client-aws-1-sandbox]
  - params:
      rebase: true
      repository: out/git
    put: git
  public: true
  serial: true
resource_types:
- name: script
  source:
    repository: cfcommunity/script-resource
  type: docker-image
- name: email
  source:
    repository: pcfseceng/email-resource
  type: docker-image
- name: slack-notification
  source:
    repository: cfcommunity/slack-notification-resource
  type: docker-image
resources:
- name: git
  source:
    branch: master
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments
  type: git
- name: code-changes
  source:
    branch: master
    paths:
    - bin/genesis
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments
  type: git
- name: client-aws-1-preprod-changes
  source:
    branch: master
    paths:
    - .genesis/cached/client-aws-1-sandbox/client.yml
    - .genesis/cached/client-aws-1-sandbox/client-aws.yml
    - .genesis/cached/client-aws-1-sandbox/client-aws-1.yml
    - client-aws-1-preprod.yml
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments
  type: git
- name: client-aws-1-prod-changes
  source:
    branch: master
    paths:
    - .genesis/cached/client-aws-1-preprod/client.yml
    - .genesis/cached/client-aws-1-preprod/client-aws.yml
    - .genesis/cached/client-aws-1-preprod/client-aws-1.yml
    - client-aws-1-prod.yml
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments
  type: git
- name: client-aws-1-sandbox-changes
  source:
    branch: master
    paths:
    - client.yml
    - client-aws.yml
    - client-aws-1.yml
    - client-aws-1-sandbox.yml
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments
  type: git
- name: slack
  source:
    url: http://127.0.0.1:1337
  type: slack-notification

EOF
# }}}
runs_ok "genesis repipe --dry-run --config ci/aws/pipeline.tests" and # {{{
runs_ok "genesis repipe --dry-run --config ci/aws/pipeline.tests >$tmp/pipeline.yml" and
yaml_is get_file("$tmp/pipeline.yml"), <<'EOF', "pipeline generated for aws/pipeline (smoke-tests, untagged)";
groups:
- jobs:
  - client-aws-1-preprod
  - client-aws-1-prod
  - client-aws-1-sandbox
  name: '*'
jobs:
- name: client-aws-1-preprod
  public: true
  serial: true
  plan:
  - aggregate:
    - get: code-changes
    - get: client-aws-1-preprod-changes
      passed:
      - client-aws-1-sandbox
  - config:
      image_resource:
        source:
          repository: starkandwayne/concourse
          tag: latest
        type: docker-image
      inputs:
      - name: code-changes
      - name: config-changes
      outputs:
      - name: out
      params:
        BOSH_CONFIG: |
          auth:
            https://preprod.example.com:25555:
              username: pp-admin
              password: Ahti2eeth3aewohnee1Phaec
          alias:
            target:
              default: https://preprod.example.com:25555
        BOSH_TARGET: default
        CURRENT_ENV: client-aws-1-preprod
        GIT_BRANCH: master
        GIT_PRIVATE_KEY: |
          -----BEGIN RSA PRIVATE KEY-----
          lol. you didn't really think that
          we'd put the key here, in a test,
          did you?!
          -----END RSA PRIVATE KEY-----
        VAULT_ADDR: https://127.0.0.1:8200
        VAULT_APP_ID: concourse
        VAULT_SKIP_VERIFY: null
        VAULT_USER_ID: aws-1
        WORKING_DIR: out/git
      platform: linux
      run:
        args:
        - ci
        - pipeline
        - stage1
        path: code/bin/genesis
    task: client-aws-1-preprod

  - put: git
    params:
      rebase: true
      repository: out/git

  # run the smoke tests against the deployment
  - task: smoke-test
    config:
      platform: linux
      image_resource:
        type: docker-image
        source:
          repository: starkandwayne/concourse
          tag:        latest

      inputs:
        - name: out

      run:
        path: out/git/bin/genesis
        args: [ci, pipeline, run-smoke-test]

      params:
        CURRENT_ENV: client-aws-1-preprod
        ERRAND_NAME: a-testing-errand-for-the-ages

        BOSH_TARGET: default
        BOSH_CONFIG: |
          auth:
            https://preprod.example.com:25555:
              username: pp-admin
              password: Ahti2eeth3aewohnee1Phaec
          alias:
            target:
              default: https://preprod.example.com:25555

- name: client-aws-1-prod
  public: true
  serial: true
  plan:
  - aggregate:
    - get: code-changes
    - get: client-aws-1-prod-changes
      passed:
      - client-aws-1-preprod
  - config:
      image_resource:
        source:
          repository: starkandwayne/concourse
          tag: latest
        type: docker-image
      inputs:
      - name: code-changes
      - name: config-changes
      outputs:
      - name: out
      params:
        BOSH_CONFIG: |
          auth:
            https://prod.example.com:25555:
              username: pr-admin
              password: eeheelod3veepaepiepee8ahc3rukaefo6equiezuapohS2u
          alias:
            target:
              default: https://prod.example.com:25555
        BOSH_TARGET: default
        CURRENT_ENV: client-aws-1-prod
        GIT_BRANCH: master
        GIT_PRIVATE_KEY: |
          -----BEGIN RSA PRIVATE KEY-----
          lol. you didn't really think that
          we'd put the key here, in a test,
          did you?!
          -----END RSA PRIVATE KEY-----
        VAULT_ADDR: https://127.0.0.1:8200
        VAULT_APP_ID: concourse
        VAULT_SKIP_VERIFY: null
        VAULT_USER_ID: aws-1
        WORKING_DIR: out/git
      platform: linux
      run:
        args:
        - ci
        - pipeline
        - stage1
        path: code/bin/genesis
    task: client-aws-1-prod
  - params:
      rebase: true
      repository: out/git
    put: git

  # run the smoke tests against the deployment
  - task: smoke-test
    config:
      platform: linux
      image_resource:
        type: docker-image
        source:
          repository: starkandwayne/concourse
          tag:        latest

      inputs:
        - name: out

      run:
        path: out/git/bin/genesis
        args: [ci, pipeline, run-smoke-test]

      params:
        CURRENT_ENV: client-aws-1-prod
        ERRAND_NAME: a-testing-errand-for-the-ages

        BOSH_TARGET: default
        BOSH_CONFIG: |
          auth:
            https://prod.example.com:25555:
              username: pr-admin
              password: eeheelod3veepaepiepee8ahc3rukaefo6equiezuapohS2u
          alias:
            target:
              default: https://prod.example.com:25555

- name: client-aws-1-sandbox
  public: true
  serial: true
  plan:
  - aggregate:
    - get: code-changes
    - get: client-aws-1-sandbox-changes
      passed: []
  - config:
      image_resource:
        source:
          repository: starkandwayne/concourse
          tag: latest
        type: docker-image
      inputs:
      - name: code-changes
      - name: config-changes
      outputs:
      - name: out
      params:
        BOSH_CONFIG: |
          auth:
            https://sandbox.example.com:25555:
              username: sb-admin
              password: PaeM2Eip
          alias:
            target:
              default: https://sandbox.example.com:25555
        BOSH_TARGET: default
        CURRENT_ENV: client-aws-1-sandbox
        GIT_BRANCH: master
        GIT_PRIVATE_KEY: |
          -----BEGIN RSA PRIVATE KEY-----
          lol. you didn't really think that
          we'd put the key here, in a test,
          did you?!
          -----END RSA PRIVATE KEY-----
        VAULT_ADDR: https://127.0.0.1:8200
        VAULT_APP_ID: concourse
        VAULT_SKIP_VERIFY: null
        VAULT_USER_ID: aws-1
        WORKING_DIR: out/git
      platform: linux
      run:
        args:
        - ci
        - pipeline
        - stage1
        path: code/bin/genesis
    task: client-aws-1-sandbox
  - params:
      rebase: true
      repository: out/git
    put: git

  # run the smoke tests against the deployment
  - task: smoke-test
    config:
      platform: linux
      image_resource:
        type: docker-image
        source:
          repository: starkandwayne/concourse
          tag:        latest

      inputs:
        - name: out

      run:
        path: out/git/bin/genesis
        args: [ci, pipeline, run-smoke-test]

      params:
        CURRENT_ENV: client-aws-1-sandbox
        ERRAND_NAME: a-testing-errand-for-the-ages

        BOSH_TARGET: default
        BOSH_CONFIG: |
          auth:
            https://sandbox.example.com:25555:
              username: sb-admin
              password: PaeM2Eip
          alias:
            target:
              default: https://sandbox.example.com:25555

resource_types:
- name: script
  type: docker-image
  source:
    repository: cfcommunity/script-resource

- name: email
  type: docker-image
  source:
    repository: pcfseceng/email-resource

- name: slack-notification
  type: docker-image
  source:
    repository: cfcommunity/slack-notification-resource

resources:
- name: git
  type: git
  source:
    branch: master
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments

- name: code-changes
  type: git
  source:
    branch: master
    paths:
    - bin/genesis
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments

- name: client-aws-1-preprod-changes
  type: git
  source:
    branch: master
    paths:
    - .genesis/cached/client-aws-1-sandbox/client.yml
    - .genesis/cached/client-aws-1-sandbox/client-aws.yml
    - .genesis/cached/client-aws-1-sandbox/client-aws-1.yml
    - client-aws-1-preprod.yml
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments

- name: client-aws-1-prod-changes
  type: git
  source:
    branch: master
    paths:
    - .genesis/cached/client-aws-1-preprod/client.yml
    - .genesis/cached/client-aws-1-preprod/client-aws.yml
    - .genesis/cached/client-aws-1-preprod/client-aws-1.yml
    - client-aws-1-prod.yml
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments

- name: client-aws-1-sandbox-changes
  type: git
  source:
    branch: master
    paths:
    - client.yml
    - client-aws.yml
    - client-aws-1.yml
    - client-aws-1-sandbox.yml
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments

- name: slack
  type: slack-notification
  source:
    url: http://127.0.0.1:1337

EOF
# }}}
runs_ok "genesis repipe --dry-run --config ci/aws/pipeline.everything" and # {{{
runs_ok "genesis repipe --dry-run --config ci/aws/pipeline.everything >$tmp/pipeline.yml" and
yaml_is get_file("$tmp/pipeline.yml"), <<'EOF', "pipeline generated for aws/pipeline (kitchen sink)";
groups:
- jobs:
  - client-aws-1-preprod
  - client-aws-1-prod
  - client-aws-1-sandbox
  name: '*'
jobs:
- name: client-aws-1-preprod
  public: true
  serial: true
  plan:
  - aggregate:
    - get: code-changes
    - get: client-aws-1-preprod-changes
      passed:
      - client-aws-1-sandbox
  - config:
      image_resource:
        source:
          repository: custom/concourse-image
          tag: rc1
        type: docker-image
      inputs:
      - name: code-changes
      - name: config-changes
      outputs:
      - name: out
      params:
        BOSH_CONFIG: |
          auth:
            https://preprod.example.com:25555:
              username: pp-admin
              password: Ahti2eeth3aewohnee1Phaec
          alias:
            target:
              default: https://preprod.example.com:25555
        BOSH_TARGET: default
        CURRENT_ENV: client-aws-1-preprod
        GIT_BRANCH: master
        GIT_PRIVATE_KEY: |
          -----BEGIN RSA PRIVATE KEY-----
          lol. you didn't really think that
          we'd put the key here, in a test,
          did you?!
          -----END RSA PRIVATE KEY-----
        VAULT_ADDR: http://myvault.myorg.com:5999
        VAULT_APP_ID: obscure-app-1
        VAULT_SKIP_VERIFY: 1
        VAULT_USER_ID: mr.awsome
        WORKING_DIR: out/git
      platform: linux
      run:
        args:
        - ci
        - pipeline
        - stage1
        path: code/bin/genesis
    task: client-aws-1-preprod
    tags: [client-aws-1-preprod]

  - put: git
    params:
      rebase: true
      repository: out/git

  # run the smoke tests against the deployment
  - task: smoke-test
    tags: [client-aws-1-preprod]
    config:
      platform: linux
      image_resource:
        type: docker-image
        source:
          repository: custom/concourse-image
          tag:        rc1

      inputs:
        - name: out

      run:
        path: out/git/bin/genesis
        args: [ci, pipeline, run-smoke-test]

      params:
        CURRENT_ENV: client-aws-1-preprod
        ERRAND_NAME: run-something-good

        BOSH_TARGET: default
        BOSH_CONFIG: |
          auth:
            https://preprod.example.com:25555:
              username: pp-admin
              password: Ahti2eeth3aewohnee1Phaec
          alias:
            target:
              default: https://preprod.example.com:25555

- name: client-aws-1-prod
  public: true
  serial: true
  plan:
  - aggregate:
    - get: code-changes
    - get: client-aws-1-prod-changes
      passed:
      - client-aws-1-preprod
  - config:
      image_resource:
        source:
          repository: custom/concourse-image
          tag: rc1
        type: docker-image
      inputs:
      - name: code-changes
      - name: config-changes
      outputs:
      - name: out
      params:
        BOSH_CONFIG: |
          auth:
            https://prod.example.com:25555:
              username: pr-admin
              password: eeheelod3veepaepiepee8ahc3rukaefo6equiezuapohS2u
          alias:
            target:
              default: https://prod.example.com:25555
        BOSH_TARGET: default
        CURRENT_ENV: client-aws-1-prod
        GIT_BRANCH: master
        GIT_PRIVATE_KEY: |
          -----BEGIN RSA PRIVATE KEY-----
          lol. you didn't really think that
          we'd put the key here, in a test,
          did you?!
          -----END RSA PRIVATE KEY-----
        VAULT_ADDR: http://myvault.myorg.com:5999
        VAULT_APP_ID: obscure-app-1
        VAULT_SKIP_VERIFY: 1
        VAULT_USER_ID: mr.awsome
        WORKING_DIR: out/git
      platform: linux
      run:
        args:
        - ci
        - pipeline
        - stage1
        path: code/bin/genesis
    task: client-aws-1-prod
    tags: [client-aws-1-prod]
  - params:
      rebase: true
      repository: out/git
    put: git

  # run the smoke tests against the deployment
  - task: smoke-test
    tags: [client-aws-1-prod]
    config:
      platform: linux
      image_resource:
        type: docker-image
        source:
          repository: custom/concourse-image
          tag:        rc1

      inputs:
        - name: out

      run:
        path: out/git/bin/genesis
        args: [ci, pipeline, run-smoke-test]

      params:
        CURRENT_ENV: client-aws-1-prod
        ERRAND_NAME: run-something-good

        BOSH_TARGET: default
        BOSH_CONFIG: |
          auth:
            https://prod.example.com:25555:
              username: pr-admin
              password: eeheelod3veepaepiepee8ahc3rukaefo6equiezuapohS2u
          alias:
            target:
              default: https://prod.example.com:25555

- name: client-aws-1-sandbox
  public: true
  serial: true
  plan:
  - aggregate:
    - get: code-changes
    - get: client-aws-1-sandbox-changes
      passed: []
  - config:
      image_resource:
        source:
          repository: custom/concourse-image
          tag: rc1
        type: docker-image
      inputs:
      - name: code-changes
      - name: config-changes
      outputs:
      - name: out
      params:
        BOSH_CONFIG: |
          auth:
            https://sandbox.example.com:25555:
              username: sb-admin
              password: PaeM2Eip
          alias:
            target:
              default: https://sandbox.example.com:25555
        BOSH_TARGET: default
        CURRENT_ENV: client-aws-1-sandbox
        GIT_BRANCH: master
        GIT_PRIVATE_KEY: |
          -----BEGIN RSA PRIVATE KEY-----
          lol. you didn't really think that
          we'd put the key here, in a test,
          did you?!
          -----END RSA PRIVATE KEY-----
        VAULT_ADDR: http://myvault.myorg.com:5999
        VAULT_APP_ID: obscure-app-1
        VAULT_SKIP_VERIFY: 1
        VAULT_USER_ID: mr.awsome
        WORKING_DIR: out/git
      platform: linux
      run:
        args:
        - ci
        - pipeline
        - stage1
        path: code/bin/genesis
    task: client-aws-1-sandbox
    tags: [client-aws-1-sandbox]
  - params:
      rebase: true
      repository: out/git
    put: git

  # run the smoke tests against the deployment
  - task: smoke-test
    tags: [client-aws-1-sandbox]
    config:
      platform: linux
      image_resource:
        type: docker-image
        source:
          repository: custom/concourse-image
          tag:        rc1

      inputs:
        - name: out

      run:
        path: out/git/bin/genesis
        args: [ci, pipeline, run-smoke-test]

      params:
        CURRENT_ENV: client-aws-1-sandbox
        ERRAND_NAME: run-something-good

        BOSH_TARGET: default
        BOSH_CONFIG: |
          auth:
            https://sandbox.example.com:25555:
              username: sb-admin
              password: PaeM2Eip
          alias:
            target:
              default: https://sandbox.example.com:25555

resource_types:
- name: script
  type: docker-image
  source:
    repository: cfcommunity/script-resource

- name: email
  type: docker-image
  source:
    repository: pcfseceng/email-resource

- name: slack-notification
  type: docker-image
  source:
    repository: cfcommunity/slack-notification-resource

resources:
- name: git
  type: git
  source:
    branch: master
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments

- name: code-changes
  type: git
  source:
    branch: master
    paths:
    - bin/genesis
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments

- name: client-aws-1-preprod-changes
  type: git
  source:
    branch: master
    paths:
    - .genesis/cached/client-aws-1-sandbox/client.yml
    - .genesis/cached/client-aws-1-sandbox/client-aws.yml
    - .genesis/cached/client-aws-1-sandbox/client-aws-1.yml
    - client-aws-1-preprod.yml
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments

- name: client-aws-1-prod-changes
  type: git
  source:
    branch: master
    paths:
    - .genesis/cached/client-aws-1-preprod/client.yml
    - .genesis/cached/client-aws-1-preprod/client-aws.yml
    - .genesis/cached/client-aws-1-preprod/client-aws-1.yml
    - client-aws-1-prod.yml
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments

- name: client-aws-1-sandbox-changes
  type: git
  source:
    branch: master
    paths:
    - client.yml
    - client-aws.yml
    - client-aws-1.yml
    - client-aws-1-sandbox.yml
    private-key: |
      -----BEGIN RSA PRIVATE KEY-----
      lol. you didn't really think that
      we'd put the key here, in a test,
      did you?!
      -----END RSA PRIVATE KEY-----
    uri: git@github.com:someco/something-deployments

- name: slack
  type: slack-notification
  source:
    url: http://127.0.0.1:1337

EOF
# }}}

output_ok "genesis describe --config ci/pipeline.all", <<EOF, "large pipelines are described properly"; # {{{
sandbox-1
  `--> dev-1
        |--> preprod-1
        |     `--> prod-1
        `--> qa-1

sandbox-2
  |--> preprod-2
  |     `--> prod-2
  `--> preprod-3
        |--> prod-3
        |--> prod-4
        `--> prod-5
EOF
# }}}
output_ok "genesis describe --config ci/aws/pipeline", <<EOF, "small pipelines are described properly"; # {{{
client-aws-1-sandbox
  `--> client-aws-1-preprod
        `--> client-aws-1-prod
EOF
# }}}

done_testing;
