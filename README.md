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


Licence
-------

3 terms  BSD Licence. ©2018 Mathieu Lecarme.