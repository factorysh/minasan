皆さん
======

Minasan is a SMTP server, linked to a Gitlab instance.

When you send a mail to user `{group}.{project}`, every gitlab's users of this project, with higher level than observator receive the mail.

Big picture
-----------

            +---------+    +------------+
    mail -> | Minasan | -> | SMTP Relay +--+-> Alice
            +---+-----+    +------------+  |
                | REST                     +-> Bob
                v                          |
            +---------+                    +-> Charly
            | Gitlab  |
            +---------+

Minasan doesn't send mails directly to users, nobody does that, but uses a relay, something like Postfix,
that handle greylist, bounce, reputation and all that voodoo dances.

Minasan uses Gitlab REST API, with a private token. Great for using, boring for testing.
Gitlab REST API doesn't expose user + project = ☆, what a pity.

Demo time
---------

You need Docker, golang build tools, a Gitlab instance with sufficient privileges.

Build

    make build

The application is here : `bin/minasan`

Launch mailhog

    make mailhog

The mailhog web interface is here : http://127.0.0.1:8025

You need a config file (or some ENVs, or some cli flags)

```yaml
---

gitlab_private_token: shmurtz
gitlab_domain: gitlab.example.com
smtp_out: 127.0.0.1:1025
smtp_domain: example.com
```

    ./bin/minasan -c config.yml

Send some mails

    ./debug_client.py factory.minasan@example.com

Usage
-----

Minasan uses [go-guerilla](https://github.com/flashmob/go-guerrilla),
[viper](https://github.com/spf13/viper) and [cobra](https://github.com/spf13/cobra)

Go-guerilla receives and routes incoming mails, Cobra provides complex documented CLI options, Viper handles ENV and configuration files.

```
$ ./bin/minasan -h
Send mail to gitlab projects

Usage:
  minasan [command]

Available Commands:
  gitlab      Ask gitlab wich mails are linked to a specific project
  help        Help about any command
  serve       Listen as a SMTP server

Flags:
  -c, --config string                 Config file
  -g, --gitlab_domain string          Gitlab domain (default "gitlab.example.com")
  -t, --gitlab_private_token string   Gitlab private token
  -h, --help                          help for minasan

Use "minasan [command] --help" for more information about a command.

```

And for the `serve` command:
```
./bin/minasan serve -h
Listen as a SMTP server

Usage:
  minasan serve [flags]

Flags:
  -h, --help                     help for serve
  -H, --metrics_address string   Prometheus probe listening address (default "127.0.0.1:8125")
  -d, --smtp_domain string       SMTP domain (default "gitlab.example.com")
  -i, --smtp_in string           SMTP input service (default "127.0.0.1:2525")
  -o, --smtp_out string          SMTP relay (default "127.0.0.1:25")

Global Flags:
  -c, --config string                 Config file
  -g, --gitlab_domain string          Gitlab domain (default "gitlab.example.com")
  -t, --gitlab_private_token string   Gitlab private token
```

API
---

Minasan is a simple SMTP server, unauthenticated.

You send a mail to `{group}.{project}@{domain}`. `domain` is the `-d` option, an arbitrary name, go-guerilla loves routing, and a domain is mandatory.

Licence
-------

3 terms  BSD Licence. ©2018 Mathieu Lecarme.