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

Go-guerilla receive and route mails, Cobra provides complex documented CLI options, Viper handles ENV and configuration files.

Licence
-------

3 terms  BSD Licence. ©2018 Mathieu Lecarme.